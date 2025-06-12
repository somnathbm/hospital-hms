package server

import (
	"context"

	pb "github.com/somnathbm/hospital-hms/microservices/opd-service/gen" // Replace with your actual module path
	"github.com/somnathbm/hospital-hms/microservices/opd-service/internal/service"
)

type OPDGRPCServer struct {
	pb.UnimplementedOPDServiceServer
	svc service.OPDService
}

func NewOPDGRPCServer(svc service.OPDService) *OPDGRPCServer {
	return &OPDGRPCServer{svc: svc}
}

func (s *OPDGRPCServer) CheckAppointmentStatus(ctx context.Context, req *pb.CheckAppointmentRequest) (*pb.CheckAppointmentResponse, error) {
	valid, msg := s.svc.CheckAppointment(req.PatientId, req.AppointmentId)
	return &pb.CheckAppointmentResponse{
		Valid:   valid,
		Message: msg,
	}, nil
}

func (s *OPDGRPCServer) StartConsultation(ctx context.Context, req *pb.StartConsultationRequest) (*pb.StartConsultationResponse, error) {
	visitID, msg := s.svc.StartConsultation(req.PatientId, req.DoctorId, req.AppointmentId)
	return &pb.StartConsultationResponse{
		VisitId: visitID,
		Message: msg,
	}, nil
}

func (s *OPDGRPCServer) RecordDiagnosis(ctx context.Context, req *pb.RecordDiagnosisRequest) (*pb.RecordDiagnosisResponse, error) {
	msg := s.svc.RecordDiagnosis(req.VisitId, req.Diagnosis)
	return &pb.RecordDiagnosisResponse{Message: msg}, nil
}

func (s *OPDGRPCServer) PrescribeTests(ctx context.Context, req *pb.PrescribeTestsRequest) (*pb.PrescribeTestsResponse, error) {
	msg := s.svc.PrescribeTests(req.VisitId, req.TestNames)
	return &pb.PrescribeTestsResponse{Message: msg}, nil
}

func (s *OPDGRPCServer) GeneratePrescription(ctx context.Context, req *pb.GeneratePrescriptionRequest) (*pb.GeneratePrescriptionResponse, error) {
	rxID, msg := s.svc.GeneratePrescription(req.VisitId, req.Medicines)
	return &pb.GeneratePrescriptionResponse{
		PrescriptionId: rxID,
		Message:        msg,
	}, nil
}

func (s *OPDGRPCServer) EndVisit(ctx context.Context, req *pb.EndVisitRequest) (*pb.EndVisitResponse, error) {
	msg := s.svc.EndVisit(req.VisitId)
	return &pb.EndVisitResponse{Message: msg}, nil
}
