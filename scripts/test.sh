#!/bin/bash
USAGE="$0"' [-h] [-v] [-t testname]
Runs tests for the project.

-h)
  Display this help msg.

-t testname)
  Runs only testname.

-v)
  Verbose (show logs)
'

# Import utils.sh
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
source ${SCRIPT_DIR}/utils.sh

# Defaults
TESTNAME=""
VERBOSE=0

# CLI options parsing
SHORT='ht:v'
OPTS="$(getopt --options $SHORT --name "$0" -- "$@")"
! [ "$?" = 0 ] && echo "$USAGE" 1>&2 && exit 1
while [[ "$#" -gt 0 ]]
do
    case "$1" in
        -h)
            echo "$USAGE" >&2
            exit 0
            ;;
        -t)
            TESTNAME="$2"
            shift
            shift
            ;;
        -v)
            VERBOSE=1
            shift
            ;;
        *)
            echo "ERROR: Unexpected argument" >&2
            exit 1
            ;;
        --)
            shift
            ;;
    esac
done

# Script
EXTRA_ARGS=( )
if [ ! -z $TESTNAME ]
then
    EXTRA_ARGS+=( "-run" $TESTNAME )
fi
if [ $VERBOSE = 1 ]
then
    EXTRA_ARGS+=( "-v" )
fi
run go test ./... ${EXTRA_ARGS[@]}
