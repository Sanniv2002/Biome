package recovery

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sanniv2002/nether/proto/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RecoveryService struct {
	storeClient store.StoreClient
}

func NewRecoveryService(storeClient store.StoreClient) *RecoveryService {
	return &RecoveryService{
		storeClient: storeClient,
	}
}

func extractPorts(mapping string) (int, int, error) {
	pattern := regexp.MustCompile(`(\d+)/tcp -> 0.0.0.0:(\d+)`)
	matches := pattern.FindStringSubmatch(mapping)
	if len(matches) != 3 {
		return 0, 0, fmt.Errorf("invalid mapping format")
	}
	containerPort, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid container port: %v", err)
	}
	hostPort, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid host port: %v", err)
	}
	return containerPort, hostPort, nil
}

func preventFalseStart(ctx context.Context, containerId string, client store.StoreClient) bool { //Function to check if container is halted and removed by autoscaling service
	data := &store.GetAllContainersRequest{}
	containers, err := client.GetAllContainers(ctx, data)
	if err != nil {
		fmt.Println("Error")
	}
	for _, container := range containers.ContainerIds {
		if containerId == container {
			return false
		}
	}
	return true
}

func CheckAndHeal(ctx context.Context, client store.StoreClient, containerID string) bool {
	data := &store.GetContainerStatsRequest{
		ContainerId: containerID,
	}
	resp, err := client.GetContainerStats(ctx, data)
	if err != nil {
		fmt.Println(err)
	}
	arg0 := "docker"
	arg1 := "stats"
	arg2 := resp.ContainerId
	arg3 := "--no-stream"
	arg4 := "--format"
	arg5 := "{{.PIDs}}"
	cmd := exec.Command(arg0, arg1, arg2, arg3, arg4, arg5)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if strings.TrimSpace(string(output)) == "0" {
		if preventFalseStart(ctx, arg2, client) {
			return false
		}
		arg := "start"
		cmd = exec.Command(arg0, arg, arg2)
		fmt.Println("Restarting container with ID:", resp.ContainerId)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return false
		}
		portCmd := exec.Command("docker", "port", resp.ContainerId)
		portOutput, err := portCmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		portMapping := strings.TrimSpace(string(portOutput))
		containerPort, hostPort, err := extractPorts(portMapping)
		portData := &store.UpdatePortMappingsRequest{
			ContainerId:   containerID,
			HostPORT:      uint32(hostPort),
			ContainerPORT: uint32(containerPort),
		}
		if err != nil {
			fmt.Println(err)
		}
		_, err = client.UpdatePortMappings(ctx, portData)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Container is healthy with ID:", resp.ContainerId)
	}
	return true
}

func (r *RecoveryService) RecoverContainers() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := r.storeClient.GetAllContainers(ctx, &store.GetAllContainersRequest{})
	if err != nil {
		log.Printf("Error getting all containers: %v", err)
		return
	}
	var wg sync.WaitGroup
	for _, container := range resp.ContainerIds {
		wg.Add(1)
		go func(container string) {
			defer wg.Done()
			CheckAndHeal(ctx, r.storeClient, container)
		}(container)
	}
	wg.Wait()
}

func (r *RecoveryService) Run(interval time.Duration) {
	for {
		r.RecoverContainers()
		time.Sleep(interval)
	}
}

func InitializeAndRun(interval time.Duration) {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := store.NewStoreClient(conn)

	RecoveryService := NewRecoveryService(client)
	RecoveryService.Run(interval)
}
