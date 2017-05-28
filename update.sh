#!/bin/bash

set -o errexit
set -o pipefail

SCRIPT_PATH="$(realpath "$0")"
SCRIPT_DIR="$(dirname "$SCRIPT_PATH")"
STEAMKIT_DIR="$SCRIPT_DIR/SteamKit"

function help() {
	local usage=$(cat <<-EOF
		Usage: %s [options] <SteamKit commit/branch/tag>

		Options:

		  -g    Generate go protobuf messages.
		  -h    Show this help.
		EOF
	)
	printf >&2 "$usage\n" "$0"
	exit 1
}

function git_init_submodules() {
	git submodules init || return $?
	git submodules update
}

function git_checkout_rev() {
	local rev="$1"

	git fetch --tags || return $?

	local head_hash="$(git rev-parse --verify "HEAD" 2> /dev/null)"
	local rev_hash="$(git rev-parse --verify "$rev" 2> /dev/null)"

	if [[ -z "$rev_hash" ]]; then
		printf >&2 "Could not resolve revision %s\n" "$rev"
		return 128
	fi

	if [[ "$head_hash" == "$rev_hash" ]]; then
		printf >&2 "Revision %s already checked out\n" "$rev"
		return 0
	fi

	git checkout -q --detach "$rev_hash" || return $?

	if [[ "$rev" == "$rev_hash" ]]; then
		fmt="Revision %s checked out"
	else
		fmt="Revision %s (%s) checked out"
	fi

	printf >&2 "$fmt\n" "$rev" "$rev_hash"
}

generate=0

while getopts "gh" opt; do
	case "$opt" in
		g) generate=1;;
		h) help ;;
	esac
done

shift $((OPTIND - 1))

revision="$1"

[[ -z "$revision" ]] && help

if [[ ! -d "$STEAMKIT_DIR" ]]; then
	git_init_submodules
fi

(
	cd "$STEAMKIT_DIR" && \
	git_checkout_rev "$revision"
)

[[ "$generate" -eq 0 ]] && exit 0

bash "$SCRIPT_DIR/generate.sh"
