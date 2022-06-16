#!/bin/sh

# A helper script to control exit code of `whoami` and other commands.
if [ "$1" == "whoami" ]
then
    exit {{.WhoAmIExitCode}}
fi
{{.Body}}
