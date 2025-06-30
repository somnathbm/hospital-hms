import pytest
from app.service import AppointmentServiceCore

@pytest.fixture
def apppointment_service():
  return AppointmentServiceCore()


def test_check_valid_appointment(apppointment_service):
  resp, msg = apppointment_service.check_appointment("p-001", "apt-123")
  assert resp == True
  assert msg == "Appointment is valid"


def test_check_invalid_appointment(apppointment_service):
  resp, msg = apppointment_service.check_appointment("p-003", "apt-333")
  assert resp == False
  assert msg == "Appointment not found or invalid"


def test_check_appointment_invalid_input(apppointment_service):
  resp, msg = apppointment_service.check_appointment("", "")
  assert resp == False
  assert msg == "Invalid input"


def test_book_appointment(apppointment_service):
  apt_id, msg = apppointment_service.book_appointment("p-001", "d-001", "2025-10-15")
  assert apt_id != ""
  assert msg == "Appointment booked"


def test_invalid_book_appointment(apppointment_service):
  apt_id, msg = apppointment_service.book_appointment("", "", "")
  assert apt_id == ""
  assert msg == "Invalid input"