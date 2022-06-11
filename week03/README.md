# 作业







# 实现

## 1、将代码传至服务器

```shell
root@deanUbuntu:~/week03homework# ls
builder  Dockerfile  myhttpserver.go
```

## 2、编写Dockerfile ,使用  scratch版本

```dockerfile
FROM golang:1.17 AS build
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=off
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver myhttpserver.go

FROM busybox COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 8360
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]
```

```dockerfile
FROM golang:1.17 AS builder


ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GORACH=amd64

WORKDIR /build
COPY . .
RUN go build -o httpserver .

FROM scratch
COPY --from=builder /builder/httpserver /
EXPOSE 8080
ENTRYPOINT ["/httpserver"]
```

目前报错，在跟老师请教中

## 3、运行

```shell
docker run -d httpserver:0.0.1
```

