APP_NAME := opd-service
PROTO_DIR := proto
SHARED_PROTO_DIR := ../../shared/proto
OPD_GEN_DIR := gen/opdpb
BILLING_GEN_DIR := gen/billingpb
APT_GEN_DIR := gen/aptpb
GO_PKG := github.com/somnathbm/hospital-hms

.PHONY: all build proto test clean run

all: build

build:
	go build -o bin/$(APP_NAME) ./cmd/server

proto-opd:
	protoc \
		--proto_path=$(SHARED_PROTO_DIR) \
		--go_out=$(OPD_GEN_DIR) \
		--go-grpc_out=$(OPD_GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(SHARED_PROTO_DIR)/opd.proto

proto-billing:
	protoc \
		--proto_path=$(SHARED_PROTO_DIR) \
		--go_out=$(BILLING_GEN_DIR) \
		--go-grpc_out=$(BILLING_GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(SHARED_PROTO_DIR)/billing.proto

proto-apt:
	protoc \
		--proto_path=$(SHARED_PROTO_DIR) \
		--go_out=$(APT_GEN_DIR) \
		--go-grpc_out=$(APT_GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(SHARED_PROTO_DIR)/appointment.proto

test:
	go test ./internal/service/... -v

run:
	go run ./cmd/server/main.go

clean:
	rm -rf bin/
	rm -rf $(GEN_DIR)/*.go
