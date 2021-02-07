#!/bin/sh

cd empty
protoc \
  --go_out=. \
  --proto_path=.. \
  empty.proto

mv ./github.com/weiwolves/protobuf/nulltypes/empty/* .
rm -rf ./github.com
cd -

cd nulldate
protoc \
  --go_out=. \
  --proto_path=.. \
  nulldate.proto
mv ./github.com/weiwolves/protobuf/nulltypes/nulldate/* .
rm -rf ./github.com
cd -

cd nullint
protoc \
  --go_out=. \
  --proto_path=.. \
  nulldate.proto
mv ./github.com/weiwolves/protobuf/nulltypes/nulldate/* .
rm -rf ./github.com
cd -

cd nullstring
protoc \
  --go_out=. \
  --proto_path=.. \
  nullstring.proto
mv ./github.com/weiwolves/protobuf/nulltypes/nullstring/* .
rm -rf ./github.com
cd -

cd timestamp
protoc \
  --go_out=. \
  --proto_path=.. \
  timestamp.proto
mv ./github.com/weiwolves/protobuf/nulltypes/timestamp/* .
rm -rf ./github.com
cd -
