#!/bin/bash

python -m grpc_tools.protoc --descriptor_set_out=true --include_imports=. --proto_path=protos --python_out=./gen/python --grpc_python_out=./gen/python ./protos/*.proto
protoc --go_out ./gen/golang --descriptor_set_out=true --include_imports=. --proto_path=protos --go_opt paths=source_relative --go-grpc_out ./gen/golang --go-grpc_opt paths=source_relative ./protos/*.proto