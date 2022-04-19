#!/bin/bash

git add bin
git add conf
git add docker
git add main
git add go.mod
git add *.go
git add map.json
git add README.md

git commit -m dev
git push
git push github