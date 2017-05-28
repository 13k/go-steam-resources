#!/bin/bash

function help() {
	echo "Usage: $0 <SteamKit tag>"
	exit 1
}

[[ -z "$1" ]] && help

steamkit_dir="SteamKit"

if [[ ! -d "$steamkit_dir" ]]; then
	git submodules init
	git submodules update
fi

pushd "$steamkit_dir"
git fetch --tags
git checkout -b "$1" "$1"
popd
