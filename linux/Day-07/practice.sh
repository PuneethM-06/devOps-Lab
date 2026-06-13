#!/usr/bin/env bash 
ENV="prod"

if [ $ENV == "test" ]; then
    echo "This is a test env"
elif [ $ENV == "dev" ]; then 
    echo "this is a dev env"
else 
    echo "This is a prod env"
fi
