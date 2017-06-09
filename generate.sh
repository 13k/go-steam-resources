#!/bin/bash

set -e

SCRIPT_PATH="$(realpath "$0")"
SCRIPT_DIR="$(dirname "$SCRIPT_PATH")"

rm -rf "./protobuf"
go run "gen.go" "$SCRIPT_DIR"
go generate "./protobuf/..."
go build "./protobuf/..."
