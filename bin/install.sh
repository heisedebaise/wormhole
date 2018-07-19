#!/bin/bash

dir=`pwd`
export GOPATH=$GOPATH:$dir
echo "using GOPATH=$GOPATH"

echo "installing github.com/google/uuid ..."
go get github.com/google/uuid
echo "installing github.com/nfnt/resize ..."
go get github.com/nfnt/resize

if [ -n "ls src/*" ]; then
    for file in src/*; do
        if [ "${file##*.}" = "go" ]; then
            name=${file%.*}
            echo "building $file to ${name/src/bin} ..."
            go build -o ${name/src/bin} $file
        fi
    done
fi
