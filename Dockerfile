FROM centos:base

ENV GOPATH=/usr/lib/go:/wormhole

RUN yum install -y git go \
    && git clone https://github.com/heisedebaise/wormhole.git \
    && cd wormhole \
    && sh install.sh

WORKDIR /wormhole
EXPOSE 8192

ENTRYPOINT [ "bin/image" ]