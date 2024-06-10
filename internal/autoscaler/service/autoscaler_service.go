package autoscaler

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/sanniv2002/nether/proto/containerinit"
	"github.com/sanniv2002/nether/proto/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AutoscalerService struct {
	storeClient     store.StoreClient
	containerClient containerinit.ContainerInitClient
}

func NewAutoscalerService(storeClient store.StoreClient, containerClient containerinit.ContainerInitClient) *AutoscalerService {
	return &AutoscalerService{
		storeClient:     storeClient,
		containerClient: containerClient,
	}
}

func (a *AutoscalerService) Run(interval time.Duration) {
	for {
		a.AutoscaleContainers()
		time.Sleep(interval)
	}
}

func parseResourceUsage(usage string) float64 {
	result, err := strconv.ParseFloat(usage, 64)
	if err != nil {
		log.Printf("Failed to parse resource usage: %v", err)
		return 0.0
	}
	return result
}

func checkResourceUsageAndCommitScaling(servers []*store.Server, presentContainers int, alias string, image string, minReplicas uint32, maxReplicas uint32, a *AutoscalerService) bool {
	if presentContainers < int(minReplicas) { //Initially Scale Cotainers to minReplicas
		a.containerClient.ScaleContainer(context.Background(), &containerinit.ScaleContainerRequest{Alias: alias, Present: 1, Target: minReplicas, Image: image})
	}
	totalCPUUsage := 0.0
	totalMEMUsage := 0.0
	var (
		minCombinedUsage     = math.MaxFloat64
		minCombinedContainer string
	)
	for _, server := range servers {
		stats, err := a.storeClient.GetContainerStats(context.Background(), &store.GetContainerStatsRequest{ContainerId: server.ContainerId})
		if err != nil {
			fmt.Println("Error getting container stats", err)
		}
		cpuUsage := parseResourceUsage(stats.CPU)
		memUsage := parseResourceUsage(stats.MEM)

		totalCPUUsage += cpuUsage
		totalMEMUsage += memUsage
		combinedUsage := cpuUsage + memUsage

		if combinedUsage < minCombinedUsage {
			minCombinedUsage = combinedUsage
			minCombinedContainer = server.ContainerId
		}
	}
	avgCPUUsage := totalCPUUsage / float64(presentContainers)
	avgMEMUsage := totalMEMUsage / float64(presentContainers)
	if avgCPUUsage >= 80 || avgMEMUsage >= 80 {
		if presentContainers < int(maxReplicas) {
			fmt.Println("Scaling up container with image:", image)
			_, err := a.containerClient.ScaleContainer(context.Background(), &containerinit.ScaleContainerRequest{Alias: alias, Present: uint32(presentContainers), Target: uint32(presentContainers + 1), Image: image})
			if err != nil {
				fmt.Println("Error scaling container", err)
			}
		}
	} else if avgCPUUsage < 20 && avgMEMUsage < 20 {
		if presentContainers > int(minReplicas) {
			fmt.Println("Suspending container with ID:", minCombinedContainer)
			_, err := a.containerClient.KillContainer(context.Background(), &containerinit.KillContainerRequest{ContainerId: minCombinedContainer})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return true
}

func (a *AutoscalerService) AutoscaleContainers() {
	var wg sync.WaitGroup
	data := &store.GetAllConfigsRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := a.storeClient.GetAllConfigs(ctx, data)
	if err != nil {
		fmt.Println("Unable to get containers")
	}
	for _, config := range resp.Configs {
		wg.Add(1)
		go func(servers []*store.Server, presentContainers int, alias string, image string, minReplicas uint32, maxReplicas uint32, a *AutoscalerService) {
			defer wg.Done()
			checkResourceUsageAndCommitScaling(servers, presentContainers, alias, image, minReplicas, maxReplicas, a)
		}(config.Servers, len(config.Servers), config.Alias, config.Image, config.MinReplicas, config.MaxReplicas, a)
	}
	wg.Wait()
}

func InitializeAndRun(interval time.Duration) {
	conn_store, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn_store.Close()
	conn_containerinit, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn_containerinit.Close()
	client_store := store.NewStoreClient(conn_store)
	client_containerinit := containerinit.NewContainerInitClient(conn_containerinit)

	autoscalerService := NewAutoscalerService(client_store, client_containerinit)
	autoscalerService.Run(interval)
}
