package main

import (
	"github.com/simabdi/authservice/config"
	pb "github.com/simabdi/authservice/proto"
	"github.com/simabdi/authservice/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	config.ConnectDatabase()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &server.AuthServiceServer{})

	log.Println("🚀 gRPC AuthService running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
