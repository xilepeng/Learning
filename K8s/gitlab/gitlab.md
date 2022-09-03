
[官方文档](https://docs.gitlab.cn/charts/quickstart/index.html)

``` s
➜  ~ helm repo add gitlab-jh https://charts.gitlab.cn/
"gitlab-jh" has been added to your repositories

➜  ~ helm repo update
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "gitlab-jh" chart repository
...Successfully got an update from the "stable" chart repository
Update Complete. ⎈Happy Helming!⎈


➜  ~ helm search repo gitlab
NAME                   	CHART VERSION	APP VERSION	DESCRIPTION
gitlab-jh/gitlab       	6.3.2        	15.3.2     	The One DevOps Platform
gitlab-jh/gitlab-runner	0.44.0       	15.3.0     	GitLab Runner
stable/gitlab-ce       	0.2.3        	9.4.1      	GitLab Community Edition
stable/gitlab-ee       	0.2.3        	9.4.1      	GitLab Enterprise Edition


helm install gitlab gitlab-jh/gitlab \
  --version 6.3.2 \
  --set global.hosts.domain=hfbpw.top \
  --set certmanager-issuer.email=lepengxi@gmail.com \
```




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

```