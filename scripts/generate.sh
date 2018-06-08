#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
PKG_REL_DIR="protobuf"

readonly green="\\e[032m"
readonly reset="\\e[0m"

function step_msg() {
  echo -e "${green} * ${reset}$*..."
}

pushd "$ROOT_DIR" > /dev/null || exit $?

step_msg "Cleaning"
rake -s distclean || exit $?

step_msg "Generating protobuf packages"
rake -s || exit $?

step_msg "Building protobuf packages"
go build "./${PKG_REL_DIR:?}/..." || exit $?

step_msg "Done!"

popd > /dev/null || exit $?
