#!/bin/bash
if  [ ! -n "$1" ] ;then
    echo "you have not input a number!"
else
    go run $1/$1.go
fi