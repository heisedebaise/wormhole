#!/bin/bash

dir=`pwd`
export GOPATH=$GOPATH:$dir
echo "using GOPATH=$GOPATH"

for name in github.com/google/uuid github.com/nfnt/resize github.com/gorilla/websocket
do
    echo "installing $name ..."
    go get $name
done

if [ -n "ls src/*" ]; then
    for file in src/*; do
        if [ "${file##*.}" = "go" ]; then
            name=${file%.*}
            echo "building $file to ${name/src/bin} ..."
            go build -o ${name/src/bin} $file
        fi
    done
fi
