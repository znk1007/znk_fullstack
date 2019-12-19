PROTO_ROOT_DIR = $(shell brew --prefix)/Cellar/protobuf/3.11.2/include
VERSION=0.0.1
default: clean stubs

## Dart requires you to manually ship all google provided proto files too.
## Dart requires you to manually ship all google provided proto files too.
stubs:
	mkdir -p lib/protos/generated
	mkdir -p lib/protos/generated/project
	protoc -I=lib/protos --dart_out=grpc:lib/protos/generated/project lib/protos/*.proto
	protoc -I$(PROTO_ROOT_DIR) --dart_out=lib/protos/generated $(PROTO_ROOT_DIR)/google/protobuf/*.proto

clean:
	rm -rf build lib/protos/generated

run-local:
	dart bin/server.dart

build-docker:
	docker build -t pambrose/flutter_module:${VERSION} .

run-docker:
	docker run -p 50051:50051 pambrose/flutter_module:${VERSION}

push-docker:
	docker push pambrose/flutter_module:${VERSION}