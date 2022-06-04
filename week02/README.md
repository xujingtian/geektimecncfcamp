## 题目

编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

- 接收客户端 request，并将 request 中带的 header 写入 response header
- 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
- Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
- 当访问 localhost/healthz 时，应返回 200

## 代码

- 代码实现：week02\myhttpserver.go
- 路由 /header 写入接收到的request中的请求
- go run main.go 路由 /herder 会写入VERSION 到 response header

> ![header](https://github.com/xujingtian/geektimecncfcamp/blob/main/week02/README.assets/image-20220604233540833.png)
>
> ![image-20220604233540833](D:\99.dean_pc\08.geektime\16.CNCF\03.code\src\github.com\geektimecncfcamp\week02\README.assets\image-20220604233540833.png)
>
> 

- Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出

> ![server](https://github.com/xujingtian/geektimecncfcamp/blob/main/week02/README.assets/image-20220604233637271.png)
>
> ![image-20220604233637271](D:\99.dean_pc\08.geektime\16.CNCF\03.code\src\github.com\geektimecncfcamp\week02\README.assets\image-20220604233637271.png)

- 当访问 localhost/healthz 时，应返回 200

> ![http200](https://github.com/xujingtian/geektimecncfcamp/blob/main/week02/README.assets/image-20220604233826725.png)
>
> ![image-20220604233826725](D:\99.dean_pc\08.geektime\16.CNCF\03.code\src\github.com\geektimecncfcamp\week02\README.assets\image-20220604233826725.png)