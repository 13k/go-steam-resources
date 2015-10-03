#!/bin/bash

function help() {
	echo "Usage: $0 <SteamKit tag>"
	exit 1
}

[[ -z "$1" ]] && help

pushd SteamKit
git fetch --tags
git checkout -b "$1" "$1"
popd
