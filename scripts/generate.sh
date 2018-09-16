#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
PROTOBUF_PKG_REL_DIR="protobuf"
STEAMLANG_PKG_REL_DIR="steamlang"

readonly green="\\e[032m"
readonly reset="\\e[0m"

function step_msg() {
	echo -e "${green} * ${reset}$*..."
}

pushd "$ROOT_DIR" >/dev/null || exit $?

step_msg "Cleaning"
rake clean || exit $?

step_msg "Generating protobuf packages"
rake generate:protobuf || exit $?

step_msg "Building protobuf packages"
go build "./${PROTOBUF_PKG_REL_DIR}/..." || exit $?

step_msg "Generating steamlang package"
rake generate:steamlang || exit $?

step_msg "Building steamlang package"
go build "./${STEAMLANG_PKG_REL_DIR}" || exit $?

step_msg "Running tests"
go test "./..." || exit $?

step_msg "Done!"

popd >/dev/null || exit $?
