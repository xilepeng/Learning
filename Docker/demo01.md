

```go
Linux电脑 安装脚本：

/bin/bash -c "$(curl -fsSL https://gitee.com/cunkai/HomebrewCN/raw/master/Homebrew.sh)"

rm Homebrew.sh ; wget https://gitee.com/cunkai/HomebrewCN/raw/master/Homebrew.sh ; bash Homebrew.sh


Linux电脑 卸载脚本：

rm HomebrewUninstall.sh ; wget https://gitee.com/cunkai/HomebrewCN/raw/master/HomebrewUninstall.sh ; bash HomebrewUninstall.sh




终端输入以下几行命令设置环境变量:

function brew() {
>     PATH="/home/linuxbrew/.linuxbrew/bin:$PATH" /home/linuxbrew/.linuxbrew/bin/brew "$@"
> }

将Homebrew加入PATH


test -d ~/.linuxbrew && eval $(~/.linuxbrew/bin/brew shellenv)
test -d /home/linuxbrew/.linuxbrew && eval $(/home/linuxbrew/.linuxbrew/bin/brew shellenv)
test -r ~/.bash_profile && echo "eval \$($(brew --prefix)/bin/brew shellenv)" >>~/.bash_profile
echo "eval \$($(brew --prefix)/bin/brew shellenv)" >>~/.profile


对于zsh用户: 还需

test -r ~/.zsh_profile && echo "eval \$($(brew --prefix)/bin/brew shellenv)" >>~/.zsh_profile
echo "export PATH=$HOME/bin:/usr/local/bin:$PATH" >> ~/.zshrc                                            
echo "eval \$($(brew --prefix)/bin/brew shellenv)" >>~/.zshrc 
source ~/.zshrc 



Homebrew-bottles 镜像使用帮助
注:该镜像是 Homebrew 二进制预编译包的镜像。本镜像站同时提供 Homebrew 的 formula 索引的镜像（即 brew update 时所更新内容），请参考 Homebrew 镜像使用帮助。

长期替换
如果你使用 bash：

echo 'export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles"' >> ~/.bash_profile
export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles"
如果你使用 zsh：

echo 'export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles"' >> ~/.zprofile
export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles"



# 对于 bash 用户
echo 'export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles"' >> ~/.bash_profile
export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles"

# 对于 zsh 用户
echo 'export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles"' >> ~/.zshrc
export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles"

```


```go

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
```



`ubuntu@x:~$ /bin/bash -c "$(curl -fsSL https://gitee.com/cunkai/HomebrewCN/raw/master/Homebrew.sh)"`

**任何工作站的Ubuntu虚拟机**

使用单个命令获取即时Ubuntu VM。多通可以启动和运行虚拟机，并像公共云一样配置它们。您的云原型在本地免费启动。

![Multipass](https://multipass.run/)


**1. 在MacOS上安装 Multipass**

![Download Multipass for MacOS](https://multipass.run/download/macos)


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
multipass delete x
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


`sudo vim /var/root/Library/Application\ Support/multipassd/multipassd-vm-instances.json`


```bash
➜  ~ multipass exec x /bin/bash
To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.

ubuntu@x:~$
```


**Ubuntu 安装 Docker**

使用官方安装脚本自动安装

`curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun`


```shell
ubuntu@x:~$ docker images
Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Get "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/images/json": dial unix /var/run/docker.sock: connect: permission denied

ubuntu@x:~$ sudo groupadd docker
groupadd: group 'docker' already exists

ubuntu@x:~$ sudo gpasswd -a ubuntu docker
Adding user ubuntu to group docker

ubuntu@x:~$ sudo service docker restart
重启 iTerm2



ubuntu@x:~$ sudo vim /etc/docker/daemon.json


{ "registry-mirrors": [
    "https://hkaofvr0.mirror.aliyuncs.com"
  ]
 }

ubuntu@x:~$ sudo systemctl daemon-reload
ubuntu@x:~$ sudo systemctl restart docker
ubuntu@x:~$ docker info

 Registry Mirrors:
  https://hkaofvr0.mirror.aliyuncs.com/


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

ubuntu@x:~$ docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
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

