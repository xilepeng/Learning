
## 系统要求

最低推荐配置:

256MB可用内存
1GB可用磁盘空间(作为一个Docker容器运行jenkins的话推荐10GB)

``` s
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

苹果系统: 安装最新版本

``` s
➜  ~ brew install jenkins


Running `brew update --auto-update`...
==> Downloading https://ghcr.io/v2/homebrew/core/openjdk/17/manifests/17.0.4.1
######################################################################## 100.0%
==> Downloading https://ghcr.io/v2/homebrew/core/openjdk/17/blobs/sha256:7850d052807931b14395f0ba01938be6718518c76c
==> Downloading from https://pkg-containers.githubusercontent.com/ghcr1/blobs/sha256:7850d052807931b14395f0ba01938b
######################################################################## 100.0%
==> Downloading https://ghcr.io/v2/homebrew/core/jenkins/manifests/2.366
######################################################################## 100.0%
==> Downloading https://ghcr.io/v2/homebrew/core/jenkins/blobs/sha256:dd5c46f2fb673efd34d024fb09f78e982f6391eefec6d
==> Downloading from https://pkg-containers.githubusercontent.com/ghcr1/blobs/sha256:dd5c46f2fb673efd34d024fb09f78e
######################################################################## 100.0%
==> Installing dependencies for jenkins: openjdk@17
==> Installing jenkins dependency: openjdk@17
==> Pouring openjdk@17--17.0.4.1.monterey.bottle.tar.gz
🍺  /usr/local/Cellar/openjdk@17/17.0.4.1: 639 files, 305.5MB
==> Installing jenkins
==> Pouring jenkins--2.366.all.bottle.tar.gz
==> Caveats
Note: When using launchctl the port will be 8080.

To restart jenkins after an upgrade:
  brew services restart jenkins
Or, if you don't want/need a background service you can just run:
  /usr/local/opt/jenkins/bin/jenkins --httpListenAddress=127.0.0.1 --httpPort=8080
==> Summary
🍺  /usr/local/Cellar/jenkins/2.366: 8 files, 92.5MB
==> Running `brew cleanup jenkins`...
Disable this behaviour by setting HOMEBREW_NO_INSTALL_CLEANUP.
Hide these hints with HOMEBREW_NO_ENV_HINTS (see `man brew`).
==> Caveats
==> jenkins
Note: When using launchctl the port will be 8080.

To restart jenkins after an upgrade:
  brew services restart jenkins
Or, if you don't want/need a background service you can just run:
  /usr/local/opt/jenkins/bin/jenkins --httpListenAddress=127.0.0.1 --httpPort=8080



➜  ~ brew services restart jenkins
==> Successfully started `jenkins` (label: homebrew.mxcl.jenkins)

```


``` s
http://127.0.0.1:8080/


➜  ~ cat /Users/x/.jenkins/secrets/initialAdminPassword
cb3ad34aebe84cc1a9cc33f5d81194ee



```







