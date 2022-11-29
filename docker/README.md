# run

```
docker run -d \
    -p 80:80 \
    --privileged=true \
    --restart=always \
    --network=local \
    -v /home/wormhole:/wormhole \
    --name=wormhole \
    alpine:base \
    /wormhole/run.sh
```