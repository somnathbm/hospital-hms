package server

import (
	"context"

	pb "github.com/somnathbm/hospital-hms/gen"
	"github.com/somnathbm/hospital-hms/internal/service"
)

type OPDGRPCServer struct {
	pb.UnimplementedOPDServiceServer
	svc service.OPDService
}

func NewOPDGRPCServer(svc service.OPDService) *OPDGRPCServer {
	return &OPDGRPCServer{svc: svc}
}

func (s *OPDGRPCServer) CheckAppointmentStatus(ctx context.Context, req *pb.AppointmentRequest) (*pb.AppointmentResponse, error) {
	valid, msg := s.svc.CheckAppointment(req.PatientId, req.AppointmentId)
	return &pb.AppointmentResponse{
		Valid:   valid,
		Message: msg,
	}, nil
}
