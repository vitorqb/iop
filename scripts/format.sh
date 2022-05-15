#!/bin/bash
USAGE="$0"' [-h] [-f]
RUns the linter.

-h)
  Display this help msg

-f)
  Fix problems.
'

# Import utils.sh
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/utils.sh

# CLI options parsing
SHORT='hf'
OPTS="$(getopt --options $SHORT --name "$0" -- "$@")"
! [ "$?" = 0 ] && echo "$USAGE" 1>&2 && exit 1
while [[ "$#" -gt 0 ]]
do
    case "$1" in
        -h)
            echo "$USAGE" >&2
            exit 0
            ;;
        -f)
            FIX=1
            shift
            ;;
        --)
            shift
            ;;
        *)
            err "Unexpected argument"
            ;;
    esac
done

# Script
ARGS=( )
if [ "$FIX" = "1" ]
then
    ARGS+=( "--fix" )
fi

run golangci-lint run "${ARGS[@]}"
