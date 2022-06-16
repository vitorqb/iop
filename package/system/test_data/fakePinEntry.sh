#!/bin/bash

# First, checks the PROMPT was set properly
read -r line
if [[ ! $line =~ SETPROMPT ]]
then
    exit 1
fi

# Secondly, checks the GETPIN was sent
read -r line
if [[ ! $line =~ ^GETPIN ]]
then
    exit 1
fi

# Thirdly, echos a pin
echo "D {{.Pin}}"
