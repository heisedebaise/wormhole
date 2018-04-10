FROM centos:base

RUN yum install -y git go \
    && git clone https://github.com/heisedebaise/wormhole.git \
    && cd wormhole \
    && sh install.sh

ENV GOPATH=/wormhole
WORKDIR /wormhole
EXPOSE 8192

ENTRYPOINT [ "go", "run" "src/image/service.go" ]