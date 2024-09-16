
## 安装 minikube、helm

``` s 
➜  ~ brew install minikube


➜  ~ minikube start
😄  Darwin 12.5.1 上的 minikube v1.26.1
🆕  Kubernetes 1.24.3 is now available. If you would like to upgrade, specify: --kubernetes-version=v1.24.3
✨  根据现有的配置文件使用 docker 驱动程序
👍  Starting control plane node minikube in cluster minikube
🚜  Pulling base image ...
💾  Downloading Kubernetes v1.22.3 preload ...
    > index.docker.io/kicbase/sta...:  355.78 MiB / 355.78 MiB  100.00% 2.32 Mi
    > preloaded-images-k8s-v18-v1...:  396.52 MiB / 396.52 MiB  100.00% 2.35 Mi
    > index.docker.io/kicbase/sta...:  0 B [___________________] ?% ? p/s 1m36s
🤷  docker "minikube" container is missing, will recreate.
🔥  Creating docker container (CPUs=2, Memory=2200MB) ...
❗  This container is having trouble accessing https://k8s.gcr.io
💡  To pull new external images, you may need to configure a proxy: https://minikube.sigs.k8s.io/docs/reference/networking/proxy/
🐳  正在 Docker 20.10.8 中准备 Kubernetes v1.22.3…
    ▪ Generating certificates and keys ...
    ▪ Booting up control plane ...
    ▪ Configuring RBAC rules ...
🔎  Verifying Kubernetes components...
    ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5
    ▪ Using image kubernetesui/dashboard:v2.6.0
    ▪ Using image kubernetesui/metrics-scraper:v1.0.8
🌟  Enabled addons: default-storageclass, storage-provisioner, dashboard

❗  /usr/local/bin/kubectl is version 1.25.0, which may have incompatibilites with Kubernetes 1.22.3.
    ▪ Want kubectl v1.22.3? Try 'minikube kubectl -- get pods -A'
🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default





➜  ~ minikube kubectl -- get pods -A
    > kubectl.sha256:  64 B / 64 B [-------------------------] 100.00% ? p/s 0s
    > kubectl:  50.63 MiB / 50.63 MiB [--------------] 100.00% 5.13 MiB p/s 10s
NAMESPACE              NAME                                         READY   STATUS    RESTARTS      AGE
kube-system            coredns-78fcd69978-txpvl                     1/1     Running   0             19m
kube-system            etcd-minikube                                1/1     Running   0             19m
kube-system            kube-apiserver-minikube                      1/1     Running   0             19m
kube-system            kube-controller-manager-minikube             1/1     Running   0             19m
kube-system            kube-proxy-lslj9                             1/1     Running   0             19m
kube-system            kube-scheduler-minikube                      1/1     Running   0             19m
kube-system            storage-provisioner                          1/1     Running   1 (19m ago)   19m
kubernetes-dashboard   dashboard-metrics-scraper-6bd447698b-m8q42   1/1     Running   0             19m
kubernetes-dashboard   kubernetes-dashboard-55f855cf74-fg6cg        1/1     Running   0             19m









➜  ~ brew install helm

➜  ~ helm repo add stable http://mirror.azure.cn/kubernetes/charts

➜  ~ helm repo list
NAME  	URL
stable	http://mirror.azure.cn/kubernetes/charts/








➜  ~ minikube dashboard
🤔  正在验证 dashboard 运行情况 ...
🚀  Launching proxy ...
🤔  正在验证 proxy 运行状况 ...
🎉  Opening http://127.0.0.1:60593/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/ in your default browser...




➜  ~ minikube start --kubernetes-version=v1.22.3

```




``` s
➜  ~ helm search repo mysql
NAME                            	CHART VERSION	APP VERSION	DESCRIPTION
stable/mysql                    	1.6.9        	5.7.30     	DEPRECATED - Fast,


➜  ~ helm show values stable/mysql
## mysql image version
## ref: https://hub.docker.com/r/library/mysql/tags/
##
image: "mysql"
imageTag: "5.7.30"




➜  ~ helm install db stable/mysql
WARNING: This chart is deprecated
NAME: db
LAST DEPLOYED: Fri Sep  2 18:23:45 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
MySQL can be accessed via port 3306 on the following DNS name from within your cluster:
db-mysql.default.svc.cluster.local

To get your root password run:

    MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace default db-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo)

To connect to your database:

1. Run an Ubuntu pod that you can use as a client:

    kubectl run -i --tty ubuntu --image=ubuntu:16.04 --restart=Never -- bash -il

2. Install the mysql client:

    $ apt-get update && apt-get install mysql-client -y

3. Connect using the mysql cli, then provide your password:
    $ mysql -h db-mysql -p

To connect to your database directly from outside the K8s cluster:
    MYSQL_HOST=127.0.0.1
    MYSQL_PORT=3306

    # Execute the following command to route the connection:
    kubectl port-forward svc/db-mysql 3306

    mysql -h ${MYSQL_HOST} -P${MYSQL_PORT} -u root -p${MYSQL_ROOT_PASSWORD}



➜  ~ helm list
NAME	NAMESPACE	REVISION	UPDATED                             	STATUS  	CHART      	APP VERSION
db  	default  	1       	2022-09-02 18:23:45.638568 +0800 CST	deployed	mysql-1.6.9	5.7.30

➜  ~ kubectl get pods
NAME                        READY   STATUS    RESTARTS       AGE
db-mysql-7f4fdddfd5-qfnnt   1/1     Running   2 (107s ago)   58m





```