#!/bin/bash
cd ../../../capsule-launcher
COUCHBASE_CLUSTER="couchbase://localhost" \
COUCHBASE_USER="admin" \
COUCHBASE_PWD="ilovepandas" \
COUCHBASE_BUCKET="wasm-data" \
go run main.go \
   -wasm=../wasm_modules/capsules-couchbase/http-module/hello.wasm \
   -mode=http \
   -httpPort=7070

