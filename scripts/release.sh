#!/bin/bash
USAGE="$0"' [-h] VERSION
Creates a gh release.

-h)
  Display this help msg
'

set -e

# Import utils.sh
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/utils.sh

# CLI options parsing
SHORT='h'
OPTS="$(getopt --options $SHORT --name "$0" -- "$@")"
! [ "$?" = 0 ] && echo "$USAGE" 1>&2 && exit 1
while [[ "$#" -gt 0 ]]
do
    echo "$1"
    case "$1" in
        -h)
            echo "$USAGE" >&2
            exit 0
            ;;
        --)
            shift
            VERSION="$1"
            shift
            ;;
        *)
            err "Unexpected argument"
            ;;
    esac
done

# Sanity check
if [[ ! $VERSION =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]
then
    err "Version should be of format x.y.z, but received $VERSION"
fi

# Script
GH_TOKEN_FILE="${SCRIPT_DIR}/../.gh_token"
if [ -r $GH_TOKEN_FILE ]
then
    msg "Reading GH_TOKEN from $GH_TOKEN_FILE"
    export GH_TOKEN="$(cat $GH_TOKEN_FILE | tr -d '\n' )"
fi
run ${SCRIPT_DIR}/build.sh
run gh release create "$VERSION" --generate-notes ./dist/iop.tar.gz
