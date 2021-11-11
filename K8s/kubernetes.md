



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



```











```shell
ubuntu@master:~$ microk8s kubectl get no
NAME     STATUS   ROLES    AGE    VERSION
master   Ready    <none>   87m    v1.22.2-3+9ad9ee77396805
node1    Ready    <none>   12m    v1.22.2-3+9ad9ee77396805
node2    Ready    <none>   100s   v1.22.2-3+9ad9ee77396805
ubuntu@master:~$ microk8s dashboard-proxy
Checking if Dashboard is running.
Dashboard will be available at https://127.0.0.1:10443
Use the following token to login:
eyJhbGciOiJSUzI1NiIsImtpZCI6ImRYdzJ2VHN4VXgxNElNXzNhSUYzRERMVVRpMGI2Z3hNdkVyaHpYdnJfQWcifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJkZWZhdWx0LXRva2VuLW44cGs5Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImRlZmF1bHQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiIzZjE1Nzg1OS0wYTA4LTQ1MTQtOWNlMi1kNTdkZmY3MTRmODciLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06ZGVmYXVsdCJ9.Vpd5VdWJcVcAlUJE9Pbk65EZ0YFYfBiSZulReFzRR22S_j8x1NO6FVjg387xhT8g_r8FI2TKo4aX-2DxoxqBd-qm0vgD2awlxFLJ6sKEIcpa0URS_E0C2UhtbZGrTMO1DIwgAmAEPSM3rorv8f9EM1QjBiKKIDjXyPulcJ5m8cxevDGHQa8njucEqAn_3LYXwewGOSAwdF7LTnQvxdRFf8f8LF5uS8njxjJtorNX5VPo1CcyOd-h1PZZhP6uw7bfOZ7IUhrUb0wxpLngztEDuAQmSSQJEz7zfOvP561vMopenqlTas9ZFB9kEPZzetdJyPJdXdPa7fOnjNToYqvTgA
E1101 23:58:41.848277   74122 portforward.go:385] error copying from local connection to remote stream: read tcp4 192.168.105.5:10443->192.168.105.1:63703: read: connection reset by peer



# 只有火狐浏览器可以访问，输入上面token
https://192.168.105.5:10443/




```








**nginx服务示例**

```shell

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




## 2、Namespace

名称空间用来隔离资源

```shell
ubuntu@master:~$ kubectl get ns
NAME              STATUS   AGE
kube-system       Active   18h
kube-public       Active   18h
kube-node-lease   Active   18h
default           Active   18h

ubuntu@master:~$ kubectl create ns hello
ubuntu@master:~$ kubectl get ns
NAME              STATUS   AGE
kube-system       Active   18h
kube-public       Active   18h
kube-node-lease   Active   18h
default           Active   18h
hello             Active   6m55s
ubuntu@master:~$ kubectl delete ns hello
namespace "hello" deleted


ubuntu@master:~$ vim hello.yaml
ubuntu@master:~$ cat hello.yaml
apiVersion: v1
kind: Namespace
metadata:
        name: hello
ubuntu@master:~$ kubectl apply -f hello.yaml
namespace/hello created
ubuntu@master:~$ kubectl get ns
NAME              STATUS   AGE
kube-system       Active   18h
kube-public       Active   18h
kube-node-lease   Active   18h
default           Active   18h
hello             Active   18s
ubuntu@master:~$ kubectl delete -f hello.yaml
namespace "hello" deleted
```


## 3、Pod

运行中的一组容器，Pod是kubernetes中应用的最小单位.

```shell

ubuntu@master:~$ kubectl run mynginx --image=nginx
pod/mynginx created
ubuntu@master:~$ kubectl get pod
NAME      READY   STATUS              RESTARTS   AGE
mynginx   0/1     ContainerCreating   0          15s

ubuntu@master:~$ kubectl get pod
NAME      READY   STATUS    RESTARTS   AGE
mynginx   1/1     Running   0          64s

ubuntu@master:~$ kubectl describe pod mynginx
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  104s  default-scheduler  Successfully assigned default/mynginx to node1
  Normal  Pulling    95s   kubelet            Pulling image "nginx"
  Normal  Pulled     47s   kubelet            Successfully pulled image "nginx" in 47.467731986s
  Normal  Created    45s   kubelet            Created container mynginx
  Normal  Started    45s   kubelet            Started container mynginx

ubuntu@master:~$ kubectl get pod -o wide
NAME      READY   STATUS    RESTARTS   AGE    IP             NODE    NOMINATED NODE   READINESS GATES
mynginx   1/1     Running   0          5m4s   10.1.166.129   node1   <none>           <none>


ubuntu@master:~$ kubectl delete pod mynginx
pod "mynginx" deleted

ubuntu@master:~$ vim nginx.yaml
ubuntu@master:~$ cat nginx.yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: mynginx
  name: mynginx
#  namespace: default
spec:
  containers:
  - image: nginx
    name: mynginx
ubuntu@master:~$ kubectl apply -f nginx.yaml
pod/mynginx created


ubuntu@master:~$ kubectl get pod -o wide
NAME      READY   STATUS    RESTARTS   AGE   IP             NODE    NOMINATED NODE   READINESS GATES
mynginx   1/1     Running   0          40s   10.1.166.130   node1   <none>           <none>


ubuntu@master:~$ kubectl delete -f nginx.yaml
pod "mynginx" deleted



ubuntu@node2:~$ kubectl logs mynginx

ubuntu@node2:~$ kubectl logs -f mynginx



ubuntu@master:~$ kubectl get pod -o wide
NAME      READY   STATUS    RESTARTS   AGE   IP           NODE    NOMINATED NODE   READINESS GATES
mynginx   1/1     Running   0          30m   10.1.104.1   node2   <none>           <none>
ubuntu@master:~$ curl 10.1.104.1
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





ubuntu@node2:~$ kubectl exec -it mynginx -- /bin/bash
root@mynginx:/# ls
bin   dev		   docker-entrypoint.sh  home  lib64  mnt  proc  run   srv  tmp  var
boot  docker-entrypoint.d  etc			 lib   media  opt  root  sbin  sys  usr
root@mynginx:/# cd /usr/share/nginx/html/
root@mynginx:/usr/share/nginx/html# ls
50x.html  index.html
root@mynginx:/usr/share/nginx/html# echo "01" > index.html
root@mynginx:/usr/share/nginx/html# exit
exit
ubuntu@node2:~$ kubectl get pod
NAME      READY   STATUS    RESTARTS   AGE
mynginx   1/1     Running   0          39m
ubuntu@node2:~$ kubectl get pod -owide
NAME      READY   STATUS    RESTARTS   AGE   IP           NODE    NOMINATED NODE   READINESS GATES
mynginx   1/1     Running   0          39m   10.1.104.1   node2   <none>           <none>
ubuntu@node2:~$ curl 10.1.104.1
01





ubuntu@master:~$ vim multicontainer-pod.yaml
ubuntu@master:~$ cat multicontainer-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: myapp
  name: myapp
spec:
  containers:
  - image: nginx
    name: nginx
  - image: tomcat:8.5.68
    name: tomcat
ubuntu@master:~$ kubectl apply -f multicontainer-pod.yaml
pod/myapp created
ubuntu@master:~$ kubectl get pod
NAME      READY   STATUS              RESTARTS   AGE
mynginx   1/1     Running             0          46m
myapp     0/2     ContainerCreating   0          9s


ubuntu@master:~$ kubectl delete pod --all
pod "mynginx" deleted
pod "myapp" deleted

```




## 4、Deployment

控制Pod，使Pod拥有多副本，自愈，扩缩容等能力

```shell
ubuntu@master:~$ kubectl create deployment mynginx --image=nginx
deployment.apps/mynginx created

ubuntu@master:~$ kubectl get pod
NAME                       READY   STATUS    RESTARTS   AGE
mynginx-5b686ccd46-xjx8p   1/1     Running   0          96s


# 启动新master 监听
ubuntu@master:~$ watch -n -1 kubectl get pod

Every 0.1s: kubectl get pod    master: Tue Nov  2 20:08:52 2021

NAME                       READY   STATUS    RESTARTS   AGE
mynginx-5b686ccd46-xjx8p   1/1     Running   0          2m3s


ubuntu@master:~$ kubectl delete pod mynginx-5b686ccd46-xjx8p
pod "mynginx-5b686ccd46-xjx8p" deleted


Every 0.1s: kubectl get pod    master: Tue Nov  2 20:09:39 2021

NAME                       READY   STATUS              RESTARTS
   AGE
mynginx-5b686ccd46-xjx8p   1/1     Terminating         0
   2m49s
mynginx-5b686ccd46-r9vf6   0/1     ContainerCreating   0
   3s


Every 0.1s: kubectl get pod    master: Tue Nov  2 20:10:01 2021

NAME                       READY   STATUS    RESTARTS   AGE
mynginx-5b686ccd46-r9vf6   1/1     Running   0          25s





ubuntu@master:~$ kubectl get pod
NAME                       READY   STATUS    RESTARTS   AGE
mynginx-5b686ccd46-r9vf6   1/1     Running   0          3m24s
ubuntu@master:~$ kubectl get deployments.apps
NAME      READY   UP-TO-DATE   AVAILABLE   AGE
mynginx   1/1     1            1           8m47s
ubuntu@master:~$ kubectl get deploy
NAME      READY   UP-TO-DATE   AVAILABLE   AGE
mynginx   1/1     1            1           9m3s

ubuntu@master:~$ kubectl delete deploy mynginx
deployment.apps "mynginx" deleted


 Every 0.1s: kubectl get pod    master: Tue Nov  2 20:15:55 2021

No resources found in default namespace.


```



**1、多副本**

```shell
ubuntu@master:~$ kubectl create deployment my-dep --image=nginx --replicas=3
deployment.apps/my-dep created
ubuntu@master:~$ kubectl get deploy
NAME     READY   UP-TO-DATE   AVAILABLE   AGE
my-dep   3/3     3            3           33s




# 启动新master 监听
ubuntu@master:~$ watch -n -1 kubectl get pod

Every 0.1s: kubectl get pod    master: Tue Nov  2 20:21:45 2021

NAME                      READY   STATUS    RESTARTS   AGE
my-dep-5b7868d854-6prsj   1/1     Running   0          2m5s
my-dep-5b7868d854-d4vtc   1/1     Running   0          2m5s
my-dep-5b7868d854-fjrnz   1/1     Running   0          2m5s


```


**访问Kubernetes仪表板**
```shell
microk8s dashboard-proxy
```

**2、扩缩容**
```shell
ubuntu@master:~$ kubectl scale deploy/my-dep --replicas=5 

ubuntu@master:~$ kubectl edit deploy my-dep

```

**3、自愈&故障转移**
● 停机
● 删除Pod
● 容器崩溃


```shell

➜  ~ multipass stop node1


ubuntu@master:~$ kubectl get pod -o wide
NAME                      READY   STATUS        RESTARTS   AGE     IP             NODE    NOMINATED NODE   READINESS GATES
my-dep-5b7868d854-5w8xq   1/1     Running       0          22m     10.1.104.14    node2   <none>           <none>
my-dep-5b7868d854-525xm   1/1     Running       0          11m     10.1.104.15    node2   <none>           <none>
my-dep-5b7868d854-f8tqv   1/1     Terminating   0          38m     10.1.166.142   node1   <none>           <none>
my-dep-5b7868d854-t8mtf   1/1     Running       0          3m28s   10.1.104.16    node2   <none>           <none>


ubuntu@master:~$ kubectl get pod -w
NAME                      READY   STATUS        RESTARTS   AGE
my-dep-5b7868d854-5w8xq   1/1     Running       0          20m
my-dep-5b7868d854-525xm   1/1     Running       0          9m10s
my-dep-5b7868d854-f8tqv   1/1     Terminating   0          35m
my-dep-5b7868d854-t8mtf   1/1     Running       0          69s

Every 0.1s: kubectl get pod            master: Tue Nov  2 23:18:44 2021

NAME                      READY   STATUS        RESTARTS   AGE
my-dep-5b7868d854-5w8xq   1/1     Running       0          24m
my-dep-5b7868d854-525xm   1/1     Running       0          12m
my-dep-5b7868d854-f8tqv   1/1     Terminating   0          39m
my-dep-5b7868d854-t8mtf   1/1     Running       0          4m52s



ubuntu@master:~$ kubectl get pod -o wide
NAME                      READY   STATUS    RESTARTS   AGE   IP            NODE    NOMINATED NODE   READINESS GATESmy-dep-5b7868d854-5w8xq   1/1     Running   0          35m   10.1.104.14   node2   <none>           <none>
my-dep-5b7868d854-525xm   1/1     Running   0          23m   10.1.104.15   node2   <none>           <none>
my-dep-5b7868d854-t8mtf   1/1     Running   0          15m   10.1.104.16   node2   <none>           <none>


```


**4、滚动更新**

```shell

ubuntu@master:~$ kubectl get deploy -oyaml

ubuntu@master:~$ watch -n 1 kubectl get pod

ubuntu@master:~$ kubectl get pod -w
NAME                      READY   STATUS    RESTARTS      AGE
my-dep-5b7868d854-t8mtf   1/1     Running   1 (36m ago)   7h14m
my-dep-5b7868d854-525xm   1/1     Running   1 (36m ago)   7h22m
my-dep-5b7868d854-5w8xq   1/1     Running   1 (36m ago)

ubuntu@master:~$ kubectl set image deploy/my-dep nginx=nginx:1.16.1 --record
Flag --record has been deprecated, --record will be removed in the future


ubuntu@master:~$ kubectl get pod -w
NAME                      READY   STATUS    RESTARTS      AGE
my-dep-5b7868d854-t8mtf   1/1     Running   1 (36m ago)   7h14m
my-dep-5b7868d854-525xm   1/1     Running   1 (36m ago)   7h22m
my-dep-5b7868d854-5w8xq   1/1     Running   1 (36m ago)   7h33m
my-dep-6b48cbf4f9-l66rj   0/1     Pending   0             0s
my-dep-6b48cbf4f9-l66rj   0/1     Pending   0             0s
my-dep-6b48cbf4f9-l66rj   0/1     ContainerCreating   0             0s
my-dep-6b48cbf4f9-l66rj   0/1     ContainerCreating   0



# 修改 kubectl edit deployment/my-dep

```


**5、版本回退**

```shell
ubuntu@master:~$ kubectl rollout history deployment/my-dep

deployment.apps/my-dep
REVISION  CHANGE-CAUSE
1         <none>
2         kubectl set image deploy/my-dep nginx=nginx:1.16.1 --record=true


ubuntu@master:~$ kubectl rollout undo deploy/my-dep --to-revision=1

deployment.apps/my-dep rolled back


ubuntu@master:~$ kubectl get deploy/my-dep -oyaml|grep image

      - image: nginx
        imagePullPolicy: Always



```
更多：
除了Deployment，k8s还有 StatefulSet 、DaemonSet 、Job  等 类型资源。我们都称为 工作负载。
有状态应用使用  StatefulSet  部署，无状态应用使用 Deployment 部署




## 5、Service

将一组 Pods 公开为网络服务的抽象方法。

```shell

ubuntu@master:~$ kubectl expose deployment my-dep --port=8000 --target-port=80
service/my-dep exposed

ubuntu@master:~$ kubectl get service
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
kubernetes   ClusterIP   10.152.183.1    <none>        443/TCP    32h
my-dep       ClusterIP   10.152.183.14   <none>        8000/TCP   41s




apiVersion: v1
kind: Service
metadata:
  labels:
    app: my-dep
  name: my-dep
spec:
  selector:
    app: my-dep
  ports:
  - port: 8000
    protocol: TCP
    targetPort: 80


#  集群内任意访问

ubuntu@master:~$ curl 10.152.183.14:8000
222
ubuntu@master:~$ curl 10.152.183.14:8000
333
ubuntu@master:~$ curl 10.152.183.14:8000
111




ubuntu@master:~$ kubectl get pod --show-labels
NAME                      READY   STATUS    RESTARTS   AGE   LABELS
my-dep-5b7868d854-h6nk4   1/1     Running   0          40m   app=my-dep,pod-template-hash=5b7868d854
my-dep-5b7868d854-nsf5s   1/1     Running   0          40m   app=my-dep,pod-template-hash=5b7868d854
my-dep-5b7868d854-blggv   1/1     Running   0          39m   app=my-dep,pod-template-hash=5b7868d854



```


**1、ClusterIP**

```shell

# 等同于没有--type的
kubectl expose deployment my-dep --port=8000 --target-port=80 --type=ClusterIP


apiVersion: v1
kind: Service
metadata:
  labels:
    app: my-dep
  name: my-dep
spec:
  ports:
  - port: 8000
    protocol: TCP
    targetPort: 80
  selector:
    app: my-dep
  type: ClusterIP
```


**2、NodePort 集群外也可以访问**

```shell

ubuntu@master:~$ kubectl get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
kubernetes   ClusterIP   10.152.183.1    <none>        443/TCP    33h
my-dep       ClusterIP   10.152.183.14   <none>        8000/TCP   61m
ubuntu@master:~$ kubectl get service
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
kubernetes   ClusterIP   10.152.183.1    <none>        443/TCP    33h
my-dep       ClusterIP   10.152.183.14   <none>        8000/TCP   61m
ubuntu@master:~$ kubectl delete service my-dep
service "my-dep" deleted

ubuntu@master:~$ kubectl expose deploy my-dep --port=8000 --target-port=80 --type=NodePort
service/my-dep exposed

ubuntu@master:~$ kubectl get svc
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kubernetes   ClusterIP   10.152.183.1    <none>        443/TCP          33h
my-dep       NodePort    10.152.183.39   <none>        8000:30689/TCP   37s



➜  ~ multipass ls
Name                    State             IPv4             Image
master                  Running           192.168.105.5    Ubuntu 20.04 LTS
                                          172.17.0.1
                                          10.1.219.64
node1                   Running           192.168.105.6    Ubuntu 20.04 LTS
                                          172.17.0.1
                                          10.1.166.128
node2                   Running           192.168.105.7    Ubuntu 20.04 LTS
                                          172.17.0.1
                                          10.1.104.0


# 外网访问
http://192.168.105.5:30689/
http://192.168.105.6:30689/
http://192.168.105.7:30689/

```


```shell
kubectl expose deployment my-dep --port=8000 --target-port=80 --type=NodePort

apiVersion: v1
kind: Service
metadata:
  labels:
    app: my-dep
  name: my-dep
spec:
  ports:
  - port: 8000
    protocol: TCP
    targetPort: 80
  selector:
    app: my-dep
  type: NodePort

```
NodePort范围在 30000-32767 之间






## 6、Ingress



```shell

ubuntu@master:~$ microk8s enable ingress



ubuntu@master:~$ microk8s.kubectl get pods --all-namespaces

ubuntu@master:~$ kubectl get pod -A

ingress       nginx-ingress-microk8s-controller-5zvc7      0/1     ImagePullBackOff   0               3m7s
ingress       nginx-ingress-microk8s-controller-xl7vh      0/1     ErrImagePull       0               3m7s
ingress       nginx-ingress-microk8s-controller-6hsvd      0/1     ErrImagePull       0               3m7s

Back-off pulling image "k8s.gcr.io/ingress-nginx/controller:v1.0.0-alpha.2"

ubuntu@master:~$ kubectl describe pods nginx-ingress-microk8s-controller -n ingress
Name:         nginx-ingress-microk8s-controller-5zvc7
Namespace:    ingress
Priority:     0
Node:         node2/192.168.105.7
Start Time:   Wed, 03 Nov 2021 09:10:06 +0800
Labels:       controller-revision-hash=7d5964757d
              name=nginx-ingress-microk8s
              pod-template-generation=1
Annotations:  cni.projectcalico.org/podIP: 10.1.104.23/32
              cni.projectcalico.org/podIPs: 10.1.104.23/32
Status:       Pending
IP:           10.1.104.23
IPs:
  IP:           10.1.104.23
Controlled By:  DaemonSet/nginx-ingress-microk8s-controller
Containers:
  nginx-ingress-microk8s:
    Container ID:
    Image:         k8s.gcr.io/ingress-nginx/controller:v1.0.0-alpha.2





registry.cn-hangzhou.aliyuncs.com/hfbpw/nginx-ingress-microk8s-controller:v1.0.0-alpha.2

registry.cn-hangzhou.aliyuncs.com/hfbpw/kube-webhook-certgen:v1.1.1


registry.cn-hangzhou.aliyuncs.com/hfbpw/nginx-ingress-microk8s-controller
registry.cn-hangzhou.aliyuncs.com/hfbpw/kube-webhook-certgen

registry-internal.cn-hangzhou.aliyuncs.com/hfbpw/nginx-ingress-microk8s-controller
registry-internal.cn-hangzhou.aliyuncs.com/hfbpw/kube-webhook-certgen



kubectl set image -n ingress /hfbpw/nginx-ingress-microk8s-controller nginx-ingress-microk8s-controller=registry.cn-hangzhou.aliyuncs.com/hfbpw/nginx-ingress-microk8s-controller:v1.0.0-alpha.2


kubectl set image -n ingress deployment/nginx-ingress-microk8s nginx-ingress-microk8s=registry.cn-hangzhou.aliyuncs.com/hfbpw/nginx-ingress-microk8s-controller



kubectl set image -n ingress deployment/kube-webhook-certgen kube-webhook-certgen=registry.cn-hangzhou.aliyuncs.com/hfbpw/kube-webhook-certgen







```



```shell

3. 将镜像推送到Registry





$ docker login --username=你那个面试咋样啊 registry.cn-hangzhou.aliyuncs.com

$ docker tag 2b2a6487f032 registry.cn-hangzhou.aliyuncs.com/hfbpw/nginx-ingress-microk8s-controller:v1.0.0-alpha.2
$ docker push registry.cn-hangzhou.aliyuncs.com/hfbpw/nginx-ingress-microk8s-controller:v1.0.0-alpha.2






kubectl set image -n ingress deployment/nginx-ingress-microk8s-controller nginx-ingress-microk8s-controller=registry.cn-hangzhou.aliyuncs.com/hfbpw/nginx-ingress-microk8s-controller:v1.0.0-alpha.2

 kubectl create deployment nginx-ingress-microk8s-controller  --image=xilepeng/nginx-ingress-microk8s-controller:v1.0.0-alpha.2 --namespace=ingress



ubuntu@master:~$  kubectl create deployment nginx-ingress-microk8s-controller  --image=xilepeng/nginx-ingress-microk8s-controller:v1.0.0-alpha.2 --namespace=ingress
deployment.apps/nginx-ingress-microk8s-controller created






kubectl set image -n ingress deployment/kube-webhook-certgen kube-webhook-certgen=registry.cn-hangzhou.aliyuncs.com/hfbpw/kube-webhook-certgen:v1.1.1



ubuntu@master:~$ kubectl create ns ingress

ubuntu@master:~$ kubectl create deployment nginx-ingress-microk8s-controller --image=wangshun1024/ingress-nginx-controller:v1.0.0-alpha.2 --namespace=ingress





ubuntu@master:~$ docker pull wangshun1024/ingress-nginx-controller:v1.0.0-alpha.2

docker tag wangshun1024/ingress-nginx-controller:v1.0.0-alpha.2 k8s.gcr.io/ingress-nginx/controller:v1.0.0-alpha.2




ubuntu@node2:~$ docker pull wangshun1024/kube-webhook-certgen:v1.1.1

k8s.gcr.io/ingress-nginx/kube-webhook-certgen:v1.1.1







$ docker tag c41e9fcadf5a registry.cn-hangzhou.aliyuncs.com/hfbpw/kube-webhook-certgen:v1.1.1
$ docker push registry.cn-hangzhou.aliyuncs.com/hfbpw/kube-webhook-certgen:v1.1.1






$ docker tag c41e9fcadf5a registry.cn-hangzhou.aliyuncs.com/hfbpw/kube-webhook-certgen:v1.1.1
$ docker push registry.cn-hangzhou.aliyuncs.com/hfbpw/kube-webhook-certgen:v1.1.1
```


microk8s.enable dashboard dns ingress istio registry storage


microk8s.stop && microk8s.start

sudo systemctl restart docker


microk8s.enable istio registry storage




















```shell

ubuntu@master:~$ kubectl create deployment web --image=nginx
deployment.apps/web created
ubuntu@master:~$ kubectl get deployments.apps
NAME   READY   UP-TO-DATE   AVAILABLE   AGE
web    1/1     1            1           32s
ubuntu@master:~$ kubectl expose deployment web --port=80 --target-port=80 --type=NodePort
service/web exposed

ubuntu@master:~$ kubectl get service
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
kubernetes   ClusterIP   10.152.183.1     <none>        443/TCP        22h
web          NodePort    10.152.183.249   <none>        80:32705/TCP   64s

浏览器访问：http://192.168.105.5:32705

Welcome to nginx!

If you see this page, the nginx web server is successfully installed and working. Further configuration is required.

For online documentation and support please refer to nginx.org.
Commercial support is available at nginx.com.

Thank you for using nginx.










```



















```shell
ubuntu@master:~$ microk8s helm3 repo add aliyuncs https://apphub.aliyuncs.com
WARNING: Kubernetes configuration file is group-readable. This is insecure. Location: /var/snap/microk8s/2551/credentials/client.config
"aliyuncs" has been added to your repositories


设置helm别名：

sudo snap alias microk8s.helm3 helm

ubuntu@master:~$ sudo snap alias microk8s.helm helm
Added:
  - microk8s.helm as helm

ubuntu@master:~$ sudo snap alias microk8s.helm3 helm



ubuntu@master:~$ helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
WARNING: Kubernetes configuration file is group-readable. This is insecure. Location: /var/snap/microk8s/2551/credentials/client.config
"ingress-nginx" has been added to your repositories
ubuntu@master:~$  helm repo update
WARNING: Kubernetes configuration file is group-readable. This is insecure. Location: /var/snap/microk8s/2551/credentials/client.config
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "ingress-nginx" chart repository
...Successfully got an update from the "aliyuncs" chart repository
Update Complete. ⎈Happy Helming!⎈


ubuntu@master:~$ helm search repo nginx-ingress
WARNING: Kubernetes configuration file is group-readable. This is insecure. Location: /var/snap/microk8s/2551/credentials/client.config
NAME                             	CHART VERSION	APP VERSION	DESCRIPTION
aliyuncs/nginx-ingress           	1.30.3       	0.28.0     	An nginx Ingress controller that uses ConfigMap...
aliyuncs/nginx-ingress-controller	5.3.4        	0.29.0     	Chart for the nginx Ingress controller
aliyuncs/nginx-lego              	0.3.1        	           	Chart for nginx-ingress-controller and kube-lego






```






