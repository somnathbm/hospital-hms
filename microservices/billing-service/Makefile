PROTO_DIR := app/proto
GEN_DIR := app/gen

proto:
	python -m grpc_tools.protoc -I$(PROTO_DIR) \
  		--python_out=$(GEN_DIR) \
  		--grpc_python_out=$(GEN_DIR) \
  		app/proto/appointment.proto