FROM golang:1.17-stretch as builder
WORKDIR /httpserver
COPY http_server.go .
COPY go.mod .
COPY go.sum .
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o httpserver

FROM debian:stretch as prod
WORKDIR /root
COPY --from=0 /httpserver/httpserver .
CMD ["./httpserver"]

