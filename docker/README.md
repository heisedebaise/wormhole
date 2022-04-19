# go:1.18

## 构建

```bash
docker build -t wormhole:latest docker/

podman build -t wormhole:latest docker/
```

## 运行

```bash
docker run -d \
    -p 8080:8080 \
    --privileged=true \
    --restart=always \
    --network=local \
    --name=wormhole \
    wormhole:latest

podman run -d \
    -p 8080:8080 \
    --privileged=true \
    --pod=local \
    --name=wormhole \
    wormhole:latest
```