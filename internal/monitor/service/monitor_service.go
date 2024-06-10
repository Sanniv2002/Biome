package monitor

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/sanniv2002/nether/proto/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MonitorService struct {
	storeClient store.StoreClient
}

func NewMonitorService(storeClient store.StoreClient) *MonitorService {
	return &MonitorService{
		storeClient: storeClient,
	}
}

func extractPercentage(s string) string {
	// Split the string by colon to separate the hash and the percentage
	parts := strings.Split(s, ":")
	if len(parts) < 2 {
		return ""
	}
	// Trim any whitespace and the '%' character from the second part
	percentage := strings.TrimSpace(parts[1])
	percentage = strings.TrimSuffix(percentage, "%")
	return percentage
}

func GetAndUpdateStats(ctx context.Context, client store.StoreClient, containerID string) bool {
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
	forCPU := "{{.Container}}: {{.CPUPerc}}"
	forMEM := "{{.Container}}: {{.MemPerc}}"
	cmd := exec.Command(arg0, arg1, arg2, arg3, arg4, forCPU)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return false
	}
	currentCPU := extractPercentage(string(output))
	cmd = exec.Command(arg0, arg1, arg2, arg3, arg4, forMEM)
	output, err = cmd.Output()
	if err != nil {
		return false
	}

	currentMEM := extractPercentage(string(output))
	updateData := &store.UpdateContainerStatsRequest{
		ContainerId: resp.ContainerId,
		Alias:       resp.Alias,
		CPU:         currentCPU,
		MEM:         currentMEM,
	}
	response, error := client.UpdateContainerStats(ctx, updateData)
	if error != nil || !response.Succes {
		fmt.Println("error updating stats", error)
	}
	return true
}

func (m *MonitorService) MonitorContainers() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Call the gRPC method to get all containers
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := store.NewStoreClient(conn)
	resp, err := client.GetAllContainers(ctx, &store.GetAllContainersRequest{})
	if err != nil {
		log.Printf("Error getting all containers: %v", err)
		return
	}
	var wg sync.WaitGroup
	for _, container := range resp.ContainerIds {
		wg.Add(1)
		go func(container string) {
			defer wg.Done()
			GetAndUpdateStats(ctx, client, container)
		}(container)
	}
	wg.Wait()
}

func (m *MonitorService) Run(interval time.Duration) {
	for {
		m.MonitorContainers()
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

	monitorService := NewMonitorService(client)
	monitorService.Run(interval)
}
