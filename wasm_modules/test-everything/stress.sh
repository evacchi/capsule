#!/bin/bash
# -n 10000 -c 1000
#hey -n 10 -c 5 -m POST \
hey -n 300 -c 100 -m POST \
-H "Content-Type: application/json" \
-d '{"message": "Golang 💚💜 wasm"}' \
"http://localhost:7070"

# Requests/sec: 30.5280
# Requests/sec: 32.1274
# Requests/sec: 31.5995