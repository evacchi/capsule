#!/bin/bash
cd ../../capsule-launcher

go run main.go \
   -wasm=../wasm_modules/capsule-http-get/hello.wasm \
   -mode=cli \
   "[GET]👋 hello world 🌍"
