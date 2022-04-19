#!/bin/sh

cd /wormhole
nohup ./wormhole </dev/null >/var/log/wormhole 2>&1 &