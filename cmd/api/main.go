package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sanniv2002/nether/internal/api/routes"
	"github.com/sanniv2002/nether/shared/config"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	conf, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Connect to gRPC server
	grpcConn, err := grpc.Dial(conf.ContainerInitServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer grpcConn.Close()

	// Register routes
	routes.RegisterRoutes(r, grpcConn)

	// Start HTTP server
	if err := r.Run(conf.ApiPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
