#!/bin/bash

dir=`pwd`
export GOPATH=$GOPATH:$dir
echo "using GOPATH=$GOPATH"
bin/image
