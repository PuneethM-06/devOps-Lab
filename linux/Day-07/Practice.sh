#!/bin/bash
echo "Hello, World!"
# Question: How to print a variable in bash?
name="Alice"
echo "Hello, $name"
date=$(date)
user=$(whoami)
echo $date
echo $user

# read command
echo "What is your name?"
read name
echo "Hello, $name!"


