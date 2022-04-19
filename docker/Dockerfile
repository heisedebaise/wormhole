FROM go:1.18

RUN git clone https://github.com/heisedebaise/wormhole.git

WORKDIR /wormhole

CMD [ "go", "run", "main/main.go", "</dev/null", ">/var/log/wormhole", "2>&1" ]