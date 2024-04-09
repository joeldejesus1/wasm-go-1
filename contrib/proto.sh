#!/bin/bash

set -e

decho(){
    1>&2 echo $@
}

PROTO_BIN="protoc"
PROTO_DIR="./contrib/proto"
PROTO_GO_DIR="./proto"



test_all(){
    decho $@
}

build_go(){
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    rm -r $PROTO_GO_DIR 2>>/dev/null || true
    $PROTO_BIN --proto_path=${PROTO_DIR} --go-grpc_out=. --go_out=. $@
    mv "github.com/joeldejesus1/wasm-go-1/proto" ./
    rm -r github.com
}

build_go $(echo $(find ./contrib/proto -type f))
