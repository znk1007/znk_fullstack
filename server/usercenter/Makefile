PROTO_ROOT_DIR = $(shell brew --prefix)/Cellar/protobuf/3.7.1/include
VERSION=0.0.3

default: clean generated

generated:
	mkdir -p model/protos/generated
	protoc -I=model/protos --go_out=plugins=grpc:model/protos/generated model/protos/*.proto

clean:
	rm -rf build model/protos/generated
