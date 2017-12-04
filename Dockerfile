FROM golang:1.9-alpine

RUN set -ex && \
    apk add --no-cache \
    git

ENV APP_DIR $GOPATH/src/github.com/h3poteto/playground-echo

ADD . ${APP_DIR}

WORKDIR ${APP_DIR}

EXPOSE 9090

RUN set -ex && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    go build

CMD ./playground-echo
