#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=. \
  *.proto


cp -rf ./github.com/weiwolves/protobuf/nulltypes/* ./
rm -rf ./github.com