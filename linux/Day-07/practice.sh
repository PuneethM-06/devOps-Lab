#!/usr/bin/env bash 
ENV="prod"

case "$ENV" in
    dev)
        echo "It is dev"
        ;;
    test)
        echo "IT is test"
        ;;
    prod)
        echo "it is prod"
        ;;
    *)
        echo "Idk"
    ;;
esac