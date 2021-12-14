

```go
ubuntu@master:~$ kubectl apply -f deploy-demo.yaml
deployment.apps/deployment-demo created

ubuntu@master:~$ kubectl apply -f service.yaml
service/myservice created
ubuntu@master:~$ kubectl get svc
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP   3d13h
myservice    ClusterIP   10.108.125.155   <none>        80/TCP    11s

ubuntu@master:~$ kubectl describe svc myservice
Name:              myservice
Namespace:         default
Labels:            <none>
Annotations:       <none>
Selector:          app=nginx
Type:              ClusterIP
IP Family Policy:  SingleStack
IP Families:       IPv4
IP:                10.108.125.155
IPs:               10.108.125.155
Port:              mynginx-http  80/TCP
TargetPort:        80/TCP
Endpoints:         10.244.1.20:80,10.244.1.91:80,10.244.1.92:80 + 1 more...
Session Affinity:  None
Events:            <none>



ubuntu@master:~$ vim deploy-demo.yaml
ubuntu@master:~$ kubectl apply -f deploy-demo.yaml
deployment.apps/deployment-demo configured
ubuntu@master:~$ vim service.yaml
ubuntu@master:~$ kubectl apply -f service.yaml
service/myservice configured
ubuntu@master:~$ kubectl describe svc myservice
Name:              myservice
Namespace:         default
Labels:            <none>
Annotations:       <none>
Selector:          app=nginx
Type:              ClusterIP
IP Family Policy:  SingleStack
IP Families:       IPv4
IP:                10.108.125.155
IPs:               10.108.125.155
Port:              mynginx-http  80/TCP
TargetPort:        nginxweb/TCP
Endpoints:         10.244.1.94:80,10.244.1.95:80,10.244.1.96:80
Session Affinity:  None
Events:            <none>


ubuntu@master:~$ kubectl run -it testservice --image=busybox /bin/sh
If you don't see a command prompt, try pressing enter.

/ # wget http://10.108.125.155:80
Connecting to 10.108.125.155:80 (10.108.125.155:80)
saving to 'index.html'
index.html           100% |******************************************|   615  0:00:00 ETA
'index.html' saved


/ # cat index.html

<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>


/ # wget -O- -q http://10.108.125.155:80

<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>



ubuntu@master:~$ kubectl get svc
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP   3d13h
myservice    ClusterIP   10.108.125.155   <none>        80/TCP    34m
ubuntu@master:~$ vim service.yaml

# 添加
type: NodePort 

ubuntu@master:~$ kubectl apply -f service.yaml
service/myservice configured
ubuntu@master:~$ kubectl get svc
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP        3d13h
myservice    NodePort    10.108.125.155   <none>        80:30937/TCP   36m


ubuntu@master:~$ kubectl describe svc myservice
Name:                     myservice
Namespace:                default
Labels:                   <none>
Annotations:              <none>
Selector:                 app=nginx
Type:                     NodePort
IP Family Policy:         SingleStack
IP Families:              IPv4
IP:                       10.108.125.155
IPs:                      10.108.125.155
Port:                     mynginx-http  80/TCP
TargetPort:               nginxweb/TCP
NodePort:                 mynginx-http  30937/TCP
Endpoints:                10.244.1.94:80,10.244.1.95:80,10.244.1.96:80
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>


ubuntu@master:~$ netstat -ntlp
(Not all processes could be identified, non-owned process info
 will not be shown, you would have to be root to see it all.)
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      -
tcp        0      0 0.0.0.0:31447           0.0.0.0:*               LISTEN      -
tcp        0      0 0.0.0.0:30937           0.0.0.0:*               LISTEN      -


http://192.168.105.5:30937/


Welcome to nginx!

If you see this page, the nginx web server is successfully installed and working. Further configuration is required.

For online documentation and support please refer to nginx.org.
Commercial support is available at nginx.com.

Thank you for using nginx.


ubuntu@node1:~$ curl http://192.168.105.5:30937/

```