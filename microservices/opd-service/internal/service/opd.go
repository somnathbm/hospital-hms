package service

import (
	"fmt"
	"sync"

	"github.com/somnathbm/hospital-hms/microservices/opd-service/internal/client"
	"github.com/somnathbm/hospital-hms/microservices/opd-service/internal/model"
)

type OPDService interface {
	CheckAppointment(patientID, appointmentID string) (bool, string)
	StartConsultation(patientID, doctorID, appointmentID string) (string, string)
	RecordDiagnosis(visitID, diagnosis string) string
	PrescribeTests(visitID string, tests []string) string
	GeneratePrescription(visitID string, medicines []string) (string, string)
	EndVisit(patientID, visitID string) (string, string)
}

type opdService struct {
	mu          sync.Mutex
	visits      map[string]*model.Visit
	appointment client.Appointment
	billing     client.Billing
}

func NewOPDService(appointment client.Appointment, billing client.Billing) OPDService {
	return &opdService{
		visits:      make(map[string]*model.Visit),
		appointment: appointment,
		billing:     billing,
	}
}

func (s *opdService) CheckAppointment(patientID, appointmentID string) (bool, string) {
	// Use the Appointment gRPC client
	return s.appointment.CheckAppointment(patientID, appointmentID)
}

func (s *opdService) StartConsultation(patientID, doctorID, appointmentID string) (string, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if valid, _ := s.appointment.CheckAppointment(patientID, appointmentID); !valid {
		return "", "Cannot start consultation: Invalid appointment"
	}

	visitID := fmt.Sprintf("visit-%d", len(s.visits)+1)
	s.visits[visitID] = &model.Visit{
		VisitID:   visitID,
		PatientID: patientID,
		DoctorID:  doctorID,
	}

	return visitID, "Consultation started"
}

func (s *opdService) RecordDiagnosis(visitID, diagnosis string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if visit, ok := s.visits[visitID]; ok {
		visit.Diagnosis = diagnosis
		return "Diagnosis recorded"
	}
	return "Visit not found"
}

func (s *opdService) PrescribeTests(visitID string, tests []string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if visit, ok := s.visits[visitID]; ok {
		visit.Tests = append(visit.Tests, tests...)
		return "Tests prescribed"
	}
	return "Visit not found"
}

func (s *opdService) GeneratePrescription(visitID string, meds []string) (string, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if visit, ok := s.visits[visitID]; ok {
		visit.Prescription = append(visit.Prescription, meds...)
		prescriptionID := fmt.Sprintf("rx-%s", visitID)
		return prescriptionID, "Prescription generated"
	}
	return "", "Visit not found"
}

func (s *opdService) EndVisit(patientID, visitID string) (string, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.visits[visitID]; !ok {
		return "", "Visit not found"
	}

	// create fixed amount for demo purpose
	billID, msg := s.billing.CreateBill(patientID, visitID, 5000.0)
	// visit.IsComplete = true
	return billID, msg
}
