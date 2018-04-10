#!/bin/bash

CUR_GOPATH=`pwd`
OLD_GOPATH="$GOPATH"
export GOPATH="$CUR_GOPATH"
echo "set GOPATH=$GOPATH"

go get github.com/nfnt/resize
echo "installing github.com/nfnt/resize ..."

if [ -n "ls src/*" ];
then
   for file in src/*
   do
       echo "installing ${file#*/} ..."
       go install ${file#*/}
   done
fi

export GOPATH="$OLD_GOPATH"
echo "reset GOPATH=$GOPATH"