# 云原生训练营模块八作业

## 1.第一部分

现在你对 Kubernetes 的控制面板的工作机制是否有了深入的了解呢？
是否对如何构建一个优雅的云上应用有了深刻的认识，那么接下来用最近学过的知识把你之前编写的 http 以优雅的方式部署起来吧，你可能需要审视之前代码是否能满足优雅上云的需求。
作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。

- 优雅启动
- 优雅终止
- 资源需求和 QoS 保证
- 探活
- 日常运维需求，日志等级
- 配置和代码分离

### 推送镜像

推送至阿里云

```shell
*****@***-mbp geektime-cloud-native-2 % docker tag sha256:c581a47a05942620656ae470213b7a0cabc6ebd445840afcce42f693d97671b9 registry.cn-beijing.aliyuncs.com/geekcncfcamp/http-server:0.0.1    
*****@***-mbp geektime-cloud-native-2 % docker push registry.cn-beijing.aliyuncs.com/geekcncfcamp/http-server:0.0.1                                                                           
The push refers to repository [registry.cn-beijing.aliyuncs.com/geekcncfcamp/http-server]
ff43a3ff622c: Layer already exists 
e265835b28ac: Layer already exists 
0.0.1: digest: sha256:b761ce057ae1cefacd34ed4e864f3a19d2d6b8f21b546434d4939b178f4e6a81 size: 740
```

为避免在生产集群上运行时出现网络问题，推送至备机备用

```shell
*****@***-mbp geektime-cloud-native-2 % docker tag sha256:c581a47a05942620656ae470213b7a0cabc6ebd445840afcce42f693d97671b9 registry.ap-southeast-1.aliyuncs.com/geekcncfcamp/http-server:0.0.1
*****@***-mbp geektime-cloud-native-2 % docker push registry.ap-southeast-1.aliyuncs.com/geekcncfcamp/http-server:0.0.1
The push refers to repository [registry.ap-southeast-1.aliyuncs.com/geekcncfcamp/http-server]
ff43a3ff622c: Pushed 
e265835b28ac: Pushed 
0.0.1: digest: sha256:b761ce057ae1cefacd34ed4e864f3a19d2d6b8f21b546434d4939b178f4e6a81 size: 740
```

### deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: httpserver
  namespace: default
  labels:
    app: httpserver
spec:
  replicas: 1
  selector:
    matchLabels:
      app: httpserver
  template:
    metadata: 
      labels:
        app: httpserver
    spec:
      containers:
        - name: httpserver   
          image: registry.cn-beijing.aliyuncs.com/geekcncfcamp/http-server:0.0.1
          ports:
            - containerPort: 8080
              protocol: TCP
```

### service
```yaml
apiVersion: v1
kind: Service
metadata:
  name: httpserver-svc
  namespace: default
  labels:
    app: httpserver
spec:
  type: ClusterIP
  ports:
    - port: 80
      targetPort: http
      protocol: TCP
      name: http
  selector:
    app: httpserver

```
### 日志等级，配置和代码分离
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-configs
  namespace: default
data:
  loglevel: debug
volumes:
# 可以在 Pod 级别设置卷，然后将其挂载到 Pod 内的容器中
  - name: config  
    configMap:    
      name: app-configs
```

## 2. 第二部分
除了将 httpServer 应用优雅的运行在 Kubernetes 之上，我们还应该考虑如何将服务发布给对内和对外的调用方。
来尝试用 Service, Ingress 将你的服务发布给集群外部的调用方吧。
在第一部分的基础上提供更加完备的部署 spec，包括（不限于）：

Service
Ingress
可以考虑的细节

如何确保整个应用的高可用。
如何通过证书保证 httpServer 的通讯安全。

### Ingress

```bash
#通过 helm 安装
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx ingress-nginx/ingress-nginx --create-namespace --namespace ingress

```

```yaml
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: httpserver
spec:
  ingressClassName: nginx
  rules:
    - host: sg.kronstadt.vip
      http:
        paths:
          - backend:
              service:
                name: httpserver-svc
                port:
                  number: 80
            path: /
            pathType: Prefix
```
