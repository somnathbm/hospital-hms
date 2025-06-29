import uuid

class BillingServiceCore:
    def __init__(self):
        # in-memory store
        self.bills = {}


    def create_bill(self, patient_id, visit_id, amount):
        if patient_id == "" or visit_id == "" or amount == "":
            return "", "Invalid input"

        bill_id = str(uuid.uuid4())
        self.bills[bill_id] = {
            "patient_id": patient_id,
            "visit_id": visit_id,
            "amount": amount
        }
        return bill_id, "Bill created successfully"
