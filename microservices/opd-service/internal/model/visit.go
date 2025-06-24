package model

type Visit struct {
	VisitID      string
	PatientID    string
	DoctorID     string
	Diagnosis    string
	Tests        []string
	Prescription []string
	IsComplete   bool
}