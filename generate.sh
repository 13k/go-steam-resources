#!/bin/bash

set -e

SCRIPT_PATH="$(realpath "$0")"
SCRIPT_DIR="$(dirname "$SCRIPT_PATH")"

green="\e[032m"
reset="\e[0m"

printf "${green}Cleaning protobuf packages ...${reset}\n"
rm -rf "./protobuf"

printf "${green} * ${reset}Generating stubs ...\n"
go run "gen.go" "$SCRIPT_DIR"

printf "${green} * ${reset}Generating protobuf packages ...\n"
go generate "./protobuf/..."

printf "${green} * ${reset}Building protobuf packages ...\n"
go build "./protobuf/..."

printf "${green} * ${reset}Done!\n"
