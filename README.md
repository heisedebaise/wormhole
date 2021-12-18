# run

```
go run main.go
```

# build

```
rm wormhole
go build -o wormhole main.go
```

# config

map.json

```
{
    ":80": "192.168.0.3:8080",
    ":8080": "192.168.0.4:8080"
}
```

- https

```
{
    ":443": "https://192.168.0.5",
    ":8443": "https://192.168.0.6:8443"
}
```

> https://github.com/FiloSottile/mkcert

```
mkcert -key-file key.pem -cert-file cert.pem  192.168.0.7
```