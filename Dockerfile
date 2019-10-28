# For development ====================
FROM golang:1.13.0-alpine as develop
WORKDIR /go/src

ENV GO111MODULE=on
RUN apk --update add --no-cache git \
  alpine-sdk \
  && go get github.com/pilu/fresh

EXPOSE 8080
CMD ["fresh"]
# ====================================

# For production =====================
FROM golang:1.13.0-alpine as build
ADD . /go/src
WORKDIR /go/src

ENV GO111MODULE=on
RUN apk --update add --no-cache git \
  mercurial
RUN go build -o PartnerAssistant main.go

FROM alpine:3.9 as release
WORKDIR /apps
COPY --from=build /go/src/PartnerAssistant /usr/local/bin/PartnerAssistant
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/PartnerAssistant"]
# ====================================
