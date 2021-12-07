

```go
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```


```go
ubuntu@master:~$ vim hpa-demo.yaml
ubuntu@master:~$ kubectl create -f hpa-demo.yaml
deployment.apps/hpa-demo created

ubuntu@master:~$ kubectl get deploy
NAME       READY   UP-TO-DATE   AVAILABLE   AGE
hpa-demo   1/1     1            1           69s

ubuntu@master:~$ kubectl autoscale deployment hpa-demo --min=1 --max=10 --cpu-percent=5
horizontalpodautoscaler.autoscaling/hpa-demo autoscaled

ubuntu@master:~$ cd /etc/kubernetes/manifests
ubuntu@master:/etc/kubernetes/manifests$ ls
etcd.yaml  kube-apiserver.yaml  kube-controller-manager.yaml  kube-scheduler.yaml
ubuntu@master:/etc/kubernetes/manifests$ sudo vim kube-controller-manager.yaml

- --horizontal-pod-autoscaler-use-rest-clients=false



```