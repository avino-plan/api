#!/bin/sh
LANGUAGE=$1
PROTO=$2
OUT=$3
mkdir -p "$OUT"
protoc -I=./proto --"$LANGUAGE"_out="$OUT" --"$LANGUAGE"_opt=paths=source_relative --"$LANGUAGE"-grpc_out="$OUT" --"$LANGUAGE"-grpc_opt=paths=source_relative "$PROTO"