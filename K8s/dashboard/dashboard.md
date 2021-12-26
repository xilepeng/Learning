
## kubernetes-dashboard









vim recommended.yaml  添加：  type: NodePort
```go

kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  type: NodePort
  ports:
    - port: 443
      targetPort: 8443
  selector:
    k8s-app: kubernetes-dashboard


```
kubectl apply -f recommended.yaml


```go
root@node1:/home/ubuntu/share# kubectl get pods --all-namespaces
NAMESPACE              NAME                                        READY   STATUS    RESTARTS      AGE
kube-system            coredns-78fcd69978-hqq4v                    1/1     Running   0             7h12m
kube-system            coredns-78fcd69978-qm6qq                    1/1     Running   0             7h12m
kube-system            etcd-master                                 1/1     Running   0             7h13m
kube-system            kube-apiserver-master                       1/1     Running   0             7h13m
kube-system            kube-controller-manager-master              1/1     Running   0             42m
kube-system            kube-flannel-ds-chrqd                       1/1     Running   0             3h59m
kube-system            kube-flannel-ds-kkqf4                       1/1     Running   0             6h1m
kube-system            kube-proxy-frgrw                            1/1     Running   0             3h59m
kube-system            kube-proxy-h5p5z                            1/1     Running   0             7h12m
kube-system            kube-scheduler-master                       1/1     Running   1 (42m ago)   42m
kubernetes-dashboard   dashboard-metrics-scraper-c45b7869d-xhr24   1/1     Running   0             3h42m
kubernetes-dashboard   kubernetes-dashboard-576cb95f94-kdts4       1/1     Running   0             3h42m

root@master:/home/ubuntu# kubectl get pods --all-namespaces
NAMESPACE              NAME                                        READY   STATUS    RESTARTS       AGE
kube-system            coredns-78fcd69978-bpljs                    1/1     Running   0              5h8m
kube-system            coredns-78fcd69978-t4slw                    1/1     Running   0              5h8m
kube-system            etcd-master                                 1/1     Running   0              5h8m
kube-system            kube-apiserver-master                       1/1     Running   0              5h8m
kube-system            kube-controller-manager-master              1/1     Running   3 (139m ago)   3h41m
kube-system            kube-flannel-ds-8vsms                       1/1     Running   0              4h44m
kube-system            kube-flannel-ds-rp96q                       1/1     Running   0              4h44m
kube-system            kube-proxy-975pb                            1/1     Running   0              5h8m
kube-system            kube-proxy-kv488                            1/1     Running   0              4h55m
kube-system            kube-scheduler-master                       1/1     Running   4 (26m ago)    3h41m
kubernetes-dashboard   dashboard-metrics-scraper-c45b7869d-6z6df   1/1     Running   0              3h44m
kubernetes-dashboard   kubernetes-dashboard-576cb95f94-7lv7g       1/1     Running   0              3h44m


root@master:/home/ubuntu# kubectl get service -n kubernetes-dashboard
NAME                        TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)         AGE
dashboard-metrics-scraper   ClusterIP   10.106.241.133   <none>        8000/TCP        4h15m
kubernetes-dashboard        NodePort    10.105.201.93    <none>        443:31447/TCP   4h15m


root@master:/home/ubuntu# kubectl describe svc kubernetes-dashboard -n kubernetes-dashboard
Name:                     kubernetes-dashboard
Namespace:                kubernetes-dashboard
Labels:                   k8s-app=kubernetes-dashboard
Annotations:              <none>
Selector:                 k8s-app=kubernetes-dashboard
Type:                     NodePort
IP Family Policy:         SingleStack
IP Families:              IPv4
IP:                       10.105.201.93
IPs:                      10.105.201.93
Port:                     <unset>  443/TCP
TargetPort:               8443/TCP
NodePort:                 <unset>  31447/TCP
Endpoints:                10.244.1.2:8443
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>


root@master:/home/ubuntu# kubectl get pods --all-namespaces -o wide
NAMESPACE              NAME                                        READY   STATUS    RESTARTS       AGE     IP              NODE     NOMINATED NODE   READINESS GATES
kube-system            coredns-78fcd69978-bpljs                    1/1     Running   0              5h39m   10.244.0.2      master   <none>           <none>
kube-system            coredns-78fcd69978-t4slw                    1/1     Running   0              5h39m   10.244.0.3      master   <none>           <none>
kube-system            etcd-master                                 1/1     Running   0              5h40m   192.168.105.5   master   <none>           <none>
kube-system            kube-apiserver-master                       1/1     Running   0              5h40m   192.168.105.5   master   <none>           <none>
kube-system            kube-controller-manager-master              1/1     Running   3 (170m ago)   4h13m   192.168.105.5   master   <none>           <none>
kube-system            kube-flannel-ds-8vsms                       1/1     Running   0              5h15m   192.168.105.6   node1    <none>           <none>
kube-system            kube-flannel-ds-rp96q                       1/1     Running   0              5h15m   192.168.105.5   master   <none>           <none>
kube-system            kube-proxy-975pb                            1/1     Running   0              5h39m   192.168.105.5   master   <none>           <none>
kube-system            kube-proxy-kv488                            1/1     Running   0              5h27m   192.168.105.6   node1    <none>           <none>
kube-system            kube-scheduler-master                       1/1     Running   4 (57m ago)    4h12m   192.168.105.5   master   <none>           <none>
kubernetes-dashboard   dashboard-metrics-scraper-c45b7869d-6z6df   1/1     Running   0              4h15m   10.244.1.3      node1    <none>           <none>
kubernetes-dashboard   kubernetes-dashboard-576cb95f94-7lv7g       1/1     Running   0              4h15m   10.244.1.2      node1    <none>           <none>
```


https://192.168.105.6:31447


root@node1:~# vim admin-account.yaml
```go

apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: admin
  namespace: kubernetes-dashboard

---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name:  cluster-admin
subjects:
  - kind: ServiceAccount
    name: admin
    namespace: kubernetes-dashboard
```



```go
从 master 拷贝到 node01
sudo vim /etc/kubernetes/admin.conf



root@node1:~# kubectl apply -f admin-account.yaml
serviceaccount/admin created
clusterrolebinding.rbac.authorization.k8s.io/admin created

ubuntu@node1:~$ kubectl get serviceaccount -n kubernetes-dashboard
NAME                   SECRETS   AGE
admin                  1         19s
default                1         6h38m
kubernetes-dashboard   1         6h38m
ubuntu@node1:~$ kubectl describe serviceaccount admin -n kubernetes-dashboard
Name:                admin
Namespace:           kubernetes-dashboard
Labels:              k8s-app=kubernetes-dashboard
Annotations:         <none>
Image pull secrets:  <none>
Mountable secrets:   admin-token-cptld
Tokens:              admin-token-cptld
Events:              <none>
ubuntu@node1:~$ kubectl get secret -n kubernetes-dashboard
NAME                               TYPE                                  DATA   AGE
admin-token-cptld                  kubernetes.io/service-account-token   3      45s
default-token-wj7kq                kubernetes.io/service-account-token   3      6h38m
kubernetes-dashboard-certs         Opaque                                0      6h38m
kubernetes-dashboard-csrf          Opaque                                1      6h38m
kubernetes-dashboard-key-holder    Opaque                                2      6h38m
kubernetes-dashboard-token-x5f4n   kubernetes.io/service-account-token   3      6h38m
ubuntu@node1:~$ kubectl describe secret admin-token-cptld -n kubernetes-dashboard
Name:         admin-token-cptld
Namespace:    kubernetes-dashboard
Labels:       <none>
Annotations:  kubernetes.io/service-account.name: admin
              kubernetes.io/service-account.uid: ffe26345-1dd4-4bfd-a8ff-5b94ebc388fa

Type:  kubernetes.io/service-account-token

Data
====
namespace:  20 bytes
token:      eyJhbGciOiJSUzI1NiIsImtpZCI6Il8zMTBqZXlMekRxZHBYcjNzNHdLVzRzTmxkdWkxZmpDcm5TQnVSRGFBUDQifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi10b2tlbi1jcHRsZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJhZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImZmZTI2MzQ1LTFkZDQtNGJmZC1hOGZmLTViOTRlYmMzODhmYSIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlcm5ldGVzLWRhc2hib2FyZDphZG1pbiJ9.BxP_FnFryfDt4K3F8Xa_rb-8JRrDlxhGV3w8cHwNct-85O3N6WDbjTes3seOva-TjRzfGJtBYjxXdAAAiF7uF6MujYrclqueD1VN1-9h-6X5hjyG3A53gjf0qrXW_Z9nMBfSlgHwmhGIeoaQZqBGGA02BGVMy-oUjyGqLtpBA0IimL7segh5HDSz4LbfHijkkGFdOTPkogMsBWJiL85Zd9ZqL5OWyz3DEOp_LPJQtpt-_3Bc8rS8Hey8iICUnz72Y433ePNxbeYnRR0V7mFGJpdSIANv0CQGOSPqFncXBHcjxDz67mAxUHW46tbA-tbdXJho_5eKmLIOEthLRC_U4A


ca.crt:     1099 bytes


```

