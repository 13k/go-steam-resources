#!/bin/bash

set -e

SCRIPT_PATH="$(realpath "$0")"
SCRIPT_DIR="$(dirname "$SCRIPT_PATH")"
STEAMKIT_DIR="$SCRIPT_DIR/SteamKit"

function help() {
	local usage=$(cat <<-EOF
		Usage: %s [options] <SteamKit tag>

		Options:

		  -g    Generate go protobuf messages.
		  -h    Show this help.
		EOF
	)
	printf >&2 "$usage\n" "$0"
	exit 1
}

function git_init_submodules() {
	git submodules init
	git submodules update
}

function git_checkout_tag() {
	local tag="$1"

	git fetch --tags

	local head_hash="$(git show-ref --verify --hash --head -- "HEAD")"
	local tag_hash="$(git show-ref --verify --hash --head -- "refs/tags/$tag")"

	if [[ "$head_hash" == "$tag_hash" ]]; then
		printf >&2 "Tag %s already checked out\n" "$tag"
		return 0
	fi

	git checkout -q --detach "refs/tags/$tag"
}

generate=0

while getopts "gh" opt; do
	case "$opt" in
		g) generate=1;;
		h) help ;;
	esac
done

shift $((OPTIND - 1))

release="$1"

[[ -z "$release" ]] && help

if [[ ! -d "$STEAMKIT_DIR" ]]; then
	git_init_submodules
fi

(
	cd "$STEAMKIT_DIR" && \
	git_checkout_tag "$release"
)

[[ "$generate" -eq 0 ]] && exit 0

bash "$SCRIPT_DIR/generate.sh"
