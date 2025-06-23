package rest

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/somnathbm/hospital-hms/microservices/opd-service/internal/service"
)

type Handler struct {
	opd service.OPDService
}

func NewHandler(opd service.OPDService) *Handler {
	return &Handler{opd: opd}
}

func (h *Handler) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/check-appointment", h.CheckAppointment)
	r.Post("/start-consultation", h.StartConsultation)

	return r
}

func (h *Handler) CheckAppointment(w http.ResponseWriter, r *http.Request) {
	patientID := r.URL.Query().Get("patient_id")
	appointmentID := r.URL.Query().Get("appointment_id")

	valid, msg := h.opd.CheckAppointment(patientID, appointmentID)

	resp := map[string]interface{}{
		"valid": valid,
		"msg":   msg,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) StartConsultation(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PatientID     string `json:"patient_id"`
		DoctorID      string `json:"doctor_id"`
		AppointmentID string `json:"appointment_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	visitID, msg := h.opd.StartConsultation(req.PatientID, req.DoctorID, req.AppointmentID)

	resp := map[string]interface{}{
		"visit_id": visitID,
		"msg":      msg,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
