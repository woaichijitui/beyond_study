#!/bin/bash

proto_file=$1



goctl rpc protoc $proto_file --go_out=. --go-grpc_out=. --zrpc_out=./