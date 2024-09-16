


## helm
``` go
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

``` go
sudo vim /etc/kubernetes/manifests/kube-apiserver.yaml

- --enable-aggregator-routing=true # 新增


```

``` go
https://github.com/kubernetes-sigs/metrics-server

kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml


https://artifacthub.io/packages/helm/metrics-server/metrics-server

https://github.com/kubernetes-sigs/metrics-server/releases/download/metrics-server-helm-chart-3.7.0/components.yaml
```




vim components.yaml

``` go

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




``` go
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









``` go
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


# 没安装 metrics-server临时设置
- --horizontal-pod-autoscaler-use-rest-clients=false

# 永久、安装 metrics-server


ubuntu@master:~$ kubectl get hpa
NAME       REFERENCE             TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
hpa-demo   Deployment/hpa-demo   0%/5%     1         10        1          64m

ubuntu@master:~$ kubectl describe hpa hpa-demo
Name:                                                  hpa-demo
Namespace:                                             default
Labels:                                                <none>
Annotations:                                           <none>
CreationTimestamp:                                     Mon, 13 Dec 2021 10:08:36 +0800
Reference:                                             Deployment/hpa-demo
Metrics:                                               ( current / target )
  resource cpu on pods  (as a percentage of request):  0% (0) / 5%
Min replicas:                                          1
Max replicas:                                          10
Deployment pods:                                       1 current / 1 desired
Conditions:
  Type            Status  Reason            Message
  ----            ------  ------            -------
  AbleToScale     True    ReadyForNewScale  recommended size matches current size
  ScalingActive   True    ValidMetricFound  the HPA was able to successfully calculate a replica count from cpu resource utilization (percentage of request)
  ScalingLimited  True    TooFewReplicas    the desired replica count is less than the minimum replica count
Events:           <none>




ubuntu@master:~$ kubectl run -i --tty test-hpa --image=busybox /bin/sh

/ # while true; do wget -q -O- http://10.244.1.6; done


ubuntu@master:~$ kubectl get hpa
NAME       REFERENCE             TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
hpa-demo   Deployment/hpa-demo   0%/5%     1         10        1          115m
ubuntu@master:~$ kubectl get hpa
NAME       REFERENCE             TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
hpa-demo   Deployment/hpa-demo   137%/5%   1         10        4          116m

# 删除 test-hpa后、自动缩容


ubuntu@master:~$ kubectl get hpa
NAME       REFERENCE             TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
hpa-demo   Deployment/hpa-demo   0%/5%     1         10        1          123m




ubuntu@master:~$ kubectl get hpa hpa-demo -o yaml
```

``` yaml

apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  annotations:
    autoscaling.alpha.kubernetes.io/conditions: '[{"type":"AbleToScale","status":"True","lastTransitionTime":"2021-12-13T02:08:52Z","reason":"ReadyForNewScale","message":"recommended
      size matches current size"},{"type":"ScalingActive","status":"True","lastTransitionTime":"2021-12-13T02:08:52Z","reason":"ValidMetricFound","message":"the
      HPA was able to successfully calculate a replica count from cpu resource utilization
      (percentage of request)"},{"type":"ScalingLimited","status":"True","lastTransitionTime":"2021-12-13T04:12:18Z","reason":"TooFewReplicas","message":"the
      desired replica count is less than the minimum replica count"}]'
    autoscaling.alpha.kubernetes.io/current-metrics: '[{"type":"Resource","resource":{"name":"cpu","currentAverageUtilization":0,"currentAverageValue":"0"}}]'
  creationTimestamp: "2021-12-13T02:08:36Z"
  name: hpa-demo
  namespace: default
  resourceVersion: "64973"
  uid: b16e5841-964e-47f2-9ee0-ee4716229ab5
spec:
  maxReplicas: 10
  minReplicas: 1
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: hpa-demo
  targetCPUUtilizationPercentage: 5
status:
  currentCPUUtilizationPercentage: 0
  currentReplicas: 1
  desiredReplicas: 1
  lastScaleTime: "2021-12-13T04:12:18Z"
```



``` yaml
ubuntu@master:~$ cat hpa-demo.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  name: hpa-demo
  labels:
    app: hpa
spec:
  revisionHistoryLimit: 15
  minReadySeconds: 5
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        resources:
          requests:
            cpu: 100m
          # limit:
          #   cpu: 200m
        ports:
        - containerPort: 80
---
apiVersion: autoscaling/v1
kind: HorizontalPodAutoscaler
metadata:
  name: hpa-demo
  namespace: default
spec:
  maxReplicas: 10
  minReplicas: 1
  scaleTargetRef:
    apiVersion: extensions/v1
    kind: Deployment
    name: hpa-demo
  targetCPUUtilizationPercentage: 5



ubuntu@master:~$ kubectl delete hpa hpa-demo
horizontalpodautoscaler.autoscaling "hpa-demo" deleted

ubuntu@master:~$ kubectl apply -f hpa-demo.yaml

Warning: resource deployments/hpa-demo is missing the kubectl.kubernetes.io/last-applied-configuration annotation which is required by kubectl apply. kubectl apply should only be used on resources created declaratively by either kubectl create --save-config or kubectl apply. The missing annotation will be patched automatically.
deployment.apps/hpa-demo configured
horizontalpodautoscaler.autoscaling/hpa-demo created


ubuntu@master:~$ kubectl get deployment
NAME       READY   UP-TO-DATE   AVAILABLE   AGE
hpa-demo   1/1     1            1           3h33m
ubuntu@master:~$ kubectl get hpa
NAME       REFERENCE             TARGETS        MINPODS   MAXPODS   REPLICAS   AGE
hpa-demo   Deployment/hpa-demo   <unknown>/5%   1         10        0          63s
```