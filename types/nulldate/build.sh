#!/bin/sh

protoc \
  --go_out=. \
  --proto_path=.. \
  nulldate.proto

mv ./github.com/weiwolves/protobuf/types/nulldate/* .
rm -rf ./github.com
