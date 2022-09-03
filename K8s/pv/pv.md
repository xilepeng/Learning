

``` go
systemctl stop firewalld
systemctl disable firewalld

安装nfs服务器端
sudo apt-get install nfs-kernel-server

安装nfs依赖
sudo apt-get update
sudo apt-get install nfs-common



ubuntu@master:/$ sudo mkdir data
ubuntu@master:/$ sudo mkdir /data/k8s

root@master:~# chmod 755 /data/k8s/


root@master:~# vim /etc/exports

/data/k8s *(rw,sync,no_root_squash)

root@master:~# systemctl start rpcbind
root@master:~# systemctl enable rpcbind


root@master:~# systemctl status rpcbind
● rpcbind.service - RPC bind portmap service
     Loaded: loaded (/lib/systemd/system/rpcbind.service; enabled; vendor preset: enabled)
     Active: active (running) since Mon 2021-12-20 10:36:27 CST; 21min ago
TriggeredBy: ● rpcbind.socket
       Docs: man:rpcbind(8)
   Main PID: 564986 (rpcbind)
      Tasks: 1 (limit: 4682)
     Memory: 668.0K
     CGroup: /system.slice/rpcbind.service
             └─564986 /sbin/rpcbind -f -w

Dec 20 10:36:27 master systemd[1]: Starting RPC bind portmap service...
Dec 20 10:36:27 master systemd[1]: Started RPC bind portmap service.



root@master:~# service nfs-server start


root@master:~# service nfs-server status
● nfs-server.service - NFS server and services
     Loaded: loaded (/lib/systemd/system/nfs-server.service; enabled; vendor preset: enabled)
    Drop-In: /run/systemd/generator/nfs-server.service.d
             └─order-with-mounts.conf
     Active: active (exited) since Mon 2021-12-20 10:36:38 CST; 29min ago
   Main PID: 565765 (code=exited, status=0/SUCCESS)
      Tasks: 0 (limit: 4682)
     Memory: 0B
     CGroup: /system.slice/nfs-server.service

Dec 20 10:36:37 master systemd[1]: Starting NFS server and services...
Dec 20 10:36:38 master systemd[1]: Finished NFS server and services.


node1节点重复以上操作


root@master:~# showmount -e master
Export list for master:
/data/k8s *

root@master:~# showmount -e 192.168.105.5
Export list for 192.168.105.5:
/data/k8s *

root@node1:~# mkdir -p kubeadm/data

root@node1:~# df -h

16ae8cbfb99d199bdbfac84b23dd3f1e1c02375/merged
192.168.105.5:/data/k8s   39G  5.0G   34G  13% /root/kubeadm/data

root@node1:~# cd kubeadm/data
root@node1:~/kubeadm/data# vim test.txt
root@node1:~/kubeadm/data# cat test.txt
nfs

root@master:~# ls /data/k8s/
test.txt

root@master:~# cat /data/k8s/test.txt
nfs

root@master:~/kubeadm# vim pv-nfs.yaml
root@master:~/kubeadm# kubectl create -f pv-nfs.yaml
persistentvolume/pv-nfs created



```