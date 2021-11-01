



**使用multipass搭建k8s多节点集群和Dashboard**

```shell
multipass launch -n master -c 1 -m 3G -d 20G
multipass launch -n node1 -c 1 -m 3G -d 20G
multipass launch -n node2 -c 1 -m 3G -d 20G


ubuntu@node1:~$ microk8s.enable registry:size=40G
Addon registry is already enabled.

➜  ~ multipass list
Name                    State             IPv4             Image
master                  Running           192.168.105.5    Ubuntu 20.04 LTS
node1                   Running           192.168.105.6    Ubuntu 20.04 LTS
node2                   Running           192.168.105.7    Ubuntu 20.04 LTS

ubuntu@node1:~$ sudo mv /etc/apt/sources.list /etc/apt/sources.list.bak
ubuntu@node1:~$ sudo vim /etc/apt/sources.list

# ubuntu 20.04(focal) 配置如下
deb http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse

deb http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse


ubuntu@master:~$ sudo apt-get update && sudo apt-get upgrade -y

```





**Ubuntu 安装 Docker**



```shell
# 使用官方安装脚本自动安装
ubuntu@master:~$ curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun

ubuntu@master:~$ docker images
Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Get "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/images/json": dial unix /var/run/docker.sock: connect: permission denied

ubuntu@master:~$ sudo groupadd docker
groupadd: group 'docker' already exists
ubuntu@master:~$ sudo gpasswd -a ubuntu docker
Adding user ubuntu to group docker
ubuntu@master:~$ sudo service docker restart
ubuntu@master:~$ sudo vim /etc/docker/daemon.json

{ "registry-mirrors": [
    "https://hkaofvr0.mirror.aliyuncs.com"
  ]
 }

ubuntu@master:~$ sudo systemctl daemon-reload
ubuntu@master:~$ sudo systemctl restart docker
# 重启 iTerm2
ubuntu@node1:~$ exit
logout
➜  ~ multipass shell node1

ubuntu@master:~$ docker info

 Registry Mirrors:
  https://hkaofvr0.mirror.aliyuncs.com/

# Install Compose on Linux systems

sudo apt install docker-compose -y

ubuntu@master:~$ sudo curl -L "https://github.com/docker/compose/releases/download/v2.0.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

ubuntu@master:~$ sudo chmod +x /usr/local/bin/docker-compose
ubuntu@master:~$ sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
ubuntu@master:~$ docker-compose --version

Docker Compose version v2.0.1


```

```shell
sudo apt-get update && sudo apt-get upgrade -y


sudo apt-get dist-upgrade -y
```


**root设置**
```shell
ubuntu@master:~$ sudo passwd root
ubuntu@master:~$ sudo passwd -dl root

ubuntu@master:~$ su root
Password:
root@master:/home/ubuntu#

root@master:/home/ubuntu# su ubuntu
ubuntu@master:~$
```




**Ubuntu 安装 Docker**



```shell
# 使用官方安装脚本自动安装
ubuntu@master:~$ curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun

ubuntu@master:~$ docker images
Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Get "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/images/json": dial unix /var/run/docker.sock: connect: permission denied

ubuntu@master:~$ sudo groupadd docker
groupadd: group 'docker' already exists
ubuntu@master:~$ sudo gpasswd -a ubuntu docker
Adding user ubuntu to group docker
ubuntu@master:~$ sudo service docker restart
ubuntu@master:~$ sudo vim /etc/docker/daemon.json

{ "registry-mirrors": [
    "https://hkaofvr0.mirror.aliyuncs.com"
  ]
 }

ubuntu@master:~$ sudo systemctl daemon-reload
ubuntu@master:~$ sudo systemctl restart docker
# 重启 iTerm2
ubuntu@node1:~$ exit
logout
➜  ~ multipass shell node1

ubuntu@master:~$ docker info

 Registry Mirrors:
  https://hkaofvr0.mirror.aliyuncs.com/

# Install Compose on Linux systems

sudo apt install docker-compose -y

ubuntu@master:~$ sudo curl -L "https://github.com/docker/compose/releases/download/v2.0.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

ubuntu@master:~$ sudo chmod +x /usr/local/bin/docker-compose
ubuntu@master:~$ sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
ubuntu@master:~$ docker-compose --version

Docker Compose version v2.0.1


```





**ubuntu 安装 microk8s**

```shell

ubuntu@master:~$ sudo snap install microk8s --classic
microk8s (1.22/stable) v1.22.2 from Canonical✓ installed

ubuntu@master:~$ microk8s status
Insufficient permissions to access MicroK8s.
You can either try again with sudo or add the user ubuntu to the 'microk8s' group:

    sudo usermod -a -G microk8s ubuntu
    sudo chown -f -R ubuntu ~/.kube

After this, reload the user groups either via a reboot or by running 'newgrp microk8s'.
ubuntu@master:~$ sudo usermod -a -G microk8s ubuntu
ubuntu@master:~$ sudo chown -f -R ubuntu ~/.kube
ubuntu@master:~$ newgrp microk8s

ubuntu@master:~$ microk8s status
microk8s is not running. Use microk8s inspect for a deeper inspection.
ubuntu@master:~$ microk8s inspect

Inspecting Certificates
Inspecting services
  Service snap.microk8s.daemon-cluster-agent is running
  Service snap.microk8s.daemon-containerd is running
  Service snap.microk8s.daemon-apiserver-kicker is running
  Service snap.microk8s.daemon-kubelite is running
  Copy service arguments to the final report tarball
Inspecting AppArmor configuration
Gathering system information
  Copy processes list to the final report tarball
  Copy snap list to the final report tarball
  Copy VM name (or none) to the final report tarball
  Copy disk usage information to the final report tarball
  Copy memory usage information to the final report tarball
  Copy server uptime to the final report tarball
  Copy current linux distribution to the final report tarball
  Copy openSSL information to the final report tarball
  Copy network configuration to the final report tarball
Inspecting kubernetes cluster
  Inspect kubernetes cluster
Inspecting juju
  Inspect Juju
Inspecting kubeflow
  Inspect Kubeflow


The change can be made persistent with: sudo apt-get install iptables-persistent
WARNING:  Docker is installed.
Add the following lines to /etc/docker/daemon.json:
{
    "insecure-registries" : ["localhost:32000"]
}
and then restart docker with: sudo systemctl restart docker
Building the report tarball
  Report tarball is at /var/snap/microk8s/2551/inspection-report-20211031_124330.tar.gz
ubuntu@master:~$ sudo iptables -P FORWARD ACCEPT
ubuntu@master:~$ sudo vim /etc/docker/daemon.json

{   "exec-opts": ["native.cgroupdriver=systemd"],
    "registry-mirrors":[
        "https://hkaofvr0.mirror.aliyuncs.com",
        "http://docker.mirrors.ustc.edu.cn"
    ],
    "insecure-registries" : ["localhost:32000"]
}

ubuntu@master:~$ sudo systemctl restart docker




ubuntu@master:~$ sudo vim /var/snap/microk8s/current/args/kubelet

--pod-infra-container-image=s7799653/pause:3.1

ubuntu@master:~$ sudo vim /var/snap/microk8s/current/args/containerd-template.toml

sandbox_image = "s7799653/pause:3.1"



 "registry-mirrors":[
        "https://hkaofvr0.mirror.aliyuncs.com",
    ]

ubuntu@master:~$ microk8s.stop && microk8s.start



设置kubectl别名：

sudo snap alias microk8s.kubectl kubectl

sudo snap alias microk8s.kubectl microk8s.kubectl




ubuntu@master:~$ kubectl get po -n kube-system
NAME                                         READY   STATUS             RESTARTS   AGE
kubernetes-dashboard-59699458b-cj6nf         1/1     Running            0          5m45s
coredns-7f9c69c78c-dhp7b                     1/1     Running            0          11m
calico-kube-controllers-78588f7f6-krq5p      1/1     Running            0          38m
calico-node-mqrqg                            1/1     Running            0          38m
metrics-server-85df567dd8-n725n              0/1     ImagePullBackOff   0          5m45s
dashboard-metrics-scraper-58d4977855-9jj7f   1/1     Running            0          5m45s


```



```shell
ubuntu@master:~$  kubectl describe pods metrics-server -n kube-system
Name:                 metrics-server-85df567dd8-n725n
Namespace:            kube-system
Priority:             2000000000
Priority Class Name:  system-cluster-critical
Node:                 master/192.168.105.5
Start Time:           Sun, 31 Oct 2021 13:12:11 +0800
Labels:               k8s-app=metrics-server
                      pod-template-hash=85df567dd8
Annotations:          cni.projectcalico.org/podIP: 10.1.219.67/32
                      cni.projectcalico.org/podIPs: 10.1.219.67/32
Status:               Pending
IP:                   10.1.219.67
IPs:
  IP:           10.1.219.67
Controlled By:  ReplicaSet/metrics-server-85df567dd8
Containers:
  metrics-server:
    Container ID:
    Image:         k8s.gcr.io/metrics-server/metrics-server:v0.5.0
    Image ID:
    Port:          443/TCP
    Host Port:     0/TCP
    Args:
      --cert-dir=/tmp
      --secure-port=443
      --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
      --kubelet-use-node-status-port
      --metric-resolution=15s
      --kubelet-insecure-tls
    State:          Waiting
      Reason:       ImagePullBackOff
    Ready:          False
    Restart Count:  0
    Requests:
      cpu:        100m
      memory:     200Mi
    Liveness:     http-get https://:https/livez delay=0s timeout=1s period=10s #success=1 #failure=3
    Readiness:    http-get https://:https/readyz delay=20s timeout=1s period=10s #success=1 #failure=3
    Environment:  <none>
    Mounts:
      /tmp from tmp-dir (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-sm5hl (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             False
  ContainersReady   False
  PodScheduled      True
Volumes:
  tmp-dir:
    Type:       EmptyDir (a temporary directory that shares a pod\'s lifetime)
    Medium:
    SizeLimit:  <unset>
  kube-api-access-sm5hl:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   Burstable
Node-Selectors:              kubernetes.io/arch=amd64
                             kubernetes.io/os=linux
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type     Reason     Age                    From               Message
  ----     ------     ----                   ----               -------
  Normal   Scheduled  7m55s                  default-scheduler  Successfully assigned kube-system/metrics-server-85df567dd8-n725n to master
  Warning  Failed     7m6s                   kubelet            Failed to pull image "k8s.gcr.io/metrics-server/metrics-server:v0.5.0": rpc error: code = Unknown desc = failed to pull and unpack image "k8s.gcr.io/metrics-server/metrics-server:v0.5.0": failed to resolve reference "k8s.gcr.io/metrics-server/metrics-server:v0.5.0": failed to do request: Head "https://k8s.gcr.io/v2/metrics-server/metrics-server/manifests/v0.5.0": dial tcp 64.233.189.82:443: i/o timeout
  Warning  Failed     7m6s                   kubelet            Error: ErrImagePull
  Normal   BackOff    7m5s                   kubelet            Back-off pulling image "k8s.gcr.io/metrics-server/metrics-server:v0.5.0"
  Warning  Failed     7m5s                   kubelet            Error: ImagePullBackOff
  Normal   Pulling    6m54s (x2 over 7m36s)  kubelet            Pulling image "k8s.gcr.io/metrics-server/metrics-server:v0.5.0"
  Warning  Failed     117s (x2 over 2m54s)   kubelet            Failed to pull image "k8s.gcr.io/metrics-server/metrics-server:v0.5.0": rpc error: code = Unknown desc = failed to pull and unpack image "k8s.gcr.io/metrics-server/metrics-server:v0.5.0": failed to resolve reference "k8s.gcr.io/metrics-server/metrics-server:v0.5.0": failed to do request: Head "https://k8s.gcr.io/v2/metrics-server/metrics-server/manifests/v0.5.0": dial tcp 64.233.189.82:443: i/o timeout
  Warning  Failed     117s (x2 over 2m54s)   kubelet            Error: ErrImagePull
  Normal   BackOff    105s (x2 over 2m54s)   kubelet            Back-off pulling image "k8s.gcr.io/metrics-server/metrics-server:v0.5.0"
  Warning  Failed     105s (x2 over 2m54s)   kubelet            Error: ImagePullBackOff
  Normal   Pulling    94s (x3 over 3m25s)    kubelet            Pulling image "k8s.gcr.io/metrics-server/metrics-server:v0.5.0"
```

```shell
好像不好用
由于网络问题会导致相关镜像下载失败，我们通常会使用VPN进行镜像下载，这里提供一个无需VPN的配置方式，
在/etc/hosts中添加以下内容：

sudo vim /etc/hosts

172.217.197.82   k8s.gcr.io
```


```shell

kubectl set image -n kube-system deployment/metrics-server metrics-server=registry.aliyuncs.com/google_containers/metrics-server:v0.5.0


ubuntu@master:~$ kubectl get po -n kube-system
NAME                                         READY   STATUS    RESTARTS      AGE
kubernetes-dashboard-59699458b-cj6nf         1/1     Running   0             19m
dashboard-metrics-scraper-58d4977855-9jj7f   1/1     Running   0             19m
calico-node-mqrqg                            1/1     Running   0             51m
coredns-7f9c69c78c-dhp7b                     1/1     Running   0             24m
calico-kube-controllers-78588f7f6-krq5p      1/1     Running   0             51m
metrics-server-75cff5db9-5bbzz               1/1     Running   1 (50s ago)   2m51s



ubuntu@master:~$ kubectl config view
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: DATA+OMITTED
    server: https://127.0.0.1:16443
  name: microk8s-cluster
contexts:
- context:
    cluster: microk8s-cluster
    user: admin
  name: microk8s
current-context: microk8s
kind: Config
preferences: {}
users:
- name: admin
  user:
    token: REDACTED


ubuntu@master:~$ kubectl config get-contexts
CURRENT   NAME       CLUSTER            AUTHINFO   NAMESPACE
*         microk8s   microk8s-cluster   admin




ubuntu@master:~$ microk8s enable dashboard dns registry istio

ubuntu@master:~$ kubectl port-forward -n kube-system service/kubernetes-dashboard 10443:443

ubuntu@master:~$ microk8s dashboard-proxy
Checking if Dashboard is running.
Dashboard will be available at https://127.0.0.1:10443
Use the following token to login:
eyJhbGciOiJSUzI1NiIsImtpZCI6IkNUVlFwaFgyUkpib3Bab0x1MDJrTEpHakhXNEZEc2lwbmFuWFdFUkxEbFUifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkZWZhdWx0LXRva2VuLWQyaHdmIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImRlZmF1bHQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJkNWU4MDQxNS04MDg0LTQ5MzYtYWY3NC0zOTRkYjI1NjBhMmMiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06ZGVmYXVsdCJ9.Rq3WsLmHTsX1_TKe7eyaCl3Otm8HPFuJWU-_xWOm7HguBYNufEIVzyj35gtfmfOHIN_Qbk5hSM69D5qAbiku_PDQEcjk6NHcqEk4fzXhvvYqqwbsDpYWT1Q10xQbcOcdFEOwk9meAxa_fqxSHczyl4WoS20JiyWk4TRY4FE4YxwL5rj4HVhGIjstK6FyKhUGloqE5HiPQ_xlxCgdzMMVM-R8eQ6-2IHMGmK6zLeDOzKl3_l9CqLiRPPTjLpGRtxRQgB40sNALQZScU_Q6HAyCY5N-nc3vFc-v0bG_nT3bH3VsHf3CVBqKWBQNNN1P9JG6JVeFigClLmqgKdKSQYNYw




ubuntu@node1:~$ microk8s config
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREekNDQWZlZ0F3SUJBZ0lVY3NUU1dXR05TekoyelQyVjRMMENSMnVGd3I0d0RRWUpLb1pJaHZjTkFRRUwKQlFBd0Z6RVZNQk1HQTFVRUF3d01NVEF1TVRVeUxqRTRNeTR4TUI0WERUSXhNVEF6TVRBME16WXhPRm9YRFRNeApNVEF5T1RBME16WXhPRm93RnpFVk1CTUdBMVVFQXd3TU1UQXVNVFV5TGpFNE15NHhNSUlCSWpBTkJna3Foa2lHCjl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUE2eW9GUDVGZFR4cUdzb09TTldHTDZqbTJZbzIxbFNCOEE4dnYKZW9OS3lFWnR3MUtuSmd3d09UdU1Lek5CMG9WSzl6ZUtTMlFrd3BhdDNKSTlicE92RWZzbElUTVBhUDMxdDRSTwp0dzBZOTNaSms2eFV3NTAvSkROalUrS0M0bzlYQk1jZ1dtcDJyb2xRVGYrTmRtUW9tZlpOOTJHMzdGQXZKQ3o1CjYwYXFhazRxVXVLZmRRNjhHY0xleU51VHByRSt0Yi9iNmJHbXpGcWFIc2Q0Ukplazl5T0F2cThSNzlOeWNXWGsKZlRDMGhIVnRNN0VhNEhUS2ZoQ2pKeExXaTdyK0RPV0U1cEF3ODVGK3d0MStxZ01KeEdCUjMrS3FYWHo0Qm9zUQpxbFhoZTd0dUJrWGpOSkZWNDRyalJjYTBBUENYUTdOQjJrYWlVMG1NL2F5ZXovZlFFd0lEQVFBQm8xTXdVVEFkCkJnTlZIUTRFRmdRVTloRS9UNEM2Y3NickRHeFVyUEhMeGxDTWxQRXdId1lEVlIwakJCZ3dGb0FVOWhFL1Q0QzYKY3NickRHeFVyUEhMeGxDTWxQRXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQwpBUUVBQWw2TVEyZWZ5aW94dmRjbjUyREdML2lMMUM2UDR0eWFJNVJQVWJEZFVNYThlRHFsVzNJd3BsTWtpcWRvCnUydG5NajFEbWJZb2pKektnL3BOWVp4dzl2cVVaREhJV2FBT0VtaWoyaUxpOHg3U3JSeEtqd2NFT3FRMmlQT2IKN0ZmU2NPaXp3RGdRYmRkSzVpb2xhdXloS3Z4MERUM2ZjZDJpemp0N0QreitZNlhaOTg0MHhOeVhhcXA2ZTBxNApKU2llcUNpYXV1R2VucU04QnpnRzV1V3YvQ1FmdmpuMHNSZVVKY0phSTl1QTAzblNBYUhyZS9XQWwxQ2JDYkVLCms1aTZrVmJta3A3LzI0NUFCeFp0WitkRkY4alZaSVhhS1JSMkhSR080V3FWTFV2cXZmREFaZ1BJdHVNbVhiR1EKSGZ6RkhzZGFKcWNLeWpXUHJCMklRUlNZMUE9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    server: https://192.168.105.6:16443
  name: microk8s-cluster
contexts:
- context:
    cluster: microk8s-cluster
    user: admin
  name: microk8s
current-context: microk8s
kind: Config
preferences: {}
users:
- name: admin
  user:
    token: OVJzY1ZKOWJiNGRiWG9sczJrdzQ4eE9kSFZJNktMT24vZnlBYysySGhGOD0K




ubuntu@node2:~$ microk8s config
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREekNDQWZlZ0F3SUJBZ0lVWlRhNUhCanczYWpkQy90bGRrSUtNZlgzb3dBd0RRWUpLb1pJaHZjTkFRRUwKQlFBd0Z6RVZNQk1HQTFVRUF3d01NVEF1TVRVeUxqRTRNeTR4TUI0WERUSXhNVEF6TVRBME16WXpNbG9YRFRNeApNVEF5T1RBME16WXpNbG93RnpFVk1CTUdBMVVFQXd3TU1UQXVNVFV5TGpFNE15NHhNSUlCSWpBTkJna3Foa2lHCjl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUEwaXUwU05jdzVLc01sY3hKdWY1blJhN0JrVzJNeFNVdFg0M3cKSEdHWmVMSHNEREE5Sm9aUWZIZUQ3bjlML05mRS81MEJ5Q1Y4K3NYSlp4dXBDeStka1M0d2FnSGFncVdYUVVuZwo5MzFTS1NGc0xOTGM4VEVlOW9ONWRGL1YvVEx0cENrSWhhNVU2ZHVmc2ZVcXFnbGR0QXFHZnRDOENWQ2dOZWJrCkIya2ovWWlCamtnZm02VDFCcFFCSXVUOHZmekFGdyt3U2JpZUNPYjVuR1NYZmFneDdBeU41U3lCY0RFeGhETDkKRW4xdTNYcWt4aGE4SzMxUWRqVW5URW5NSTNPK01DR3FJZGd3WUpCemNwK1hDTVBkaE5Pd1hxRDF2NjY3N29icQpiWGh6VS9CTlZzLzUyVStvdkV4ekZLRktHSEErM3pLV0dEZlV6TDZaNS85ZEl1ZDFWd0lEQVFBQm8xTXdVVEFkCkJnTlZIUTRFRmdRVTNEV1ZrWUR5cXhBNjUyNzZaUFQzcDJTSEV0WXdId1lEVlIwakJCZ3dGb0FVM0RXVmtZRHkKcXhBNjUyNzZaUFQzcDJTSEV0WXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQwpBUUVBU2svRWwyWkJDT2JlRHBNRFpETm5jZHRHVWpVU3kzcCszVWM1Z2hJdTV2YVp5MVFHS3B1am9lTUpqakxoCkQxc0FJQ3RUNTZKVGplY3VMVFNSVEtBeW9UQVZDeDVRQ2x3YkVzamFadTJ3TTcwSElKS0MvaVZSdkdXVFlJelgKTURQYjVFVExEUnZ4NTlCWHgvcHRkM3VrZnUxNDA1Q1hua1U5dXlMODlNOFpOYlgwZzFNYUFTRHByUGcrVWJCWAowMmhWOGptaE5ITi9UTEFIa2hBajJzQkIzczVDTjQ1WVZTOVVTTHIzK09PRC9Lb0dvcU8vd3RMQVhMbjBvVm1SCkhFMHNaUUlVSC9Mb2lQbVh0OGF2VWYzTkhOK29KSXU2S0lOR2JUaUhBc1htb2l4S1d1WUVOaWNDRVNodXFoK3EKMW9idWFVa1FwOVZjTEtVcmRrcU82WFplenc9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    server: https://192.168.105.7:16443
  name: microk8s-cluster
contexts:
- context:
    cluster: microk8s-cluster
    user: admin
  name: microk8s
current-context: microk8s
kind: Config
preferences: {}
users:
- name: admin
  user:
    token: SG16QUZIS2M0bjRhQ2pCd2xpay9uSVNDVElob3V0K0lSU0ZhWVR6Tll4az0K




ubuntu@master:~$ microk8s add-node
From the node you wish to join to this cluster, run the following:
microk8s join 192.168.105.5:25000/d18d10d874acfd2b539e490e93439f2f/0e3a1ca4a9ea

If the node you are adding is not reachable through the default interface you can use one of the following:
 microk8s join 192.168.105.5:25000/d18d10d874acfd2b539e490e93439f2f/0e3a1ca4a9ea
 microk8s join 172.17.0.1:25000/d18d10d874acfd2b539e490e93439f2f/0e3a1ca4a9ea





ubuntu@node1:~$ microk8s join 192.168.105.5:25000/cad1bfa9289242ab68e41c3996db7ffa/0e3a1ca4a9ea
Contacting cluster at 192.168.105.5
Waiting for this node to finish joining the cluster. .. .. .. .. .. .. .. .. .. ..


ubuntu@node2:~$ microk8s join 192.168.105.5:25000/03d641eb4d7790ddca34bbdb7d27e184/0e3a1ca4a9ea
Contacting cluster at 192.168.105.5
Waiting for this node to finish joining the cluster. ..


ubuntu@master:~$ microk8s kubectl get no
NAME     STATUS   ROLES    AGE     VERSION
master   Ready    <none>   3h17m   v1.22.2-3+9ad9ee77396805
node1    Ready    <none>   34m     v1.22.2-3+9ad9ee77396805
node2    Ready    <none>   27m     v1.22.2-3+9ad9ee77396805




multipass exec master -- sudo /snap/bin/microk8s kubectl port-forward -n kube-system service/kubernetes-dashboard 10443:443 --address 0.0.0.0




ubuntu@master:~$ vim pod_nginx.yml

apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  containers:
  - name: nginx
    image: nginx
    ports:
    - containerPort: 80








ubuntu@master:~$ sudo kubectl get pods -o wide
NAME    READY   STATUS    RESTARTS   AGE   IP             NODE    NOMINATED NODE   READINESS GATES
nginx   1/1     Running   0          29m   10.1.166.130   node1   <none>           <none>



ubuntu@master:~$ sudo kubectl describe pods nginx
Name:         nginx
Namespace:    default
Priority:     0
Node:         node1/192.168.105.6
Start Time:   Sun, 31 Oct 2021 18:35:58 +0800
Labels:       app=nginx
Annotations:  cni.projectcalico.org/podIP:
              cni.projectcalico.org/podIPs:
Status:       Running
IP:           10.1.166.130
IPs:
  IP:  10.1.166.130
Containers:
  nginx:
    Container ID:   containerd://fe8c7412fee20b1e6863b4a068b95791b387f4213e04cd1f3c031fe555e8ac8e
    Image:          nginx
    Image ID:       docker.io/library/nginx@sha256:644a70516a26004c97d0d85c7fe1d0c3a67ea8ab7ddf4aff193d9f301670cf36
    Port:           80/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Sun, 31 Oct 2021 18:44:25 +0800
    Ready:          True



➜  demo cat rc_nginx.yaml

apiVersion: v1
kind: ReplicationController
metadata:
  name: nginx # 随机产生一个名字
spec:
  replicas: 3 # 创建三个副本，也就是三个pod
  selector:
    app: nginx
  template: # 这个就是在定义一个pod
    metadata:
      name: nginx
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        ports:
        - containerPort: 80 # 暴露80端口


ubuntu@master:~$ sudo kubectl create -f rc_nginx.yaml
replicationcontroller/nginx created




ubuntu@master:~$ sudo kubectl get rc
NAME    DESIRED   CURRENT   READY   AGE
nginx   3         3         0       2m19s
ubuntu@master:~$ sudo kubectl get pods
NAME          READY   STATUS              RESTARTS        AGE
nginx         1/1     Running             3 (3m53s ago)   15h
nginx-vxd9w   0/1     ContainerCreating   0               69s
nginx-x8s4t   0/1     ContainerCreating   0               63s




ubuntu@master:~$ sudo kubectl delete pods nginx-x8s4t
pod "nginx-x8s4t" deleted
ubuntu@master:~$ sudo kubectl get rc
NAME    DESIRED   CURRENT   READY   AGE
nginx   3         3         3       5m44s
ubuntu@master:~$ sudo kubectl get pods
NAME          READY   STATUS    RESTARTS       AGE
nginx         1/1     Running   3 (7m1s ago)   15h
nginx-vxd9w   1/1     Running   0              4m17s
nginx-jfbhz   1/1     Running   0              21s


ubuntu@master:~$ sudo kubectl scale rc nginx --replicas=2
replicationcontroller/nginx scaled
ubuntu@master:~$ sudo kubectl get pods
NAME          READY   STATUS    RESTARTS        AGE
nginx         1/1     Running   3 (9m38s ago)   15h
nginx-vxd9w   1/1     Running   0               6m54s

ubuntu@master:~$ sudo kubectl get rc
NAME    DESIRED   CURRENT   READY   AGE
nginx   2         2         2       8m54s

ubuntu@master:~$ sudo kubectl scale rc nginx --replicas=3
replicationcontroller/nginx scaled
ubuntu@master:~$ sudo kubectl get rc
NAME    DESIRED   CURRENT   READY   AGE
nginx   3         3         2       17m
ubuntu@master:~$ sudo kubectl get pods
NAME          READY   STATUS              RESTARTS      AGE
nginx-vxd9w   1/1     Running             0             15m
nginx         1/1     Running             3 (18m ago)   16h
nginx-65c4h   0/1     ContainerCreating   0             6s


ubuntu@master:~$ sudo kubectl get pods -o wide
NAME          READY   STATUS    RESTARTS      AGE   IP             NODE    NOMINATED NODE   READINESS GATES
nginx-vxd9w   1/1     Running   0             16m   10.1.104.1     node2   <none>           <none>
nginx         1/1     Running   3 (19m ago)   16h   10.1.166.133   node1   <none>           <none>
nginx-65c4h   1/1     Running   0             47s   10.1.166.136   node1   <none>           <none>

ubuntu@master:~$ sudo kubectl delete -f rc_nginx.yaml
replicationcontroller "nginx" deleted
ubuntu@master:~$ sudo kubectl get pods
NAME          READY   STATUS        RESTARTS      AGE
nginx-vxd9w   1/1     Terminating   0             19m
nginx         1/1     Terminating   3 (22m ago)   16h
nginx-65c4h   1/1     Terminating   0             3m49s


ubuntu@master:~$ cat rc_nginx.yaml
apiVersion: apps/v1
kind: ReplicationController
metadata:
  name: nginx # 随机产生一个名字
spec:
  replicas: 3 # 创建三个副本，也就是三个pod
  selector:
    app: nginx
  template: # 这个就是在定义一个pod
    metadata:
      name: nginx
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        ports:
        - containerPort: 80 # 暴露80端口






```

















**k8s**

```shell

ubuntu@master:~$ sudo snap install kubeadm --classic
kubeadm 1.22.3 from Canonical✓ installed
ubuntu@master:~$ sudo snap install kubelet --classic
kubelet 1.22.3 from Canonical✓ installed
ubuntu@master:~$ sudo snap install kubectl --classic
kubectl 1.22.3 from Canonical✓ installed


ubuntu@master:~$ sudo timedatectl set-timezone Asia/Shanghai
ubuntu@master:~$ sudo systemctl restart rsyslog
ubuntu@master:~$ sudo swapoff -a
ubuntu@master:~$ free -m

ubuntu@master:~$ sudo vim /etc/docker/daemon.json


{
    "registry-mirrors":[
        "https://hkaofvr0.mirror.aliyuncs.com",
        "http://docker.mirrors.ustc.edu.cn",
        "http://hub-mirror.c.163.com"
    ],
    "insecure-registries":[
        "registry.docker-cn.com",
        "docker.mirrors.ustc.edu.cn"
    ],
    "debug":true,
    "experimental":true
}




ubuntu@master:~$ docker run hello-world

Hello from Docker!

ubuntu@master:~$ sudo service docker restart
ubuntu@master:~$ docker info







sudo apt-get update && sudo apt-get install -y ca-certificates curl software-properties-common apt-transport-https curl


ubuntu@master:~$ sudo vim /etc/apt/sources.list.d/kubernetes.list

deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main

sudo apt-get update && sudo apt-get install -y kubelet kubeadm kubectl


由于网络原因，提前需要准备的镜像：

ubuntu@master:~$ sudo kubeadm config images list
k8s.gcr.io/kube-apiserver:v1.22.3
k8s.gcr.io/kube-controller-manager:v1.22.3
k8s.gcr.io/kube-scheduler:v1.22.3
k8s.gcr.io/kube-proxy:v1.22.3
k8s.gcr.io/pause:3.5
k8s.gcr.io/etcd:3.5.0-0
k8s.gcr.io/coredns/coredns:v1.8.4


通过阿里云源下载上面镜像：

sudo docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:v1.22.3

sudo docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:v1.22.3

sudo docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:v1.22.3

sudo docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:v1.22.3

sudo docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.5

sudo docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:3.5.0-0

sudo docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:v1.8.4


sudo docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver:v1.22.3 k8s.gcr.io/kube-apiserver:v1.22.3

sudo docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager:v1.22.3 k8s.gcr.io/kube-controller-manager:v1.22.3

sudo docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler:v1.22.3 k8s.gcr.io/kube-scheduler:v1.22.3

sudo docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy:v1.22.3 k8s.gcr.io/kube-proxy:v1.22.3

sudo docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.5 k8s.gcr.io/pause:3.5

sudo docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/etcd:3.5.0-0 k8s.gcr.io/etcd:3.5.0-0

sudo docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:v1.8.4 k8s.gcr.io/coredns/coredns:v1.8.4






ubuntu@master:~$ docker images
REPOSITORY                                                                    TAG       IMAGE ID       CREATED        SIZE
k8s.gcr.io/kube-apiserver                                                     v1.22.3   53224b502ea4   3 days ago     128MB
registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver            v1.22.3   53224b502ea4   3 days ago     128MB
k8s.gcr.io/kube-scheduler                                                     v1.22.3   0aa9c7e31d30   3 days ago     52.7MB
registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler            v1.22.3   0aa9c7e31d30   3 days ago     52.7MB
k8s.gcr.io/kube-controller-manager                                            v1.22.3   05c905cef780   3 days ago     122MB
registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager   v1.22.3   05c905cef780   3 days ago     122MB
k8s.gcr.io/kube-proxy                                                         v1.22.3   6120bd723dce   3 days ago     104MB
registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy                v1.22.3   6120bd723dce   3 days ago     104MB
hello-world                                                                   latest    feb5d9fea6a5   5 weeks ago    13.3kB
k8s.gcr.io/etcd                                                               3.5.0-0   004811815584   4 months ago   295MB
registry.cn-hangzhou.aliyuncs.com/google_containers/etcd                      3.5.0-0   004811815584   4 months ago   295MB
k8s.gcr.io/coredns/coredns                                                    v1.8.4    8d147537fb7d   5 months ago   47.6MB
registry.cn-hangzhou.aliyuncs.com/google_containers/coredns                   v1.8.4    8d147537fb7d   5 months ago   47.6MB
k8s.gcr.io/pause                                                              3.5       ed210e3e4a5b   7 months ago   683kB
registry.cn-hangzhou.aliyuncs.com/google_containers/pause                     3.5       ed210e3e4a5b   7 months ago   683kB



**初始化所有节点**

sudo kubeadm reset
sudo rm -rf $HOME/.kube/config


[kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10248/healthz' failed with error: Get "http://localhost:10248/healthz": dial tcp 127.0.0.1:10248: connect: connection refused.

cgroup驱动问题。默认情况下Kubernetes cgroup驱动程序设置为system，但docker设置为systemd。我们需要更改Docker cgroup驱动，通过创建配置文件/etc/docker/daemon.json并添加以下行：

{"exec-opts": ["native.cgroupdriver=systemd"]}


方法一：
ubuntu@master:~$ sudo vim /etc/docker/daemon.json

{   "exec-opts": ["native.cgroupdriver=systemd"],
    "registry-mirrors":[
        "https://hkaofvr0.mirror.aliyuncs.com",
        "http://docker.mirrors.ustc.edu.cn",
        "http://hub-mirror.c.163.com"
    ],
    "insecure-registries":[
        "registry.docker-cn.com",
        "docker.mirrors.ustc.edu.cn"
    ],
    "debug":true,
    "experimental":true
}



注意：命令将会重写/etc/docker/daemon.json
然后，为使配置生效，你必须重启docker和kubelet。


sudo systemctl daemon-reload
sudo systemctl restart docker
sudo systemctl restart kubelet


**初始化 master 节点:**

ubuntu@master:~$ sudo rm -rf $HOME/.kube/config
ubuntu@master:~$ sudo kubeadm init --pod-network-cidr 10.193.165.1/24 \
>     --image-repository registry.cn-hangzhou.aliyuncs.com/google_containers





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

sudo kubeadm join 192.168.105.5:6443 --token w9h8zb.s4btu28r5ea7acjd \
	--discovery-token-ca-cert-hash sha256:53109f3d8a8ab8dca55bd551d1dbc19dec7926eb7b6a5106f26b10f52e407a52



ubuntu@master:~$ mkdir -p $HOME/.kube
ubuntu@master:~$ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
ubuntu@master:~$ sudo chown $(id -u):$(id -g) $HOME/.kube/config
ubuntu@master:~$ export KUBECONFIG=/etc/kubernetes/admin.conf



ubuntu@node1:~$ sudo kubeadm join 192.168.105.5:6443 --token w9h8zb.s4btu28r5ea7acjd \
> --discovery-token-ca-cert-hash sha256:53109f3d8a8ab8dca55bd551d1dbc19dec7926eb7b6a5106f26b10f52e407a52
[preflight] Running pre-flight checks
[preflight] Reading configuration from the cluster...
[preflight] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Starting the kubelet
[kubelet-start] Waiting for the kubelet to perform the TLS Bootstrap...

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the control-plane to see this node join the cluster.


ubuntu@node2:~$ sudo kubeadm join 192.168.105.5:6443 --token w9h8zb.s4btu28r5ea7acjd \
> --discovery-token-ca-cert-hash sha256:53109f3d8a8ab8dca55bd551d1dbc19dec7926eb7b6a5106f26b10f52e407a52
[preflight] Running pre-flight checks
[preflight] Reading configuration from the cluster...
[preflight] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[kubelet-start] Writing kubelet environment file with flags to file "/var/lib/kubelet/kubeadm-flags.env"
[kubelet-start] Starting the kubelet
[kubelet-start] Waiting for the kubelet to perform the TLS Bootstrap...

This node has joined the cluster:
* Certificate signing request was sent to apiserver and a response was received.
* The Kubelet was informed of the new secure connection details.

Run 'kubectl get nodes' on the control-plane to see this node join the cluster.





ubuntu@master:~$ sudo kubectl get nodes
The connection to the server localhost:8080 was refused - did you specify the right host or port?


ubuntu@master:~$ systemctl status kubelet

kubelet.go:2337] "Container runtime network not ready



```

出现这个问题的原因是kubectl命令需要使用kubernetes-admin来运行，解决方法如下，将主节点中的【/etc/kubernetes/admin.conf】文件拷贝到从节点相同目录下，然后配置环境变量：
echo "export KUBECONFIG=/etc/kubernetes/admin.conf" >> ~/.bash_profile

立即生效：
source ~/.bash_profile

ubuntu@master:~$ sudo vim /etc/kubernetes/admin.conf

```bash
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvakNDQWVhZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJeE1UQXpNVEF6TVRFek5Wb1hEVE14TVRBeU9UQXpNVEV6TlZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBS0FHCk90Ny9rbTZhaHlTcmdrQ2w5OXJMek9nSUUxTG9sbkIxK3BEaFhzRlh3WW5URSt4NGtGcjRvamJ0SlF6eEkxVEEKeHlaNFFCMFI5OG5jMEhBY3JjRjA1K1FaT3VST2xRenhXYmMwZlREaXpwQy83TnhzeG5hN201bGpVQTFaaHNHdApURlhobEpmb2JqL0tjRTJIZXRWVElFNU1XR1ZKcWpXVE1obHBPZVhCL2ZVV0dNMGZkZkljc2NUWlZYaXVremoxCndmdlhnQ3JiYklKVUFrTEdocjY4TFh1ZnVyVzk5NUJscjd6UXJjeUVaMmRwMVV4ZE5Va0dtU0o1L2MzaExaWGYKLy9JUTlWYzMwVjJhcUtRbHh2VEtVc1NzMHFyVGltOTZ3UlhZbGdDQzgzeEk1cmJyUTlmNmFTRkNnZFJtN2tmSQpBaHM0bVEzMmpsVm9lU2xGVGZVQ0F3RUFBYU5aTUZjd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZNODl3c3lSbUpBZmdJRGFyMXdyWDVsMGRCcmhNQlVHQTFVZEVRUU8KTUF5Q0NtdDFZbVZ5Ym1WMFpYTXdEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBR3RhYXA0emNQMmU2NzRnL2lScApuZEdTb1puMG9UVW1sN1BpRmtLcjZ1T0dHRnF6Mk9oRENkQi9kQmNPTnpGUmZYdVFzejlobVpGYS9CTmNBUUtDCjZZV2M4WXd6a1BhclIvK2I3b3pKMy9hQ1RRRC9ySzZNbWZjd0NUaVVJR3VadkRVeHdHTzZGU3pQeWtTU1pPUkQKcVFpemxYb1BwZGszbVlndlFEbFI3ZlJ3Z0JpN0VKN1dzYWxKaFpmRDZCWVJtSGxHSlB0YXpsUHA0eDN2SHhnagpyR1ZIbW83am1RbCthdktheVRlZmZEejVJMHdMRU5SNHhoNGxISWN3WjRnOU5aY2hyMXNydVBwS3hqbmdId0RwCnNPQm5RL2xMVXNMN3pHT2l0TnJUOGpKODRFRWpsNjdKOU9wcE13c1ZCemFWaHVJYUZaUG9DK3VVMWplcmFacUgKbjBRPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    server: https://192.168.105.5:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURJVENDQWdtZ0F3SUJBZ0lJZkpHQ0wwYkIzOWN3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TVRFd016RXdNekV4TXpWYUZ3MHlNakV3TXpFd016RXhNemRhTURReApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1Sa3dGd1lEVlFRREV4QnJkV0psY201bGRHVnpMV0ZrCmJXbHVNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXZBZmpVaGk1a2dKT0tkUkgKdnh1VjdwVDF1YnZNS1A1Q0R6L0JjMGZ1cCtyUm5oTjZZdEhxMWNhRitDRGZwMVVNUHVKU3A5MmdiaUE3NXduWgpRZTg1M2U4Y05MckQzQWs3TVJyUDloOVFXK2FaY05PQWRYN05CMExiOUV6c2JuNDZ1NVR5eHF1SzdMZE1WUTNDCmFQY3c3S0RNVDkrNWZ1Q09vc1FicUNEU1kvTjJ2WUJHSit4ZHBSUjIrNkF6RzJqL1cyb0lCc3FUODRaWklGTGsKWVdMYjRubmJuMFJaMVl3TzJQamJLaE90cHFGNmRLZVhsdlNkaDlrKzlTOFRVVDJ6SFlXdmp2UTlVVUJMblJidQozYVB6SVl3b0t4N05TYWFLTitXMlZYdWNYTkJhV3pLOTM3SXJmNUZhYjM5RXBhV0IxT3dYekRxMFVpM2pWZGFLCm8yYXZCUUlEQVFBQm8xWXdWREFPQmdOVkhROEJBZjhFQkFNQ0JhQXdFd1lEVlIwbEJBd3dDZ1lJS3dZQkJRVUgKQXdJd0RBWURWUjBUQVFIL0JBSXdBREFmQmdOVkhTTUVHREFXZ0JUUFBjTE1rWmlRSDRDQTJxOWNLMStaZEhRYQo0VEFOQmdrcWhraUc5dzBCQVFzRkFBT0NBUUVBSVdPc25tdlRqK3JmTThLRktHWU9UNTJKN241SVJ3TW91WlczCndpYllLeGpkL3A2R3VTRHhUU0ttSmlhdmtWa3FFZmdobXg2a1krRDVHR0dyalZKTjVSbmc5NmNTeFdxS09UV0sKL2JYRkIvVzh3TkhvbkFNTUdjWHdvZzBCQks2cXR0Nm10Rkt3UE9rMUd6WU9HSHhoVjJzUHp1V2tmQlh1MkFJVwpiT0RhNy9QVGExY1B4ejhjVE1YZ3V1SFlWMGQ0QUNreTZBSEEwTW05b3poR0MwMk1Gb1E3VXJUcEVmb0ErbHhpCnZSMEhLNWxzSXV5ODZhY1JkbnhlWk1sWmYzVUlXb1A1SXozQ3U3NUcwWkFkRUxtSG5BRVNqaTBKaXNrMTMyTDAKTkhMQUZBYmVMcU0xWFVFdmNUekFFcWFEak1SSU54TTZkQ0lKaGRQZ2diMTZOenZxS1E9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBdkFmalVoaTVrZ0pPS2RSSHZ4dVY3cFQxdWJ2TUtQNUNEei9CYzBmdXArclJuaE42Cll0SHExY2FGK0NEZnAxVU1QdUpTcDkyZ2JpQTc1d25aUWU4NTNlOGNOTHJEM0FrN01SclA5aDlRVythWmNOT0EKZFg3TkIwTGI5RXpzYm40NnU1VHl4cXVLN0xkTVZRM0NhUGN3N0tETVQ5KzVmdUNPb3NRYnFDRFNZL04ydllCRwpKK3hkcFJSMis2QXpHMmovVzJvSUJzcVQ4NFpaSUZMa1lXTGI0bm5ibjBSWjFZd08yUGpiS2hPdHBxRjZkS2VYCmx2U2RoOWsrOVM4VFVUMnpIWVd2anZROVVVQkxuUmJ1M2FQeklZd29LeDdOU2FhS04rVzJWWHVjWE5CYVd6SzkKMzdJcmY1RmFiMzlFcGFXQjFPd1h6RHEwVWkzalZkYUtvMmF2QlFJREFRQUJBb0lCQUJtR3RKeVRrTFc2ckdQUApWc0loKzVPOUV6TFl4Tm1YMHQ0QkdNRU90ZDZENlZzZFo5TDhqblhoRXBUaDJacjgxOEc5dGR6bUlINXl6ODhJCnJqN2VQQ3ZkNXlWZGFYTjRxVmw4TzFrOVhRcHMzczNMV0xnYXM3alZvY2lqbk1GUFQ2ZmhpNmZlNStoRTVnN3gKNUQrZHR6Mllnc1FoS2l5SVRiVFpES3doY0k1eTVmQVVmT2lRN2hRMGsyalVHbmJKakNEa2xqMHgvdExtNTI0ZwpZNUQrZDBQNmsyNkZ4L3JjQTdIc1dRaFV4TGVMN1ROUUpJR1pqRGV6WEhnTExqNGt4T0prY1BZdFZRd3NuV2N4CnVGYUJVaVBGQW0wSmRJeDlWd1pzZGFKQVRCdW9iQXVlWlhvRVdGcGN6WER1YjVad2dyVXRTekdDcjVFclRkRjEKOXFxM0xrMENnWUVBOVBoMGVjMGllaEV2VmQ4SXFCOE4vZWtIdExwQTNnbFU5TVJFa0FGMDgwZ21QZlZFdStCQgpnbnJjZ2lCTmZZak1HWWFvZENXQnYzTmhLRldwVWZSa1BFSVM3TGo0TmxtMDJGVkNpWkVQV08yTlUvaEJKWVdVCk9zcFhiVnVrSHdneHVTQ2ttTUcrOXBpd0l5bHI1aDViNldRcjRTSHRwMWM1Zk04eGJWK1JLbzhDZ1lFQXhIOGsKZFNzc09CVVIvNm82K0NnUC9yTmRCdmZ2dG9XaTg5TVpJT3FQYjlyOFhMQ3ErcDYyRDFHWHlkY3F1TmNlTHZpNApLS2liUGdpQzV3VUwzSXRuc1MxTTJ4cG1aYlJQYVdOT0YrUXJuZittMVVwdldVWFF4eE56aGhtR0RLN0FwaDU4Cjl0MmQ5OW8yL3YvOGdnbnlPL0VjUjJOdlNnclhiUHBoRnVwTFp5c0NnWUVBeHJYWGZZbkp3cXkzOGZjV1JaSFQKUHJqTldHdVEvTXNqRDkwNUpIc2FDOUhsdTNKK0M4eFFOM0JEK0lZTmRad0d1MW9Bd1I0L2pqWWdoS0JmMWF2aAp0d045NkdudVplQUVrMHN0eXZ0Uk1MaVpZSmpLOTQzUWZKUHZzVFNaLzlZY3gyZlVQRmxFOXpGS0IydXRLNkJLCnZCcFdnYXVNNUtZYmFzT0MrNDUrNmI4Q2dZQU5jejZGbnF6cU1kY0hqeGwyUHhmSk1YSGV0ajM4V20xckR6b2IKSExNd3p1YVRXYXg2ZHo2ekVieTRIamhZYXJFd09lb3hMa2xpRjZjT05UUTRwbm8wa0l0QVBEQmlLZE45dlVSdApNanNpRjR2TjhjaHBiZlA2aklSTjI1Vm5iTmJYQ0NNNXFyWFRiMlp3VWdsd3hVbUZmTjJZclYvQ3k3Y3ZTSHc0CjIxVEovUUtCZ1FES3lVL0tBREhEYVVFMXhRdFBHcTZ3WGhPZysxRE91dkJPU2laVWhsQjVQUG1VaTJIQ0tyVlcKMEt2OTViZE5kNTc1eDFjU2VTbW1BcVlScWdrcVhJUnZHLzNwRmdtMXBpdkpmbGFYWngvVmI2NWRrQnNUSE5zSwp0ZDZRTWR3U0JEYlVGSDhLRmtPci8xWVZYL2lSNVAwaDVSWmRnQitmdUN2VEVKazRJVlJGNFE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
```


