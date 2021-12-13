


## helm
```go
删除默认的源



✅微软仓库
helm repo add stable http://mirror.azure.cn/kubernetes/charts/

国内镜像源
helm repo add stable https://burdenbear.github.io/kube-charts-mirror/
阿里云仓库
helm repo add stable https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts

helm repo update

helm env

查看helm源添加情况

helm repo list


helm install stable/redis --generate-name

```


## metrics-server

```go
sudo vim /etc/kubernetes/manifests/kube-apiserver.yaml

- --enable-aggregator-routing=true # 新增


```

```go
https://github.com/kubernetes-sigs/metrics-server

kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml


https://artifacthub.io/packages/helm/metrics-server/metrics-server

https://github.com/kubernetes-sigs/metrics-server/releases/download/metrics-server-helm-chart-3.7.0/components.yaml
```




vim components.yaml

```go

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: metrics-server
  name: metrics-server
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    k8s-app: metrics-server
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
    rbac.authorization.k8s.io/aggregate-to-edit: "true"
    rbac.authorization.k8s.io/aggregate-to-view: "true"
  name: system:aggregated-metrics-reader
rules:
- apiGroups:
  - metrics.k8s.io
  resources:
  - pods
  - nodes
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  labels:
    k8s-app: metrics-server
  name: system:metrics-server
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - nodes
  - nodes/stats
  - namespaces
  - configmaps
  verbs:
  - get
  - list
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    k8s-app: metrics-server
  name: metrics-server-auth-reader
  namespace: kube-system
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: extension-apiserver-authentication-reader
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    k8s-app: metrics-server
  name: metrics-server:system:auth-delegator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  labels:
    k8s-app: metrics-server
  name: system:metrics-server
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:metrics-server
subjects:
- kind: ServiceAccount
  name: metrics-server
  namespace: kube-system
---
apiVersion: v1
kind: Service
metadata:
  labels:
    k8s-app: metrics-server
  name: metrics-server
  namespace: kube-system
spec:
  ports:
  - name: https
    port: 443
    protocol: TCP
    targetPort: https
  selector:
    k8s-app: metrics-server
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    k8s-app: metrics-server
  name: metrics-server
  namespace: kube-system
spec:
  selector:
    matchLabels:
      k8s-app: metrics-server
  strategy:
    rollingUpdate:
      maxUnavailable: 0
  template:
    metadata:
      labels:
        k8s-app: metrics-server
    spec:
      containers:
      - args:
        - --cert-dir=/tmp
        - --secure-port=4443
        - --kubelet-preferred-address-types=InternalIP  # 删掉 ,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s
        - --kubelet-insecure-tls                    # 加上该启动参数
          # image: k8s.gcr.io/metrics-server/metrics-server:v0.5.2
        image: registry.aliyuncs.com/google_containers/metrics-server:v0.5.2  # 添加可访问镜像
        imagePullPolicy: IfNotPresent
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /livez
            port: https
            scheme: HTTPS
          periodSeconds: 10
        name: metrics-server
        ports:
        - containerPort: 4443
          name: https
          protocol: TCP
        readinessProbe:
          failureThreshold: 3
          httpGet:
            path: /readyz
            port: https
            scheme: HTTPS
          initialDelaySeconds: 20
          periodSeconds: 10
        resources:
          requests:
            cpu: 100m
            memory: 200Mi
        securityContext:
          readOnlyRootFilesystem: true
          runAsNonRoot: true
          runAsUser: 1000
        volumeMounts:
        - mountPath: /tmp
          name: tmp-dir
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: system-cluster-critical
      serviceAccountName: metrics-server
      volumes:
      - emptyDir: {}
        name: tmp-dir
---
apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  labels:
    k8s-app: metrics-server
  name: v1beta1.metrics.k8s.io
spec:
  group: metrics.k8s.io
  groupPriorityMinimum: 100
  insecureSkipTLSVerify: true
  service:
    name: metrics-server
    namespace: kube-system
  version: v1beta1
  versionPriority: 100
```




```go
kubectl apply -f components.yaml

root@master:/home/ubuntu# kubectl get pod -n kube-system | grep metrics-server
metrics-server-77bdbd79-bghfk    1/1     Running   0             4m47s


root@master:/home/ubuntu# kubectl describe svc metrics-server -n kube-system
Name:              metrics-server
Namespace:         kube-system
Labels:            k8s-app=metrics-server
Annotations:       <none>
Selector:          k8s-app=metrics-server
Type:              ClusterIP
IP Family Policy:  SingleStack
IP Families:       IPv4
IP:                10.111.246.16
IPs:               10.111.246.16
Port:              https  443/TCP
TargetPort:        https/TCP
Endpoints:         10.244.1.5:4443
Session Affinity:  None
Events:            <none>
root@master:/home/ubuntu# ping 10.244.1.5
PING 10.244.1.5 (10.244.1.5) 56(84) bytes of data.
64 bytes from 10.244.1.5: icmp_seq=1 ttl=63 time=6.94 ms
64 bytes from 10.244.1.5: icmp_seq=2 ttl=63 time=0.535 ms
64 bytes from 10.244.1.5: icmp_seq=3 ttl=63 time=0.653 ms

root@master:/home/ubuntu# kubectl top nodes
NAME     CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%
master   468m         23%    1570Mi          40%
node1    92m          4%     1136Mi          29%


root@master:/home/ubuntu# kubectl top pod --all-namespaces
NAMESPACE              NAME                                        CPU(cores)   MEMORY(bytes)
kube-system            coredns-78fcd69978-bpljs                    4m           11Mi
kube-system            coredns-78fcd69978-t4slw                    4m           12Mi
kube-system            etcd-master                                 51m          45Mi
kube-system            kube-apiserver-master                       163m         245Mi
kube-system            kube-controller-manager-master              69m          42Mi
kube-system            kube-flannel-ds-8vsms                       4m           10Mi
kube-system            kube-flannel-ds-rp96q                       6m           10Mi
kube-system            kube-proxy-975pb                            1m           10Mi
kube-system            kube-proxy-kv488                            1m           10Mi
kube-system            kube-scheduler-master                       10m          16Mi
kube-system            metrics-server-77bdbd79-bghfk               13m          15Mi
kubernetes-dashboard   dashboard-metrics-scraper-c45b7869d-6z6df   1m           4Mi
kubernetes-dashboard   kubernetes-dashboard-576cb95f94-7lv7g       40m          18Mi


ubuntu@master:~$ kubectl get pods --all-namespaces
NAMESPACE              NAME                                        READY   STATUS    RESTARTS      AGE
kube-system            coredns-78fcd69978-bpljs                    1/1     Running   0             6h57m
kube-system            coredns-78fcd69978-t4slw                    1/1     Running   0             6h57m
kube-system            etcd-master                                 1/1     Running   0             6h58m
kube-system            kube-apiserver-master                       1/1     Running   0             24m
kube-system            kube-controller-manager-master              1/1     Running   4 (24m ago)   5h31m
kube-system            kube-flannel-ds-8vsms                       1/1     Running   0             6h33m
kube-system            kube-flannel-ds-rp96q                       1/1     Running   0             6h33m
kube-system            kube-proxy-975pb                            1/1     Running   0             6h57m
kube-system            kube-proxy-kv488                            1/1     Running   0             6h45m
kube-system            kube-scheduler-master                       1/1     Running   5 (24m ago)   5h30m
kube-system            metrics-server-77bdbd79-bghfk               1/1     Running   0             13m
kubernetes-dashboard   dashboard-metrics-scraper-c45b7869d-6z6df   1/1     Running   0             5h33m
kubernetes-dashboard   kubernetes-dashboard-576cb95f94-7lv7g       1/1     Running   0             5h33m

```









```go
ubuntu@master:~$ vim hpa-demo.yaml
ubuntu@master:~$ kubectl create -f hpa-demo.yaml
deployment.apps/hpa-demo created

ubuntu@master:~$ kubectl get deploy
NAME       READY   UP-TO-DATE   AVAILABLE   AGE
hpa-demo   1/1     1            1           69s

ubuntu@master:~$ kubectl autoscale deployment hpa-demo --min=1 --max=10 --cpu-percent=5
horizontalpodautoscaler.autoscaling/hpa-demo autoscaled

ubuntu@master:~$ cd /etc/kubernetes/manifests
ubuntu@master:/etc/kubernetes/manifests$ ls
etcd.yaml  kube-apiserver.yaml  kube-controller-manager.yaml  kube-scheduler.yaml
ubuntu@master:/etc/kubernetes/manifests$ sudo vim kube-controller-manager.yaml

- --horizontal-pod-autoscaler-use-rest-clients=false





```