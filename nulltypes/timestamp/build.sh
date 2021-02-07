#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=. \
  timestamp.proto

mv ./github.com/weiwolves/protobuf/nulltypes/timestamp/* .
rm -rf ./github.com
