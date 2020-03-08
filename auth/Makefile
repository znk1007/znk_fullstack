PROTO_ROOT_DIR = $(shell brew --prefix)/Cellar/protobuf/3.11.2/include
VERSION=0.0.1

default: clean generate_proto 

generate_proto:
	mkdir -p lib/other/protos/generated
	mkdir -p lib/other/protos/generated/project
	protoc -I=lib/other/protos --dart_out=grpc:lib/other/protos/generated/project lib/auth/protos/*.proto
	protoc -I$(PROTO_ROOT_DIR) --dart_out=lib/other/protos/generated $(PROTO_ROOT_DIR)/google/protobuf/*.proto

clean:
	rm -rf build /lib/other/protos/generated