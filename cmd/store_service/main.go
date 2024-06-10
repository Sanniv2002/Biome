package main

import (
	"log"
	"net"

	"github.com/sanniv2002/nether/internal/store/service"
	"github.com/sanniv2002/nether/proto/store"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	store.RegisterStoreServer(s, service.NewStoreService())

	log.Println("Server is running on port 50052...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
