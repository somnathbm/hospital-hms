package client

type Appointment interface {
	CheckAppointment(patientId, appointmentId string) (bool, string)
}

type Billing interface {
	CreateBill(patientId, visitId string, amount float64) (string, string)
}
