#!/bin/sh
[ -f compile ] && go build -o compile compile.go
./compile