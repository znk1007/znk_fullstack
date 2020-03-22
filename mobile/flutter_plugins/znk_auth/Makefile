PROTO_ROOT = $(shell brew --prefix)/Cellar/protobuf/3.11.2/include
VERSION = 0.0.2
default: clean generate

clean:
	rm -rf build lib/model/protos/generated
## Dart requires you to manually ship all google provided proto files too.
generate:
	mkdir -p lib/model/protos/generated
	mkdir -p lib/model/protos/generated/auth
	protoc -I=lib/model/protos --dart_out=grpc:lib/model/protos/generated/auth lib/model/protos/*.proto
	protoc -I$(PROTO_ROOT) --dart_out=lib/model/protos/generated $(PROTO_ROOT)/google/protobuf/*.proto