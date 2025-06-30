import uuid

class AppointmentServiceCore:
    def __init__(self):
        # in-memory store for demo
        self.appointments = {
            "apt-123": {"patient_id": "p-001", "valid": True},
            "apt-456": {"patient_id": "p-002", "valid": True},
        }


    def check_appointment(self, patient_id, appointment_id):
        if patient_id == "" or appointment_id == "":
            return False, "Invalid input"
        appt = self.appointments.get(appointment_id)
        if appt and appt["valid"] and appt["patient_id"] == patient_id:
            return True, "Appointment is valid"
        return False, "Appointment not found or invalid"
    

    def book_appointment(self, patient_id, doctor_id, slot):
        if patient_id == "" or doctor_id == "" or slot == "":
            return "", "Invalid input"
        
        new_id = str(uuid.uuid4())
        self.appointments[new_id] = {
            "patient_id": patient_id,
            "doctor_id": doctor_id,
            "slot": slot,
            "valid": True
        }
        return new_id, "Appointment booked"