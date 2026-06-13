#!/usr/bin/env bash 
ENV="prod"
echo "Hey There!"

file="practice.sh"
for file in *.sh
do
    bash $file 
done