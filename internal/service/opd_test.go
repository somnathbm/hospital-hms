package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOPDService(t *testing.T) {
	svc := NewOPDService()

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

	// t.Run("Start consultation with invalid appointment", func(t *testing.T) {
	// 	visitID, msg := svc.StartConsultation("p-002", "d-102", "apt-999")
	// 	assert.Empty(t, visitID)
	// 	assert.Equal(t, "Cannot start consultation: Invalid appointment", msg)
	// })

	// t.Run("Record diagnosis", func(t *testing.T) {
	// 	visitID, _ := svc.StartConsultation("p-003", "d-103", "apt-456")
	// 	msg := svc.RecordDiagnosis(visitID, "Fever")
	// 	assert.Equal(t, "Diagnosis recorded", msg)
	// })

	// t.Run("Prescribe tests", func(t *testing.T) {
	// 	visitID, _ := svc.StartConsultation("p-004", "d-104", "apt-456")
	// 	msg := svc.PrescribeTests(visitID, []string{"CBC", "X-Ray"})
	// 	assert.Equal(t, "Tests prescribed", msg)
	// })

	// t.Run("Generate prescription", func(t *testing.T) {
	// 	visitID, _ := svc.StartConsultation("p-005", "d-105", "apt-123")
	// 	rxID, msg := svc.GeneratePrescription(visitID, []string{"Paracetamol", "Antibiotic"})
	// 	assert.Contains(t, rxID, "rx-visit")
	// 	assert.Equal(t, "Prescription generated", msg)
	// })

	// t.Run("End visit", func(t *testing.T) {
	// 	visitID, _ := svc.StartConsultation("p-006", "d-106", "apt-456")
	// 	msg := svc.EndVisit(visitID)
	// 	assert.Equal(t, "Visit completed", msg)
	// })
}
