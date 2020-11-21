#!/bin/bash

BASEDIR=$(dirname "$0")

protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ${BASEDIR}/weight_tracker.proto