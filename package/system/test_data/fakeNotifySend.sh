#!/bin/bash
title="$1"
body="$2"
echo -n "title=$title;body=$body" >{{.OutputFile}}
