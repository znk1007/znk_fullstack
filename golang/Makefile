PROTO_ROOT_DIR = $(shell brew --prefix)/Cellar/protobuf/3.7.1/include
VERSION=0.1.0

default: clean stubs

## Dart requires you to manually ship all google provided proto files too.
stubs:
	mkdir -p protos/generated
	protoc -I=protos --gogofaster_out=plugins=grpc:protos/generated protos/*.proto

clean:
	rm -rf build protos/generated

run-local:
	dart bin/server.dart

build-docker:
	docker build -t pambrose/flutter_module:${VERSION} .

run-docker:
	docker run -p 50051:50051 pambrose/flutter_module:${VERSION}

push-docker:
	docker push pambrose/flutter_module:${VERSION}
