FROM golang:1.15-alpine

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

RUN apk update && \
    apk add --no-cache git make gcc sudo vim

RUN go get github.com/360EntSecGroup-Skylar/excelize/v2
RUN go get github.com/BurntSushi/toml
