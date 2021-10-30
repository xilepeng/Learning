



****
```shell
# 设置kubectl别名：
ubuntu@master:~$ sudo snap alias microk8s.kubectl kubectl
Added:
  - microk8s.kubectl as kubectl


Warning: spec.template.spec.nodeSelector[beta.kubernetes.io/os]: deprecated since v1.14; use "kubernetes.io/os" instead
deployment.apps/kubernetes-dashboard created
service/dashboard-metrics-scraper created
deployment.apps/dashboard-metrics-scraper created

If RBAC is not enabled access the dashboard using the default token retrieved with:

token=$(microk8s kubectl -n kube-system get secret | grep default-token | cut -d " " -f1)
microk8s kubectl -n kube-system describe secret $token

In an RBAC enabled setup (microk8s enable RBAC) you need to create a user with restricted
permissions as shown in:
https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md



sudo usermod -a -G microk8s ubuntu
sudo chown -f -R ubuntu ~/.kube


WARNING:  IPtables FORWARD policy is DROP. Consider enabling traffic forwarding with: sudo iptables -P FORWARD ACCEPT
The change can be made persistent with: sudo apt-get install iptables-persistent
WARNING:  Docker is installed.
Add the following lines to /etc/docker/daemon.json:
{
    "insecure-registries" : ["localhost:32000"]
}
and then restart docker with: sudo systemctl restart docker
Building the report tarball
  Report tarball is at /var/snap/microk8s/2551/inspection-report-20211030_125322.tar.gz




kubectl set image -n kube-system deployment/metrics-server metrics-server=phperall/metrics-server:v0.5.0


kubectl describe pods metrics-server -n kube-system

metrics-server:
    Container ID:
    Image:         k8s.gcr.io/metrics-server/metrics-server:v0.5.0
    Image ID:
    Port:          443/TCP
    Host Port:     0/TCP
    Args:
      --cert-dir=/tmp
      --secure-port=443
      --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
      --kubelet-use-node-status-port
      --metric-resolution=15s
      --kubelet-insecure-tls
    State:          Waiting
      Reason:       ImagePullBackOff
    Ready:          False



kubelet  Failed to pull image "k8s.gcr.io/metrics-server/metrics-server:v0.5.0": rpc error: code = Unknown desc = failed to pull and unpack image "k8s.gcr.io/metrics-server/metrics-server:v0.5.0"


kubectl set image -n kube-system deployment/metrics-server metrics-server=registry.aliyuncs.com/google_containers/metrics-server-amd64:v0.3.6

 
 microk8s.stop&&microk8s.start

kubectl apply -f components.yaml

kubectl get pods --all-namespaces | grep metrics

```