package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/go-chi/chi"
	pb "github.com/somnathbm/hospital-hms/microservices/opd-service/gen/opdpb"
	"github.com/somnathbm/hospital-hms/microservices/opd-service/internal/client"
	"github.com/somnathbm/hospital-hms/microservices/opd-service/internal/rest"
	"github.com/somnathbm/hospital-hms/microservices/opd-service/internal/server"
	"github.com/somnathbm/hospital-hms/microservices/opd-service/internal/service"
)

func main() {
	grpcPort := ":50051"
	restPort := ":8080"

	// Create gRPC clients for Appointment & Billing services
	appointmentClient := client.NewAppointmentClient("appointment-service:50052")
	billingClient := client.NewBillingClient("billing-service:50052")

	// Create OPD service with these clients
	opdSvc := service.NewOPDService(appointmentClient, billingClient)

	// gRPC setup
	grpcServer := grpc.NewServer()
	pb.RegisterOPDServiceServer(grpcServer, server.NewOPDGRPCServer(opdSvc))
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Printf("✅ OPD gRPC server listening on %s", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// REST setup
	restHandler := rest.NewHandler(opdSvc)
	r := chi.NewRouter()
	r.Mount("/opd", restHandler.Routes())

	log.Printf("✅ OPD REST server listening on %s", restPort)
	if err := http.ListenAndServe(restPort, r); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}
}
