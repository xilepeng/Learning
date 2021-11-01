


```bash

一、Multipass介绍
        Multipass是一种简单的虚拟机工具。它不仅使启用虚拟机变得快速简易，还使管理那些虚拟机变得异常简单，因此可以立即开始针对云、边缘、物联网或任何一种类型的技术进行开发。实际上，Multipass包含一个系统任务栏工具，你只要点击一下就可以启动和停止虚拟机，甚至进入虚拟机的外壳。支持Linux、Windows系统等。
二、Multipass环境搭建
1.Ubuntu系统搭建
sudo apt update                   
sudo apt install snapd           #安装snapd
sudo snap install multipass      #安装multipass
 
2.Centos系统搭建
sudo yum install epel-release
sudo yum install snapd
sudo systemctl enable --now snapd.socket
sudo ln -s /var/lib/snapd/snap /snap
sudo snap install multipass
 
三、Multipass常用命令
1.查找镜像
multipass find
 
2.创建虚拟机
语法：multipass launch -n 虚拟机名称 
-n, --name: 名称
-c, --cpus: cpu核心数, 默认: 1
-m, --mem: 内存大小, 默认: 1G
-d, --disk: 硬盘大小, 默认: 5G

multipass launch -n ubuntu-lts -c 4 -m 4G -d 40G
 
3.进入虚拟机
语法：multipass shell 虚拟机名称
multipass shell ubuntu-lts
 
4.直接使用虚拟机
语法：multipass exec 虚拟机名称  --命令
multipass exec ubuntu-lts -- ls
 
5.查看虚拟机列表
multipass ls
multipass list
 
6.查看虚拟机信息
语法：multipass info 虚拟机名称
multipass info ubuntu-lts
 
7.重启虚拟机
语法：multipass restart 虚拟机名称
multipass restart ubuntu-lts
 
8.删除虚拟机
语法：multipass delete 虚拟机名称
--purge  彻底删除
multipass delete ubuntu-lts
multipass delete --purge ubuntu-lts 彻底删除
 
9.恢复删除虚拟机
语法：multipass recover 虚拟机名称
multipass recover ubuntu-lts
 
10.启动虚拟机
语法：multipass start 虚拟机名称
multipass start ubuntu-lts
 
11.暂停虚拟机
语法：multipass stop 虚拟机名称
multipass stop ubuntu-lts
 
12.宿主机挂载虚拟机
语法：multipass mount 宿主机目录 虚拟机名称：虚拟机目录
multipass mount /mnt ubuntu-python
multipass mount /mnt ubuntu-python:/mnt
 
13.宿主机卸载虚拟机
语法：multipass unmount 虚拟机名称
multipass unmount ubuntu-lts
 
14.挂起虚拟机
语法：multipass suspend 虚拟机名称
multipass suspend ubuntu-lts
 
15.获取版本信息
multipass version
 
16.帮助
multipass help


17.主机和容器数据交互

第一种 使用挂载数据卷的方式: multipass mount $HOME 容器名

# 或者指定容器目录
multipass mount $HOME 容器名:目录名

如果要卸载数据卷: multipass umount 容器名

第二种 transfer 进行文件复制传输

multipass transfer 主机文件 容器名:容器目录

➜  ~ multipass mount /Users/x/Shared x:/home/ubuntu/Shared

➜  ~ multipass mount /mnt ubuntu:/mnt
Source path "/mnt" does not exist
➜  ~ multipass mount /mnt ubuntu:/mnt
Source path "/mnt" does not exist
➜  ~ multipass info x
Name:           x
State:          Running
IPv4:           192.168.105.4
                172.17.0.1
                192.168.0.1
                172.38.0.1
Release:        Ubuntu 20.04.3 LTS
Image hash:     efd98e5a2269 (Ubuntu 20.04 LTS)
Load:           0.24 0.17 0.17
Disk usage:     6.6G out of 38.6G
Memory usage:   361.0M out of 3.8G
Mounts:         /Users/x/Shared => /home/ubuntu/Shared
                    UID map: 501:default
                    GID map: 20:default

18. multipass shell
进入一个与宿主机隔离的 Linux 容器！
multipass 会自动创建并运行一个名为 Primary 的容器（如果还没有创建或运行的话），这个容器也会自动挂载宿主机的 Home 目录，就是这么省心省力。

➜  ~ multipass shell
Launched: primary
Mounted '/Users/x' into 'primary:Home'
Welcome to Ubuntu 20.04.3 LTS (GNU/Linux 5.4.0-89-generic x86_64)

➜  ~ multipass list
Name                    State             IPv4             Image
primary                 Running           192.168.105.2    Ubuntu 20.04 LTS
x                       Running           192.168.105.4    Ubuntu 20.04 LTS

➜  ~ ping 192.168.105.2
PING 192.168.105.2 (192.168.105.2): 56 data bytes
64 bytes from 192.168.105.2: icmp_seq=0 ttl=64 time=0.457 ms
64 bytes from 192.168.105.2: icmp_seq=1 ttl=64 time=0.302 ms

➜  ~ ping 192.168.105.4
PING 192.168.105.4 (192.168.105.4): 56 data bytes
64 bytes from 192.168.105.4: icmp_seq=0 ttl=64 time=41.990 ms
64 bytes from 192.168.105.4: icmp_seq=1 ttl=64 time=0.866 ms

multipass shell x

```





**任何工作站的Ubuntu虚拟机**

使用单个命令获取即时Ubuntu VM。多通可以启动和运行虚拟机，并像公共云一样配置它们。您的云原型在本地免费启动。

[Multipass](https://multipass.run/)

**1. 在MacOS上安装 Multipass**

[Download Multipass for MacOS](https://multipass.run/download/macos)

```shell
brew install multipass
```



**2. 如何启动LTS实例**

使用Multipass的前五分钟让您知道随身携带轻量级云是多么容易。让我们启动几个LTS实例，列出它们，执行命令，使用云init并清理旧实例以开始。

启动实例（默认情况下，您将获得当前的Ubuntu LTS）

`multipass launch --name x`

在这种情况下运行命令，尝试运行bash（退出登录或ctrl-d退出）

`multipass exec x /bin/bash`

查看您的实例
`multipass list`

停止并启动实例

```bash
multipass stop x 
multipass start x

```

清理你不需要的东西
```
multipass delete master
multipass delete --all
multipass purge

```

查找使用多通启动的备用图像

`multipass find`

获取协助
`multipass help`






**安装完成，查看版本：**

```go
➜  ~ multipass version
multipass  1.7.2+mac
multipassd 1.7.2+mac
```


**创建 Ubuntu 虚拟机**

```go
➜  ~ multipass launch -n x -c 4 -m 4G -d 40G
Launched: x
```
-n, --name: 名称
-c, --cpus: cpu核心数, 默认: 1
-m, --mem: 内存大小, 默认: 1G
-d, --disk: 硬盘大小, 默认: 5G

# 如果需要修改配置文件

```bash
# 停用multipassd
sudo launchctl unload /Library/LaunchDaemons/com.canonical.multipassd.plist
# 修改配置文件
sudo vi '/var/root/Library/Application Support/multipassd/multipassd-vm-instances.json'
# 启动multipassd
sudo launchctl load /Library/LaunchDaemons/com.canonical.multipassd.plist
```









**帮助命令**

```
docker version

docker -v

docker info

docker --help
```

**镜像命令**

```bash
➜  ~ docker images --help

Usage:  docker images [OPTIONS] [REPOSITORY[:TAG]]

List images

Options:
  -a, --all             Show all images (default

  -q, --quiet           Only show image IDs


➜  ~ docker images
REPOSITORY                              TAG                                                     IMAGE ID       CREATED         SIZE
mysql                                   latest                                                  ecac195d15af   2 days ago      516MB


➜  ~ docker images -q
0deb7380d708


➜  ~ docker search redis
NAME                             DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
redis                            Redis is an open source key-value store that…   10055     [OK]


➜  ~ docker search redis -f=stars=500
NAME      DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
redis     Redis is an open source key-value store that…   10055     [OK]
```


**Mac 安装 Docker**
```shell
brew install --cask --appdir=/Applications docker

```


**snap 安装docker无法使用 swarm**
```shell
# 无法使用 swarm: mkdir /var/lib/docker: read-only file system
ubuntu@master:~$ sudo snap install docker
ubuntu@master:~$ sudo vim /var/snap/docker/1125/config/daemon.json

{
    "log-level":        "error",
    "storage-driver":   "overlay2",
    "registry-mirrors": ["https://hkaofvr0.mirror.aliyuncs.com"]

}


ubuntu@master:~$ sudo systemctl daemon-reload
ubuntu@master:~$ sudo snap restart docker
ubuntu@master:~$ sudo snap start docker
```



```shell
➜  ~ multipass list
Name                    State             IPv4             Image
master                  Running           192.168.105.5    Ubuntu 20.04 LTS
                                          172.17.0.1
node1                   Running           192.168.105.6    Ubuntu 20.04 LTS
                                          172.17.0.1
node2                   Running           192.168.105.7    Ubuntu 20.04 LTS
                                          172.17.0.1

# 在master添加hosts

ubuntu@master:~$ sudo vim /etc/host

192.168.105.5 master
192.168.105.6 node1
192.168.105.7 node2 


在所有节点添加
ubuntu@master:~$ sudo vim /etc/hosts

192.168.105.5 master
192.168.105.6 node1
192.168.105.7 node2 


ubuntu@node1:~$ sudo vim /etc/hosts
192.168.105.5 master
192.168.105.6 node1
192.168.105.7 node2 

ubuntu@node2:~$ sudo vim /etc/hosts
192.168.105.5 master
192.168.105.6 node1
192.168.105.7 node2 

# iTerm2多个窗口同时输入命令
打开这个功能的快捷键就是：
⌘(command) + ⇧(shift) + i  
会弹出告警信息，点OK确认。  关闭其实也很简单。再次输入刚刚打开的那个命令就行了。

# 关闭防火墙
ubuntu@master:~$ sudo apt-get install ufw

ubuntu@master:~$ sudo ufw disable
Firewall stopped and disabled on system startup
ubuntu@master:~$ sudo ufw status
Status: inactive

# 关闭selinux

ubuntu@master:~$ sudo vim /etc/selinux/config

SELINUX=disabled


# 关闭swap
ubuntu@master:~$ sudo sed -ri 's/.*swap.*/#&/' /etc/fstab


# 将桥接的IPv4流量传递到iptables的链
ubuntu@master:~$ sudo vim /etc/sysctl.d/k8s.conf
 
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1


# 生效
ubuntu@master:~$ sudo sysctl --system


# 时间同步
ubuntu@master:~$ sudo timedatectl set-timezone Asia/Shanghai


ubuntu@master:~$ sudo apt-get update
ubuntu@master:~$ sudo apt-get install virtualbox -y
```







**MicroK8s在几秒钟内安装一个节点，CNCF认证的Kubernetes集群**

MicroK8s是一款适用于Linux、Windows和macOS的轻量级零操作Kubernetes。单个命令将安装所有上游Kubernetes服务及其依赖项。通过支持x86和ARM64，MicroK8从本地工作站运行到边缘和物联网设备。

```shell
# 在Linux上安装MicroK8s
ubuntu@master:~$ sudo snap install microk8s --classic
# 将您的用户添加到microk8s管理组
# MicroK8s创建一个组，以无缝使用需要管理员权限的命令。使用以下命令加入群组：
ubuntu@master:~$ sudo usermod -a -G microk8s $USER
ubuntu@master:~$ sudo chown -f -R $USER ~/.kube
# 您还需要重新进入会话才能进行群组更新：
ubuntu@master:~$ su - $USER





# 在Kubernetes启动时检查状态
ubuntu@master:~$ microk8s status --wait-ready



ubuntu@master:~$ microk8s status
microk8s is not running. Use microk8s inspect for a deeper inspection.



# 1. 修改pod的sandbox
pod的sandbox 默认是 k8s.gcr.io/pause:3.1，这个镜像是无法获取的。需要将sandbox修改为国内可以获取的镜像。
ubuntu@master:~$ sudo vim /var/snap/microk8s/current/args/kubelet

--pod-infra-container-image=s7799653/pause:3.1

# 2. 配置 microk8s 内置 docker 的 registry.mirrors


ubuntu@master:~$ sudo vim /var/snap/microk8s/current/args/containerd.toml

    sandbox_image = "s7799653/pause:3.1"

     [plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
        endpoint = ["https://hkaofvr0.mirror.aliyuncs.com",
                "https://docker.mirrors.ustc.edu.cn",
                "https://hub-mirror.c.163.com",
                "https://mirror.ccs.tencentyun.com",
                "https://registry-1.docker.io" ]

        endpoint = ["https://hkaofvr0.mirror.aliyuncs.com",
                "https://docker.mirrors.ustc.edu.cn",
                "https://hub-mirror.c.163.com",
                "https://mirror.ccs.tencentyun.com",
                "https://registry-1.docker.io", ]


ubuntu@master:~$ sudo vim /var/snap/microk8s/current/args/containerd-template.toml

    sandbox_image = "s7799653/pause:3.1"
        endpoint = ["https://hkaofvr0.mirror.aliyuncs.com",
                "https://docker.mirrors.ustc.edu.cn",
                "https://hub-mirror.c.163.com",
                "https://mirror.ccs.tencentyun.com",
                "https://registry-1.docker.io", ]


ubuntu@master:~$ microk8s.stop&&microk8s.start

ubuntu@master:~$ microk8s stop
Stopped.
ubuntu@master:~$ microk8s start
Started.
ubuntu@master:~$ microk8s status
microk8s is running












# 打开你想要的服务
microk8s enable dashboard dns ingress
# 尝试microk8s enable --help列出可用服务和可选功能。microk8s disable ‹name›关闭服务。

# 开始使用Kubernetes
microk8s kubectl get all --all-namespaces
如果您主要使用MicroK8s，您可以将我们的kubectl作为命令行上的默认库布克特l，alias mkctl=”microk8s kubectl”由于它是标准的上游kubectl，您还可以通过“--kubeconfig”参数指向相应的kubeconfig文件来驱动其他Kubernetes集群。

# 访问Kubernetes仪表板
microk8s dashboard-proxy

# 启动和停止Kubernetes以节省电池

Kubernetes是一系列系统服务，一直在相互交谈。如果您不需要它们在后台运行，那么您将通过停止它们来节省电池。
microk8s start和microk8s stop将为您工作


多节点集群

Charmed Kubernetes跨云安装CNCF认证的Kubernetes集群

Charmed Kubernetes是一种完全自动化的模型驱动的方法，用于从裸金属到云安装和管理Kubernetes。从头开始构建Kubernetes云，将其与您最喜欢的工具集成，并创建多云拓扑。







```





**docker swarm**

```shell
ubuntu@master:~$ docker swarm init --advertise-addr=192.168.105.5

Swarm initialized: current node (qkdbdlpsu3ld9r8vwv3ngr4ey) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-27j725aoejjj5o2ejyh4b0ea1s3bk4y3eqg5k9ns7grscv541t-0zabcyioutjgybsw291lprzx9 192.168.105.5:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.



ubuntu@node1:~$ docker swarm join --token SWMTKN-1-27j725aoejjj5o2ejyh4b0ea1s3bk4y3eqg5k9ns7grscv541t-0zabcyioutjgybsw291lprzx9 192.168.105.5:2377
This node joined a swarm as a worker.

ubuntu@node2:~$ docker swarm join --token SWMTKN-1-27j725aoejjj5o2ejyh4b0ea1s3bk4y3eqg5k9ns7grscv541t-0zabcyioutjgybsw291lprzx9 192.168.105.5:2377
This node joined a swarm as a worker.

ubuntu@master:~$ docker node ls
ID                            HOSTNAME   STATUS    AVAILABILITY   MANAGER STATUS   ENGINE VERSION
qkdbdlpsu3ld9r8vwv3ngr4ey *   master     Ready     Active         Leader           20.10.8
svuvyka6s8lquq5gfmcqtwp0m     node1      Ready     Active                          20.10.8
mctvuj0sf77y7d52jwbhbbirt     node2      Ready     Active                          20.10.8


ubuntu@master:~$ docker service create --name demo busybox sh -c "while true;do sleep 3600;done"
plzhrzwk957zveodenc75o5rb
overall progress: 1 out of 1 tasks
1/1: running
verify: Service converged
ubuntu@master:~$ docker service ls

ID             NAME      MODE         REPLICAS   IMAGE            PORTS
plzhrzwk957z   demo      replicated   1/1        busybox:latest

ubuntu@master:~$ docker service ps demo
ID             NAME      IMAGE            NODE      DESIRED STATE   CURRENT STATE            ERROR     PORTS
xvgclxgnc1js   demo.1    busybox:latest   master    Running         Running 41 seconds ago


ubuntu@master:~$ docker ps
CONTAINER ID   IMAGE            COMMAND                  CREATED              STATUS              PORTS     NAMES
61bdd70e034d   busybox:latest   "sh -c 'while true;d…"   About a minute ago   Up About a minute             demo.1.xvgclxgnc1jszgmsus0ndscz7


ubuntu@master:~$ docker service scale demo-5
Invalid scale specifier 'demo-5'.
See 'docker service scale --help'.

Usage:  docker service scale SERVICE=REPLICAS [SERVICE=REPLICAS...]

Scale one or multiple replicated services
ubuntu@master:~$ docker service scale demo=5
demo scaled to 5
overall progress: 5 out of 5 tasks
1/5: running   [==================================================>]
2/5: running   [==================================================>]
3/5: running   [==================================================>]
4/5: running   [==================================================>]
5/5: running   [==================================================>]
verify: Service converged
ubuntu@master:~$ docker service ls
ID             NAME      MODE         REPLICAS   IMAGE            PORTS
plzhrzwk957z   demo      replicated   5/5        busybox:latest
ubuntu@master:~$ docker service ps demo
ID             NAME      IMAGE            NODE      DESIRED STATE   CURRENT STATE                ERROR     PORTS
xvgclxgnc1js   demo.1    busybox:latest   master    Running         Running 5 minutes ago
v5snefltyca7   demo.2    busybox:latest   node1     Running         Running 42 seconds ago
cb6sfs86mbs9   demo.3    busybox:latest   node2     Running         Running 25 seconds ago
nnl3634ts1d5   demo.4    busybox:latest   node2     Running         Running 25 seconds ago
jy9o6i5jr4x7   demo.5    busybox:latest   master    Running         Running about a minute ago

ubuntu@node1:~$ docker ps
CONTAINER ID   IMAGE            COMMAND                  CREATED         STATUS              PORTS     NAMES
dddeb7cb957f   busybox:latest   "sh -c 'while true;d…"   2 minutes ago   Up About a minute             demo.2.v5snefltyca7cpc5tln5f139t


ubuntu@node2:~$ docker ps
CONTAINER ID   IMAGE            COMMAND                  CREATED         STATUS              PORTS     NAMES
594b0aec3700   busybox:latest   "sh -c 'while true;d…"   2 minutes ago   Up About a minute             demo.4.nnl3634ts1d5kpuwaprmytsr9
6afbe3eeb282   busybox:latest   "sh -c 'while true;d…"   2 minutes ago   Up About a minute             demo.3.cb6sfs86mbs9o5caa5spkvytk

ubuntu@node2:~$ docker rm -f 6afbe3eeb282
6afbe3eeb282


ubuntu@master:~$ docker service ls
ID             NAME      MODE         REPLICAS   IMAGE            PORTS
plzhrzwk957z   demo      replicated   5/5        busybox:latest
ubuntu@master:~$ docker service ps demo
ID             NAME         IMAGE            NODE      DESIRED STATE   CURRENT STATE            ERROR                         PORTS
xvgclxgnc1js   demo.1       busybox:latest   master    Running         Running 9 minutes ago
v5snefltyca7   demo.2       busybox:latest   node1     Running         Running 4 minutes ago
h7wi2biidqtg   demo.3       busybox:latest   node1     Running         Running 29 seconds ago
cb6sfs86mbs9    \_ demo.3   busybox:latest   node2     Shutdown        Failed 38 seconds ago    "task: non-zero exit (137)"
nnl3634ts1d5   demo.4       busybox:latest   node2     Running         Running 4 minutes ago
jy9o6i5jr4x7   demo.5       busybox:latest   master    Running         Running 5 minutes ago



ubuntu@node1:~$ docker ps
CONTAINER ID   IMAGE            COMMAND                  CREATED              STATUS              PORTS     NAMES
c92419e48347   busybox:latest   "sh -c 'while true;d…"   About a minute ago   Up About a minute             demo.3.h7wi2biidqtghd4a0iityh04v
dddeb7cb957f   busybox:latest   "sh -c 'while true;d…"   6 minutes ago        Up 5 minutes                  demo.2.v5snefltyca7cpc5tln5f139t



ubuntu@master:~$ docker service rm demo
demo
ubuntu@master:~$ docker service ps demo
no such service: demo



```



**实战**

```shell
ubuntu@master:~$ sudo systemctl stop docker
Warning: Stopping docker.service, but it can still be activated by:
  docker.socket
  
ubuntu@node1:~$ docker node ls
Error response from daemon: This node is not a swarm manager. Worker nodes can\'t be used to view or modify cluster state. Please run this command on a manager node or promote the current node to a manager.

ubuntu@master:~$ sudo systemctl start docker


ubuntu@node2:~$ docker swarm leave
Node left the swarm.

ubuntu@master:~$ docker node ls
ID                            HOSTNAME   STATUS    AVAILABILITY   MANAGER STATUS   ENGINE VERSION
qn19qqxq66mxqbsti24f2mn3k *   master     Ready     Active         Leader           20.10.10
sfwf6nhluw5ohecrn7ufuqrxq     node1      Ready     Active                          20.10.10
r8yf610jdyn2lmnwh4tq1ygza     node2      Down      Active                          20.10.10

ubuntu@master:~$ docker swarm join-token manager
To add a manager to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-11h76824w8y8g2wk866zmqmnk681a20rekj9novh8oaqxxasiu-6565rved6au530fwkx77wjytn 192.168.105.5:2377
    
ubuntu@node2:~$ docker swarm join --token SWMTKN-1-11h76824w8y8g2wk866zmqmnk681a20rekj9novh8oaqxxasiu-6565rved6au530fwkx77wjytn 192.168.105.5:2377
This node joined a swarm as a manager.

ubuntu@node2:~$ docker node ls
ID                            HOSTNAME   STATUS    AVAILABILITY   MANAGER STATUS   ENGINE VERSION
qn19qqxq66mxqbsti24f2mn3k     master     Ready     Active         Leader           20.10.10
sfwf6nhluw5ohecrn7ufuqrxq     node1      Ready     Active                          20.10.10
r8yf610jdyn2lmnwh4tq1ygza     node2      Down      Active                          20.10.10
z8n0qyhc7isk0y11h3lfkii29 *   node2      Ready     Active         Reachable        20.10.10

# 集群可用，3个主节点，大于一台管理节点存活
# Raft协议：保证大多数节点存活，才可以使用，高可用
# 弹性、扩缩容

ubuntu@master:~$ docker service create -p 8888:80 --name my-nginx nginx

mzzf4c0bvvubhsmec0g5vgcnq
overall progress: 1 out of 1 tasks
1/1: running
verify: Service converged

ubuntu@master:~$ docker service ps my-nginx
ID             NAME         IMAGE          NODE      DESIRED STATE   CURRENT STATE         ERROR     PORTS
7g80iphcw4bh   my-nginx.1   nginx:latest   node1     Running         Running 2 hours ago

ubuntu@master:~$ docker service ls
ID             NAME       MODE         REPLICAS   IMAGE                   PORTS
mzzf4c0bvvub   my-nginx   replicated   1/1        nginx:latest            *:8888->80/tcp


ubuntu@master:~$ docker service update --replicas 3 my-nginx
my-nginx
overall progress: 3 out of 3 tasks
1/3: running
2/3: running
3/3: running
verify: Service converged
ubuntu@master:~$ docker service ps my-nginx
ID             NAME         IMAGE          NODE      DESIRED STATE   CURRENT STATE            ERROR     PORTS
7g80iphcw4bh   my-nginx.1   nginx:latest   node1     Running         Running 2 hours ago
l4fn8a3ubibw   my-nginx.2   nginx:latest   node2     Running         Running 28 seconds ago
tk2ewthlrpdh   my-nginx.3   nginx:latest   master    Running         Running 30 seconds ago



http://192.168.105.6:8888

Welcome to nginx!

If you see this page, the nginx web server is successfully installed and working. Further configuration is required.

For online documentation and support please refer to nginx.org.
Commercial support is available at nginx.com.

Thank you for using nginx.

ubuntu@master:~$ docker service scale my-nginx=5
my-nginx scaled to 5
overall progress: 5 out of 5 tasks
1/5: running
2/5: running
3/5: running
4/5: running
5/5: running
verify: Service converged
ubuntu@master:~$ docker service ps my-nginx
ID             NAME         IMAGE          NODE      DESIRED STATE   CURRENT STATE            ERROR     PORTS
7g80iphcw4bh   my-nginx.1   nginx:latest   node1     Running         Running 2 hours ago
lsgi6ipstv28   my-nginx.2   nginx:latest   master    Running         Running 10 seconds ago
si2xexb5ct94   my-nginx.3   nginx:latest   node2     Running         Running 9 seconds ago
5ngt0ses5hyk   my-nginx.4   nginx:latest   node2     Running         Running 8 seconds ago
qzz62rlmz93k   my-nginx.5   nginx:latest   node1     Running         Running 11 seconds ago

ubuntu@master:~$ docker service rm my-nginx
my-nginx
ubuntu@master:~$ docker service ls
ID             NAME      MODE         REPLICAS   IMAGE                   PORTS


```












```bash

➜  ~ docker pull redis
Using default tag: latest  
latest: Pulling from library/redis
7d63c13d9b9b: Pull complete # 分层下载，docker image的核心，联合文件系统
a2c3b174c5ad: Pull complete
283a10257b0f: Pull complete
7a08c63a873a: Pull complete
0531663a7f55: Pull complete
9bf50efb265c: Pull complete
Digest: sha256:a89cb097693dd354de598d279c304a1c73ee550fbfff6d9ee515568e0c749cfe
Status: Downloaded newer image for redis:latest
docker.io/library/redis:latest  #真实地址


# 等价
docker pull redis
docker pull docker.io/library/redis:latest


➜  ~ docker pull mysql:5.7
5.7: Pulling from library/mysql
b380bbd43752: Already exists
f23cbf2ecc5d: Already exists
30cfc6c29c0a: Already exists
b38609286cbe: Already exists
8211d9e66cd6: Already exists
2313f9eeca4a: Already exists
7eb487d00da0: Already exists
a71aacf913e7: Pull complete
393153c555df: Pull complete
06628e2290d7: Pull complete
ff2ab8dac9ac: Pull complete
Digest: sha256:2db8bfd2656b51ded5d938abcded8d32ec6181a9eae8dfc7ddf87a656ef97e97
Status: Downloaded newer image for mysql:5.7
docker.io/library/mysql:5.7


# 删除 images id=938b57d64674 

➜  ~ docker rmi -f 938b57d64674
Untagged: mysql:5.7
Untagged: mysql@sha256:2db8bfd2656b51ded5d938abcded8d32ec6181a9eae8dfc7ddf87a656ef97e97
Deleted: sha256:938b57d64674c4a123bf8bed384e5e057be77db934303b3023d9be331398b761
Deleted: sha256:d81fc74bcfc422d67d8507aa0688160bc4ca6515e0a1c8edcdb54f89a0376ff1
Deleted: sha256:a6a530ba6d8591630a1325b53ef2404b8ab593a0775441b716ac4175c14463e6
Deleted: sha256:2a503984330e2cec317bc2ef793f5d4d7b3fd8d50009a4f673026c3195460200
Deleted: sha256:e2a4585c625da1cf4909cdf89b8433dd89ed5c90ebdb3a979d068b161513de90

# 全部删除
➜  ~ docker rmi -f $(docker images -aq)
```


**容器命令**

```bash
➜  ~ docker pull ubuntu

# 启动并进入容器
➜  ~ docker run -it ubuntu /bin/bash
root@658d9ac2bb23:/#

root@658d9ac2bb23:/# exit
exit
➜  ~

# 查看运行的容器

➜  ~ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED       STATUS       PORTS     NAMES
440e5b6b3a30   ubuntu    "bash"    9 hours ago   Up 2 hours             ubuntu

# 曾经运行过的

➜  ~ docker ps -a
CONTAINER ID   IMAGE          COMMAND                  CREATED         STATUS                          PORTS                                        NAMES
658d9ac2bb23   ubuntu         "/bin/bash"              3 minutes ago   Exited (0) About a minute ago                                                strange_pascal
440e5b6b3a30   ubuntu         "bash"                   9 hours ago     Up 2 hours                                                                   ubuntu
c96cefc26fb1   10d7504ea271   "/whoami"                20 hours ago    Exited (255) 4 hours ago        80/tcp                                       demo-whoami-1
937064d706ca   traefik:v2.5   "/entrypoint.sh --ap…"   20 hours ago    Exited (255) 4 hours ago        0.0.0.0:80->80/tcp, 0.0.0.0:8080->8080/tcp   demo-reverse-proxy-1
073ed0c2fc09   mysql          "docker-entrypoint.s…"   43 hours ago    Exited (255) 4 hours ago        0.0.0.0:3306->3306/tcp, 33060/tcp            mysql


➜  ~ docker ps -n=1
CONTAINER ID   IMAGE     COMMAND       CREATED          STATUS                      PORTS     NAMES
658d9ac2bb23   ubuntu    "/bin/bash"   16 minutes ago   Exited (0) 14 minutes ago             strange_pascal

➜  ~ docker ps -aq
658d9ac2bb23
440e5b6b3a30
c96cefc26fb1
937064d706ca
073ed0c2fc09


# 停止容器并退出
➜  ~ docker run -it ubuntu /bin/bash
root@0e8a4299ceb0:/# exit
exit
➜  ~ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES

# contrl + P + Q 容器不停止退出


```


**删除容器**

```bash
docker rm 容器id


# 删除所有容器

docker rm -f $(docker ps -aq)

docker ps -a -q | xargs docker rm 
```


**启动停止容器**

```
docker start 容器id

docker stop 容器id

docker kill id


```



**常用命令**

**后台启动容器**
```bash
➜  Learning git:(main) ✗ docker run -d ubuntu
fef5ab64b693ad7084987563c845feae76415d524bf422a85d1e3c5fd6523b81
➜  Learning git:(main) ✗ docker ps               
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES

# docker 发现没有提供服务，立即停止
```

**查看日志**

```bash
➜  ~ docker run -d ubuntu /bin/sh -c "while true;do echo mojoman;sleep 1;done"
410f051aa755c818d8fd65f4a170fa69fbd87ddfae05ad9448df41f9d5e00ab6
➜  ~ docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS     NAMES
410f051aa755   ubuntu    "/bin/sh -c 'while t…"   4 seconds ago   Up 3 seconds             charming_swartz
3ede536d2906   ubuntu    "/bin/bash"              7 minutes ago   Up 7 minutes             dreamy_jang
➜  ~ docker logs -tf --tail 10 410f051aa755
2021-10-21T11:07:44.171459600Z mojoman
2021-10-21T11:07:45.175328700Z mojoman
2021-10-21T11:07:46.177583400Z mojoman
2021-10-21T11:07:47.180573300Z mojoman
2021-10-21T11:07:48.184208500Z mojoman
2021-10-21T11:07:49.190032200Z mojoman
2021-10-21T11:07:50.196649000Z mojoman
```

**查看容器中进程信息**
```bash
➜  ~ docker top 410f051aa755
UID                 PID                 PPID                C                   STIME               TTY                 TIME                CMD
root                22945               22919               0                   11:07               ?                   00:00:01            /bin/sh -c while true;do echo mojoman;sleep 1;done
root                50997               22945               0                   11:53               ?                   00:00:00            sleep 1
```

**查看镜像元数据**
```bash
➜  ~ docker inspect 410f051aa755
[
    {
        "Id": "410f051aa755c818d8fd65f4a170fa69fbd87ddfae05ad9448df41f9d5e00ab6",
        "Created": "2021-10-21T11:07:21.2302738Z",
        "Path": "/bin/sh",
        "Args": [
            "-c",
            "while true;do echo mojoman;sleep 1;done"
        ],
        "State": {
            "Status": "running",
            "Running": true,
            ...
]
```

**进入当前正在运行的容器**
```bash
➜  ~ docker start fd68d763d455
fd68d763d455
➜  ~ docker ps
CONTAINER ID   IMAGE     COMMAND       CREATED          STATUS         PORTS     NAMES
fd68d763d455   ubuntu    "/bin/bash"   58 seconds ago   Up 3 seconds             lucid_newton


# 进入容器打开新终端
➜  ~ docker exec -it fd68d763d455 /bin/bash
root@fd68d763d455:/#


# 进入容器正在执行的终端，不会打开新终端
➜  ~ docker attach fd68d763d455
root@fd68d763d455:/#

```

**从容器内拷贝文件到主机**
```bash
root@625456221cb4:/home# touch x.go
root@625456221cb4:/home# ls
x.go
root@625456221cb4:/home# exit
exit

➜  demo docker cp 625456221cb4:/home/x.go /Users/X/demo
➜  demo ls
docker-compose.yml mojo.go            x.go
```


**部署 nginx**

暴露端口3344 `-p 3344:80`

```shell
➜  / docker pull nginx
Using default tag: latest
latest: Pulling from library/nginx
b380bbd43752: Already exists
fca7e12d1754: Pull complete
745ab57616cb: Pull complete
a4723e260b6f: Pull complete
1c84ebdff681: Pull complete
858292fd2e56: Pull complete
Digest: sha256:644a70516a26004c97d0d85c7fe1d0c3a67ea8ab7ddf4aff193d9f301670cf36
Status: Downloaded newer image for nginx:latest
docker.io/library/nginx:latest


➜  / docker run -d -p 3344:80 nginx
de55b21c29a68f0591ffea5bf760cb08bf2eb40aa54e95ba45709bf558d8f0cb

➜  / docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED         STATUS         PORTS                  NAMES
de55b21c29a6   nginx     "/docker-entrypoint.…"   7 seconds ago   Up 5 seconds   0.0.0.0:3344->80/tcp   pedantic_khorana

➜  / curl  0.0.0.0:3344


➜  / docker exec -it de55b21c29a6 /bin/bash
root@de55b21c29a6:/# whereis nginx
nginx: /usr/sbin/nginx /usr/lib/nginx /etc/nginx /usr/share/nginx
root@de55b21c29a6:/# /etc/nginx
bash: /etc/nginx: Is a directory
root@de55b21c29a6:/# ls
bin   dev		   docker-entrypoint.sh  home  lib64  mnt  proc  run   srv  tmp  var
boot  docker-entrypoint.d  etc			 lib   media  opt  root  sbin  sys  usr
root@de55b21c29a6:/# cd /etc/nginx
root@de55b21c29a6:/etc/nginx# ls
conf.d	fastcgi_params	mime.types  modules  nginx.conf  scgi_params  uwsgi_params

root@de55b21c29a6:/etc/nginx# exit
exit
➜  / docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED        STATUS        PORTS                  NAMES
de55b21c29a6   nginx     "/docker-entrypoint.…"   11 hours ago   Up 11 hours   0.0.0.0:3344->80/tcp   pedantic_khorana

➜  / docker stop de55b21c29a6
de55b21c29a6

➜  / curl  0.0.0.0:3344
curl: (7) Failed to connect to 0.0.0.0 port 3344: Connection refused
```

docker web可视化管理工具

安装portainer 
```shell

docker run -d -p 8088:9000 --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v --privileged=true portainer/portainer

```


**commit 镜像**

```shell
➜  ~ docker commit --help

Usage:  docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]

Create a new image from a container's changes

Options:
  -a, --author string    Author (e.g., "John Hannibal
                         Smith <hannibal@a-team.com>")
  -c, --change list      Apply Dockerfile instruction
                         to the created image
  -m, --message string   Commit message
  -p, --pause            Pause container during
                         commit (default true)
```

```shell

```



**已入门 Docker**
------


**容器数据卷**

使用数据卷

方式一：直接使用命令来挂载 -v

**双向同步数据**
好处：只需本地修改、容器内会自动同步
```shell
docker run it -v 主机目录：容器目录


➜  ~ docker run -it -v /Users/Shared/test:/home ubuntu /bin/bash
root@d84239a1cc5a:/# ls
bin   dev  home  lib32  libx32  mnt  proc  run   srv  tmp  var
boot  etc  lib   lib64  media   opt  root  sbin  sys  usr
root@d84239a1cc5a:/# cd home
root@d84239a1cc5a:/home# ls
root@d84239a1cc5a:/home#

➜  ~ docker inspect d84239a1cc5a

        "Mounts": [
            {
                "Type": "bind",
                "Source": "/Users/Shared/test",
                "Destination": "/home",
                "Mode": "",
                "RW": true,
                "Propagation": "rprivate"
            }
        ],
root@d84239a1cc5a:/home# touch main.go

➜  test ls
main.go




root@d84239a1cc5a:/home# exit
exit
➜  ~ docker ps
CONTAINER ID   IMAGE     COMMAND                  CREATED      STATUS       PORTS                               NAMES


➜  test vim main.go
➜  test cat main.go
package main

func mian(){

}


➜  ~ docker run -it -v /Users/Shared/test:/home ubuntu /bin/bash
root@f310a6debc99:/# cd home
root@f310a6debc99:/home# ls
main.go
root@f310a6debc99:/home# bat main.go
bash: bat: command not found
root@f310a6debc99:/home# cat main.go
package main

func mian(){

}


```





**安装mysql同步数据**

```sql
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:tag

docker exec -it some-mysql bash

mysql -uroot -p123456
```

```sql
➜  ~ docker pull mysql

➜  ~ docker images
REPOSITORY            TAG       IMAGE ID       CREATED        SIZE
mysql                 latest    ecac195d15af   3 days ago     516MB


➜  ~ docker run -d -p 3310:3306 -v /Users/Shared/mysql/conf:/etc/mysql/conf.d -v /Users/Shared/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 --name mysql01 mysql:latest
ac84769de66024ca380c4336b0e77dba130e7463d73d414afa742a63f444d1de
 
➜  ~ docker exec -it mysql01 bash
root@ac84769de660:/# mysql -uroot -p123456
mysql> create database xn;
Query OK, 1 row affected (0.04 sec)


➜  ~ /Users/Shared/mysql/data
➜  data ls
#ib_16384_0.dblwr  binlog.index       ib_logfile0        performance_schema undo_001
#ib_16384_1.dblwr  ca-key.pem         ib_logfile1        private_key.pem    undo_002

➜  data ls
 xn


➜  ~ docker rm -f mysql01
mysql01
➜  ~ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES

```
数据已同步到本地，将mysql删除、本地数据不会丢失，这就实现**容器数据持久化**


**匿名卷挂载**

```shell
➜  ~ docker run -d -P --name nginx01 -v /ect/nginx nginx
5b6c7c9b76e0ecd138472c6e3c488e4f76722671f4546c8842cbbce48ac4fd2a
➜  ~ docker volume ls
DRIVER    VOLUME NAME
local     3a740520a550c70da6f631c9e2c3a591357d9c69798e336a741519cd21dcc842
```

**具名卷挂载**

```shell
➜  ~ docker run -d -P --name nginx02 -v juming-nginx:/ect/nginx nginx
549c7be291dd94965425c2400797404ae02e215e9e069baa20b76b6c06c307fe

➜  ~ docker volume ls
DRIVER    VOLUME NAME
local     juming-nginx



➜  ~ docker volume inspect juming-nginx
[
    {
        "CreatedAt": "2021-10-22T08:14:44Z",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/juming-nginx/_data",
        "Name": "juming-nginx",
        "Options": null,
        "Scope": "local"
    }
]
```




```
docker run -itd --name ubuntu-test ubuntu

docker exec -it ubuntu-test /bin/bash


➜  ~ docker run -itd --name ubuntu-test ubuntu
a19189acd61de8300ff2731c5f27af084dfaa53015429fa6c61b50a4aa82ffe6

➜  ~ docker exec -it ubuntu-test /bin/bash
root@a19189acd61d:/#

```








```dockerfile
ubuntu@x:~/docker-volume$ pwd
/home/ubuntu/docker-volume

ubuntu@x:~/docker-volume$ ls
dockerfile1

ubuntu@x:~/docker-volume$ vim dockerfile1

ubuntu@x:~/docker-volume$ cat dockerfile1

FROM ubuntu

VOLUME ["volume01","volume02"]

CMD echo "end ..."
CMD /bin/bash



ubuntu@x:~/docker-volume$ docker build -f dockerfile1 -t x/ubuntu .
Sending build context to Docker daemon  2.048kB
Step 1/4 : FROM ubuntu
latest: Pulling from library/ubuntu
7b1a6ab2e44d: Pull complete
Digest: sha256:626ffe58f6e7566e00254b638eb7e0f3b11d4da9675088f4781a50ae288f3322
Status: Downloaded newer image for ubuntu:latest
 ---> ba6acccedd29
Step 2/4 : VOLUME ["volume01","volume02"]
 ---> Running in de9067bf26c6
Removing intermediate container de9067bf26c6
 ---> 4788d0658835
Step 3/4 : CMD echo "end ..."
 ---> Running in 29cd8dfb21f9
Removing intermediate container 29cd8dfb21f9
 ---> 45efc4c25149
Step 4/4 : CMD /bin/bash
 ---> Running in 677b5764a520
Removing intermediate container 677b5764a520
 ---> 220e6a6dea5d
Successfully built 220e6a6dea5d
Successfully tagged x/ubuntu:latest


ubuntu@x:~/docker-volume$ docker images
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
x/ubuntu     latest    220e6a6dea5d   4 minutes ago   72.8MB
mysql        latest    ecac195d15af   4 days ago      516MB
ubuntu       latest    ba6acccedd29   7 days ago      72.8MB



ubuntu@x:~/docker-volume$ docker build -f /home/ubuntu/docker-volume/dockerfile1 -t x/ubuntu:1.0 .
Sending build context to Docker daemon  2.048kB
Step 1/4 : FROM ubuntu
 ---> ba6acccedd29
Step 2/4 : VOLUME ["volume01","volume02"]
 ---> Using cache
 ---> 4788d0658835
Step 3/4 : CMD echo "end ..."
 ---> Using cache
 ---> 45efc4c25149
Step 4/4 : CMD /bin/bash
 ---> Using cache
 ---> 220e6a6dea5d
Successfully built 220e6a6dea5d
Successfully tagged x/ubuntu:1.0
ubuntu@x:~/docker-volume$ docker images
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
x/ubuntu     1.0       220e6a6dea5d   6 minutes ago   72.8MB
x/ubuntu     latest    220e6a6dea5d   6 minutes ago   72.8MB
mysql        latest    ecac195d15af   4 days ago      516MB
ubuntu       latest    ba6acccedd29   7 days ago      72.8MB

```


**启动自己镜像**

```shell
ubuntu@x:~/docker-volume$ docker images
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
x/ubuntu     1.0       220e6a6dea5d   6 minutes ago   72.8MB
x/ubuntu     latest    220e6a6dea5d   6 minutes ago   72.8MB
mysql        latest    ecac195d15af   4 days ago      516MB
ubuntu       latest    ba6acccedd29   7 days ago      72.8MB
ubuntu@x:~/docker-volume$ docker run -it 220e6a6dea5d /bin/bash
root@263a442288cb:/# ls
bin   etc   lib32   media  proc  sbin  tmp  volume01
boot  home  lib64   mnt    root  srv   usr  volume02
dev   lib   libx32  opt    run   sys   var

 
```


## 找到volume目录


```

ubuntu@x:~/docker-volume$ docker inspect 565a3b003ace


        "Mounts": [
            {
                "Type": "volume",
                "Name": "4dcd24a360cd377189c0fbc454c09736548d7be54673b5573ce6851bb71ceff9",
                "Source": "/var/lib/docker/volumes/4dcd24a360cd377189c0fbc454c09736548d7be54673b5573ce6851bb71ceff9/_data",
                "Destination": "volume02",
                "Driver": "local",
                "Mode": "",
                "RW": true,
                "Propagation": ""
            },
            {
                "Type": "volume",
                "Name": "a3f1451a671f9ce8b5be5388e7e068963bc30cebd2c2617472b2d42cff401e3d",
                "Source": "/var/lib/docker/volumes/a3f1451a671f9ce8b5be5388e7e068963bc30cebd2c2617472b2d42cff401e3d/_data",
                "Destination": "volume01",
                "Driver": "local",
                "Mode": "",
                "RW": true,
                "Propagation": ""
            }
        ],

```

**查看数据同步成功**
```bash
ubuntu@x:~$ cd /var/lib/docker/volumes/a3f1451a671f9ce8b5be5388e7e068963bc30cebd2c2617472b2d42cff401e3d/_data
ubuntu@x:/var/lib/docker/volumes/a3f1451a671f9ce8b5be5388e7e068963bc30cebd2c2617472b2d42cff401e3d/_data$ ls
main.go
```






**容器间数据共享卷：数据双向拷贝**

`docker run -it --name docker02 --volumes-from docker01 x/ubuntu`


```
ubuntu@x:~$ docker images
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
x/ubuntu     latest    25d88dec65e5   3 seconds ago   72.8MB

ubuntu@x:~$ docker images
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
x/ubuntu     latest    25d88dec65e5   3 seconds ago   72.8MB
mysql        latest    ecac195d15af   4 days ago      516MB
ubuntu       latest    ba6acccedd29   7 days ago      72.8MB

ubuntu@x:~$ docker run -it --name docker01 x/ubuntu

root@356ee34cc6ca:/# cd volume01
root@356ee34cc6ca:/volume01# ls
root@356ee34cc6ca:/volume01# touch docker01
root@356ee34cc6ca:/volume01# ls
docker01



➜  ~ multipass exec x /bin/bash
ubuntu@x:~$ docker ps
CONTAINER ID   IMAGE      COMMAND                  CREATED         STATUS         PORTS     NAMES
356ee34cc6ca   x/ubuntu   "/bin/sh -c /bin/bash"   2 minutes ago   Up 2 minutes             docker01


ubuntu@x:~$ docker run -it --name docker02 --volumes-from docker01 x/ubuntu

root@79924290a42f:/# cd volume01
root@79924290a42f:/volume01# ls
docker01



ubuntu@x:~$ docker run -it --name docker03 --volumes-from docker01 x/ubuntu
root@5445e0fa4ffc:/# cd volume01
root@5445e0fa4ffc:/volume01# ls
docker01



删除 docker01

ubuntu@x:~$ docker rm -f 356ee34cc6ca
356ee34cc6ca
ubuntu@x:~$ docker ps -a
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS                   PORTS     NAMES
5445e0fa4ffc   x/ubuntu       "/bin/sh -c /bin/bash"   5 minutes ago    Up 5 minutes                       docker03
79924290a42f   x/ubuntu       "/bin/sh -c /bin/bash"   16 minutes ago   Up 16 minutes                      docker02


# docker02容器同步文件依然在
ubuntu@x:~$ docker run -it --name docker02 --volumes-from docker01 x/ubuntu

root@79924290a42f:/volume01# ls
docker01  docker03
```



**数据卷容器的生命周期一直持续到没有容器使用为止**





```bash
ubuntu@x:~$ docker run -d -p 3306:3306 -v /etc/mysql/conf.d -v /var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 --name mysql mysql:latest
c9cbb8634a619d014ccbcde50c79fb899b7d0404bece3b644de2ebf03aa2a377

ubuntu@x:~$ docker run -d -p 3307:3306 -v /etc/mysql/conf.d -v /var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 --name mysql02 --volumes-from mysql mysql:latest
a8f5e85ae1d71397881b4494599aaf91ac08dce9b8d5fd4ca096e5805686ca8c

ubuntu@x:~$ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS          PORTS                                                  NAMES
a8f5e85ae1d7   mysql:latest   "docker-entrypoint.s…"   58 seconds ago   Up 56 seconds   33060/tcp, 0.0.0.0:3307->3306/tcp, :::3307->3306/tcp   mysql02
c9cbb8634a61   mysql:latest   "docker-entrypoint.s…"   9 minutes ago    Up 9 minutes    0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   mysql


ubuntu@x:~$ docker exec -it mysql /bin/bash
root@c9cbb8634a61:/# mysql -uroot -p123456
mysql>
```



## dockerfile


```dockerfile
ubuntu@x:~/dockerfile$ vim mydockerfile

ubuntu@x:~/dockerfile$ cat mydockerfile
FROM ubuntu

MAINTAINER x<lepengxi@gmail.com>
ENV MYPATH /usr/local

WORKDIR $MYPATH

RUN apt-get update && apt-get -y install vim
RUN apt-get install net-tools

EXPOSE 80

CMD echo $MYPATH
CMD echo "end..."
CMD /bin/bash



ubuntu@x:~/dockerfile$ docker build -f mydockerfile -t myubuntu:0.1 .

Successfully built bf0b925ea7b9
Successfully tagged myubuntu:0.1



ubuntu@x:~/dockerfile$ docker images
REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
myubuntu     0.1       a5b0a9ddf497   2 minutes ago    174MB

ubuntu@x:~/dockerfile$ docker run -it myubuntu:0.1
root@e203287ef30c:/usr/local# vim
root@e203287ef30c:/usr/local# pwd
/usr/local
root@e203287ef30c:/usr/local# ifconfig



ubuntu@x:~$ docker history a5b0a9ddf497
IMAGE          CREATED          CREATED BY                                      SIZE      COMMENT
a5b0a9ddf497   8 minutes ago    /bin/sh -c #(nop)  CMD ["/bin/sh" "-c" "/bin…   0B
98824b54de5c   8 minutes ago    /bin/sh -c #(nop)  CMD ["/bin/sh" "-c" "echo…   0B
9873d7d204e2   8 minutes ago    /bin/sh -c #(nop)  CMD ["/bin/sh" "-c" "echo…   0B
7142880a49da   8 minutes ago    /bin/sh -c #(nop)  EXPOSE 80                    0B
9b8b31f54e7f   8 minutes ago    /bin/sh -c apt-get install net-tools            1.52MB
e42076e6e636   13 minutes ago   /bin/sh -c apt-get update && apt-get -y inst…   99.3MB
24262deb811e   35 minutes ago   /bin/sh -c #(nop) WORKDIR /usr/local            0B
40cc68ef3e9c   35 minutes ago   /bin/sh -c #(nop)  ENV MYPATH=/usr/local        0B
35a3f981c519   35 minutes ago   /bin/sh -c #(nop)  MAINTAINER x<lepengxi@gma…   0B
ba6acccedd29   7 days ago       /bin/sh -c #(nop)  CMD ["bash"]                 0B
<missing>      7 days ago       /bin/sh -c #(nop) ADD file:5d68d27cc15a80653…   72.8MB
```





```dockerfile
ubuntu@x:~/dockerfile$ vim docker-cmd
ubuntu@x:~/dockerfile$ cat docker-cmd
FROM ubuntu

CMD ["ls","-a"]

ubuntu@x:~/dockerfile$ docker build -f docker-cmd -t cmd .
Sending build context to Docker daemon  3.072kB
Step 1/2 : FROM ubuntu
 ---> ba6acccedd29
Step 2/2 : CMD ["ls","-a"]
 ---> Running in b7693c9daae7
Removing intermediate container b7693c9daae7
 ---> 765d6aa3b132
Successfully built 765d6aa3b132
Successfully tagged cmd:latest
ubuntu@x:~/dockerfile$ docker run 765d6aa3b132
.
..
.dockerenv
bin
boot
dev
etc
home


# 不能追加，必须全部替换

ubuntu@x:~/dockerfile$ docker run 765d6aa3b132 -l
docker: Error response from daemon: OCI runtime create failed: container_linux.go:380: starting container process caused: exec: "-l": executable file not found in $PATH: unknown.


ubuntu@x:~/dockerfile$ docker run 765d6aa3b132 ls -al
total 56
drwxr-xr-x   1 root root 4096 Oct 23 14:26 .
drwxr-xr-x   1 root root 4096 Oct 23 14:26 ..
-rwxr-xr-x   1 root root    0 Oct 23 14:26 .dockerenv
lrwxrwxrwx   1 root root    7 Oct  6 16:47 bin -> usr/bin
drwxr-xr-x   2 root root 4096 Apr 15  2020 boot
drwxr-xr-x   5 root root  340 Oct 23 14:26 dev
drwxr-xr-x   1 root root 4096 Oct 23 14:26 etc
drwxr-xr-x   2 root root 4096 Apr 15  2020 home




# 可以追加命令

ubuntu@x:~/dockerfile$ vim docker-entrypoint

ubuntu@x:~/dockerfile$ cat docker-entrypoint
FROM ubuntu

ENTRYPOINT ["ls","-a"]

ubuntu@x:~/dockerfile$ docker build -f docker-entrypoint -t entrypoint .

ubuntu@x:~/dockerfile$ docker run 228e3d259c17
.
..
.dockerenv
bin
boot
dev
etc
home

ubuntu@x:~/dockerfile$ docker run 228e3d259c17 -l
total 56
drwxr-xr-x   1 root root 4096 Oct 23 14:35 .
drwxr-xr-x   1 root root 4096 Oct 23 14:35 ..
-rwxr-xr-x   1 root root    0 Oct 23 14:35 .dockerenv
lrwxrwxrwx   1 root root    7 Oct  6 16:47 bin -> usr/bin
drwxr-xr-x   2 root root 4096 Apr 15  2020 boot
drwxr-xr-x   5 root root  340 Oct 23 14:35 dev
drwxr-xr-x   1 root root 4096 Oct 23 14:35 etc
drwxr-xr-x   2 root root 4096 Apr 15  2020 home
```











**镜像发布**

```dockerfile
ubuntu@x:~/dockerfile$ docker login -u x

ubuntu@x:~$ docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don not have a Docker ID, head over to https://hub.docker.com to create one.
Username: xilepeng
Password:
WARNING! Your password will be stored unencrypted in /home/ubuntu/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded




ubuntu@x:~/dockerfile$ cat mydockerfile
FROM ubuntu

MAINTAINER x<lepengxi@gmail.com>
ENV MYPATH /usr/local

WORKDIR $MYPATH

RUN apt-get update && apt-get -y install vim
RUN apt-get install net-tools

EXPOSE 80

CMD echo $MYPATH
CMD echo "end..."
CMD /bin/bash
ubuntu@x:~/dockerfile$ docker build -f mydockerfile -t xilepeng/ubuntu .


Successfully built a5b0a9ddf497
Successfully tagged xilepeng/ubuntu:latest
ubuntu@x:~/dockerfile$ docker images

xilepeng/ubuntu   latest    a5b0a9ddf497   16 hours ago   174MB

ubuntu@x:~/dockerfile$ docker push xilepeng/ubuntu
Using default tag: latest
The push refers to repository [docker.io/xilepeng/ubuntu]
82e18c734c40: Pushed
b68cb5c1d13d: Pushed
9f54eef41275: Mounted from library/ubuntu
latest: digest: sha256:dc71bd31b77150560a90d0b7faaaecc9b37977df22f9c90ecbfe838e4452e977 size: 951


```




**实践 gin**
```dockerfile
ubuntu@x:~/go-demo$ go env -w GO111MODULE=on
ubuntu@x:~/go-demo$ go env -w GOPROXY=https://goproxy.io,direct

ubuntu@x:~/go-demo$ go mod init go-demo
go: creating new go.mod: module go-demo
go: to add module requirements and sums:
	go mod tidy
ubuntu@x:~/go-demo$ ls
Dockerfile  go.mod  main.go
ubuntu@x:~/go-demo$ go mod tidy


ubuntu@x:~/go-demo$ ls
Dockerfile  go.mod  go.sum  main.go
ubuntu@x:~/go-demo$ cat main.go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}


ubuntu@x:~/go-demo$ cat Dockerfile
FROM golang:1.17.2-alpine AS builder
MAINTAINER x<lepengxi@gmail.com>
# 编译工作目录
WORKDIR /app
# 拷贝本地所有文件到编译工作目录
COPY . .

# RUN go get -u github.com/gin-gonic/gin
ENV GOPROXY https://goproxy.cn

# 编译go文件,设置不调用Cgo
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:latest
# 二进制文件存放目录
WORKDIR /app
# builder别称镜像目录生成文件放入工作目录
COPY --from=builder /app/server /app
# 打开8080端口
EXPOSE 8080
# 从工作目录运行server二进制可执行文件
ENTRYPOINT [ "/app/server" ]



ubuntu@x:~/go-demo$ docker build -t xilepeng/gin-ping .

Successfully built 0d83808031c5
Successfully tagged xilepeng/gin-ping:latest

ubuntu@x:~/go-demo$ docker images
REPOSITORY          TAG             IMAGE ID       CREATED             SIZE
xilepeng/gin-ping   latest          0d83808031c5   52 seconds ago      14.7MB

ubuntu@x:~/go-demo$ docker run -d xilepeng/gin-ping
f073baba719c8a06d41246f2efc6c581eb202494d5974c8ea5b13ae391e3d01e
ubuntu@x:~/go-demo$ docker ps
CONTAINER ID   IMAGE               COMMAND         CREATED         STATUS         PORTS      NAMES
f073baba719c   xilepeng/gin-ping   "/app/server"   5 seconds ago   Up 4 seconds   8080/tcp   trusting_mclaren
ubuntu@x:~/go-demo$ docker inspect f073baba719c

        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "506ee3a90afaac210fde2a8da6c29d76aee7e267d10a87e215fc78e78ecc66bd",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "8080/tcp": null
            },
            "SandboxKey": "/var/run/docker/netns/506ee3a90afa",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "d6a958b9fa5dcb9a333862768ae038faf6f6f2eb45937dd4caac7933104a0da0",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:02",


ubuntu@x:~/go-demo$ curl http://172.17.0.2:8080/ping
{"message":"pong"}


```





```dockerfile

删除所有容器
ubuntu@x:~$ docker rm -f $(docker ps -aq)


删除所有镜像

ubuntu@x:~$ docker rmi -f $(docker images -a)

```

**Docker 网络**

```shell
ubuntu@x:~$ ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: enp0s2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 22:03:01:e1:1a:a5 brd ff:ff:ff:ff:ff:ff
    inet 192.168.105.4/24 brd 192.168.105.255 scope global dynamic enp0s2
       valid_lft 65514sec preferred_lft 65514sec
    inet6 fdaa:9e07:41a7:e34f:2003:1ff:fee1:1aa5/64 scope global dynamic mngtmpaddr noprefixroute
       valid_lft 2591990sec preferred_lft 604790sec
    inet6 fe80::2003:1ff:fee1:1aa5/64 scope link
       valid_lft forever preferred_lft forever
3: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default
    link/ether 02:42:64:36:10:43 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:64ff:fe36:1043/64 scope link
       valid_lft forever preferred_lft forever
```

**自定义网络**

```shell
ubuntu@x:~$ docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
18d7aa9c30f4   bridge    bridge    local
9f67d7222dee   host      host      local
23dd5a314f78   none      null      local
```

**网络模式：**

- bridge: 桥接 docker 默认
- host：和宿主机共享网络
- none: 不配置网络
- container: 容器内网络连通（用的少），局限很大

```shell
ubuntu@x:~$ docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 mynet
f7fbb7667e3afcb9cfb7cbc2d72e1772872c94f3a626e8f8d44f2d8a974962a4
ubuntu@x:~$ docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
18d7aa9c30f4   bridge    bridge    local
9f67d7222dee   host      host      local
f7fbb7667e3a   mynet     bridge    local
```



```shell
ubuntu@x:~$ docker inspect mynet
[
    {
        "Name": "mynet",
        "Id": "f7fbb7667e3afcb9cfb7cbc2d72e1772872c94f3a626e8f8d44f2d8a974962a4",
        "Created": "2021-10-25T12:43:42.411601753+08:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "192.168.0.0/16",
                    "Gateway": "192.168.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {},
        "Options": {},
        "Labels": {}
    }
]
```



```shell
ubuntu@x:~$ docker inspect mynet
[
    {
        "Name": "mynet",
        "Id": "f7fbb7667e3afcb9cfb7cbc2d72e1772872c94f3a626e8f8d44f2d8a974962a4",
        "Created": "2021-10-25T12:43:42.411601753+08:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "192.168.0.0/16",
                    "Gateway": "192.168.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "40c30f43fef30554c4f3bcb6448517cee224b18bd589ed2849fc068852f36ece": {
                "Name": "ubuntu-01",
                "EndpointID": "5b9db3c9d28421d447e3077357c265fbc6cc0467962073dd73d7e93833f47074",
                "MacAddress": "02:42:c0:a8:00:02",
                "IPv4Address": "192.168.0.2/16",
                "IPv6Address": ""
            },
            "5d757ceecfaefb0acf520535f2a42a14a39ea5b1f473c2c9a09279fb6adb04a6": {
                "Name": "ubuntu-02",
                "EndpointID": "6e0682e52454bcf9cab3a3ac45c991b9f37267ab5458597e405cc7f26e6146dd",
                "MacAddress": "02:42:c0:a8:00:03",
                "IPv4Address": "192.168.0.3/16",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {}
    }
]
```



```shell
ubuntu@x:~$ docker run -itd --name ubuntu-01 --network mynet ubuntu /bin/bash
40c30f43fef30554c4f3bcb6448517cee224b18bd589ed2849fc068852f36ece
ubuntu@x:~$ docker run -itd --name ubuntu-02 --network mynet ubuntu /bin/bash
5d757ceecfaefb0acf520535f2a42a14a39ea5b1f473c2c9a09279fb6adb04a6
ubuntu@x:~$ docker ps
CONTAINER ID   IMAGE     COMMAND       CREATED          STATUS          PORTS     NAMES
5d757ceecfae   ubuntu    "/bin/bash"   20 seconds ago   Up 19 seconds             ubuntu-02
40c30f43fef3   ubuntu    "/bin/bash"   29 seconds ago   Up 27 seconds             ubuntu-01

ubuntu@x:~$ docker exec -it ubuntu-01 /bin/bash
root@40c30f43fef3:/# apt-get update
root@40c30f43fef3:/# apt install iputils-ping -y

root@40c30f43fef3:/# ping 192.168.0.3
PING 192.168.0.3 (192.168.0.3) 56(84) bytes of data.
64 bytes from 192.168.0.3: icmp_seq=1 ttl=64 time=2.13 ms
64 bytes from 192.168.0.3: icmp_seq=2 ttl=64 time=0.499 ms

root@40c30f43fef3:/# ping ubuntu-02
PING ubuntu-02 (192.168.0.3) 56(84) bytes of data.
64 bytes from ubuntu-02.mynet (192.168.0.3): icmp_seq=1 ttl=64 time=0.170 ms
64 bytes from ubuntu-02.mynet (192.168.0.3): icmp_seq=2 ttl=64 time=0.208 ms

ubuntu@x:~$ docker exec -it ubuntu-01 ping 192.168.0.3
PING 192.168.0.3 (192.168.0.3) 56(84) bytes of data.
64 bytes from 192.168.0.3: icmp_seq=1 ttl=64 time=0.099 ms
64 bytes from 192.168.0.3: icmp_seq=2 ttl=64 time=0.197 ms

ubuntu@x:~$ docker exec -it ubuntu-01 ping ubuntu-02
PING ubuntu-02 (192.168.0.3) 56(84) bytes of data.
64 bytes from ubuntu-02.mynet (192.168.0.3): icmp_seq=1 ttl=64 time=0.097 ms
64 bytes from ubuntu-02.mynet (192.168.0.3): icmp_seq=2 ttl=64 time=0.183 ms
```

**网络连通**

```shell
ubuntu@x:~$ docker network connect --help

Usage:  docker network connect [OPTIONS] NETWORK CONTAINER

ubuntu@x:~$ docker run -itd --name ubuntu-no ubuntu /bin/bash
ubuntu@x:~$ docker run -itd --name ubuntu ubuntu /bin/bash

ubuntu@x:~$ docker network connect mynet ubuntu
ubuntu@x:~$ docker network inspect mynet


"Containers": {
            "40c30f43fef30554c4f3bcb6448517cee224b18bd589ed2849fc068852f36ece": {
                "Name": "ubuntu-01",
                "EndpointID": "5b9db3c9d28421d447e3077357c265fbc6cc0467962073dd73d7e93833f47074",
                "MacAddress": "02:42:c0:a8:00:02",
                "IPv4Address": "192.168.0.2/16",
                "IPv6Address": ""
            },
            "5d757ceecfaefb0acf520535f2a42a14a39ea5b1f473c2c9a09279fb6adb04a6": {
                "Name": "ubuntu-02",
                "EndpointID": "6e0682e52454bcf9cab3a3ac45c991b9f37267ab5458597e405cc7f26e6146dd",
                "MacAddress": "02:42:c0:a8:00:03",
                "IPv4Address": "192.168.0.3/16",
                "IPv6Address": ""
            },
            "7cc44de2433366275e3a4e7ddf584aac6f5c3a8fac411de49a1ee495525bdc94": {
                "Name": "ubuntu",
                "EndpointID": "49a8e7b224758f0a3d5b152f50285e156964f242f613dbcc3de0029969221511",
                "MacAddress": "02:42:c0:a8:00:04",
                "IPv4Address": "192.168.0.4/16",
                "IPv6Address": ""
            }
        },
        

```

**连通之后，就是将 ubuntu 加入 mynet 网络下**

```shell
ubuntu@x:~$ docker exec -it ubuntu-01 ping ubuntu
PING ubuntu (192.168.0.4) 56(84) bytes of data.
64 bytes from ubuntu.mynet (192.168.0.4): icmp_seq=1 ttl=64 time=1.74 ms
64 bytes from ubuntu.mynet (192.168.0.4): icmp_seq=2 ttl=64 time=0.221 ms

ubuntu@x:~$ docker exec -it ubuntu-01 ping ubuntu-no
ping: ubuntu-no: Temporary failure in name resolution
```



**redis 集群**



```shell
ubuntu@x:~$ docker network create redis --subnet 172.38.0.0/16
67839a9ee88e776b4b921b981bff45791bed321a3d271d6bdce1ba0ac4870477
ubuntu@x:~$ docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
18d7aa9c30f4   bridge    bridge    local
9f67d7222dee   host      host      local
f7fbb7667e3a   mynet     bridge    local
23dd5a314f78   none      null      local
67839a9ee88e   redis     bridge    local
ubuntu@x:~$ docker network inspect redis
[
    {
        "Name": "redis",
        "Id": "67839a9ee88e776b4b921b981bff45791bed321a3d271d6bdce1ba0ac4870477",
        "Created": "2021-10-25T16:53:18.712494047+08:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "172.38.0.0/16"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {},
        "Options": {},
        "Labels": {}
    }
]



```



```shell
for port in $(seq 1 6); \
do \
mkdir -p /mydata/redis/node-${port}/conf
touch /mydata/redis/node-${port}/conf/redis.conf
cat << EOF >/mydata/redis/node-${port}/conf/redis.conf
port 6379
bind 0.0.0.0
cluster-enabled yes 
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip 172.38.0.1${port}
cluster-announce-port 6379
cluster-announce-bus-port 16379
appendonly yes
EOF
done

```



```shell
ubuntu@x:~$ sudo passwd root
New password:
Retype new password:
passwd: password updated successfully
ubuntu@x:~$ su
Password:
root@x:/home/ubuntu# for port in $(seq 1 6); \
> do \
> mkdir -p /mydata/redis/node-${port}/conf
> touch /mydata/redis/node-${port}/conf/redis.conf
> cat << EOF >/mydata/redis/node-${port}/conf/redis.conf
> port 6379
> bind 0.0.0.0
> cluster-enabled yes
> cluster-config-file nodes.conf
> cluster-node-timeout 5000
> cluster-announce-ip 172.38.0.1${port}
> cluster-announce-port 6379
> cluster-announce-bus-port 16379
> appendonly yes
> EOF
> done

root@x:/mydata/redis/node-1/conf# cat redis.conf
port 6379
bind 0.0.0.0
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip 172.38.0.11
cluster-announce-port 6379
cluster-announce-bus-port 16379
appendonly yes

root@x:~# sudo passwd -dl root
```







```shell
docker run -p 6371:6379 -p 16371:16379 --name redis-1 \
-v /mydata/redis/node-1/data:/data \
-v /mydata/redis/node-1/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.11 redis:latest redis-server /etc/redis/redis.conf


docker run -p 6372:6379 -p 16372:16379 --name redis-2 \
-v /mydata/redis/node-2/data:/data \
-v /mydata/redis/node-2/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.12 redis:latest redis-server /etc/redis/redis.conf


docker run -p 6373:6379 -p 16373:16379 --name redis-3 \
-v /mydata/redis/node-3/data:/data \
-v /mydata/redis/node-3/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.13 redis:latest redis-server /etc/redis/redis.conf


docker run -p 6374:6379 -p 16374:16379 --name redis-4 \
-v /mydata/redis/node-4/data:/data \
-v /mydata/redis/node-4/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.14 redis:latest redis-server /etc/redis/redis.conf


docker run -p 6375:6379 -p 16375:16379 --name redis-5 \
-v /mydata/redis/node-5/data:/data \
-v /mydata/redis/node-5/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.15 redis:latest redis-server /etc/redis/redis.conf


docker run -p 6376:6379 -p 16376:16379 --name redis-6 \
-v /mydata/redis/node-6/data:/data \
-v /mydata/redis/node-6/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.16 redis:latest redis-server /etc/redis/redis.conf






ubuntu@x:~$ docker run -p 6371:6379 -p 16371:16379 --name redis-1 \
> -v /mydata/redis/node-1/data:/data \
> -v /mydata/redis/node-1/conf/redis.conf:/etc/redis/redis.conf \
> -d --net redis --ip 172.38.0.11 redis:latest redis-server /etc/redis/redis.conf
45ff6db22e8bb592a68db715295779e38ea4b2ccfef1e0902eb5be1258853483
ubuntu@x:~$ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS         PORTS                                                                                      NAMES
45ff6db22e8b   redis:latest   "docker-entrypoint.s…"   12 seconds ago   Up 9 seconds   0.0.0.0:6371->6379/tcp, :::6371->6379/tcp, 0.0.0.0:16371->16379/tcp, :::16371->16379/tcp   redis-1

...


ubuntu@x:~$ docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED              STATUS              PORTS                                                                                      NAMES
c95b5b3188ed   redis:latest   "docker-entrypoint.s…"   6 seconds ago        Up 4 seconds        0.0.0.0:6376->6379/tcp, :::6376->6379/tcp, 0.0.0.0:16376->16379/tcp, :::16376->16379/tcp   redis-6
c9e43c0f9a61   redis:latest   "docker-entrypoint.s…"   39 seconds ago       Up 38 seconds       0.0.0.0:6375->6379/tcp, :::6375->6379/tcp, 0.0.0.0:16375->16379/tcp, :::16375->16379/tcp   redis-5
fe26323bd9e7   redis:latest   "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:6374->6379/tcp, :::6374->6379/tcp, 0.0.0.0:16374->16379/tcp, :::16374->16379/tcp   redis-4
1660039034fe   redis:latest   "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:6373->6379/tcp, :::6373->6379/tcp, 0.0.0.0:16373->16379/tcp, :::16373->16379/tcp   redis-3
1bdd587793d6   redis:latest   "docker-entrypoint.s…"   2 minutes ago        Up 2 minutes        0.0.0.0:6372->6379/tcp, :::6372->6379/tcp, 0.0.0.0:16372->16379/tcp, :::16372->16379/tcp   redis-2
45ff6db22e8b   redis:latest   "docker-entrypoint.s…"   6 minutes ago        Up 6 minutes        0.0.0.0:6371->6379/tcp, :::6371->6379/tcp, 0.0.0.0:16371->16379/tcp, :::16371->16379/tcp   redis-1



ubuntu@x:~$ docker exec -it redis-1 /bin/bash
root@45ff6db22e8b:/data# ls
appendonly.aof	nodes.conf
root@45ff6db22e8b:/data#


root@45ff6db22e8b:/data# redis-cli --cluster create 172.38.0.11:6379 172.38.0.12:6379 172.38.0.13:6379 172.38.0.14:6379 172.38.0.15:6379 172.38.0.16:6379 --cluster-replicas 1
>>> Performing hash slots allocation on 6 nodes...
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
Adding replica 172.38.0.15:6379 to 172.38.0.11:6379
Adding replica 172.38.0.16:6379 to 172.38.0.12:6379
Adding replica 172.38.0.14:6379 to 172.38.0.13:6379
M: 6ef16d62c844287639af6fa558aa83db2a255d1b 172.38.0.11:6379
   slots:[0-5460] (5461 slots) master
M: 5a95e4469496ce73488dd41174d15efb3fb0df55 172.38.0.12:6379
   slots:[5461-10922] (5462 slots) master
M: 62e4718bdaac5f559f37a168d0e3ce05585591bc 172.38.0.13:6379
   slots:[10923-16383] (5461 slots) master
S: 78c2893501257bd70c1adadea57ddc4c9c2e3388 172.38.0.14:6379
   replicates 62e4718bdaac5f559f37a168d0e3ce05585591bc
S: 62e1d1303c232979fedda8ae779eae774a214daf 172.38.0.15:6379
   replicates 6ef16d62c844287639af6fa558aa83db2a255d1b
S: 94055bd4d435f97e33f12496e63d9f5043de75c2 172.38.0.16:6379
   replicates 5a95e4469496ce73488dd41174d15efb3fb0df55
Can I set the above configuration? (type 'yes' to accept): yes
>>> Nodes configuration updated
>>> Assign a different config epoch to each node
>>> Sending CLUSTER MEET messages to join the cluster
Waiting for the cluster to join
.
>>> Performing Cluster Check (using node 172.38.0.11:6379)
M: 6ef16d62c844287639af6fa558aa83db2a255d1b 172.38.0.11:6379
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
M: 5a95e4469496ce73488dd41174d15efb3fb0df55 172.38.0.12:6379
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: 94055bd4d435f97e33f12496e63d9f5043de75c2 172.38.0.16:6379
   slots: (0 slots) slave
   replicates 5a95e4469496ce73488dd41174d15efb3fb0df55
S: 62e1d1303c232979fedda8ae779eae774a214daf 172.38.0.15:6379
   slots: (0 slots) slave
   replicates 6ef16d62c844287639af6fa558aa83db2a255d1b
S: 78c2893501257bd70c1adadea57ddc4c9c2e3388 172.38.0.14:6379
   slots: (0 slots) slave
   replicates 62e4718bdaac5f559f37a168d0e3ce05585591bc
M: 62e4718bdaac5f559f37a168d0e3ce05585591bc 172.38.0.13:6379
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.



root@45ff6db22e8b:/data# redis-cli -c
127.0.0.1:6379> cluster info

cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:6
cluster_size:3
cluster_current_epoch:6
cluster_my_epoch:1
cluster_stats_messages_ping_sent:173
cluster_stats_messages_pong_sent:176
cluster_stats_messages_sent:349
cluster_stats_messages_ping_received:171
cluster_stats_messages_pong_received:173
cluster_stats_messages_meet_received:5
cluster_stats_messages_received:349


127.0.0.1:6379> cluster nodes
6ef16d62c844287639af6fa558aa83db2a255d1b 172.38.0.11:6379@16379 myself,master - 0 1635156827000 1 connected 0-5460
5a95e4469496ce73488dd41174d15efb3fb0df55 172.38.0.12:6379@16379 master - 0 1635156828817 2 connected 5461-10922
94055bd4d435f97e33f12496e63d9f5043de75c2 172.38.0.16:6379@16379 slave 5a95e4469496ce73488dd41174d15efb3fb0df55 0 1635156829247 2 connected
62e1d1303c232979fedda8ae779eae774a214daf 172.38.0.15:6379@16379 slave 6ef16d62c844287639af6fa558aa83db2a255d1b 0 1635156828000 1 connected
78c2893501257bd70c1adadea57ddc4c9c2e3388 172.38.0.14:6379@16379 slave 62e4718bdaac5f559f37a168d0e3ce05585591bc 0 1635156828505 3 connected
62e4718bdaac5f559f37a168d0e3ce05585591bc 172.38.0.13:6379@16379 master - 0 1635156828000 3 connected 10923-16383

172.38.0.15:6379> set k1 v1
-> Redirected to slot [12706] located at 172.38.0.14:6379
OK

ubuntu@x:~$ docker stop redis-4

172.38.0.14:6379> get k1
Error: Server closed the connection
172.38.0.14:6379> get k1
Could not connect to Redis at 172.38.0.14:6379: No route to host
(34.23s)
not connected>

127.0.0.1:6379> get k1
-> Redirected to slot [12706] located at 172.38.0.13:6379
"v1"

172.38.0.13:6379> cluster nodes
94055bd4d435f97e33f12496e63d9f5043de75c2 172.38.0.16:6379@16379 master - 0 1635159704000 8 connected 5461-10922
5a95e4469496ce73488dd41174d15efb3fb0df55 172.38.0.12:6379@16379 slave 94055bd4d435f97e33f12496e63d9f5043de75c2 0 1635159704945 8 connected
62e4718bdaac5f559f37a168d0e3ce05585591bc 172.38.0.13:6379@16379 myself,master - 0 1635159704000 10 connected 10923-16383
78c2893501257bd70c1adadea57ddc4c9c2e3388 172.38.0.14:6379@16379 master,fail - 1635159512948 1635159510872 7 connected
6ef16d62c844287639af6fa558aa83db2a255d1b 172.38.0.11:6379@16379 slave 62e1d1303c232979fedda8ae779eae774a214daf 0 1635159705377 9 connected
62e1d1303c232979fedda8ae779eae774a214daf 172.38.0.15:6379@16379 master - 0 1635159705099 9 connected 0-5460
```



**微服务打包docker镜像**

```shell
$ docker build -t memory .

Successfully built 915f17114664
Successfully tagged memory:latest

ubuntu@x:~/Shared/Memorization-App/account$ docker images
REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
memory       latest    915f17114664   2 hours ago   15.2MB

ubuntu@x:~$ docker run -d -P --name memo memory
5cc4fe17edd5caead6be3ee28222ff218aae77d4d6609e941f0ac2d8b9323228

ubuntu@x:~$ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED          STATUS          PORTS                                         NAMES
5cc4fe17edd5   memory    "./run"   18 seconds ago   Up 15 seconds   0.0.0.0:49156->8080/tcp, :::49156->8080/tcp   memo

ubuntu@x:~$ curl localhost:49156/api/account
{"hello":"world"}ubuntu@x:~$ curl 0.0.0.0:49156/api/account
{"hello":"world"}ubuntu@x:~$
```



## Docker Compose

Ubuntu 20.04.3 LTS 安装 Docker Compose

```shell
ubuntu@x:~$ sudo curl -L "https://github.com/docker/compose/releases/download/v2.0.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   633  100   633    0     0     38      0  0:00:16  0:00:16 --:--:--   170
100 24.7M  100 24.7M    0     0  1246k      0  0:00:20  0:00:20 --:--:-- 7767k

ubuntu@x:~$ sudo chmod +x /usr/local/bin/docker-compose
ubuntu@x:~$ sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
ubuntu@x:~$ docker-compose --version
Docker Compose version v2.0.1

```



```

