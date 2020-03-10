#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  nullint.proto

mv ./github.com/weiwolves/protobuf/types/nullint/* .
rm -rf ./github.com
