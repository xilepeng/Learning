
```go
ubuntu@master:~/jenkins$ vim jenkins.yaml
ubuntu@master:~/jenkins$ vim pvc.yaml
ubuntu@master:~/jenkins$ vim rbac.yaml
ubuntu@master:~/jenkins$ ls
jenkins.yaml  pvc.yaml  rbac.yaml
ubuntu@master:~/jenkins$ kubectl create -f .

```

```shell
ubuntu@master:~/jenkins$ kubectl apply -f .
Warning: resource deployments/jenkins is missing the kubectl.kubernetes.io/last-applied-configuration annotation which is required by kubectl apply. kubectl apply should only be used on resources created declaratively by either kubectl create --save-config or kubectl apply. The missing annotation will be patched automatically.
deployment.apps/jenkins configured
Warning: resource services/jenkins is missing the kubectl.kubernetes.io/last-applied-configuration annotation which is required by kubectl apply. kubectl apply should only be used on resources created declaratively by either kubectl create --save-config or kubectl apply. The missing annotation will be patched automatically.
Warning: resource persistentvolumes/opspv is missing the kubectl.kubernetes.io/last-applied-configuration annotation which is required by kubectl apply. kubectl apply should only be used on resources created declaratively by either kubectl create --save-config or kubectl apply. The missing annotation will be patched automatically.
Warning: resource persistentvolumeclaims/opspvc is missing the kubectl.kubernetes.io/last-applied-configuration annotation which is required by kubectl apply. kubectl apply should only be used on resources created declaratively by either kubectl create --save-config or kubectl apply. The missing annotation will be patched automatically.
Warning: resource serviceaccounts/jenkins is missing the kubectl.kubernetes.io/last-applied-configuration annotation which is required by kubectl apply. kubectl apply should only be used on resources created declaratively by either kubectl create --save-config or kubectl apply. The missing annotation will be patched automatically.
serviceaccount/jenkins configured
role.rbac.authorization.k8s.io/jenkins created
rolebinding.rbac.authorization.k8s.io/jenkins created
Error from server (Invalid): error when applying patch:
{"metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"Service\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"jenkins\"},\"name\":\"jenkins\",\"namespace\":\"kube-ops\"},\"spec\":{\"ports\":[{\"name\":\"web\",\"nodePort\":31002,\"port\":31000,\"targetPort\":\"web\"},{\"name\":\"agent\",\"port\":31001,\"targetPort\":\"agent\"}],\"selector\":{\"app\":\"jenkins\"},\"type\":\"NodePort\"}}\n"}},"spec":{"$setElementOrder/ports":[{"port":31000},{"port":31001}],"ports":[{"name":"web","nodePort":31002,"port":31000,"targetPort":"web"},{"name":"agent","port":31001,"targetPort":"agent"}]}}
to:
Resource: "/v1, Resource=services", GroupVersionKind: "/v1, Kind=Service"
Name: "jenkins", Namespace: "kube-ops"
for: "jenkins.yaml": Service "jenkins" is invalid: [spec.ports[2].name: Duplicate value: "web", spec.ports[3].name: Duplicate value: "agent"]
Error from server (Invalid): error when applying patch:
{"metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"PersistentVolume\",\"metadata\":{\"annotations\":{},\"name\":\"opspv\"},\"spec\":{\"accessModes\":[\"ReadWriteMany\"],\"capacity\":{\"storage\":\"20Gi\"},\"nfs\":{\"path\":\"/mnt/jenkins\",\"server\":\"192.168.105.5\"},\"persistentVolumeReclaimPolicy\":\"Delete\"}}\n"}},"spec":{"nfs":{"path":"/mnt/jenkins"}}}
to:
Resource: "/v1, Resource=persistentvolumes", GroupVersionKind: "/v1, Kind=PersistentVolume"
Name: "opspv", Namespace: ""
for: "pvc.yaml": PersistentVolume "opspv" is invalid: spec.persistentvolumesource: Forbidden: spec.persistentvolumesource is immutable after creation
  core.PersistentVolumeSource{
  	... // 2 identical fields
  	HostPath:  nil,
  	Glusterfs: nil,
  	NFS: &core.NFSVolumeSource{
  		Server:   "192.168.105.5",
- 		Path:     "/mnt/jenkins",
+ 		Path:     "/data/k8s",
  		ReadOnly: false,
  	},
  	RBD:     nil,
  	Quobyte: nil,
  	... // 15 identical fields
  }

Error from server (Invalid): error when applying patch:
{"metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"v1\",\"kind\":\"PersistentVolumeClaim\",\"metadata\":{\"annotations\":{},\"name\":\"opspvc\",\"namespace\":\"kube-ops\"},\"spec\":{\"accessModes\":[\"ReadWriteMany\"],\"resources\":{\"requests\":{\"storage\":\"1Gi\"}}}}\n"}},"spec":{"resources":{"requests":{"storage":"1Gi"}}}}
to:
Resource: "/v1, Resource=persistentvolumeclaims", GroupVersionKind: "/v1, Kind=PersistentVolumeClaim"
Name: "opspvc", Namespace: "kube-ops"
for: "pvc.yaml": PersistentVolumeClaim "opspvc" is invalid: spec.resources.requests.storage: Forbidden: field can not be less than previous value

```


```shell
ubuntu@master:~/jenkins$ kubectl get svc -n kube-ops
NAME      TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)                          AGE
jenkins   NodePort   10.98.225.203   <none>        8080:30001/TCP,50000:32708/TCP   24h

ubuntu@master:~/jenkins$ kubectl describe svc jenkins -n kube-ops
Name:                     jenkins
Namespace:                kube-ops
Labels:                   app=jenkins
Annotations:              <none>
Selector:                 app=jenkins
Type:                     NodePort
IP Family Policy:         SingleStack
IP Families:              IPv4
IP:                       10.98.225.203
IPs:                      10.98.225.203
Port:                     web  8080/TCP
TargetPort:               web/TCP
NodePort:                 web  30001/TCP
Endpoints:
Port:                     agent  50000/TCP
TargetPort:               agent/TCP
NodePort:                 agent  32708/TCP
Endpoints:
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>

```




```shell

安装nfs服务器端
sudo apt-get install nfs-kernel-server

安装nfs依赖
sudo apt-get update
sudo apt-get install nfs-common



ubuntu@master:/$ sudo mkdir /mnt/jenkins


root@master:~# chmod 755 /mnt/jenkins


root@master:~# vim /etc/exports

/mnt *(rw,sync,no_root_squash)



cat > /etc/exports << EOF
/mnt *(rw,no_root_squash)
EOF


root@master:/home/ubuntu# mkdir /mnt/jenkins

root@master:/home/ubuntu#  chmod 755 /mnt/jenkins
root@master:/home/ubuntu# vim /etc/exports
root@master:/home/ubuntu# systemctl start rpcbind
root@master:/home/ubuntu# systemctl enable rpcbind


Synchronizing state of rpcbind.service with SysV service script with /lib/systemd/systemd-sysv-install.
Executing: /lib/systemd/systemd-sysv-install enable rpcbind
root@master:/home/ubuntu# systemctl status rpcbind
● rpcbind.service - RPC bind portmap service
     Loaded: loaded (/lib/systemd/system/rpcbind.service; enabled; vendor preset: enabled)
     Active: active (running) since Wed 2021-12-29 14:55:12 CST; 2 days ago
TriggeredBy: ● rpcbind.socket
       Docs: man:rpcbind(8)
   Main PID: 568 (rpcbind)
      Tasks: 1 (limit: 4682)
     Memory: 2.1M
     CGroup: /system.slice/rpcbind.service
             └─568 /sbin/rpcbind -f -w

Dec 29 14:55:12 master systemd[1]: Starting RPC bind portmap service...
Dec 29 14:55:12 master systemd[1]: Started RPC bind portmap service.


root@master:/home/ubuntu# service nfs-server start

```