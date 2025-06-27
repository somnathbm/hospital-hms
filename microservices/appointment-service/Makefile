SHARED_PROTO_DIR := ../../shared/proto
GEN_DIR := app/gen

proto:
	python -m grpc_tools.protoc -I$(SHARED_PROTO_DIR) \
  		--python_out=$(GEN_DIR) \
  		--grpc_python_out=$(GEN_DIR) \
  		$(SHARED_PROTO_DIR)/appointment.proto