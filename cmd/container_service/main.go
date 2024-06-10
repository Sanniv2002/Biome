package main

import (
	"log"
	"net"

	"github.com/sanniv2002/nether/internal/containerinit/service"
	"github.com/sanniv2002/nether/proto/containerinit"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	containerinit.RegisterContainerInitServer(s, service.NewContainerService())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
