package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/somnathbm/hospital-hms/microservices/opd-service/gen"
	"github.com/somnathbm/hospital-hms/microservices/opd-service/internal/server"
	"github.com/somnathbm/hospital-hms/microservices/opd-service/internal/service"
)

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Initialize the business logic
	opdService := service.NewOPDService()

	// Initialize gRPC server
	grpcServer := grpc.NewServer()

	// Register the gRPC server
	pb.RegisterOPDServiceServer(grpcServer, server.NewOPDGRPCServer(opdService))

	log.Println("✅ OPD gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
