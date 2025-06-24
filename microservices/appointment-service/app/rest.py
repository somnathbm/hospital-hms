from fastapi import FastAPI
from pydantic import BaseModel

def create_rest_api(core):
    app = FastAPI()

    class CheckAppointmentRequest(BaseModel):
        patient_id: str
        appointment_id: str


    class BookAppointmentRequest(BaseModel):
        patient_id: str
        doctor_id: str
        slot: str

    @app.get("/health")
    def health():
        return {"status": "ok"}
    
    
    @app.post("/appointments/check")
    def check_appt(req: CheckAppointmentRequest):
        valid, msg = core.check_appointment(req.patient_id, req.appointment_id)
        return {"valid": valid, "msg": msg}
    

    @app.post("/appointments/book")
    def book_appt(req: BookAppointmentRequest):
        appt_id, msg = core.book_appointment(req.patient_id, req.doctor_id, req.slot)
        return {"appointment_id": appt_id, "msg": msg}
    
    return app
    