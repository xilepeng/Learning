

``` go
ubuntu@master:~$ echo -n "admin" | base64
YWRtaW4=
ubuntu@master:~$ echo -n "admin321" | base64
YWRtaW4zMjE=


ubuntu@master:~$ vim secret.yaml

ubuntu@master:~$ kubectl create -f secret.yaml
secret/mysecret created
ubuntu@master:~$ kubectl get secret
NAME                  TYPE                                  DATA   AGE
default-token-7kt56   kubernetes.io/service-account-token   3      4d19h
mysecret              Opaque                                2      52s

ubuntu@master:~$ kubectl describe secret mysecret
Name:         mysecret
Namespace:    default
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
password:  8 bytes
username:  5 bytes

ubuntu@master:~$ kubectl get secret mysecret -o yaml
apiVersion: v1
data:
  password: YWRtaW4zMjE=
  username: YWRtaW4=
kind: Secret
metadata:
  creationTimestamp: "2021-12-16T04:13:25Z"
  name: mysecret
  namespace: default
  resourceVersion: "234463"
  uid: 41fc2dc7-3ff2-4c9a-8504-57514ccc8689
type: Opaque




```