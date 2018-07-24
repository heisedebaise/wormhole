# HTTPS

```
openssl genrsa -out wormhole.key 2048
openssl req -new -x509 -key wormhole.key -out wormhole.crt -days 365
```