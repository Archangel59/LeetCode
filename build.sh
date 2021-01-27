#!/bin/bash
echo "enter the question number."
read number
mkdir $number
touch $number/$number.go
touch $number/$number.md
echo "package main" > $number/$number.go
touch $number/$number.cpp
echo -e "#include<bits/stdc++.h>\nusing namespace std;" > $number/$number.cpp
echo "build $number success !"