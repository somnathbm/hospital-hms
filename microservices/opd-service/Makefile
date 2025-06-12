APP_NAME := opd-service
PROTO_DIR := proto
GEN_DIR := gen
GO_PKG := github.com/somnathbm/hospital-hms

.PHONY: all build proto test clean run

all: build

build:
	go build -o bin/$(APP_NAME) ./cmd/server

proto:
	protoc -I. \
		-I$(PROTO_DIR) \
	  -I$(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2) \
	  -I$(shell go list -f '{{ .Dir }}' -m google.golang.org/protobuf) \
	  -I$(shell go list -f '{{ .Dir }}' -m github.com/googleapis/googleapis) \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GEN_DIR) \
		--go-grpc_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(GEN_DIR) --grpc-gateway_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto

test:
	go test ./internal/service/... -v

run:
	go run ./cmd/server/main.go

clean:
	rm -rf bin/
	rm -rf $(GEN_DIR)/*.go
