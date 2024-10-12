
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



```shell
brew install etcd

brew install cfssl

brew install goreman  ## 安装进程管理工具
```


```shell
 ~  etcd --version
etcd Version: 3.5.16
Git SHA: f20bbadd4
Go Version: go1.23.1
Go OS/Arch: darwin/amd64


 ~  etcdctl put key value
OK
 ~  etcdctl get key
key
value


 ~  etcdctl --endpoints=127.0.0.1:2379 put username x
OK
 ~  etcdctl --endpoints=127.0.0.1:2379 get username
username
x


➜  ~  go install github.com/mattn/goreman@latest  ## 安装进程管理工具

➜  ~ vim local-cluster-profile

```



local-cluster-profile


```shell
etcd1: etcd --name infra1 --listen-client-urls http://127.0.0.1:12379 --advertise-client-urls http://127.0.0.1:12379 --listen-peer-urls http://127.0.0.1:12380 --initial-advertise-peer-urls http://127.0.0.1:12380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr

etcd2: etcd --name infra2 --listen-client-urls http://127.0.0.1:22379 --advertise-client-urls http://127.0.0.1:22379 --listen-peer-urls http://127.0.0.1:22380 --initial-advertise-peer-urls http://127.0.0.1:22380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr

etcd3: etcd --name infra3 --listen-client-urls http://127.0.0.1:32379 --advertise-client-urls http://127.0.0.1:32379 --listen-peer-urls http://127.0.0.1:32380 --initial-advertise-peer-urls http://127.0.0.1:32380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr


```



动态服务发现，需要不同主机实现，单节点无法实现
```shell
# etcd1: 
etcd --name infra1 --data-dir /data/etcd \
--listen-client-urls http://192.168.105.15:12379 \
--advertise-client-urls http://192.168.105.15:12379 \
--listen-peer-urls http://192.168.105.15:12380 \
--initial-advertise-peer-urls http://192.168.105.15:12380 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster 'infra1=http://192.168.105.15:12380,infra2=http://127.0.0.1:22380' \
--initial-cluster-state new \
--discovery  https://discovery.etcd.io/0261ffeecde45258413d87c310672ec5
# --enable-pprof \
# --logger=zap \
# --log-outputs=stderr \

# etcd2: 
etcd --name infra2 --data-dir /data/etcd \
--listen-client-urls http://127.0.0.1:22379 \
--advertise-client-urls http://127.0.0.1:22379 \
--listen-peer-urls http://127.0.0.1:22380 \
--initial-advertise-peer-urls http://127.0.0.1:22380 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' \
--initial-cluster-state new \
--enable-pprof \
--logger=zap \
--log-outputs=stderr \
--discovery  https://discovery.etcd.io/0261ffeecde45258413d87c310672ec5

# etcd3: 
etcd --name infra3 --data-dir /data/etcd \
--listen-client-urls http://127.0.0.1:32379 \
--advertise-client-urls http://127.0.0.1:32379 \
--listen-peer-urls http://127.0.0.1:32380 \
--initial-advertise-peer-urls http://127.0.0.1:32380 \
--initial-cluster-token etcd-cluster-1 \
--initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' \
--initial-cluster-state new \
--enable-pprof \
--logger=zap \
--log-outputs=stderr \
--discovery  https://discovery.etcd.io/0261ffeecde45258413d87c310672ec5

```




```shell
➜ ~ goreman -f ./local-cluster-profile start

➜ ~ etcdctl --endpoints=localhost:12379 member list

➜ ~ etcd git:(main) ✗ etcdctl --endpoints=localhost:12379 member list
8211f1d0f64f3269, started, infra1, http://127.0.0.1:12380, http://127.0.0.1:12379, false
91bc3c398fb3c146, started, infra2, http://127.0.0.1:22380, http://127.0.0.1:22379, false
fd422379fda50e48, started, infra3, http://127.0.0.1:32380, http://127.0.0.1:32379, false


➜  ~ goreman run stop etcd1
➜  ~ goreman run restart etcd1

➜  ~ etcdctl put name x --endpoints=localhost:12379
OK
➜  ~ etcdctl get name --endpoints=localhost:12379
name
x
➜  ~ etcdctl get name --endpoints=localhost:22379
name
x



➜  ~ curl https://discovery.etcd.io/new\?size\=3
https://discovery.etcd.io/46db8aaec9f3fc018c61ed6cbfb25fc5


{
  "action": "get",
  "node": {
    "key": "/46db8aaec9f3fc018c61ed6cbfb25fc5",
    "dir": true,
    "modifiedIndex": 171733191,
    "createdIndex": 171733191
  }
}

```


## etcd 客户端 etcdctl 对 key 增删改查操作



``` shell
➜  etcd git:(main) ✗ etcdctl put /name1 x
OK
➜  etcd git:(main) ✗ etcdctl put /name2 y
OK
➜  etcd git:(main) ✗ etcdctl put /name3 z
OK
➜  etcd git:(main) ✗ etcdctl get /name1
/name1
x
➜  etcd git:(main) ✗ etcdctl get --prefix /name
/name1
x
/name2
y
/name3
z
➜  etcd git:(main) ✗ etcdctl get --prefix /name --limit 2
/name1
x
/name2
y

# 按顺序获取
➜  etcd git:(main) ✗ etcdctl get --from-key /name2
/name2
y
/name3
z
key
value
name
x
username
x

username
x
➜  etcd git:(main) ✗ etcdctl put username x2
OK
➜  etcd git:(main) ✗ etcdctl get username -w=json
{"header":{"cluster_id":14841639068965178418,"member_id":10276657743932975437,"revision":9,"raft_term":80},"kvs":[{"key":"dXNlcm5hbWU=","create_revision":5,"mod_revision":9,"version":2,"value":"eDI="}],"count":1}
➜  etcd git:(main) ✗ etcdctl get username --rev=8
username
x


➜  etcd git:(main) ✗ etcdctl get --prefix /name
/name1
x
/name2
y
/name3
z
➜  etcd git:(main) ✗ etcdctl del /name3
1
➜  etcd git:(main) ✗ etcdctl get /name3
➜  etcd git:(main) ✗ etcdctl get --prefix /name
/name1
x
/name2
y

➜  etcd git:(main) ✗ etcdctl del /name3 --prev-kv
1
/name3
z




➜  ~ etcdctl put /name xn
OK
➜  ~ etcdctl put /name X
OK
➜  ~ etcdctl del /name
1

➜  ~ etcdctl watch /name
PUT
/name
xn
PUT
/name
X
DELETE
/name



➜  ~ etcdctl lease grant 300
lease 694d92366935f918 granted with TTL(300s)
➜  ~ etcdctl put --lease=694d92366935f918 /name xn
OK
➜  ~ etcdctl lease keep-alive 694d92366935f918
lease 694d92366935f918 keepalived with TTL(300)
➜  ~ etcdctl lease timetolive 694d92366935f918
lease 694d92366935f918 granted with TTL(300s), remaining(170s)

➜  ~ etcdctl lease timetolive --keys 694d92366935f918
lease 694d92366935f918 already expired


➜  ~ etcdctl lease grant 300
lease 694d92366935f91c granted with TTL(300s)
➜  ~ etcdctl put --lease=694d92366935f91c /
name xn
OK
➜  ~ etcdctl put --lease=694d92366935f91c /name1 x
OK
➜  ~ etcdctl lease timetolive --keys 694d92366935f91c
lease 694d92366935f91c granted with TTL(300s), remaining(232s), attached keys([/name /name1])

➜  ~ etcdctl lease revoke 694d92366935f91c
```




ETCD 权限管理

``` shell
➜  ~ etcdctl user add root
Password of root:
Type password of root again for confirmation:
User root created



➜  ~ etcdctl auth enable --user=root:0000
Authentication Enabled
➜  ~ etcdctl auth status --user=root:0000

```





``` shell

docker-compose up -d


```
























Unix 关闭查看进程，杀掉进程
```shell
lsof -c etcd
lsof -i:2379
kill pid
```

