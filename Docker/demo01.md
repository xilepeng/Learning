




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

```
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
root@de55b21c29a6:/etc/nginx#


```