

```shell
brew install etcd

brew install cfssl

brew install goreman
```


``` shell 
mkdir -p /opt/etcd/{bin,cfg,ssl}
cd /opt/etcd

 /opt/etcd  ls
bin cfg ssl
 /opt/etcd  cd cfg
 /opt/etcd/cfg  vim docker-compose.yml
```

docker-compose.yml

``` yml
version: "3.0"

networks:
  etcd-net:           # 网络
    driver: bridge    # 桥接模式

volumes:
  etcd1_data:         # 挂载到本地的数据卷名
    driver: local
  etcd2_data:
    driver: local
  etcd3_data:
    driver: local
###
### etcd 其他环境配置见：https://doczhcn.gitbook.io/etcd/index/index-1/configuration
###
services:
  etcd1:
    image: bitnami/etcd:latest  # 镜像
    container_name: etcd1       # 容器名 --name
    restart: always             # 总是重启
    networks:
      - etcd-net                # 使用的网络 --network
    ports:                      # 端口映射 -p
      - "20000:2379"
      - "20001:2380"
    environment:                # 环境变量 --env
      - ALLOW_NONE_AUTHENTICATION=yes                       # 允许不用密码登录
      - ETCD_NAME=etcd1                                     # etcd 的名字
      - ETCD_INITIAL_ADVERTISE_PEER_URLS=http://etcd1:2380  # 列出这个成员的伙伴 URL 以便通告给集群的其他成员
      - ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380           # 用于监听伙伴通讯的URL列表
      - ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379         # 用于监听客户端通讯的URL列表
      - ETCD_ADVERTISE_CLIENT_URLS=http://etcd1:2379        # 列出这个成员的客户端URL，通告给集群中的其他成员
      - ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster             # 在启动期间用于 etcd 集群的初始化集群记号
      - ETCD_INITIAL_CLUSTER=etcd1=http://etcd1:2380,etcd2=http://etcd2:2380,etcd3=http://etcd3:2380 # 为启动初始化集群配置
      - ETCD_INITIAL_CLUSTER_STATE=new                      # 初始化集群状态
    volumes:
      - etcd1_data:/bitnami/etcd                            # 挂载的数据卷

  etcd2:
    image: bitnami/etcd:latest
    container_name: etcd2
    restart: always
    networks:
      - etcd-net
    ports:
      - "20002:2379"
      - "20003:2380"
    environment:
      - ALLOW_NONE_AUTHENTICATION=yes
      - ETCD_NAME=etcd2
      - ETCD_INITIAL_ADVERTISE_PEER_URLS=http://etcd2:2380
      - ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380
      - ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379
      - ETCD_ADVERTISE_CLIENT_URLS=http://etcd2:2379
      - ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster
      - ETCD_INITIAL_CLUSTER=etcd1=http://etcd1:2380,etcd2=http://etcd2:2380,etcd3=http://etcd3:2380
      - ETCD_INITIAL_CLUSTER_STATE=new
    volumes:
      - etcd2_data:/bitnami/etcd

  etcd3:
    image: bitnami/etcd:latest
    container_name: etcd3
    restart: always
    networks:
      - etcd-net
    ports:
      - "20004:2379"
      - "20005:2380"
    environment:
      - ALLOW_NONE_AUTHENTICATION=yes
      - ETCD_NAME=etcd3
      - ETCD_INITIAL_ADVERTISE_PEER_URLS=http://etcd3:2380
      - ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380
      - ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379
      - ETCD_ADVERTISE_CLIENT_URLS=http://etcd3:2379
      - ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster
      - ETCD_INITIAL_CLUSTER=etcd1=http://etcd1:2380,etcd2=http://etcd2:2380,etcd3=http://etcd3:2380
      - ETCD_INITIAL_CLUSTER_STATE=new
    volumes:
      - etcd3_data:/bitnami/etcd
```




``` shell
docker-compose up -d











```




``` shell
 ~  etcdctl put name x
OK
 ~  etcdctl get name
name
x


cd /opt/etcd/cfg
 /opt/etcd/cfg  vim etcd.yaml
```


``` yaml
# file: etcd.yaml
---
apiVersion: v1
kind: Service
metadata:
  name: etcd
  namespace: default
spec:
  type: ClusterIP
  clusterIP: None
  selector:
    app: etcd
  ##
  ## Ideally we would use SRV records to do peer discovery for initialization.
  ## Unfortunately discovery will not work without logic to wait for these to
  ## populate in the container. This problem is relatively easy to overcome by
  ## making changes to prevent the etcd process from starting until the records
  ## have populated. The documentation on statefulsets briefly talk about it.
  ##   https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#stable-network-id
  publishNotReadyAddresses: true
  ##
  ## The naming scheme of the client and server ports match the scheme that etcd
  ## uses when doing discovery with SRV records.
  ports:
  - name: etcd-client
    port: 2379
  - name: etcd-server
    port: 2380
  - name: etcd-metrics
    port: 8080
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  namespace: default
  name: etcd
spec:
  ##
  ## The service name is being set to leverage the service headlessly.
  ## https://kubernetes.io/docs/concepts/services-networking/service/#headless-services
  serviceName: etcd
  ##
  ## If you are increasing the replica count of an existing cluster, you should
  ## also update the --initial-cluster-state flag as noted further down in the
  ## container configuration.
  replicas: 3
  ##
  ## For initialization, the etcd pods must be available to eachother before
  ## they are "ready" for traffic. The "Parallel" policy makes this possible.
  podManagementPolicy: Parallel
  ##
  ## To ensure availability of the etcd cluster, the rolling update strategy
  ## is used. For availability, there must be at least 51% of the etcd nodes
  ## online at any given time.
  updateStrategy:
    type: RollingUpdate
  ##
  ## This is label query over pods that should match the replica count.
  ## It must match the pod template's labels. For more information, see the
  ## following documentation:
  ##   https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
  selector:
    matchLabels:
      app: etcd
  ##
  ## Pod configuration template.
  template:
    metadata:
      ##
      ## The labeling here is tied to the "matchLabels" of this StatefulSet and
      ## "affinity" configuration of the pod that will be created.
      ##
      ## This example's labeling scheme is fine for one etcd cluster per
      ## namespace, but should you desire multiple clusters per namespace, you
      ## will need to update the labeling schema to be unique per etcd cluster.
      labels:
        app: etcd
      annotations:
        ##
        ## This gets referenced in the etcd container's configuration as part of
        ## the DNS name. It must match the service name created for the etcd
        ## cluster. The choice to place it in an annotation instead of the env
        ## settings is because there should only be 1 service per etcd cluster.
        serviceName: etcd
    spec:
      ##
      ## Configuring the node affinity is necessary to prevent etcd servers from
      ## ending up on the same hardware together.
      ##
      ## See the scheduling documentation for more information about this:
      ##   https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity
      affinity:
        ## The podAntiAffinity is a set of rules for scheduling that describe
        ## when NOT to place a pod from this StatefulSet on a node.
        podAntiAffinity:
          ##
          ## When preparing to place the pod on a node, the scheduler will check
          ## for other pods matching the rules described by the labelSelector
          ## separated by the chosen topology key.
          requiredDuringSchedulingIgnoredDuringExecution:
          ## This label selector is looking for app=etcd
          - labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values:
                - etcd
            ## This topology key denotes a common label used on nodes in the
            ## cluster. The podAntiAffinity configuration essentially states
            ## that if another pod has a label of app=etcd on the node, the
            ## scheduler should not place another pod on the node.
            ##   https://kubernetes.io/docs/reference/labels-annotations-taints/#kubernetesiohostname
            topologyKey: "kubernetes.io/hostname"
      ##
      ## Containers in the pod
      containers:
      ## This example only has this etcd container.
      - name: etcd
        image: quay.io/coreos/etcd:v3.5.15
        imagePullPolicy: IfNotPresent
        ports:
        - name: etcd-client
          containerPort: 2379
        - name: etcd-server
          containerPort: 2380
        - name: etcd-metrics
          containerPort: 8080
        ##
        ## These probes will fail over TLS for self-signed certificates, so etcd
        ## is configured to deliver metrics over port 8080 further down.
        ##
        ## As mentioned in the "Monitoring etcd" page, /readyz and /livez were
        ## added in v3.5.12. Prior to this, monitoring required extra tooling
        ## inside the container to make these probes work.
        ##
        ## The values in this readiness probe should be further validated, it
        ## is only an example configuration.
        readinessProbe:
          httpGet:
            path: /readyz
            port: 8080
          initialDelaySeconds: 10
          periodSeconds: 5
          timeoutSeconds: 5
          successThreshold: 1
          failureThreshold: 30
        ## The values in this liveness probe should be further validated, it
        ## is only an example configuration.
        livenessProbe:
          httpGet:
            path: /livez
            port: 8080
          initialDelaySeconds: 15
          periodSeconds: 10
          timeoutSeconds: 5
          failureThreshold: 3
        env:
        ##
        ## Environment variables defined here can be used by other parts of the
        ## container configuration. They are interpreted by Kubernetes, instead
        ## of in the container environment.
        ##
        ## These env vars pass along information about the pod.
        - name: K8S_NAMESPACE
          valueFrom:
            fieldRef:
             fieldPath: metadata.namespace
        - name: HOSTNAME
          valueFrom:
            fieldRef:
             fieldPath: metadata.name
        - name: SERVICE_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.annotations['serviceName']
        ##
        ## Configuring etcdctl inside the container to connect to the etcd node
        ## in the container reduces confusion when debugging.
        - name: ETCDCTL_ENDPOINTS
          value: $(HOSTNAME).$(SERVICE_NAME):2379
        ##
        ## TLS client configuration for etcdctl in the container.
        ## These files paths are part of the "etcd-client-certs" volume mount.
        # - name: ETCDCTL_KEY
        #   value: /etc/etcd/certs/client/tls.key
        # - name: ETCDCTL_CERT
        #   value: /etc/etcd/certs/client/tls.crt
        # - name: ETCDCTL_CACERT
        #   value: /etc/etcd/certs/client/ca.crt
        ##
        ## Use this URI_SCHEME value for non-TLS clusters.
        - name: URI_SCHEME
          value: "http"
        ## TLS: Use this URI_SCHEME for TLS clusters.
        # - name: URI_SCHEME
        # value: "https"
        ##
        ## If you're using a different container, the executable may be in a
        ## different location. This example uses the full path to help remove
        ## ambiguity to you, the reader.
        ## Often you can just use "etcd" instead of "/usr/local/bin/etcd" and it
        ## will work because the $PATH includes a directory containing "etcd".
        command:
        - /usr/local/bin/etcd
        ##
        ## Arguments used with the etcd command inside the container.
        args:
        ##
        ## Configure the name of the etcd server.
        - --name=$(HOSTNAME)
        ##
        ## Configure etcd to use the persistent storage configured below.
        - --data-dir=/data
        ##
        ## In this example we're consolidating the WAL into sharing space with
        ## the data directory. This is not ideal in production environments and
        ## should be placed in it's own volume.
        - --wal-dir=/data/wal
        ##
        ## URL configurations are parameterized here and you shouldn't need to
        ## do anything with these.
        - --listen-peer-urls=$(URI_SCHEME)://0.0.0.0:2380
        - --listen-client-urls=$(URI_SCHEME)://0.0.0.0:2379
        - --advertise-client-urls=$(URI_SCHEME)://$(HOSTNAME).$(SERVICE_NAME):2379
        ##
        ## This must be set to "new" for initial cluster bootstrapping. To scale
        ## the cluster up, this should be changed to "existing" when the replica
        ## count is increased. If set incorrectly, etcd makes an attempt to
        ## start but fail safely.
        - --initial-cluster-state=new
        ##
        ## Token used for cluster initialization. The recommendation for this is
        ## to use a unique token for every cluster. This example parameterized
        ## to be unique to the namespace, but if you are deploying multiple etcd
        ## clusters in the same namespace, you should do something extra to
        ## ensure uniqueness amongst clusters.
        - --initial-cluster-token=etcd-$(K8S_NAMESPACE)
        ##
        ## The initial cluster flag needs to be updated to match the number of
        ## replicas configured. When combined, these are a little hard to read.
        ## Here is what a single parameterized peer looks like:
        ##   etcd-0=$(URI_SCHEME)://etcd-0.$(SERVICE_NAME):2380
        - --initial-cluster=etcd-0=$(URI_SCHEME)://etcd-0.$(SERVICE_NAME):2380,etcd-1=$(URI_SCHEME)://etcd-1.$(SERVICE_NAME):2380,etcd-2=$(URI_SCHEME)://etcd-2.$(SERVICE_NAME):2380
        ##
        ## The peer urls flag should be fine as-is.
        - --initial-advertise-peer-urls=$(URI_SCHEME)://$(HOSTNAME).$(SERVICE_NAME):2380
        ##
        ## This avoids probe failure if you opt to configure TLS.
        - --listen-metrics-urls=http://0.0.0.0:8080
        ##
        ## These are some configurations you may want to consider enabling, but
        ## should look into further to identify what settings are best for you.
        # - --auto-compaction-mode=periodic
        # - --auto-compaction-retention=10m
        ##
        ## TLS client configuration for etcd, reusing the etcdctl env vars.
        # - --client-cert-auth
        # - --trusted-ca-file=$(ETCDCTL_CACERT)
        # - --cert-file=$(ETCDCTL_CERT)
        # - --key-file=$(ETCDCTL_KEY)
        ##
        ## TLS server configuration for etcdctl in the container.
        ## These files paths are part of the "etcd-server-certs" volume mount.
        # - --peer-client-cert-auth
        # - --peer-trusted-ca-file=/etc/etcd/certs/server/ca.crt
        # - --peer-cert-file=/etc/etcd/certs/server/tls.crt
        # - --peer-key-file=/etc/etcd/certs/server/tls.key
        ##
        ## This is the mount configuration.
        volumeMounts:
        - name: etcd-data
          mountPath: /data
        ##
        ## TLS client configuration for etcdctl
        # - name: etcd-client-tls
        #   mountPath: "/etc/etcd/certs/client"
        #   readOnly: true
        ##
        ## TLS server configuration
        # - name: etcd-server-tls
        #   mountPath: "/etc/etcd/certs/server"
        #   readOnly: true
      volumes:
      ##
      ## TLS client configuration
      # - name: etcd-client-tls
      #   secret:
      #     secretName: etcd-client-tls
      #     optional: false
      ##
      ## TLS server configuration
      # - name: etcd-server-tls
      #   secret:
      #     secretName: etcd-server-tls
      #     optional: false
  ##
  ## This StatefulSet will uses the volumeClaimTemplate field to create a PVC in
  ## the cluster for each replica. These PVCs can not be easily resized later.
  volumeClaimTemplates:
  - metadata:
      name: etcd-data
    spec:
      accessModes: ["ReadWriteOnce"]
      ##
      ## In some clusters, it is necessary to explicitly set the storage class.
      ## This example will end up using the default storage class.
      # storageClassName: ""
      resources:
        requests:
          storage: 1Gi

```



``` shell


```