#!/bin/bash
if  [ ! -n "$1" ] ;then
    echo "you have not input a number!"
else
    g++ $1/$1.cpp
    ./a.out
    rm a.out
fi