package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/somnathbm/hospital-hms/microservices/opd-service/gen/billing"
)

type BillingClient struct {
	client pb.BillingServiceClient
}

func NewBillingClient(address string) *BillingClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Billing Service: %v", err)
	}
	c := pb.NewBillingServiceClient(conn)
	return &BillingClient{client: c}
}

func (b *BillingClient) CreateBill(patientID, visitID string, amount float32) (string, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req := &pb.CreateBillRequest{
		PatientId: patientID,
		VisitId:   visitID,
		Amount:    amount,
	}
	resp, err := b.client.CreateBill(ctx, req)
	if err != nil {
		log.Printf("Error calling CreateBill: %v", err)
		return "", "Error contacting Billing service"
	}
	return resp.BillId, resp.Msg
}