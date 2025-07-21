# billing-service

This is the Billing service.
Like happens in a hospital, different departents such as insurance, labs, OPD or IPD or emergency etc. relies on this services. The service now exposes `/metrics` for APM monitoring.

Here other services calls this services via gRPC and REST.

## ⛩️ Architecture diagram
![billing-service-diagram](../../docs/diagram/billing-service-diagram.svg)