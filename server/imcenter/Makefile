PROTO_ROOT_DIR = $(shell brew --prefix)/Cellar/protobuf/3.7.1/include
VERSION=0.0.3

default: clean generated

generated:
	mkdir -p source/model/protos/generated
	protoc -I=source/model/protos --go_out=plugins=grpc:source/model/protos/generated source/model/protos/*.proto

clean:
	rm -rf build source/model/protos/generated
