
GOPATH:=$(shell go env GOPATH)


.PHONY: proto build test
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/consignment/consignment.proto

build: proto
	GOOS=linux GOARCH=amd64 go build -o consignment-srv main.go
	docker build -t consignment-srv .

test:
	go test -v ./... -cover