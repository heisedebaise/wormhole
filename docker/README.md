# run

```
docker run -d \
    -p 8080:8080 \
    --privileged=true \
    --restart=always \
    --network=local \
    -v /home/wormhole:/wormhole \
    --name=wormhole \
    alpine:base \
    /wormhole/run.sh
```