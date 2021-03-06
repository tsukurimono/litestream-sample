FROM golang:1.15.7-alpine as dev

ENV ROOT=/go/src/litestream-sample
ENV GOMODCACHE=/go/cache/go_mod
ENV GO111MODULE=on
ENV CGO_ENABLED 1
WORKDIR ${ROOT}

RUN apk update && apk add git \
    vim \
    bash \
    gcc \
    libc-dev \
    sqlite \
    sqlite-dev

CMD ["bash"]
