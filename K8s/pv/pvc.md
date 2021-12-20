

```go
root@master:~/kubeadm# vim nfs-pvc.yaml
root@master:~/kubeadm# kubectl create -f nfs-pvc.yaml
persistentvolumeclaim/pvc-nfs created
root@master:~/kubeadm# kubectl get pvc
NAME      STATUS   VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS   AGE
pvc-nfs   Bound    pv01     1Gi        RWO                           41s

root@master:~/kubeadm# kubectl get pv
NAME     CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM             STORAGECLASS   REASON   AGE
pv-nfs   1Gi        RWO            Recycle          Available                                             169m
pv01     1Gi        RWO            Recycle          Bound       default/pvc-nfs                           6h23m
pv02     1Gi        RWO            Recycle          Available                                             6h23m



```