FROM golang:1.20-buster

ENV APP_DIR /tracker
ENV GO111MODULE on

# Installing esentials
RUN apt-get update && \
    apt-get install -qy --no-install-recommends \
        curl \
        vim \
        xz-utils

COPY . ${APP_DIR}
WORKDIR ${APP_DIR}

RUN make deps && \
    GO111MODULE=off make daemon && \
    go install -tags ' mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1
