from fastapi import FastAPI
from pydantic import BaseModel

def create_rest_api(core):
    app = FastAPI()

    class CreateBillRequest(BaseModel):
        patient_id: str
        visit_id: str
        amount: float

    @app.get("/health")
    def health():
        return {"status": "ok"}


    @app.post("/bills")
    def create_bill(req: CreateBillRequest):
        bill_id, _ = core.create_bill(req.patient_id, req.visit_id, req.amount)
        return {"bill_id": bill_id}
    return app
