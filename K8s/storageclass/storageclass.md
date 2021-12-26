

[https://github.com/kubernetes-retired/external-storage/tree/master/nfs-client](https://github.com/kubernetes-retired/external-storage/tree/master/nfs-client)


```go
ubuntu@master:~$ mkdir storageclass
ubuntu@master:~$ cd storageclass/
ubuntu@master:~/storageclass$ vim rbac.yaml
ubuntu@master:~/storageclass$ vim deployment.yaml
ubuntu@master:~/storageclass$ vim class.yaml
ubuntu@master:~/storageclass$ ls
class.yaml  deployment.yaml  rbac.yaml

ubuntu@master:~/storageclass$ kubectl create -f .

storageclass.storage.k8s.io/managed-nfs-storage created
deployment.apps/nfs-client-provisioner created
serviceaccount/nfs-client-provisioner created
clusterrole.rbac.authorization.k8s.io/nfs-client-provisioner-runner created
clusterrolebinding.rbac.authorization.k8s.io/run-nfs-client-provisioner created
role.rbac.authorization.k8s.io/leader-locking-nfs-client-provisioner created
rolebinding.rbac.authorization.k8s.io/leader-locking-nfs-client-provisioner created


ubuntu@master:~/storageclass$ kubectl get sc
NAME                  PROVISIONER      RECLAIMPOLICY   VOLUMEBINDINGMODE   ALLOWVOLUMEEXPANSION   AGE
managed-nfs-storage   fuseim.pri/ifs   Delete          Immediate           false                  69s


```