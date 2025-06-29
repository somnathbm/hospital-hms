import pytest
from app.service import BillingServiceCore


@pytest.fixture
def billing_service():
    return BillingServiceCore()


def test_create_bill_success(billing_service):
    bill_id, message = billing_service.create_bill(
        "p-001", "visit-001", "d-101")
    # assert bill_id == "bill-visit-001"
    assert message == "Bill created successfully"


def test_create_bill_invalid_input(billing_service):
    bill_id, message = billing_service.create_bill("", "", "")
    assert bill_id == ""
    assert message == "Invalid input"
