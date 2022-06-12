# 作业

- 构建本地镜像
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
- 将镜像推送至 docker 官方镜像仓库
- 通过 docker 命令本地启动 httpserver
- 通过 nsenter 进入容器查看 IP 配置





# 实现

## 1、构建本地镜像

```shell
sudo apt install docker
```



## 2、将代码传至服务器

```shell
root@deanUbuntu:~/week03homework# ls
builder  Dockerfile  myhttpserver.go
```

## 3、编写Dockerfile 

记录了2个版本（busybox、scratch），实际使用  scratch版本,隐含要求， 多阶段构建

第一个From编译，第二个打包，中间会产生dangling images,可删除

```dockerfile
FROM golang:1.17 AS builder


ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GORACH=amd64

WORKDIR /builder
COPY . .
RUN go build -o httpserver .

FROM scratch
COPY --from=builder /builder/httpserver /
EXPOSE 8080
ENTRYPOINT ["/httpserver"]
```



```dockerfile
FROM golang:1.17 AS build
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=off
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver myhttpserver.go

FROM busybox COPY --from=build /httpserver/httpserver /httpserver/httpserver ##此处报错提示需要one or three arguements，没有跟老师请教，要是老师看到，可以指导一下
EXPOSE 8360
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]
//CMD ["./httpserver""]
```

注意使用CMD 启动，有可能会造成僵尸进程





## 4、编译

```shell
docker build -t httpserver:0.1 -f Dockerfile .
```

目前报错，在跟老师请教中

```shell
Step 7/9 : COPY --from=builder /builder/httpserver /
COPY failed: stat builder/httpserver: file does not exis
```

原因是工作目录和copy from命令中的目录不一致

```sh
Successfully built ba62ebacf4e8
Successfully tagged httpserver:0.1
root@deanUbuntu:~/week03homework# docker images
```

顺便删除一下dangling images

```
docker image prune
```

```
root@deanUbuntu:~/week03homework# docker images
REPOSITORY   TAG       IMAGE ID       CREATED              SIZE
httpserver   0.1       03414a8001d2   About a minute ago   6.11MB
<none>       <none>    46c4ed7fa315   About a minute ago   962MB
golang       1.17      0e970dee9aa4   10 days ago          941MB
root@deanUbuntu:~/week03homework# docker image prune
WARNING! This will remove all dangling images.
Are you sure you want to continue? [y/N] y
Deleted Images:
deleted: sha256:46c4ed7fa3154f70dcd184ad2c3b61a89ccf56886a6c14f4fa6551de535031ba
deleted: sha256:ab46ed41d06aa0974f79afeec2c4f85136abfa3cda89381eb21c50460c686b6c
deleted: sha256:b1dfc1532eb3d6a6379aa71ccef17e75908ce3e9a4dbc165efd8126776f5aa91
deleted: sha256:7e937e35275c80e94800fc39bbefb7fdd5c62ed334caff1a2c7ce094e3b8b4b5
deleted: sha256:cd3b7bc6ef3a32e06575562537e1ea22cb2aab8135fb3673749fd6f941ff6792
deleted: sha256:3661be74d370639400d6abe0eb0dc19da2599cfab52b7ce70f1a1ee34e89a364
deleted: sha256:50e8e4e2c8fb95158c797896207cfd70bfa6f36cec2d3b69eeab223be0763fbe

Total reclaimed space: 21.04MB
root@deanUbuntu:~/week03homework# docker images
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
httpserver   0.1       03414a8001d2   2 minutes ago   6.11MB
golang       1.17      0e970dee9aa4   10 days ago     941MB
root@deanUbuntu:~/week03homework# 
```



## 5、运行

```shell
docker run -d httpserver:0.1
```



```shell
root@deanUbuntu:~/week03homework# docker run -d httpserver:0.1
59624f80abb6b011049f291b6a12bb955ef8f0aecfefd31e98eec5414f2c0587
root@deanUbuntu:~/week03homework# docker ps
CONTAINER ID   IMAGE            COMMAND         CREATED          STATUS          PORTS      NAMES
59624f80abb6   httpserver:0.1   "/httpserver"   14 seconds ago   Up 14 seconds   8080/tcp   stupefied_noether
```



## 6、镜像推送

未处理



## 7、进入容器

通过 nsenter 进入容器查看 IP 配置

先查看docker信息

```
root@deanUbuntu:~/week03homework# docker ps
CONTAINER ID   IMAGE            COMMAND         CREATED          STATUS          PORTS      NAMES
59624f80abb6   httpserver:0.1   "/httpserver"   14 seconds ago   Up 14 seconds   8080/tcp   stupefied_noether
```

根据docker NAMES获取PID

```
root@deanUbuntu:~/week03homework# PID=$(docker inspect --format "{{ .State.Pid }}" stupefied_noether)
root@deanUbuntu:~/week03homework# echo $PID
61987
root@deanUbuntu:~/week03homework# nsenter -t 61987 -n ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
19: eth0@if20: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```

