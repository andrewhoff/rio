all: proto build

proto:
	protoc -I helloproto/ helloproto/hello.proto --go_out=plugins=grpc:helloproto

build:
	go build .