#!/bin/bash
if  [ ! -n "$1" ] ;then
    echo "you have not input a number!"
else
    g++ -std=c++14 $1/$1.cpp
    ./a.out
    rm a.out
fi