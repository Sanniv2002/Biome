# Paths
PROTO_DIR = .       # Directory where your .proto files are located
OUT_DIR = .         # Directory where you want to generate the Go code

# Protobuf files
PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)

# Generate Go code from .proto files
.PHONY: generate
generate:
    @for proto in $(PROTO_FILES); do \
        protoc -I=$(PROTO_DIR) --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) $$proto; \
    done

# Clean generated files
.PHONY: clean
clean:
    @rm -f $(OUT_DIR)/*.pb.go

# Default target
.PHONY: all
all: generate
