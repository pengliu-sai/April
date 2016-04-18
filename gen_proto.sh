#!/bin/sh


protoc --plugin=protoc-gen-go=./bin/protoc-gen-go --go_out=. ./src/protos/msgid/*.proto
protoc --plugin=protoc-gen-go=./bin/protoc-gen-go --go_out=. ./src/protos/system/*.proto
protoc --plugin=protoc-gen-go=./bin/protoc-gen-go --go_out=. ./src/protos/admin/*.proto
protoc --plugin=protoc-gen-go=./bin/protoc-gen-go --go_out=. ./src/protos/gate/*.proto
protoc --plugin=protoc-gen-go=./bin/protoc-gen-go --go_out=. ./src/protos/world/*.proto
protoc --plugin=protoc-gen-go=./bin/protoc-gen-go --go_out=. ./src/protos/db/*.proto
protoc --plugin=protoc-gen-go=./bin/protoc-gen-go --go_out=. ./src/protos/config/*.proto

cd src/protos/game
protoc --plugin=protoc-gen-go=./../../../bin/protoc-gen-go --go_out=. game.proto
