

## 配置文件

```go
ubuntu@node1:~$ systemctl status kubelet
● kubelet.service - kubelet: The Kubernetes Node Agent
     Loaded: loaded (/lib/systemd/system/kubelet.service; enabled; vendor preset: enabled)
    Drop-In: /etc/systemd/system/kubelet.service.d
             └─10-kubeadm.conf
     Active: active (running) since Thu 2021-12-02 20:37:46 CST; 16h ago
       Docs: https://kubernetes.io/docs/home/
   Main PID: 635 (kubelet)
      Tasks: 16 (limit: 4682)
     Memory: 134.6M
     CGroup: /system.slice/kubelet.service
             └─635 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.con>


sudo vim /etc/systemd/system/kubelet.service.d/10-kubeadm.conf

Environment="KUBELET_SYSTEM_PODS_ARGS=--pod-manifest-path=/etc/kubernetes/manifests --allow-privileged=true"

sudo cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf


ubuntu@node1:/etc/kubernetes/manifests$ kubectl get pod
NAME               READY   STATUS    RESTARTS   AGE
static-pod-node1   1/1     Running   0          3h38m
static-web         1/1     Running   0          4h54m

ubuntu@node1:/etc/kubernetes/manifests$ kubectl delete pod static-pod-node1
pod "static-pod-node1" deleted

ubuntu@node1:/etc/kubernetes/manifests$ kubectl get pod
NAME               READY   STATUS    RESTARTS   AGE
static-pod-node1   0/1     Pending   0          2s
static-web         1/1     Running   0          4h56m

ubuntu@node1:/etc/kubernetes/manifests$ sudo mv static-pod.yaml /tmp/
ubuntu@node1:/etc/kubernetes/manifests$ kubectl get pod
NAME         READY   STATUS    RESTARTS   AGE
static-web   1/1     Running   0          5h

ubuntu@node1:/etc/kubernetes/manifests$ kubectl delete pod static-web
pod "static-web" deleted
ubuntu@node1:/etc/kubernetes/manifests$ docker ps |grep static

ubuntu@node1:/etc/kubernetes/manifests$ sudo mv /tmp/static-pod.yaml  .
ubuntu@node1:/etc/kubernetes/manifests$ kubectl get pods
NAME               READY   STATUS    RESTARTS   AGE
static-pod-node1   0/1     Pending   0          5s

ubuntu@master:~$ cd /etc/kubernetes/manifests
ubuntu@master:/etc/kubernetes/manifests$ ls
etcd.yaml  kube-apiserver.yaml  kube-controller-manager.yaml  kube-scheduler.yaml


```

