
swagger:
	swag init -g ./app/cmd/app/main.go -o ./app/docs

migrate:

PROTO_DIR = app/sdk/proto
OUT_DIR   = .

PROTOC      = protoc
PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GO_GRPC = protoc-gen-go-grpc

PROTO_FILES = $(shell find $(PROTO_DIR) -name '*.proto')

proto:
	@echo ">> Generating Go code from proto files..."
	$(PROTOC) \
		--go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)

clean:
	@echo ">> Cleaning generated files..."
	@find $(PROTO_DIR) -name '*.pb.go' -delete
