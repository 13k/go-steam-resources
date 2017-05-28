#!/bin/bash

SCRIPT_PATH="$(realpath "$0")"
SCRIPT_DIR="$(dirname "$SCRIPT_PATH")"

go run "gen.go" "$SCRIPT_DIR" && \
	go build ./protobuf
