#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  nullstring.proto

mv ./github.com/weiwolves/protobuf/nulltypes/nullstring/* .
rm -rf ./github.com
