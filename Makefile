.PHONY: client server
DEP_CACHE=$(PWD)/.dep/

all: client

clean:
	rm -rf vendor
	find . -name *.test -delete

pull:
	docker pull bearstech/golang-dev:latest
	docker pull bearstech/golang-dep:latest

client: bin vendor
	go build -ldflags "-X version.GitVersion=$(shell git rev-parse HEAD || echo '🐔')" \
		-o bin/garc github.com/factorysh/gitlab-authenticated-rpc-client

protoc:
	go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		gitlab.proto --go_out=plugins=grpc:rpc_gitlab
	protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	    auth.proto --go_out=plugins=grpc:rpc_auth

vendor:
	make dep

dep:
	mkdir -p $(GOPATH)/src
	dep ensure -v

bin:
	mkdir -p bin

docker-client: bin
	rm -f bin/gar*
	docker run --rm \
		-v `pwd`:/go/src/github.com/factorysh/gitlab-authenticated-rpc-client \
		-v /tmp:/.cache \
		-w /go/src/github.com/factorysh/gitlab-authenticated-rpc-client \
		-u `id -u` \
		bearstech/golang-dep:latest \
		make client
