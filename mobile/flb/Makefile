PROTO_ROOT = $(shell brew --prefix)/Cellar/protobuf/3.12.3/include
VERSION = 0.0.2
default: clean generate
#3.11.4
clean:
	rm -rf build lib/util/db/protos/generated
## Dart requires you to manually ship all google provided proto files too.
generate:
	mkdir -p lib/util/db/protos/generated
	mkdir -p lib/util/db/protos/generated/user
	protoc -I=lib/util/db/protos --dart_out=grpc:lib/util/db/protos/generated/user lib/util/db/protos/*.proto
	# protoc -I$(PROTO_ROOT) --dart_out=lib/util/db/protos/generated $(PROTO_ROOT)/google/protobuf/*.proto
