#!/bin/bash

python -m grpc_tools.protoc --descriptor_set_out=true --include_imports=. --proto_path=protos --python_out=./gen/pyalloff --grpc_python_out=./gen/pyalloff ./protos/*.proto
protoc --go_out ./gen/goalloff --descriptor_set_out=true --include_imports=. --proto_path=protos --go_opt paths=source_relative --go-grpc_out ./gen/goalloff --go-grpc_opt paths=source_relative ./protos/*.proto