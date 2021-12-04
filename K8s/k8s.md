

## Ubuntu 安装 k8s


准备虚拟机：
```go
multipass launch -n master -c 2 -m 4G -d 40G
multipass shell master

multipass launch -n node1 -c 2 -m 4G -d 40G
multipass shell node1

```



```go


# 在master添加hosts
sudo vim /etc/hosts

192.168.105.5 master
192.168.105.6 node1


# 关闭防火墙
systemctl stop firewalld
systemctl disable firewalld

# 关闭防火墙
ubuntu@master:~$ sudo apt-get install ufw

ubuntu@master:~$ sudo ufw disable
Firewall stopped and disabled on system startup
ubuntu@master:~$ sudo ufw status
Status: inactive


# 关闭selinux

setenforce 0

cat /etc/selinux/config

# 关闭selinux

ubuntu@master:~$ sudo vim /etc/selinux/config

SELINUX=disabled


# 关闭swap
ubuntu@master:~$ sudo sed -ri 's/.*swap.*/#&/' /etc/fstab


# 开启路由转发

vim /etc/sysctl.d/k8s.conf


net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1


# 生效
sysctl -p /etc/sysctl.d/k8s.conf

```






## 允许 iptables 检查桥接流量

确保 br_netfilter 模块被加载。这一操作可以通过运行 lsmod | grep br_netfilter 来完成。若要显式加载该模块，可执行 sudo modprobe br_netfilter。

为了让你的 Linux 节点上的 iptables 能够正确地查看桥接流量，你需要确保在你的 sysctl 配置中将 net.bridge.bridge-nf-call-iptables 设置为 1。例如：
```go


cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
sudo sysctl --system
```


1. 更新 apt 包索引并安装使用 Kubernetes apt 仓库所需要的包：
```go
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl
```

2. 下载 Google Cloud 公开签名秘钥：
```go
sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg
```

3. 添加 Kubernetes apt 仓库：
```go
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
```


4. 更新 apt 包索引，安装 kubelet、kubeadm 和 kubectl，并锁定其版本：
```go

sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl

```






```go


ubuntu@master:~$ docker info

Cgroup Driver: cgroupfs

ubuntu@master:~$ sudo vim /etc/docker/daemon.json

{ "registry-mirrors": ["https://hkaofvr0.mirror.aliyuncs.com"],
        "exec-opts": ["native.cgroupdriver=systemd"]
}

systemctl restart docker && systemctl status docker



vim /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

Environment="KUBELET_EXTAR_ARGS=--fail-swap-on=false"



systemctl daemon-reload

重新启动kubelet.service
		1.systemctl daemon-reload
		2.systemctl restart kubelet.service



ubuntu@master:~$ kubeadm config images list
k8s.gcr.io/kube-apiserver:v1.22.4
k8s.gcr.io/kube-controller-manager:v1.22.4
k8s.gcr.io/kube-scheduler:v1.22.4
k8s.gcr.io/kube-proxy:v1.22.4
k8s.gcr.io/pause:3.5
k8s.gcr.io/etcd:3.5.0-0
k8s.gcr.io/coredns/coredns:v1.8.4


node1:



kubeadm config images pull --image-repository=registry.aliyuncs.com/google_containers

ubuntu@master:~$ kubeadm config images pull --image-repository=registry.aliyuncs.com/google_containers
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-apiserver:v1.22.4
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-controller-manager:v1.22.4
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-scheduler:v1.22.4
[config/images] Pulled registry.aliyuncs.com/google_containers/kube-proxy:v1.22.4
[config/images] Pulled registry.aliyuncs.com/google_containers/pause:3.5
[config/images] Pulled registry.aliyuncs.com/google_containers/etcd:3.5.0-0
[config/images] Pulled registry.aliyuncs.com/google_containers/coredns:v1.8.4



ubuntu@master:~$ docker images
REPOSITORY                                                        TAG       IMAGE ID       CREATED        SIZE
registry.aliyuncs.com/google_containers/kube-apiserver            v1.22.4   8a5cc299272d   13 days ago    128MB
registry.aliyuncs.com/google_containers/kube-controller-manager   v1.22.4   0ce02f92d3e4   13 days ago    122MB
registry.aliyuncs.com/google_containers/kube-scheduler            v1.22.4   721ba97f54a6   13 days ago    52.7MB
registry.aliyuncs.com/google_containers/kube-proxy                v1.22.4   edeff87e4802   13 days ago    104MB
registry.aliyuncs.com/google_containers/etcd                      3.5.0-0   004811815584   5 months ago   295MB
registry.aliyuncs.com/google_containers/coredns                   v1.8.4    8d147537fb7d   6 months ago   47.6MB
registry.aliyuncs.com/google_containers/pause                     3.5       ed210e3e4a5b   8 months ago   683kB






images=(  # 下面的镜像应该去除"k8s.gcr.io/"的前缀，版本换成上面获取到的版本
    kube-apiserver:v1.22.4
    kube-controller-manager:v1.22.4
    kube-scheduler:v1.22.4
    kube-proxy:v1.22.4
    pause:3.5
    etcd:3.5.0-0
    coredns:v1.8.4
)


for imageName in ${images[@]} ; do
    docker tag registry.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
    docker rmi registry.aliyuncs.com/google_containers/$imageName
done

docker tag k8s.gcr.io/coredns:v1.8.4 k8s.gcr.io/coredns/coredns:v1.8.4
docker rmi k8s.gcr.io/coredns:v1.8.4






ubuntu@master:~$ docker images
REPOSITORY                           TAG       IMAGE ID       CREATED        SIZE
k8s.gcr.io/kube-apiserver            v1.22.4   8a5cc299272d   13 days ago    128MB
k8s.gcr.io/kube-controller-manager   v1.22.4   0ce02f92d3e4   13 days ago    122MB
k8s.gcr.io/kube-scheduler            v1.22.4   721ba97f54a6   13 days ago    52.7MB
k8s.gcr.io/kube-proxy                v1.22.4   edeff87e4802   13 days ago    104MB
k8s.gcr.io/etcd                      3.5.0-0   004811815584   5 months ago   295MB
k8s.gcr.io/coredns/coredns           v1.8.4    8d147537fb7d   6 months ago   47.6MB
k8s.gcr.io/pause                     3.5       ed210e3e4a5b   8 months ago   683kB






sudo kubeadm init --kubernetes-version=v1.22.4 --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=192.168.105.5

或：
sudo kubeadm init --kubernetes-version=v1.22.4 --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=192.168.105.5 --image-repository registry.aliyuncs.com/google_containers



Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.105.5:6443 --token 8dtoxq.mc8y47svyn1qbs3o \
	--discovery-token-ca-cert-hash sha256:276f67ebefc068beeb31005935889d6874a36cc26c12fd3663d6d2aea1c15e0d



echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> /etc/profile

source /etc/profile




For Kubernetes v1.17+ 

https://github.com/flannel-io/flannel


kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

kubectl apply -f kube-flannel.yml


root@master:/home/ubuntu# kubectl apply -f kube-flannel.yml
Warning: policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
podsecuritypolicy.policy/psp.flannel.unprivileged created
clusterrole.rbac.authorization.k8s.io/flannel created
clusterrolebinding.rbac.authorization.k8s.io/flannel created
serviceaccount/flannel created
configmap/kube-flannel-cfg created
daemonset.apps/kube-flannel-ds created

kubectl get pods --all-namespaces



下载到本地
https://github.com/flannel-io/flannel/releases/download/v0.15.1/flanneld-v0.15.1-amd64.docker


docker load < flanneld-v0.15.1-amd64.docker

root@master:~# docker images
REPOSITORY                                       TAG             IMAGE ID       CREATED        SIZE
k8s.gcr.io/kube-apiserver                        v1.22.4         8a5cc299272d   2 weeks ago    128MB
k8s.gcr.io/kube-scheduler                        v1.22.4         721ba97f54a6   2 weeks ago    52.7MB
k8s.gcr.io/kube-controller-manager               v1.22.4         0ce02f92d3e4   2 weeks ago    122MB
k8s.gcr.io/kube-proxy                            v1.22.4         edeff87e4802   2 weeks ago    104MB
quay.io/coreos/flannel                           v0.15.1         e6ea68648f0c   2 weeks ago    69.5MB
quay.io/coreos/flannel                           v0.15.1-amd64   e6ea68648f0c   2 weeks ago    69.5MB
rancher/mirrored-flannelcni-flannel-cni-plugin   v1.0.0          cd5235cd7dc2   4 weeks ago    9.03MB
k8s.gcr.io/etcd                                  3.5.0-0         004811815584   5 months ago   295MB
k8s.gcr.io/coredns/coredns                       v1.8.4          8d147537fb7d   6 months ago   47.6MB
k8s.gcr.io/pause                                 3.5             ed210e3e4a5b   8 months ago   683kB
root@master:~# kubectl get pods --all-namespaces
NAMESPACE     NAME                             READY   STATUS    RESTARTS      AGE
kube-system   coredns-78fcd69978-hqq4v         1/1     Running   0             97m
kube-system   coredns-78fcd69978-qm6qq         1/1     Running   0             97m
kube-system   etcd-master                      1/1     Running   0             98m
kube-system   kube-apiserver-master            1/1     Running   0             98m
kube-system   kube-controller-manager-master   1/1     Running   0             98m
kube-system   kube-flannel-ds-kkqf4            1/1     Running   0             26m
kube-system   kube-proxy-h5p5z                 1/1     Running   0             97m
kube-system   kube-scheduler-master            1/1     Running   1 (86m ago)   98m

root@master:~# kubectl get nodes
NAME     STATUS   ROLES                  AGE    VERSION
master   Ready    control-plane,master   102m   v1.22.4


```

## flannel
```go
https://github.com/flannel-io/flannel


chrome下载到本地：
https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml

vim kube-flannel.yml

kubectl apply -f kube-flannel.yml
```





## node1 Kubernetes Dashboard
```go
kubectl apply -f recommended.yaml

vim recommended.yaml
chrome下载到本地：
https://raw.githubusercontent.com/kubernetes/dashboard/v2.4.0/aio/deploy/recommended.yaml

```

## ## node1 加入集群
```go

kubeadm join 192.168.105.5:6443 --token 8dtoxq.mc8y47svyn1qbs3o \
	--discovery-token-ca-cert-hash sha256:276f67ebefc068beeb31005935889d6874a36cc26c12fd3663d6d2aea1c15e0d

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the control-plane to see this node join the cluster.
```



## scheduler   Unhealthy 

```go
root@master:~# kubectl get cs
Warning: v1 ComponentStatus is deprecated in v1.19+
NAME                 STATUS      MESSAGE                                                                                       ERROR
scheduler            Unhealthy   Get "http://127.0.0.1:10251/healthz": dial tcp 127.0.0.1:10251: connect: connection refused
controller-manager   Healthy     ok
etcd-0               Healthy     {"health":"true","reason":""}



vim /etc/kubernetes/manifests/kube-controller-manager.yaml

vim /etc/kubernetes/manifests/kube-scheduler.yaml

注释掉
# - --port=0


root@master:~#  kubectl get cs
Warning: v1 ComponentStatus is deprecated in v1.19+
NAME                 STATUS    MESSAGE                         ERROR
scheduler            Healthy   ok
controller-manager   Healthy   ok
etcd-0               Healthy   {"health":"true","reason":""}
```



```go
root@master:~# kubectl get pods --all-namespaces -o wide
NAMESPACE              NAME                                        READY   STATUS    RESTARTS        AGE     IP              NODE     NOMINATED NODE   READINESS GATES
kube-system            coredns-78fcd69978-hqq4v                    1/1     Running   0               6h37m   10.244.0.3      master   <none>           <none>
kube-system            coredns-78fcd69978-qm6qq                    1/1     Running   0               6h37m   10.244.0.2      master   <none>           <none>
kube-system            etcd-master                                 1/1     Running   0               6h37m   192.168.105.5   master   <none>           <none>
kube-system            kube-apiserver-master                       1/1     Running   0               6h37m   192.168.105.5   master   <none>           <none>
kube-system            kube-controller-manager-master              1/1     Running   0               7m4s    192.168.105.5   master   <none>           <none>
kube-system            kube-flannel-ds-chrqd                       1/1     Running   0               3h24m   192.168.105.6   node1    <none>           <none>
kube-system            kube-flannel-ds-kkqf4                       1/1     Running   0               5h25m   192.168.105.5   master   <none>           <none>
kube-system            kube-proxy-frgrw                            1/1     Running   0               3h24m   192.168.105.6   node1    <none>           <none>
kube-system            kube-proxy-h5p5z                            1/1     Running   0               6h37m   192.168.105.5   master   <none>           <none>
kube-system            kube-scheduler-master                       1/1     Running   1 (6m29s ago)   6m38s   192.168.105.5   master   <none>           <none>
kubernetes-dashboard   dashboard-metrics-scraper-c45b7869d-xhr24   1/1     Running   0               3h6m    10.244.1.2      node1    <none>           <none>
kubernetes-dashboard   kubernetes-dashboard-576cb95f94-kdts4       1/1     Running   0               3h6m    10.244.1.3      node1    <none>           <none>

```


## 错误❌解决：

```go

[wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
[kubelet-check] Initial timeout of 40s passed.
[kubelet-check] It seems like the kubelet isn't running or healthy.
[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10248/healthz' failed with error: Get "http://localhost:10248/healthz": dial tcp [::1]:10248: connect: connection refused.


	Unfortunately, an error has occurred:
		timed out waiting for the condition

	This error is likely caused by:
		- The kubelet is not running
		- The kubelet is unhealthy due to a misconfiguration of the node in some way (required cgroups disabled)

	If you are on a systemd-powered system, you can try to troubleshoot the error with the following commands:
		- 'systemctl status kubelet'
		- 'journalctl -xeu kubelet'

	Additionally, a control plane component may have crashed or exited when started by the container runtime.
	To troubleshoot, list all containers using your preferred container runtimes CLI.



Failed to run kubelet" err="failed to run Kubelet: misconfiguration: kubelet cgroup driver: \"systemd\" is different from docker cgroup driver: \"cgroupfs\""


sudo vim /etc/docker/daemon.json

{
  "exec-opts": ["native.cgroupdriver=systemd"]
}

systemctl restart docker
systemctl status docker


sudo kubeadm reset

sudo rm -rf /etc/kubernetes/manifests/kube-apiserver.yaml
sudo rm -rf /etc/kubernetes/manifests/kube-controller-manager.yaml
sudo rm -rf /etc/kubernetes/manifests/kube-scheduler.yaml
sudo rm -rf /etc/kubernetes/manifests/etcd.yaml
sudo rm -rf /var/lib/etcd

sudo rm -rf $HOME/.kube/config

sudo kubeadm init --kubernetes-version=v1.22.4 --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=192.168.105.5



```







## 在Docker中下载并运行Jenkins （在macOS和Linux上）
```go

docker run \
  -u root \
  --rm \
  -d \
  -p 8080:8080 \
  -p 50000:50000 \
  -v jenkins-data:/var/jenkins_home \
  -v /var/run/docker.sock:/var/run/docker.sock \
  jenkinsci/blueocean
```







```go
root@master:~# kubectl get nodes
NAME     STATUS   ROLES                  AGE     VERSION
master   Ready    control-plane,master   6h54m   v1.22.4
node1    Ready    <none>                 3h41m   v1.22.4

root@master:~# kubectl get pods --all-namespaces
NAMESPACE              NAME                                        READY   STATUS    RESTARTS      AGE
kube-system            coredns-78fcd69978-hqq4v                    1/1     Running   0             6h55m
kube-system            coredns-78fcd69978-qm6qq                    1/1     Running   0             6h55m
kube-system            etcd-master                                 1/1     Running   0             6h55m
kube-system            kube-apiserver-master                       1/1     Running   0             6h55m
kube-system            kube-controller-manager-master              1/1     Running   0             25m
kube-system            kube-flannel-ds-chrqd                       1/1     Running   0             3h42m
kube-system            kube-flannel-ds-kkqf4                       1/1     Running   0             5h43m
kube-system            kube-proxy-frgrw                            1/1     Running   0             3h42m
kube-system            kube-proxy-h5p5z                            1/1     Running   0             6h55m
kube-system            kube-scheduler-master                       1/1     Running   1 (24m ago)   24m
kubernetes-dashboard   dashboard-metrics-scraper-c45b7869d-xhr24   1/1     Running   0             3h24m
kubernetes-dashboard   kubernetes-dashboard-576cb95f94-kdts4       1/1     Running   0             3h24m
```



## kubernetes-dashboard



vim recommended.yaml  添加：  type: NodePort
```go

kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  type: NodePort
  ports:
    - port: 443
      targetPort: 8443
  selector:
    k8s-app: kubernetes-dashboard


```
kubectl apply -f recommended.yaml


```go
root@node1:/home/ubuntu/share# kubectl get pods --all-namespaces
NAMESPACE              NAME                                        READY   STATUS    RESTARTS      AGE
kube-system            coredns-78fcd69978-hqq4v                    1/1     Running   0             7h12m
kube-system            coredns-78fcd69978-qm6qq                    1/1     Running   0             7h12m
kube-system            etcd-master                                 1/1     Running   0             7h13m
kube-system            kube-apiserver-master                       1/1     Running   0             7h13m
kube-system            kube-controller-manager-master              1/1     Running   0             42m
kube-system            kube-flannel-ds-chrqd                       1/1     Running   0             3h59m
kube-system            kube-flannel-ds-kkqf4                       1/1     Running   0             6h1m
kube-system            kube-proxy-frgrw                            1/1     Running   0             3h59m
kube-system            kube-proxy-h5p5z                            1/1     Running   0             7h12m
kube-system            kube-scheduler-master                       1/1     Running   1 (42m ago)   42m
kubernetes-dashboard   dashboard-metrics-scraper-c45b7869d-xhr24   1/1     Running   0             3h42m
kubernetes-dashboard   kubernetes-dashboard-576cb95f94-kdts4       1/1     Running   0             3h42m


root@node1:~# kubectl get service -n kubernetes-dashboard
NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)         AGE
dashboard-metrics-scraper   ClusterIP   10.101.186.25   <none>        8000/TCP        3h44m
kubernetes-dashboard        NodePort    10.107.59.8     <none>        443:31996/TCP   3h44m



root@node1:~# kubectl describe svc kubernetes-dashboard -n kubernetes-dashboard
Name:                     kubernetes-dashboard
Namespace:                kubernetes-dashboard
Labels:                   k8s-app=kubernetes-dashboard
Annotations:              <none>
Selector:                 k8s-app=kubernetes-dashboard
Type:                     NodePort
IP Family Policy:         SingleStack
IP Families:              IPv4
IP:                       10.107.59.8
IPs:                      10.107.59.8
Port:                     <unset>  443/TCP
TargetPort:               8443/TCP
NodePort:                 <unset>  31996/TCP
Endpoints:                10.244.1.3:8443
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```


https://192.168.105.6:31996


root@node1:~# vim admin-account.yaml
```go

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: admin
  namespace: kubernetes-dashboard

---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name:  cluster-admin
subjects:
  - kind: ServiceAccount
    name: admin
    namespace: kubernetes-dashboard
```



```go

root@node1:~# kubectl apply -f admin-account.yaml
serviceaccount/admin created
clusterrolebinding.rbac.authorization.k8s.io/admin created


root@node1:~# kubectl get serviceaccount -n kubernetes-dashboard
NAME                   SECRETS   AGE
admin                  1         71s
default                1         4h34m
kubernetes-dashboard   1         4h34m


root@node1:~# kubectl describe serviceaccount admin -n kubernetes-dashboard
Name:                admin
Namespace:           kubernetes-dashboard
Labels:              k8s-app=kubernetes-dashboard
Annotations:         <none>
Image pull secrets:  <none>
Mountable secrets:   admin-token-4x6nk
Tokens:              admin-token-4x6nk
Events:              <none>



root@node1:~# kubectl get secret -n kubernetes-dashboard
NAME                               TYPE                                  DATA   AGE
admin-token-4x6nk                  kubernetes.io/service-account-token   3      4m50s
default-token-l6mdd                kubernetes.io/service-account-token   3      4h37m
kubernetes-dashboard-certs         Opaque                                0      4h37m
kubernetes-dashboard-csrf          Opaque                                1      4h37m
kubernetes-dashboard-key-holder    Opaque                                2      4h37m
kubernetes-dashboard-token-mqsn8   kubernetes.io/service-account-token   3      4h37m


root@node1:~# kubectl describe secret admin-token-4x6nk -n kubernetes-dashboard
Name:         admin-token-4x6nk
Namespace:    kubernetes-dashboard
Labels:       <none>
Annotations:  kubernetes.io/service-account.name: admin
              kubernetes.io/service-account.uid: f74424ed-f98f-49eb-b170-7846ca815fc7

Type:  kubernetes.io/service-account-token

Data
====
ca.crt:     1099 bytes
namespace:  20 bytes
token:      eyJhbGciOiJSUzI1NiIsImtpZCI6IlgzSFJwa1VFZnNUZ3R2Wnk1YUZlUVdwdTFqeXJNM2RHY1RkOXNuNVJSMU0ifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi10b2tlbi00eDZuayIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJhZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImY3NDQyNGVkLWY5OGYtNDllYi1iMTcwLTc4NDZjYTgxNWZjNyIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlcm5ldGVzLWRhc2hib2FyZDphZG1pbiJ9.vtHqmnKIU_Vfe1_PDuXga4y0Xe1N5mkA-gUhD398V77_pYIyKqAWwLhaM7JdDRI5f_EsIAP3UcOkhf2pIBZ4jAlKYi-Ad908y-7zheVoyW_RBCMRZ1dGpMjYwYxGhCmTRAO1Z7EGajqz8DDAhXryCO3qdJ7AkesfY4fHPwtSaLiXr5K4088MI-nDcb8zPblaDQM9io2qEAaf45-lNA8754NCOyZ-KKpL0XLSwijgOC-EV973lAeZs7uOG0LI2xDNGN-mHNiCy5ccXS7UIrqd76_jPGGqinPXJVAD7KMYRE_VDUHPIZ9dfCFn_Wcfh_imkuamZE-78be6Vak2N0iPUA



```

