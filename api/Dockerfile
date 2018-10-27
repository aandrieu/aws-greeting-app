FROM golang:1.9

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/webserver
WORKDIR /go/src/webserver

COPY ./src/webserver/Gopkg.lock ./Gopkg.lock
COPY ./src/webserver/Gopkg.toml ./Gopkg.toml
RUN dep ensure -vendor-only

COPY . /go