#!/bin/bash

git add go.mod
git add main.go
git add map.json
git add push.sh
git add README.md
git add run.sh

git commit -m dev
git push
git push github