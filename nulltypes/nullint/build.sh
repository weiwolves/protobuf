#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  nullint.proto

mv ./github.com/weiwolves/protobuf/nulltypes/nullint/* .
rm -rf ./github.com
