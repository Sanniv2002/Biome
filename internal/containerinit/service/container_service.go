package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/sanniv2002/nether/proto/containerinit"
	"github.com/sanniv2002/nether/proto/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ContainerService struct {
	containerinit.UnimplementedContainerInitServer
}

func NewContainerService() *ContainerService {
	return &ContainerService{}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
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

func KillContainer(containerId string) bool {
	arg0 := "docker"
	agr1 := "kill"
	cmd := exec.Command(arg0, agr1, containerId)
	if _, err := cmd.Output(); err != nil {
		return false
	}
	arg2 := "rm"
	cmd = exec.Command(arg0, arg2, containerId)
	if _, err := cmd.Output(); err != nil {
		return false
	}
	return true
}

func SpawnChild(ctx context.Context, client store.StoreClient, image string, alias string) bool {
	app := "docker"
	arg0 := "run"
	arg1 := "-d"
	arg2 := "-P"
	cmd := exec.Command(app, arg0, arg1, arg2, image)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	containerID := strings.TrimSpace(string(output))
	portCmd := exec.Command("docker", "port", containerID)
	portOutput, err := portCmd.Output()
	if err != nil {
		return false
	}
	portMapping := strings.TrimSpace(string(portOutput))

	//Extract container and host ports
	containerPORT, hostPort, err := extractPorts(portMapping)
	if err != nil {
		return false
	}
	data := &store.AddChildContainerRequest{
		Alias:         alias,
		ContainerId:   containerID,
		HostPORT:      uint32(hostPort),
		ContainerPORT: uint32(containerPORT),
	}
	client.AddChildContainer(ctx, data)
	return true
}

func (s *ContainerService) StartContainer(ctx context.Context, req *containerinit.StartContainerRequest) (*containerinit.StartContainerResponse, error) {
	//privilege := "sudo"	//Use this and add to cmd and portCmd for linux machines
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := store.NewStoreClient(conn)

	app := "docker"
	arg0 := "run"
	arg1 := "-d"
	arg2 := "-P"
	arg3 := string(req.Image)

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %w", err)
	}
	containerID := strings.TrimSpace(string(output))

	// Get the port mapping
	portCmd := exec.Command("docker", "port", containerID)
	portOutput, err := portCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get port mapping: %w", err)
	}

	portMapping := strings.TrimSpace(string(portOutput))

	//Extract container and host ports
	containerPort, hostPort, err := extractPorts(portMapping)
	if err != nil {
		return nil, fmt.Errorf("failed to extract ports: %w", err)
	}

	//Create an alias for the container
	alias := RandStringBytes(12)

	data := &store.AddParentContainerRequest{
		Alias: alias,
		Info: &store.Server{
			ContainerId:   containerID,
			HostPORT:      uint32(hostPort),
			ContainerPORT: uint32(containerPort),
		},
		Image:       req.Image,
		MinReplicas: req.MinContainers,
		MaxReplicas: req.MaxContainers,
	}

	//Add the Parent Container to the Store
	client.AddParentContainer(context.Background(), data)

	return &containerinit.StartContainerResponse{
		ContainerId:   containerID,
		ContainerPORT: uint32(containerPort),
		HostPORT:      uint32(hostPort),
		Alias:         alias,
	}, nil
}

func (s *ContainerService) ScaleContainer(ctx context.Context, req *containerinit.ScaleContainerRequest) (*containerinit.ScaleContainerResponse, error) {
	var wg sync.WaitGroup
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := store.NewStoreClient(conn)
	for i := req.Present; i < req.Target; i++ { //All Spawn Commands will run in seperate go routines
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := SpawnChild(ctx, client, req.Image, req.Alias)
			if !result {
				fmt.Println("Fatal error starting child")
			}
		}()
	}

	wg.Wait()

	return &containerinit.ScaleContainerResponse{
		Success: true,
	}, nil
}

func (s *ContainerService) KillContainer(ctx context.Context, req *containerinit.KillContainerRequest) (*containerinit.KillContainerResponse, error) {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := store.NewStoreClient(conn)
	arg0 := "docker"
	agr1 := "kill"
	arg2 := req.ContainerId
	cmd := exec.Command(arg0, agr1, arg2)
	if _, err := cmd.Output(); err != nil {
		return nil, fmt.Errorf("failed to kill container: %w", err)
	}

	arg3 := "rm"
	cmd = exec.Command(arg0, arg3, req.ContainerId)
	if _, err := cmd.Output(); err != nil {
		fmt.Println("Failed to remove container", err)
	}

	data := &store.RemoveContainerRequest{
		ContainerId: req.ContainerId,
	}

	_, err = client.RemoveContainer(ctx, data)
	if err != nil {
		fmt.Println(err)
	}
	return &containerinit.KillContainerResponse{
		Success: true,
	}, nil
}

func (s *ContainerService) KillConfig(ctx context.Context, req *containerinit.KillConfigRequest) (*containerinit.KillConfigResponse, error) {
	var wg sync.WaitGroup
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := store.NewStoreClient(conn)
	configs, err := client.GetAllConfigs(ctx, &store.GetAllConfigsRequest{})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	for _, config := range configs.Configs {
		if config.Alias == req.Alias {
			for _, server := range config.Servers {
				wg.Add(1)
				go func(server *store.Server) {
					defer wg.Done()
					KillContainer(server.ContainerId)
				}(server)
				wg.Wait()
			}
			break
		}
	}
	return &containerinit.KillConfigResponse{
		Success: true,
	}, nil
}
