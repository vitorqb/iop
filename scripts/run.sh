#!/bin/bash
USAGE="$0"' [-h] [args...]
Runs the project.

-h)
  Display this help msg
'

# Import utils.sh
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/utils.sh

# CLI options parsing
SHORT='h'
OPTS="$(getopt --options $SHORT --name "$0" -- "$@")"
! [ "$?" = 0 ] && echo "$USAGE" 1>&2 && exit 1
while [[ "$#" -gt 0 ]]
do
    case "$1" in
        -h)
            echo "$USAGE" >&2
            exit 0
            ;;
        --)
            # Leave it for the go run cmd
            shift
            break
            ;;
        *)
            echo "ERROR: Unexpected argument" >&2
            exit 1
            ;;
    esac
done

# Script
run go run main.go "$@"
