#!/bin/sh

cd empty
./build.sh
cd -

cd nulldate
./build.sh
cd -

cd nullint
./build.sh
cd -

cd nullstring
./build.sh
cd -

cd timestamp
./build.sh
cd -
