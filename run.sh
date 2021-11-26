#!/bin/bash

chmod +x wormhole

nohup ./wormhole </dev/null >/var/log/wormhole 2>&1 &