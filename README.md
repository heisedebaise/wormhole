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

http/s to http/s

```
{
    "http://:8085": "http://192.168.0.5",
    "https://:8086": "http://192.168.0.6",
    "http://:8087": "https://192.168.0.7",
    "https://:8088": "http://192.168.0.8"
}
```

use https://github.com/FiloSottile/mkcert to generate tls

```
mkcert -key-file key.pem -cert-file cert.pem  localhost
```