

``` s
sudo docker run --detach \
  --hostname localhost \
  --publish 443:443 --publish 8929:8929 --publish 2224:22 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \
  registry.gitlab.cn/omnibus/gitlab-jh:latest
```




``` yml

# docker-compose.yml
version: '3.7'
services:
  gitlab:
    image: 'registry.gitlab.cn/omnibus/gitlab-jh:latest'
    restart: always
    hostname: 'localhost'
    container_name: gitlab
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'http://localhost:8929'
        gitlab_rails['gitlab_shell_ssh_port'] = 2224
    ports:
      - '8929:8929'
      - '8443:443'
      - '2224:2224'
    volumes:
      - '$GITLAB_HOME/config:/etc/gitlab'
      - '$GITLAB_HOME/logs:/var/log/gitlab'
      - '$GITLAB_HOME/data:/var/opt/gitlab'
    networks:
      - gitlab
  gitlab-runner:
    image: gitlab/gitlab-runner:alpine
    container_name: gitlab-runner    
    restart: always
    depends_on:
      - gitlab
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - '$GITLAB_HOME/gitlab-runner:/etc/gitlab-runner'
    networks:
      - gitlab

networks:
  gitlab:
    name: gitlab-network
```








``` yml

version: "3.6"
services:
  gitlab:
    image: registry.gitlab.cn/omnibus/gitlab-jh:latest
    ports:
      - "22:22"
      - "80:80"
      - "443:443"
    volumes:
      - $GITLAB_HOME/data:/var/opt/gitlab
      - $GITLAB_HOME/logs:/var/log/gitlab
      - $GITLAB_HOME/config:/etc/gitlab
    shm_size: '256m'
    environment:
      GITLAB_OMNIBUS_CONFIG: "from_file('/omnibus_config.rb')"
    configs:
      - source: gitlab
        target: /omnibus_config.rb
    secrets:
      - gitlab_root_password
  gitlab-runner:
    image: gitlab/gitlab-runner:alpine
    deploy:
      mode: replicated
      replicas: 4
configs:
  gitlab:
    file: ./gitlab.rb
secrets:
  gitlab_root_password:
    file: ./root_password.txt



```


``` s
docker logs -f gitlab


```



docker-compose.yml

``` yml

# docker-compose.yml
version: '3.7'
services:
  web:
    image: 'registry.gitlab.cn/omnibus/gitlab-jh:latest'
    restart: always
    hostname: 'localhost'
    container_name: gitlab
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'http://localhost'
    ports:
      - '8080:80'
      - '8443:443'
    volumes:
      - '$GITLAB_HOME/config:/etc/gitlab'
      - '$GITLAB_HOME/logs:/var/log/gitlab'
      - '$GITLAB_HOME/data:/var/opt/gitlab'
    networks:
      - gitlab
  gitlab-runner:
    image: gitlab/gitlab-runner:alpine
    container_name: gitlab-runner    
    restart: always
    depends_on:
      - web
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - '$GITLAB_HOME/gitlab-runner:/etc/gitlab-runner'
    networks:
      - gitlab

networks:
  gitlab:
    name: gitlab-network
```





``` s
vim .zshrc

export GITLAB_HOME=$HOME/gitlab


sudo docker run --detach \
  --hostname localhost \
  --publish 8929:8929 --publish 2289:22 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \     
  registry.gitlab.cn/omnibus/gitlab-jh:latest








sudo docker run --detach \
  --hostname localhost \
  --publish 443:443 --publish 8929:8929 --publish 2289:22 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \
  registry.gitlab.cn/omnibus/gitlab-jh:latest




```






``` s

sudo docker run --detach \
  --hostname localhost\
  --publish 8929:8929 --publish 2289:22 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \     
  registry.gitlab.cn/omnibus/gitlab-jh:latest
  gitlab/gitlab-ce:latest

```


``` s
➜  gitlab mkdir config logs data
➜  gitlab chmod 777 config logs data

➜  gitlab ll
total 0
drwxrwxrwx  2 x  staff    64B  9  4 16:06 config
drwxrwxrwx  2 x  staff    64B  9  4 16:07 data
drwxrwxrwx  2 x  staff    64B  9  4 16:08 logs


```


``` s
sudo docker run --detach \
  --hostname localhost \
  --env GITLAB_OMNIBUS_CONFIG="external_url 'localhost'; gitlab_rails['lfs_enabled'] = true;" \
  --publish 443:443 --publish 8929:8929 --publish 2224:22 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \
  registry.gitlab.cn/omnibus/gitlab-jh:latest
```





























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
    hostname: 'localhost'
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'https://localhost'
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
    hostname: 'localhost'
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'http://localhost:8929'
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




``` s
docker exec -it gitlab-ce grep 'Password:' /etc/gitlab/initial_root_password

Password: 2G24STEl7EnIWOhdNNlzIQjKs8mov6uGngMrer4wpTk=

sudo docker logs -f gitlab-ce

```

gitlab-runner | ERROR: Failed to load config stat /etc/gitlab-runner/config.toml: no such file or directory  builds=0


``` yml
version: '3.6'
services:
  web:
    image: 'registry.gitlab.cn/omnibus/gitlab-jh:latest'
    restart: always
    hostname: 'localhost'
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'https://localhost'
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
    image: 'registry.gitlab.cn/omnibus/gitlab-jh:latest'
    restart: always
    hostname: 'localhost'
    environment:
      GITLAB_OMNIBUS_CONFIG: |
        external_url 'http://localhost:8929'
        gitlab_rails['gitlab_shell_ssh_port'] = 2224
    ports:
      - '8929:8929'
      - '2224:22'
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









```

