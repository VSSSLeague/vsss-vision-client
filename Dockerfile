FROM golang:1.14-alpine AS build_go
WORKDIR /go/src/github.com/VSSSLeague/vsss-vision-client
COPY . .
RUN go get -v -t -d ./...
RUN go get -v github.com/gobuffalo/packr/packr
WORKDIR cmd/
RUN GOOS=linux GOARCH=amd64 packr build -o ../../release/vsss-vision-client_linux_amd64
