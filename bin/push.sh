#!/bin/bash

git add bin
git add docker
git add main
git add go.mod
git add *.go
git add config.json
git add README.md

git commit -m dev
git push
git push github