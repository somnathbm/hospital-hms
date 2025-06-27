package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAppointmentClient struct{}
type mockBillingClient struct {
	billedVisitIds []string
}

func (m *mockAppointmentClient) CheckAppointment(patientId, appointmentId string) (bool, string) {
	if appointmentId == "apt-123" || appointmentId == "apt-456" {
		return true, "Appointment is valid"
	}
	return false, "Invalid appointment"
}

func (m *mockBillingClient) CreateBill(patientId, visitId string, amount float64) (string, string) {
	m.billedVisitIds = append(m.billedVisitIds, visitId)
	return "bill-999", "Bill generated successfully. Visit completed"
}

func TestOPDService(t *testing.T) {
	appointmentClient := &mockAppointmentClient{}
	billingClient := &mockBillingClient{}
	svc := NewOPDService(appointmentClient, billingClient)

	t.Run("Check valid appointment", func(t *testing.T) {
		valid, msg := svc.CheckAppointment("p-001", "apt-123")
		assert.True(t, valid)
		assert.Equal(t, "Appointment is valid", msg)
	})

	t.Run("Start consultation with valid appointment", func(t *testing.T) {
		visitID, msg := svc.StartConsultation("p-001", "d-101", "apt-123")
		assert.NotEmpty(t, visitID)
		assert.Equal(t, "Consultation started", msg)
	})

	t.Run("Start consultation with invalid appointment", func(t *testing.T) {
		visitID, msg := svc.StartConsultation("p-002", "d-102", "apt-999")
		assert.Empty(t, visitID)
		assert.Equal(t, "Cannot start consultation: Invalid appointment", msg)
	})

	t.Run("Record diagnosis", func(t *testing.T) {
		visitID, _ := svc.StartConsultation("p-001", "d-101", "apt-123")
		msg := svc.RecordDiagnosis(visitID, "Fever")
		assert.Equal(t, "Diagnosis recorded", msg)
	})

	t.Run("Prescribe tests", func(t *testing.T) {
		visitID, _ := svc.StartConsultation("p-001", "d-101", "apt-123")
		msg := svc.PrescribeTests(visitID, []string{"CBC", "X-Ray"})
		assert.Equal(t, "Tests prescribed", msg)
	})

	t.Run("Generate prescription", func(t *testing.T) {
		visitID, _ := svc.StartConsultation("p-001", "d-101", "apt-123")
		rxID, msg := svc.GeneratePrescription(visitID, []string{"Paracetamol", "Antibiotic"})
		assert.Contains(t, rxID, "rx-visit")
		assert.Equal(t, "Prescription generated", msg)
	})

	t.Run("End visit", func(t *testing.T) {
		visitID, _ := svc.StartConsultation("p-001", "d-101", "apt-123")
		_, msg := svc.EndVisit("p-001", visitID)
		assert.Equal(t, "Bill generated successfully. Visit completed", msg)
	})
}
