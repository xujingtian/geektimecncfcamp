# 手动创建证书
```shell
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=haylin Inc./CN=*.haylin.test' -keyout haylin.test.key -out haylin.test.crt
kubectl create -n istio-system secret tls haylin-credential --key=haylin.test.key --cert=haylin.test.crt
```

# 创建namespace,给namespace加上label
```shell
kubectl apply -f httpserver-ns.yaml
kubectl label ns httpserver istio-injection=enabled
```

# 部署httpserver
```shell
kubectl apply -f httpserver-configmap.yaml
kubectl apply -f httpserver-deployment.yaml
kubectl apply -f httpserver-svc.yaml
```

# Http
配置Gateway/VirtualService
```shell
kubectl apply -f istio-gw-http.yaml
kubectl apply -f istio-vs-http.yaml
```
验证效果
```shell
[root@i-5v9r2v9h 4.sidecar]# curl -H "Host: httpserver.haylin.test" 10.98.60.64/healthz -v
*   Trying 10.98.60.64...
* TCP_NODELAY set
* Connected to 10.98.60.64 (10.98.60.64) port 80 (#0)
> GET /healthz HTTP/1.1
> Host: httpserver.haylin.test
> User-Agent: curl/7.61.1
> Accept: */*
>
< HTTP/1.1 200 OK
< date: Sun, 10 Jul 2022 14:51:55 GMT
< content-length: 3
< content-type: text/plain; charset=utf-8
< x-envoy-upstream-service-time: 153
< server: istio-envoy
<
* Connection #0 to host 10.98.60.64 left intact

```

# Https
配置Gateway/VirtualService
```shell
kubectl apply -f istio-gw-https.yaml
kubectl apply -f istio-vs-https.yaml
```
验证效果
```shell
[root@i-5v9r2v9h ~]# curl --resolve httpsserver.haylin.test:443:10.98.60.64 https://httpsserver.haylin.test/healthz -v -k
* Added httpsserver.haylin.test:443:10.98.60.64 to DNS cache
* Hostname httpsserver.haylin.test was found in DNS cache
*   Trying 10.98.60.64...
* TCP_NODELAY set
* Connected to httpsserver.haylin.test (10.98.60.64) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/pki/tls/certs/ca-bundle.crt
  CApath: none
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, [no content] (0):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, [no content] (0):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_256_GCM_SHA384
* ALPN, server accepted to use h2
* Server certificate:
*  subject: O=haylin Inc.; CN=*.haylin.test
*  start date: Jul 10 14:26:59 2022 GMT
*  expire date: Jul 10 14:26:59 2023 GMT
*  issuer: O=haylin Inc.; CN=*.haylin.test
*  SSL certificate verify result: self signed certificate (18), continuing anyway.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* TLSv1.3 (OUT), TLS app data, [no content] (0):
* TLSv1.3 (OUT), TLS app data, [no content] (0):
* TLSv1.3 (OUT), TLS app data, [no content] (0):
* Using Stream ID: 1 (easy handle 0x5601b3694460)
* TLSv1.3 (OUT), TLS app data, [no content] (0):
> GET /healthz HTTP/2
> Host: httpsserver.haylin.test
> User-Agent: curl/7.61.1
> Accept: */*
>
* TLSv1.3 (IN), TLS handshake, [no content] (0):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS app data, [no content] (0):
* Connection state changed (MAX_CONCURRENT_STREAMS == 2147483647)!
* TLSv1.3 (OUT), TLS app data, [no content] (0):
< HTTP/2 503
< content-length: 19
< content-type: text/plain
< date: Sun, 10 Jul 2022 15:20:14 GMT
< server: istio-envoy
<
* Connection #0 to host httpsserver.haylin.test left intact

```

# Tracing
```shell
istioctl dashboard jaeger --address 0.0.0.0
```
截图见screenshots文件夹