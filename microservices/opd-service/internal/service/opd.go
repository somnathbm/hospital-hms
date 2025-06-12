package service

import (
	"fmt"
	"sync"
)

type OPDService interface {
	CheckAppointment(patientID, appointmentID string) (bool, string)
	StartConsultation(patientID, doctorID, appointmentID string) (string, string)
	RecordDiagnosis(visitID, diagnosis string) string
	PrescribeTests(visitID string, tests []string) string
	GeneratePrescription(visitID string, medicines []string) (string, string)
	EndVisit(visitID string) string
}

type opdService struct {
	mu          sync.Mutex
	visits      map[string]*Visit
	appointment map[string]bool
}

func NewOPDService() OPDService {
	return &opdService{
		visits: make(map[string]*Visit),
		appointment: map[string]bool{
			"apt-123": true,
		},
	}
}

type Visit struct {
	VisitID      string
	PatientID    string
	DoctorID     string
	Diagnosis    string
	Tests        []string
	Prescription []string
	IsComplete   bool
}

func (s *opdService) CheckAppointment(patientID, appointmentID string) (bool, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	valid, exists := s.appointment[appointmentID]
	if exists && valid {
		return true, "Appointment is valid"
	}
	return false, "Appointment not found or invalid"
}

func (s *opdService) StartConsultation(patientID, doctorID, appointmentID string) (string, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if valid, _ := s.CheckAppointment(patientID, appointmentID); !valid {
		return "", "Cannot start consultation: Invalid appointment"
	}

	visitID := fmt.Sprintf("visit-%d", len(s.visits)+1)
	s.visits[visitID] = &Visit{
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

func (s *opdService) EndVisit(visitID string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if visit, ok := s.visits[visitID]; ok {
		visit.IsComplete = true
		return "Visit completed"
	}
	return "Visit not found"
}
