


```go
ubuntu@master:~$ vim pod-hook1.yaml
ubuntu@master:~$ kubectl apply -f pod-hook1.yaml
pod/hook-demo1 created
ubuntu@master:~$ kubectl get pods
NAME         READY   STATUS    RESTARTS   AGE
hook-demo1   1/1     Running   0          51s


ubuntu@master:~$ kubectl exec hook-demo1 -i -t /bin/bash
kubectl exec [POD] [COMMAND] is DEPRECATED and will be removed in a future version. Use kubectl exec [POD] -- [COMMAND] instead.
root@hook-demo1:/# cat /usr/share/message
hello from the postStart Handler

强制删除

ubuntu@master:~$ kubectl delete pod hook-demo1 --grace-period=0 --force
warning: Immediate deletion does not wait for confirmation that the running resource has been terminated. The resource may continue to run on the cluster indefinitely.
pod "hook-demo1" force deleted

ubuntu@master:~$ kubectl get pods
No resources found in default namespace.



```

