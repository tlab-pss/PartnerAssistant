FROM golang:1.13.0-alpine

WORKDIR /go/src

ENV GO111MODULE=on

RUN apk add --no-cache \
  alpine-sdk \
  git \
  && go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
