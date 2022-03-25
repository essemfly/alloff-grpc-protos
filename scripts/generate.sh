#!/bin/bash

cd protos

python -m grpc_tools.protoc --descriptor_set_out=true --include_imports=.. --proto_path=. --python_out=../gen/python --grpc_python_out=../gen/python ./*.proto
protoc --go_out ../gen/golang --descriptor_set_out=true --include_imports=.. --go_opt paths=source_relative --go-grpc_out ../gen/golang --go-grpc_opt paths=source_relative ./*.proto