
```go
ubuntu@master:~$ vim configmap.yaml
ubuntu@master:~$ kubectl apply -f configmap.yaml
configmap/demo1 created
ubuntu@master:~$ kubectl get configmap
NAME               DATA   AGE
demo1              3      29s
kube-root-ca.crt   1      3d15h


ubuntu@master:~$ kubectl describe configmap demo1
Name:         demo1
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
config:
----
property.1=value-1
property.2=value-2
property.3=value-3

data.1:
----
hello
data.2:
----
world

BinaryData
====

Events:  <none>



ubuntu@master:~$ kubectl create configmap cm-demo2 --from-file=testcm
configmap/cm-demo2 created
ubuntu@master:~$ cat testcm/mysql.conf
host=127.0.0.1
port=3306
ubuntu@master:~$ cat testcm/redis.conf
host=127.0.0.1
port=6379


ubuntu@master:~$ kubectl get configmap
NAME               DATA   AGE
cm-demo2           2      88s
demo1              3      45m
kube-root-ca.crt   1      3d16h

ubuntu@master:~$ ls ~/testcm/
mysql.conf  redis.conf

ubuntu@master:~$ kubectl describe configmap cm-demo2
Name:         cm-demo2
Namespace:    default
Labels:       <none>
Annotations:  <none>

Data
====
mysql.conf:
----
host=127.0.0.1
port=3306

redis.conf:
----
host=127.0.0.1
port=6379


BinaryData
====

Events:  <none>



ubuntu@master:~$ kubectl get configmap cm-demo2 -o yaml
apiVersion: v1
data:
  mysql.conf: |
    host=127.0.0.1
    port=3306
  redis.conf: |
    host=127.0.0.1
    port=6379
kind: ConfigMap
metadata:
  creationTimestamp: "2021-12-15T00:58:02Z"
  name: cm-demo2
  namespace: default
  resourceVersion: "168981"
  uid: 97654e99-c8ee-4e92-9b32-cf4d71f288e4



ubuntu@master:~$ kubectl create configmap cm-demo3 --from-file=testcm/redis.conf
configmap/cm-demo3 created
ubuntu@master:~$ kubectl get configmap
NAME               DATA   AGE
cm-demo2           2      45m
cm-demo3           1      22s
demo1              3      89m
kube-root-ca.crt   1      3d17h
ubuntu@master:~$ kubectl get configmap cm-demo3 -o yaml
apiVersion: v1
data:
  redis.conf: |
    host=127.0.0.1
    port=6379
kind: ConfigMap
metadata:
  creationTimestamp: "2021-12-15T01:43:10Z"
  name: cm-demo3
  namespace: default
  resourceVersion: "171281"
  uid: 1893bfe5-a29e-41d3-8855-50ad790274ce


ubuntu@master:~$ kubectl create configmap cm-demo4 --from-literal=db.host=localhost --from-literal=db.port=3306
configmap/cm-demo4 created
ubuntu@master:~$ kubectl get configmap
NAME               DATA   AGE
cm-demo2           2      133m
cm-demo3           1      88m
cm-demo4           2      12s
demo1              3      176m
kube-root-ca.crt   1      3d18h
ubuntu@master:~$ kubectl get configmap cm-demo4 -o yaml
apiVersion: v1
data:
  db.host: localhost
  db.port: "3306"
kind: ConfigMap
metadata:
  creationTimestamp: "2021-12-15T03:10:59Z"
  name: cm-demo4
  namespace: default
  resourceVersion: "177956"
  uid: 610ff089-abba-425a-bd32-5c141969a1fa


ubuntu@master:~$ vim cmtest1-pod.yaml
ubuntu@master:~$ kubectl create -f cmtest1-pod.yaml
pod/testcm1-pod created
ubuntu@master:~$ kubectl get pods
NAME                               READY   STATUS             RESTARTS      AGE
deployment-demo-58684549dc-84ksm   1/1     Running            1 (52m ago)   15h
deployment-demo-58684549dc-mkcdr   1/1     Running            1 (52m ago)   15h
deployment-demo-58684549dc-zqkh9   1/1     Running            1 (52m ago)   15h
hpa-demo-c9ddb6864-tldck           1/1     Running            1 (52m ago)   2d7h
job-demo--1-6lgx5                  0/1     Completed          0             2d5h
testcm1-pod                        0/1     CrashLoopBackOff   2 (18s ago)   85s
testservice                        1/1     Running            2 (52m ago)   15h


ubuntu@master:~$ kubectl logs testcm1-pod
MYSERVICE_SERVICE_HOST=10.108.125.155
KUBERNETES_SERVICE_PORT=443
KUBERNETES_PORT=tcp://10.96.0.1:443
HOSTNAME=testcm1-pod
DB_PORT=3306
SHLVL=1
HOME=/root
MYSERVICE_SERVICE_PORT=80
MYSERVICE_PORT=tcp://10.108.125.155:80
mysql.conf=host=127.0.0.1
port=3306

MYSERVICE_PORT_80_TCP_ADDR=10.108.125.155
MYSERVICE_PORT_80_TCP_PORT=80
MYSERVICE_PORT_80_TCP_PROTO=tcp
MYSERVICE_SERVICE_PORT_MYNGINX_HTTP=80
redis.conf=host=127.0.0.1
port=6379

KUBERNETES_PORT_443_TCP_ADDR=10.96.0.1
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
KUBERNETES_PORT_443_TCP_PORT=443
KUBERNETES_PORT_443_TCP_PROTO=tcp
MYSERVICE_PORT_80_TCP=tcp://10.108.125.155:80
KUBERNETES_PORT_443_TCP=tcp://10.96.0.1:443
KUBERNETES_SERVICE_PORT_HTTPS=443
KUBERNETES_SERVICE_HOST=10.96.0.1
PWD=/
DB_HOST=localhost





ubuntu@master:~$ vim cmtest2-pod.yaml
ubuntu@master:~$ kubectl apply -f cmtest2-pod.yaml
pod/testcm2-pod created
ubuntu@master:~$ kubectl get pods
NAME                               READY   STATUS              RESTARTS       AGE
deployment-demo-58684549dc-84ksm   1/1     Running             1 (158m ago)   17h
deployment-demo-58684549dc-mkcdr   1/1     Running             1 (158m ago)   17h
deployment-demo-58684549dc-zqkh9   1/1     Running             1 (158m ago)   17h
hpa-demo-c9ddb6864-tldck           1/1     Running             1 (158m ago)   2d9h
job-demo--1-6lgx5                  0/1     Completed           0              2d6h
testcm1-pod                        0/1     CrashLoopBackOff    7 (2m ago)     15m
testcm2-pod                        0/1     ContainerCreating   0              13s
testservice                        1/1     Running             2 (158m ago)   17h


ubuntu@master:~$ kubectl logs testcm2-pod
localhost 3306





```