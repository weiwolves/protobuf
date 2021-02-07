#!/bin/sh

protoc -I . -I $(go list -m -f "{{.Dir}}" github.com/srikrsna/protoc-gen-gotag) \
  --go_out=plugins=grpc:../domain \
  *.proto

# protoc-go-tags --dir=../domain
