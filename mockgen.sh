#!/bin/bash
for filepath in ./api/v1/doer/*
do
    if [[ $filepath == *.pb.go ]]
    then
        echo "mocking " $filepath
        filename="$(basename -- $filepath)"
        mockgen -source=$filepath -destination=./mocks/$filename -package mocks
    fi
done
