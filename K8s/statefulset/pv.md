
```go
ubuntu@master:~$ vim pv-demo.yaml
ubuntu@master:~$ kubectl create -f pv-demo.yaml
persistentvolume/pv01 created
persistentvolume/pv02 created
ubuntu@master:~$ kubectl get pv
NAME   CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS   REASON   AGE
pv01   1Gi        RWO            Recycle          Available                                   34s
pv02   1Gi        RWO            Recycle          Available                                   34s
```