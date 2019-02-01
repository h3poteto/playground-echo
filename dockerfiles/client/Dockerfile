FROM golang:1.11-alpine

RUN set -ex && \
    apk add --no-cache \
    git

ENV APP_DIR $GOPATH/src/github.com/h3poteto/playground-echo/client

ADD ./client ${APP_DIR}

WORKDIR ${APP_DIR}

RUN set -ex && \
    go build

CMD ./client
