.PHONY: all clean update-deps check-version gen gen-to gen-reset

SHELL:=/bin/bash

all: docker-build

clean:
	rm -rf go
	mkdir go

update-deps:
	go get -u github.com/golang/protobuf/proto
	go get -d -u github.com/golang/protobuf/protoc-gen-go
	git -C $(shell go env GOPATH)/src/github.com/golang/protobuf checkout $(PROTOBUF_VERSION) # checkout as branch to prevent detached head
	go install github.com/golang/protobuf/protoc-gen-go
	git -C $(shell go env GOPATH)/src/github.com/golang/protobuf checkout master # switch back to master for dep operation
	go get -u google.golang.org/grpc

gen:
	for a in $(shell find ./proto -name *.proto); do protoc -I=./proto --go_out=plugins=grpc:${GOPATH}/src $$a; done

# deprecate for services using go mod, add the following instead
# replace github.com/yougroupteam/u-protocol => <local path>/u-protocol
gen-to: gen
	cp -Rf ./go ../${PROJ}/vendor/github.com/tyeryan/l-protocol/

# deprecate for services using go mod, add the following instead
# replace github.com/yougroupteam/u-protocol => <local path>/u-protocol
gen-reset:
	git -C ../${PROJ}/vendor/github.com/tyeryan/l-protocol/ reset --hard HEAD

docker-build:
	docker build -t l-protocol .
	docker run --rm -v ${CURDIR}:/go/src/github.com/tyeryan/l-protocol l-protocol
