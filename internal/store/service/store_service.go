package service

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sanniv2002/nether/prisma/db"
	"github.com/sanniv2002/nether/proto/store"
)

type ContainerIdResult struct {
	ContainerId string `json:"containerId"`
}

type StoreService struct {
	store.UnimplementedStoreServer
}

func NewStoreService() *StoreService {
	return &StoreService{}
}

// Function to add the parent container in the store
func (s *StoreService) AddParentContainer(ctx context.Context, req *store.AddParentContainerRequest) (*store.AddParentContainerResponse, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println("Error Connecting Prisma Client")
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	key := req.Alias

	createdConfig, err := client.Config.CreateOne(
		db.Config.Alias.Set(key),
		db.Config.Image.Set(req.Image),
		db.Config.MinReplicas.Set(int(req.MinReplicas)),
		db.Config.MaxReplicas.Set(int(req.MaxReplicas)),
	).Exec(ctx)
	if err != nil {
		fmt.Println("error")
	}

	_, err = client.Server.CreateOne(
		db.Server.ContainerID.Set(req.Info.ContainerId),
		db.Server.HostPORT.Set(int(req.Info.HostPORT)),
		db.Server.ContainerPORT.Set(int(req.Info.ContainerPORT)),
		db.Server.Config.Link(
			db.Config.ID.Equals(createdConfig.ID),
		),
	).Exec(ctx)

	if err != nil {
		fmt.Println("db error")
	}

	err = rdb.HMSet(ctx, req.Info.ContainerId, "alias", req.Alias, "CPU", 0, "MEM", 0).Err()
	if err != nil {
		fmt.Println("redis error")
	}

	return &store.AddParentContainerResponse{
		Success: true,
	}, nil
}

// Function to check and add more children to a parent container
func (s *StoreService) AddChildContainer(ctx context.Context, req *store.AddChildContainerRequest) (*store.AddChildContainerResponse, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println("Error Connecting Prisma Client")
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	if rdb == nil {
		return nil, fmt.Errorf("redis client is not initialized")
	}

	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}

	if req.ContainerId == "" {
		return nil, fmt.Errorf("container ID is empty")
	}

	config, err := client.Config.FindUnique(
		db.Config.Alias.Equals(req.Alias),
	).Exec(ctx)
	if err != nil {
		fmt.Println("Error")
	}

	_, er := client.Server.CreateOne(
		db.Server.ContainerID.Set(req.ContainerId),
		db.Server.HostPORT.Set(int(req.HostPORT)),
		db.Server.ContainerPORT.Set(int(req.ContainerPORT)),
		db.Server.Config.Link(
			db.Config.ID.Equals(config.ID),
		),
	).Exec(ctx)

	if er != nil {
		fmt.Println("error")
	}

	err = rdb.HMSet(ctx, req.ContainerId, "alias", req.Alias, "CPU", 0, "MEM", 0).Err()
	if err != nil {
		fmt.Println("redis error")
	}

	return &store.AddChildContainerResponse{
		Success: true,
	}, nil
}

func (s *StoreService) RemoveContainer(ctx context.Context, req *store.RemoveContainerRequest) (*store.RemoveContainerResponse, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println("Error Connecting Prisma Client")
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	err := rdb.Del(ctx, req.ContainerId)
	if err != nil {
		fmt.Println("error", err)
	}

	_, er := client.Server.FindUnique(
		db.Server.ContainerID.Equals(req.ContainerId),
	).Delete().Exec(ctx)
	if er != nil {
		return &store.RemoveContainerResponse{
			Success: false,
		}, nil
	}
	return &store.RemoveContainerResponse{
		Success: true,
	}, nil
}

func (s *StoreService) GetAllContainers(ctx context.Context, req *store.GetAllContainersRequest) (*store.GetAllContainersResponse, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println("Error Connecting Prisma Client")
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	query := `SELECT "containerId" FROM "Server";`
	var containerIdResults []ContainerIdResult

	if err := client.Prisma.QueryRaw(query).Exec(ctx, &containerIdResults); err != nil {
		fmt.Println("error", err)
	}

	containerIds := make([]string, len(containerIdResults))
	for i, result := range containerIdResults {
		containerIds[i] = result.ContainerId
	}

	return &store.GetAllContainersResponse{
		ContainerIds: containerIds,
	}, nil
}

func (s *StoreService) UpdateContainerStats(ctx context.Context, req *store.UpdateContainerStatsRequest) (*store.UpdateContainerStatsResponse, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.HMSet(ctx, req.ContainerId, "alias", req.Alias, "CPU", req.CPU, "MEM", req.MEM)
	if err != nil {
		fmt.Println("error", err)
	}

	return &store.UpdateContainerStatsResponse{
		Succes: true,
	}, nil
}

func (s *StoreService) GetContainerStats(ctx context.Context, req *store.GetContainerStatsRequest) (*store.GetContainerStatsResponse, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	CPU := rdb.HGet(ctx, req.ContainerId, "CPU")
	MEM := rdb.HGet(ctx, req.ContainerId, "MEM")
	ALIAS := rdb.HGet(ctx, req.ContainerId, "alias")

	cpuValue, err := CPU.Result()
	if err != nil {
		fmt.Println(err)
	}

	memValue, err := MEM.Result()
	if err != nil {
		fmt.Println(err)
	}
	alaisValue, err := ALIAS.Result()
	if err != nil {
		fmt.Println(err)
	}
	data := &store.GetContainerStatsResponse{
		ContainerId: req.ContainerId,
		Alias:       alaisValue,
		CPU:         cpuValue,
		MEM:         memValue,
	}
	return &store.GetContainerStatsResponse{
		ContainerId: data.ContainerId,
		Alias:       data.Alias,
		CPU:         data.CPU,
		MEM:         data.MEM,
	}, nil
}

func (s *StoreService) GetContainerLimits(ctx context.Context, req *store.GetContainerLimitsRequest) (*store.GetContainerLimitsResponse, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println("Error Connecting Prisma Client")
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	config, err := client.Config.FindUnique(
		db.Config.Alias.Equals(req.Alias),
	).Exec(context.Background())
	if err != nil {
		fmt.Println("Failed to query Config for alias", req.Alias, err)
	}

	return &store.GetContainerLimitsResponse{
		MinReplicas: uint32(config.MinReplicas),
		MaxReplicas: uint32(config.MaxReplicas),
	}, nil
}

func (s *StoreService) UpdatePortMappings(ctx context.Context, req *store.UpdatePortMappingsRequest) (*store.UpdatePortMappingsResponse, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println("Error Connecting Prisma Client")
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	_, err := client.Server.FindUnique(
		db.Server.ContainerID.Equals(req.ContainerId),
	).Update(
		db.Server.HostPORT.Set(int(req.HostPORT)),
		db.Server.ContainerPORT.Set(int(req.ContainerPORT)),
	).Exec(context.Background())
	if err != nil {
		fmt.Println("error finding contianer", err)
	}
	return &store.UpdatePortMappingsResponse{
		Success: "updated",
	}, nil
}

func (s *StoreService) GetAllConfigs(ctx context.Context, req *store.GetAllConfigsRequest) (*store.GetAllConfigsResponse, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println("Error Connecting Prisma Client")
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	configs, err := client.Config.FindMany().
		With(
			db.Config.Servers.Fetch(),
		).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch configs: %w", err)
	}
	var configResponses []*store.Config
	for _, config := range configs {
		servers := config.Servers()
		var serverResponses []*store.Server
		for _, server := range servers {
			serverResponse := &store.Server{
				ContainerId:   server.ContainerID,
				HostPORT:      uint32(server.HostPORT),
				ContainerPORT: uint32(server.ContainerPORT),
			}
			serverResponses = append(serverResponses, serverResponse)
		}

		configResponse := &store.Config{
			Alias:       config.Alias,
			Image:       config.Image,
			MinReplicas: uint32(config.MinReplicas),
			MaxReplicas: uint32(config.MaxReplicas),
			Servers:     serverResponses,
		}
		configResponses = append(configResponses, configResponse)
	}

	return &store.GetAllConfigsResponse{
		Configs: configResponses,
	}, nil
}
