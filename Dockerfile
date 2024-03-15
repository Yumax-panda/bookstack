FROM golang:1.22.1-alpine AS build

RUN mkdir /storage

WORKDIR /go/src/github.com/Yumax-panda/bookstack

COPY . .

RUN apk upgrade --update && apk --no-cache add git

