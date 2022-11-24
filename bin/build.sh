#!/bin/bash

rm -rf wormhole
go build -o wormhole main/main.go

dir=/mnt/hgfs/share/wormhole
mkdir -p $dir
rm -rf $dir/wormhole
cp wormhole $dir/