#!/bin/bash
echo "enter the question number."
read number
mkdir $number
touch $number/$number.go
touch $number/$number.md
echo "package main" > $number/$number.go
echo "build $number success !"