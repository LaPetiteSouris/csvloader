FROM golang:1.13-alpine as build-base
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make
ADD . /go/src/github.com/LaPetiteSouris/csvloader
WORKDIR /go/src/github.com/LaPetiteSouris/csvloader
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOARCH amd64
ENV GOOS linux
RUN go build -mod=readonly -o csvloader

FROM alpine:latest
RUN apk add --no-cache --upgrade bash
COPY --from=build-base /go/src/github.com/LaPetiteSouris/csvloader /usr/local/bin/csvloader

CMD ["sleep", "100d"]
