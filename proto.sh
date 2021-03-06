#!/bin/bash

export IMPORT_PATH=./third_party:.
export GENERATOR="gofast_out"
export OUTPUT_DIR="."
export PROTO_FILES="./*.proto"

export OPTIONS_API="Mgoogle/api/annotations.proto=."
export OPTIONS_API_V2="Mmodels.proto=github.com/otsimo/otsimopb,Mmessages.proto=github.com/otsimo/otsimopb,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types"
export OPTIONS_GW="Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,Mmodels.proto=github.com/otsimo/otsimopb,Mmessages.proto=github.com/otsimo/otsimopb,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types"

protoc --proto_path=$IMPORT_PATH --${GENERATOR}=${OPTIONS_API},plugins=grpc:${OUTPUT_DIR} $PROTO_FILES
protoc --proto_path=$IMPORT_PATH --${GENERATOR}=${OPTIONS_API_V2},plugins=grpc:.. ./v2/*.proto
protoc --proto_path=$IMPORT_PATH --grpc-gateway_out=${OPTIONS_API},logtostderr=true:${OUTPUT_DIR}/.. ./v2/*.proto

protoc --proto_path=$IMPORT_PATH --js-fetch_out=${OUTPUT_DIR} $PROTO_FILES
protoc --proto_path=$IMPORT_PATH --js-fetch_out=${OUTPUT_DIR} ./v2/*.proto