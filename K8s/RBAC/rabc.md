

``` go
ubuntu@master:~/certs$ openssl genrsa -out hfbpw.key 2048
Generating RSA private key, 2048 bit long modulus (2 primes)
......................+++++
..............+++++
e is 65537 (0x010001)
ubuntu@master:~/certs$ ls
hfbpw.key
ubuntu@master:~/certs$ openssl req -new -key hfbpw.key -out hfbpw.csr -subj "/CN=hfbpw/O=x"
ubuntu@master:~/certs$ ls
hfbpw.csr  hfbpw.key

ubuntu@master:~/certs$ ls /etc/kubernetes/pki/
apiserver-etcd-client.crt     apiserver-kubelet-client.key  ca.crt  front-proxy-ca.crt      front-proxy-client.key
apiserver-etcd-client.key     apiserver.crt                 ca.key  front-proxy-ca.key      sa.key


ubuntu@master:~/certs$ su root
Password:
root@master:/home/ubuntu/certs# openssl x509 -req -in hfbpw.csr -CA /etc/kubernetes/pki/ca.crt -CAkey /etc/kubernetes/pki/ca.key -CAcreateserial -out hfbpw.crt -days 500
Signature ok
subject=CN = hfbpw, O = x
Getting CA Private Key

root@master:/home/ubuntu/certs# ls
hfbpw.crt  hfbpw.csr  hfbpw.key


root@master:/home/ubuntu/certs# kubectl config set-credentials hfbpw --client-certificate=hfbpw.crt --client-key=hfbpw.key

User "hfbpw" set.

root@master:/home/ubuntu/certs# kubectl config set-context hfbpw-context --cluster=kubernetes --namespace=kube-system --user=hfbpw

Context "hfbpw-context" created.

ubuntu@master:~$ kubectl get pods --context=hfbpw-context
Error from server (Forbidden): pods is forbidden: User "hfbpw" cannot list resource "pods" in API group "" in the namespace "kube-system"


ubuntu@master:~$ mkdir useraccount
ubuntu@master:~$ cd useraccount/

ubuntu@master:~/useraccount$ vim hfbpw-role.yaml
ubuntu@master:~/useraccount$ kubectl create -f hfbpw-role.yaml 
role.rbac.authorization.k8s.io/hfbpw-role created

ubuntu@master:~/useraccount$ kubectl get role -n kube-system
NAME                                             CREATED AT
extension-apiserver-authentication-reader        2021-12-11T08:18:00Z
hfbpw-role                                       2021-12-16T14:49:39Z

ubuntu@master:~/useraccount$ vim hfbpw-rolebinding.yaml
ubuntu@master:~/useraccount$ kubectl create -f hfbpw-rolebinding.yaml
rolebinding.rbac.authorization.k8s.io/hfbpw-rolebinding created

ubuntu@master:~/useraccount$ kubectl get rolebinding -n kube-system
NAME                                                ROLE                                                  AGE
hfbpw-rolebinding                                   Role/hfbpw-role                                       8m55s

ubuntu@master:~/useraccount$ kubectl get pods --context=hfbpw-context

```



``` go
ubuntu@master:~/certs$ mkdir serviceaccount
ubuntu@master:~/certs$ ls
hfbpw.crt  hfbpw.csr  hfbpw.key  serviceaccount
ubuntu@master:~/certs$ cd serviceaccount/

ubuntu@master:~/certs/serviceaccount$ kubectl create sa hfbpw-sa -n kube-system
serviceaccount/hfbpw-sa created

ubuntu@master:~/certs/serviceaccount$ kubectl get sa -n kube-system

hfbpw-sa                             1         4m5s


ubuntu@master:~/certs/serviceaccount$ vim hfbpw-sa-role.yaml
ubuntu@master:~/certs/serviceaccount$ kubectl create -f hfbpw-sa-role.yaml
role.rbac.authorization.k8s.io/hfbpw-sa-role created
ubuntu@master:~/certs/serviceaccount$ kubectl get role -n kube-system
NAME                                             CREATED AT
hfbpw-role                                       2021-12-16T14:49:39Z
hfbpw-sa-role                                    2021-12-17T04:44:01Z


ubuntu@master:~/certs/serviceaccount$ vim hfbpw-sa-rolebinding.yaml
ubuntu@master:~/certs/serviceaccount$ kubectl create -f hfbpw-sa-rolebinding.yaml
rolebinding.rbac.authorization.k8s.io/hfbpw-sa-rolebinding created
ubuntu@master:~/certs/serviceaccount$ kubectl get rolebinding -n kube-system
NAME                                                ROLE                                                  AGE
hfbpw-rolebinding                                   Role/hfbpw-role                                       125m
hfbpw-sa-rolebinding                                Role/hfbpw-sa-role                                    43s





```