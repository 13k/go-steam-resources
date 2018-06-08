#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
STEAMDB_PROTO_DIR="$ROOT_DIR/SteamDatabase/Protobufs"

function help() {
	local usage

  usage="$(cat <<-EOF
		Usage: %s [options]

		Options:

		  -g    Generate go protobuf messages.
		  -h    Show this help.
		EOF
	)"

  # shellcheck disable=SC2059
	printf >&2 "$usage\\n" "$0"
	exit 1
}

function git_init_submodules() {
	git submodules init && \
  git submodules update
}

function git_checkout_rev() {
	local rev="$1" head_hash rev_hash

	git fetch --tags || return $?

	head_hash="$(git rev-parse --verify "HEAD" 2> /dev/null)" || return $?
	rev_hash="$(git rev-parse --verify "$rev" 2> /dev/null)" || return $?

	if [[ -z "$rev_hash" ]]; then
		printf >&2 'Could not resolve revision %s\n' "$rev"
		return 128
	fi

	if [[ "$head_hash" == "$rev_hash" ]]; then
		printf >&2 'Revision %s already checked out\n' "$rev"
		return 0
	fi

	git checkout -q --detach "$rev_hash" || return $?

	if [[ "$rev" == "$rev_hash" ]]; then
		fmt="Revision %s checked out"
	else
		fmt="Revision %s (%s) checked out"
	fi

  # shellcheck disable=SC2059
	printf >&2 "$fmt\\n" "$rev" "$rev_hash"
}

generate=0

while getopts "gh" opt; do
	case "$opt" in
		g) generate=1;;
		h) help ;;
    *) help ;;
	esac
done

shift $((OPTIND - 1))

if [[ ! -d "$STEAMDB_PROTO_DIR" ]]; then
	git_init_submodules || exit $?
fi

(
	cd "$STEAMDB_PROTO_DIR" && \
	git_checkout_rev "origin/master"
) || exit $?

[[ "$generate" -eq 0 ]] && exit 0

exec bash "$SCRIPT_DIR/generate.sh"
