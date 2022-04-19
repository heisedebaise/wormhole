#!/bin/bash

git add bin
git add conf
git add main
git add go.mod
git add *.go
git add map.json
git add push.sh
git add README.md
git add run.sh

git commit -m dev
git push
git push github