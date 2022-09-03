































## docker 安装 gitlab

``` s
➜  ~ docker version
Client:
 Cloud integration: v1.0.29
 Version:           20.10.17
 API version:       1.41
 Go version:        go1.17.11
 Git commit:        100c701
 Built:             Mon Jun  6 23:04:45 2022
 OS/Arch:           darwin/amd64
 Context:           default
 Experimental:      true

➜  ~ docker-compose version
Docker Compose version 2.10.2


➜  local sudo mkdir docker
Password:
➜  local ls
Caskroom   Frameworks bin        etc        lib        sbin       ssl        var
Cellar     Homebrew   docker     include    opt        share      texlive
➜  local cd docker

➜  docker sudo mkdir gitlab_docker
➜  docker ls
gitlab_docker
➜  docker pwd
/usr/local/docker


➜  docker sudo vim docker-compose.yml



docker search gitlab-ce

NAME                                     DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
gitlab/gitlab-ce                         GitLab Community Edition docker image based …   3690                 [OK]


docker pull gitlab/gitlab-ce



```







docker-compose.yml


``` yml

version: '3.6'
services:
  web:
    image: 'registry.gitlab.cn/omnibus/gitlab-jh:latest'
    restart: always
    hostname: 'gitlab.example.com'
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'https://gitlab.example.com'
        # Add any other gitlab.rb configuration here, each on its own line
    ports:
      - '80:80'
      - '443:443'
      - '22:22'
    volumes:
      - '$GITLAB_HOME/config:/etc/gitlab'
      - '$GITLAB_HOME/logs:/var/log/gitlab'
      - '$GITLAB_HOME/data:/var/opt/gitlab'
    shm_size: '256m'
```


``` yml
version: '3.6'
services:
  web:
    image: 'gitlab/gitlab-ce:latest'
    restart: always
    hostname: 'gitlab'
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'http://127.0.0.1/:8929'
        gitlab_rails['gitlab_shell_ssh_port'] = 2224
    ports:
      - '8929:8929'
      - '2224:2224'
    volumes:
      - '$GITLAB_HOME/config:/etc/gitlab'
      - '$GITLAB_HOME/logs:/var/log/gitlab'
      - '$GITLAB_HOME/data:/var/opt/gitlab'
    shm_size: '256m'
```



``` s




helm install gitlab gitlab-jh/gitlab \
  --version 6.3.2 \
  --set global.hosts.domain=127.0.0.1 \
  --set certmanager-issuer.email=lepengxi@gmail.com \



```





















## helm 安装 gitlab

[官方文档](https://docs.gitlab.cn/charts/quickstart/index.html)

``` s
➜  ~ helm repo add gitlab-jh https://charts.gitlab.cn/
"gitlab-jh" has been added to your repositories

➜  ~ helm repo update
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "gitlab-jh" chart repository
...Successfully got an update from the "stable" chart repository
Update Complete. ⎈Happy Helming!⎈

helm install gitlab gitlab-jh/gitlab \
  --version 6.3.2 \
  --set global.hosts.domain=localhost \
  --set global.hosts.externalIP=127.0.0.1 \
  --set certmanager-issuer.email=lepengxi@gmail.ocm 





➜  ~ helm search repo gitlab
NAME                   	CHART VERSION	APP VERSION	DESCRIPTION
gitlab-jh/gitlab       	6.3.2        	15.3.2     	The One DevOps Platform
gitlab-jh/gitlab-runner	0.44.0       	15.3.0     	GitLab Runner
stable/gitlab-ce       	0.2.3        	9.4.1      	GitLab Community Edition
stable/gitlab-ee       	0.2.3        	9.4.1      	GitLab Enterprise Edition


helm install gitlab gitlab-jh/gitlab \
  --version 6.3.2 \
  --set global.hosts.domain=127.0.0.1 \
  --set certmanager-issuer.email=lepengxi@gmail.com 
```




## docker-compose 安装 gitlab


``` s

➜  ~ export GITLAB_HOME=$HOME/gitlab
➜  ~ mkdir gitlab
➜  ~ ls
Applications Documents    Library      Music        Public       go           share
Desktop      Downloads    Movies       Pictures     gitlab       node_modules x
➜  ~ cd gitlab


➜  gitlab vim docker-compose.yml
➜  gitlab cat docker-compose.yml


```


docker-compose.yml

``` yml
version: '3.6'
services:
  web:
    image: 'registry.gitlab.cn/omnibus/gitlab-jh:latest'
    restart: always
    hostname: 'gitlab'
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'http://127.0.0.1/'
        # Add any other gitlab.rb configuration here, each on its own line
    ports:
      - '80:80'
      - '443:443'
      - '22:22'
    volumes:
      - '$GITLAB_HOME/config:/etc/gitlab'
      - '$GITLAB_HOME/logs:/var/log/gitlab'
      - '$GITLAB_HOME/data:/var/opt/gitlab'
    shm_size: '256m'
```


``` yml
version: '3.6'
services:
  web:
    image: 'gitlab/gitlab-ce:latest'
    restart: always
    hostname: 'gitlab'
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'http://127.0.0.1/:8929'
        gitlab_rails['gitlab_shell_ssh_port'] = 2224
    ports:
      - '8929:8929'
      - '2224:2224'
    volumes:
      - '$GITLAB_HOME/config:/etc/gitlab'
      - '$GITLAB_HOME/logs:/var/log/gitlab'
      - '$GITLAB_HOME/data:/var/opt/gitlab'
    shm_size: '256m'
```





``` s 
➜  gitlab docker-compose up -d 


[+] Running 7/9
 ⠹ web Pulling                                                                                              308.2s
   ⠿ 3b65ec22a9e9 Pull complete                                                                              24.2s
   ⠿ ff0b65fed8c6 Pull complete                                                                              30.6s
   ⠿ 4a9560beaf66 Pull complete                                                                              30.8s
   ⠿ 0b894e0a6870 Pull complete                                                                              31.0s
   ⠿ f8b6f0b15af6 Pull complete                                                                              31.2s
   ⠿ 4cd8165e8ffb Pull complete                                                                              31.5s
   ⠿ f00fbf1d60e2 Pull complete                                                                              31.8s
   ⠋ 86646388c934 Downloading [=================================================> ] ...                     303.9s





sudo docker run --detach \
  --hostname gitlab.example.com \
  --publish 8929:8929 --publish 2224:2224 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \     
  gitlab/gitlab-ce:latest







```

