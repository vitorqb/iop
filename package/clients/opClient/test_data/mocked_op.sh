#!/bin/sh

# A helper script to control exit code of `whoami` and other commands.
if [ "$1" == "account" ] && [ "$2" == "get" ]
then
    exit {{.AccountGetExitCode}}
fi
{{.Body}}
