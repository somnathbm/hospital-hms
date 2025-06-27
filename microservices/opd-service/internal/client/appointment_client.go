package client

import (
	"context"
	"log"
	"time"

	pb "github.com/somnathbm/hospital-hms/microservices/opd-service/gen/aptpb"
	"google.golang.org/grpc"
)

type AppointmentClient struct {
	client pb.AppointmentServiceClient
}

func NewAppointmentClient(address string) *AppointmentClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to appointment service: %v", err)
	}
	c := pb.NewAppointmentServiceClient(conn)
	return &AppointmentClient{client: c}
}

func (a *AppointmentClient) CheckAppointment(patientID, appointmentID string) (bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req := &pb.CheckAppointmentRequest{
		PatientId:     patientID,
		AppointmentId: appointmentID,
	}

	resp, err := a.client.CheckAppointment(ctx, req)
	if err != nil {
		log.Printf("Error calling CheckAppointment: %v", err)
		return false, "Error contacting Appointment service"
	}
	return resp.Valid, resp.Msg
}
