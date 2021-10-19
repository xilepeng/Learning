
docker redis
```bash
docker exec -it redis /bin/bash
root@74db1bab4c23:/data# redis-cli
127.0.0.1:6379>
```



```bash
$ brew install redis

To start redis:
  brew services start redis
Or, if you don't want/need a background service you can just run:
  /usr/local/opt/redis/bin/redis-server /usr/local/etc/redis.conf

redis-cli

brew services list

brew services start redis

brew services stop redis
```


批量删除进程

```sql
➜  config ps -ef | grep redis-server
  501 93578     1   0  5:16下午 ??         3:37.51 redis-server *:7000 [cluster]
  501 95111     1   0  5:21下午 ??         3:31.00 redis-server *:7001 [cluster]
  501 95136     1   0  5:21下午 ??         3:30.11 redis-server *:7002 [cluster]
  501 95162     1   0  5:21下午 ??         3:26.85 redis-server *:7003 [cluster]
  501 95182     1   0  5:21下午 ??         3:25.01 redis-server *:7004 [cluster]
  501 95206     1   0  5:21下午 ??         3:24.43 redis-server *:7005 [cluster]
  501 67987 44918   0 12:47下午 ttys001    0:00.00 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn --exclude-dir=.idea --exclude-dir=.tox redis-server
➜  config kill 93578
➜  config kill 95111
```

```sql
➜  config ps -ef | grep redis-server | grep 700 | awk '{print $2}'
95136
95162
95182
95206

➜  config ps -ef | grep redis-server | grep 700 | awk '{print $2}' | xargs kill
```


```sql
➜  config ps -ef | grep redis-server
  501 68676 44918   0 12:50下午 ttys001    0:00.00 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn --exclude-dir=.idea --exclude-dir=.tox redis-server
```



```bash
redis-server xconfig/redis.conf
redis-cli

127.0.0.1:6379>
```

1. [第2章 API的理解和使用](#第2章-api的理解和使用)
1. [第3章 Redis客户端的使用](#第3章-redis客户端的使用)
1. [第4章 瑞士军刀Redis其他功能](#第4章-瑞士军刀redis其他功能)
1. [第5章 Redis持久化的取舍和选择](#第5章-redis持久化的取舍和选择)
1. [第6章 常见的持久化开发运维问题](#第6章-常见的持久化开发运维问题)
1. [第7章 Redis复制的原理与优化](#第7章-redis复制的原理与优化)
1. [第8章 Redis Sentinel](#第8章-redis-sentinel)
1. [第9章 初识Redis Cluster](#第9章-初识redis-cluster)
1. [第10章 深入Redis Cluster](#第10章-深入redis-cluster)
1. [第11章 缓存设计与优化](#第11章-缓存设计与优化)
1. [redis cli命令](#redis-cli命令)


## 第2章 API的理解和使用

Redis提供的5种数据结构字符串（string）、哈希（hash）、列表（list）、集合（set）、有序集合（zset）的数据模型、常用命令、典型应用场景。















## 第3章 Redis客户端的使用

```go
package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	fmt.Println("连接成功！")

	ok, err := conn.Do("SET", "hello", "world")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ok)

	v, err := redis.String(conn.Do("GET", "hello"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	defer conn.Close()
}

```











## 第4章 瑞士军刀Redis其他功能 

除了5种数据结构外，Redis还提供了诸如慢查询、Pipeline、Bitmap、HyperLogLog、发布订阅、GEO等附加功能，在这些功能的帮助下，Redis的应用场景更加丰富。



**HyperLogLog 基数统计**

```sql
127.0.0.1:6379> pfadd mykey a b c d e f g
(integer) 1
127.0.0.1:6379> PFCOUNT mykey
(integer) 7
127.0.0.1:6379> pfadd mykey2 f g j l h
(integer) 1
127.0.0.1:6379> PFCOUNT mykey2
(integer) 5
127.0.0.1:6379> PFMERGE mykey3 mykey mykey2
OK
127.0.0.1:6379> PFCOUNT mykey3
(integer) 10
```


**Bitmap 位图**

```sql
127.0.0.1:6379> setbit sign 0 0
(integer) 0
127.0.0.1:6379> setbit sign 1 1
(integer) 0
127.0.0.1:6379> setbit sign 2 1
(integer) 0
127.0.0.1:6379> setbit sign 3 1
(integer) 0
127.0.0.1:6379> setbit sign 4 1
(integer) 0
127.0.0.1:6379> setbit sign 5 1
(integer) 0
127.0.0.1:6379> setbit sign 6 0
(integer) 0

127.0.0.1:6379> getbit sign 6
(integer) 0

127.0.0.1:6379> bitcount sign
(integer) 5
```

**事务 multi**




Redis 事务本质：一组命令的集合，一个事务中的所有命令都会被序列化，在事务执行过程中，会按照顺序执行。

**一次性、顺序性、排他性**

```sql
--- 队列 SET SET SET 执行
```


1. Redis 单条命令保持原子性，但事务不保证原子性！
1. Redis 事务没有隔离级别的概念（所有的命令在事务中，并没有直接被执行，只有发起执行命令才会被执行）

**Redis事务：**
  - 开启事务: multi
  - 命令入队: set/get ...
  - 执行事务: exec


```sql
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> set k1 v1
QUEUED
127.0.0.1:6379(TX)> set k2 v2
QUEUED
127.0.0.1:6379(TX)> get k2
QUEUED
127.0.0.1:6379(TX)> set k3 v3
QUEUED
127.0.0.1:6379(TX)> exec
1) OK
2) OK
3) "v2"
4) OK
```


discard 取消事务

```sql
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> set k4 v4
QUEUED
127.0.0.1:6379(TX)> discard 
OK
127.0.0.1:6379> get k4
(nil)
127.0.0.1:6379>
```

编译型异常（代码有问题，命令有错）事务中所有命令都不会被执行
```sql
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> set k1 v1 k2 v2
QUEUED
127.0.0.1:6379(TX)> getset k3
(error) ERR wrong number of arguments for 'getset' command
127.0.0.1:6379(TX)> exec
(error) EXECABORT Transaction discarded because of previous errors.
127.0.0.1:6379> get k1
(nil)
```

运行时异常（1/0),如果存在语法性错误，其他命令正常执行，错误命令抛出异常
```sql
127.0.0.1:6379> set k1 "v1"
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> incr k1
QUEUED
127.0.0.1:6379(TX)> set k2 v2
QUEUED
127.0.0.1:6379(TX)> get k2
QUEUED
127.0.0.1:6379(TX)> exec
1) (error) ERR value is not an integer or out of range
2) OK
3) "v2"
```


**监控 watch**

**悲观锁**
  - 很悲观，什么时候都会出问题，无论做什么都会加锁

**乐观锁**
  - 很乐观，什么时候都不会出问题，所以不会加锁，更新数据的时候去判断一下，在此期间是否有人修改过这个数据
  - 获取version
  - 更新的时候比较version

Redis 监视测试

```sql
127.0.0.1:6379> set money 100
OK
127.0.0.1:6379> set out 0
OK
127.0.0.1:6379> watch money
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> decrby money 20
QUEUED
127.0.0.1:6379(TX)> incrby out 20
QUEUED
127.0.0.1:6379(TX)> exec
1) (integer) 80
2) (integer) 20
```


测试多线程修改值，监视失败，使用 watch 可以当做 redis 的乐观锁操作 


```bash
127.0.0.1:6379> watch money
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> decrby money 10
QUEUED
127.0.0.1:6379(TX)> incrby out 10
QUEUED
127.0.0.1:6379(TX)> exec #执行之前，另一个线程修改数据
(nil)

127.0.0.1:6379> unwatch
OK
127.0.0.1:6379> watch money
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> decrby money 1
QUEUED
127.0.0.1:6379(TX)> incrby money 1
QUEUED
127.0.0.1:6379(TX)> exec
1) (integer) 999
2) (integer) 1000
```

```bash
127.0.0.1:6379> get money
"80"
127.0.0.1:6379> set money 1000
OK
```


**订阅发布**

订阅端
```sql
127.0.0.1:6379> subscribe mojo # 订阅一个频道mojo
Reading messages... (press Ctrl-C to quit)
1) "subscribe"
2) "mojo"
3) (integer) 1

1) "message" # 消息
2) "mojo"    # 消息频道
3) "hello"   # 消息详情

1) "message"
2) "mojo"
3) "X"
```

发送端
```sql
127.0.0.1:6379> publish mojo "hello" # 发布者发送消息到频道
(integer) 1
127.0.0.1:6379> publish mojo "X"
(integer) 1
```



## 第5章 Redis持久化的取舍和选择

Redis的持久化功能有效避免因进程退出造成的数据丢失问题，本章将介绍介绍RDB和AOF两种持久化配置和运行流程，以及选择策略

**RDB (Redis Database)**


```sql
127.0.0.1:6379> config get dir
1) "dir"
2) "/usr/local/var/db/redis"
```


**AOF（Append Only File)** 将我们所有命令都记录下来


## 第6章 常见的持久化开发运维问题


## 第7章 Redis复制的原理与优化

**主从复制的作用：**
1. 数据冗余：
1. 故障恢复：
1. 负载均衡：
1. 高可用（集群）基石：


只配置从库，不配置主库

查看当前库的信息
```sql
127.0.0.1:6379> info replication
# Replication
role:master # 角色 
connected_slaves:0  # 没有从机
master_failover_state:no-failover
master_replid:1a7bde8f15fed3be8389358a001ddb8470b2c724
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0
```


redis-6379.conf 去掉所有注释和空格

```sql
/usr/local/etc/config $  cat redis-6379.conf | grep -v "#" | grep -v "^$"

bind 127.0.0.1
protected-mode yes
port 6379
tcp-backlog 511
timeout 0
tcp-keepalive 300
daemonize yes
supervised no
pidfile /var/run/redis_6379.pid
loglevel notice
logfile "6379.log"
databases 16
always-show-logo yes
stop-writes-on-bgsave-error yes
rdbcompression yes
rdbchecksum yes
dbfilename dump-6379.rdb
rdb-del-sync-files no
dir /opt/data

replica-serve-stale-data yes
replica-read-only yes
repl-diskless-sync no
repl-diskless-sync-delay 5
repl-diskless-load disabled
repl-disable-tcp-nodelay no
replica-priority 100
acllog-max-len 128
lazyfree-lazy-eviction no
lazyfree-lazy-expire no
lazyfree-lazy-server-del no
replica-lazy-flush no
lazyfree-lazy-user-del no
appendonly no
appendfilename "appendonly.aof"
appendfsync everysec
no-appendfsync-on-rewrite no
auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb
aof-load-truncated yes
aof-use-rdb-preamble yes
lua-time-limit 5000
slowlog-log-slower-than 10000
slowlog-max-len 128
latency-monitor-threshold 0
notify-keyspace-events ""
hash-max-ziplist-entries 512
hash-max-ziplist-value 64
list-max-ziplist-size -2
list-compress-depth 0
set-max-intset-entries 512
zset-max-ziplist-entries 128
zset-max-ziplist-value 64
hll-sparse-max-bytes 3000
stream-node-max-bytes 4096
stream-node-max-entries 100
activerehashing yes
client-output-buffer-limit normal 0 0 0
client-output-buffer-limit replica 256mb 64mb 60
client-output-buffer-limit pubsub 32mb 8mb 60
hz 10
dynamic-hz yes
aof-rewrite-incremental-fsync yes
rdb-save-incremental-fsync yes
jemalloc-bg-thread yes
```



```sql
/usr/local/etc/config $  cat redis-6380.conf | grep -v "#" | grep -v "^$"

bind 127.0.0.1
protected-mode yes
port 6380
tcp-backlog 511
timeout 0
tcp-keepalive 300
daemonize yes
supervised no
pidfile /var/run/redis_6380.pid
loglevel notice
logfile "6380.log"
databases 16
always-show-logo yes
stop-writes-on-bgsave-error yes
rdbcompression yes
rdbchecksum yes
dbfilename dump-6380.rdb
rdb-del-sync-files no
dir /opt/data
replicaof 127.0.0.1 6379

replica-serve-stale-data yes
replica-read-only yes
repl-diskless-sync no
repl-diskless-sync-delay 5
repl-diskless-load disabled
repl-disable-tcp-nodelay no
replica-priority 100
acllog-max-len 128
lazyfree-lazy-eviction no
lazyfree-lazy-expire no
lazyfree-lazy-server-del no
replica-lazy-flush no
lazyfree-lazy-user-del no
appendonly no
appendfilename "appendonly.aof"
appendfsync everysec
no-appendfsync-on-rewrite no
auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb
aof-load-truncated yes
aof-use-rdb-preamble yes
lua-time-limit 5000
slowlog-log-slower-than 10000
slowlog-max-len 128
latency-monitor-threshold 0
notify-keyspace-events ""
hash-max-ziplist-entries 512
hash-max-ziplist-value 64
list-max-ziplist-size -2
list-compress-depth 0
set-max-intset-entries 512
zset-max-ziplist-entries 128
zset-max-ziplist-value 64
hll-sparse-max-bytes 3000
stream-node-max-bytes 4096
stream-node-max-entries 100
activerehashing yes
client-output-buffer-limit normal 0 0 0
client-output-buffer-limit replica 256mb 64mb 60
client-output-buffer-limit pubsub 32mb 8mb 60
hz 10
dynamic-hz yes
aof-rewrite-incremental-fsync yes
rdb-save-incremental-fsync yes
jemalloc-bg-thread yes
```



启动3个 Redis 服务
```sql
redis-server /usr/local/etc/config/redis-6379.conf
redis-server /usr/local/etc/config/redis-6380.conf
redis-server /usr/local/etc/config/redis-6381.conf

ps -ef|grep redis

  501 65774     1   0  1:47下午 ??         0:02.01 /usr/local/opt/redis/bin/redis-server 127.0.0.1:6379
  501 79628     1   0  3:36下午 ??         0:00.19 redis-server 127.0.0.1:6380
  501 79653     1   0  3:36下午 ??         0:00.17 redis-server 127.0.0.1:6381
  501 79788 68861   0  3:37下午 ttys003    0:00.01 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn --exclude-dir=.idea --exclude-dir=.tox redis
```

**一主二从**

星型、 链表型（中间依旧是从机）
   


**redis-cli 主**
```sql
$ redis-cli -p 6379

127.0.0.1:6379> ping
PONG

127.0.0.1:6379> info replication
# Replication
role:master
connected_slaves:2
slave0:ip=127.0.0.1,port=6380,state=online,offset=1134,lag=1
slave1:ip=127.0.0.1,port=6381,state=online,offset=1134,lag=1
master_failover_state:no-failover
master_replid:deabae32f1c3410f7ecb4d2e2a34489e39491fcf
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:1134
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:1134
```


**redis-cli 从**
```sql
$ redis-cli -p 6380

127.0.0.1:6380> ping
PONG

127.0.0.1:6380> info replication
# Replication
role:slave
master_host:127.0.0.1
master_port:6379
master_link_status:up
master_last_io_seconds_ago:0
master_sync_in_progress:0
slave_read_repl_offset:1176
slave_repl_offset:1176
slave_priority:100
slave_read_only:1
replica_announced:1
connected_slaves:0
master_failover_state:no-failover
master_replid:deabae32f1c3410f7ecb4d2e2a34489e39491fcf
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:1176
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:1176
```


**redis-cli 从**

```sql
redis-cli -p 6381

127.0.0.1:6381> ping
PONG

127.0.0.1:6381> info replication
# Replication
role:slave
master_host:127.0.0.1
master_port:6379
master_link_status:up
master_last_io_seconds_ago:4
master_sync_in_progress:0
slave_read_repl_offset:1190
slave_repl_offset:1190
slave_priority:100
slave_read_only:1
replica_announced:1
connected_slaves:0
master_failover_state:no-failover
master_replid:deabae32f1c3410f7ecb4d2e2a34489e39491fcf
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:1190
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:15
repl_backlog_histlen:1176
```


`SLAVEOF 127.0.0.1 6382` **暂时添加从机**
 
```SQL
$ redis-server /usr/local/etc/config/redis-6382.conf

$ redis-cli -p 6382

127.0.0.1:6382> info replication

# Replication
role:master
connected_slaves:0
master_failover_state:no-failover
master_replid:fc22e6c54c67eef8bfbf82e6d6f30d32ce978a0b
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0

127.0.0.1:6382> SLAVEOF 127.0.0.1 6382
OK

127.0.0.1:6382> info replication
# Replication
role:slave
master_host:127.0.0.1
master_port:6382
master_link_status:down
master_last_io_seconds_ago:-1
master_sync_in_progress:0
slave_read_repl_offset:0
slave_repl_offset:0
master_link_down_since_seconds:-1
slave_priority:100
slave_read_only:1
replica_announced:1
connected_slaves:0
master_failover_state:no-failover
master_replid:fc22e6c54c67eef8bfbf82e6d6f30d32ce978a0b
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:0
second_repl_offset:-1
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0
```

**主机写**

```sql
127.0.0.1:6379> set k1 v1
OK
```

**从机读**
```sql
127.0.0.1:6380> keys *
1) "k1"

127.0.0.1:6380> get k1
"v1"

127.0.0.1:6380> set k2 v2
(error) READONLY You can't write against a read only replica.
```

**从机读**
```sql
127.0.0.1:6381> get k1
"v1"
```

**主机关闭：从机依然是从机，只能读不能写**

**主机回来：依然是主机写从机读**
```sql
127.0.0.1:6379> shutdown
not connected> exit
$ ps -ef|grep redis

  501 79628     1   0  3:36下午 ??         0:21.07 redis-server 127.0.0.1:6380
  501 79653     1   0  3:36下午 ??         0:21.12 redis-server 127.0.0.1:6381
```

```sql
127.0.0.1:6380> info replication

# Replication
role:slave
master_host:127.0.0.1
master_port:6379
master_link_status:up
master_last_io_seconds_ago:3
master_sync_in_progress:0
slave_read_repl_offset:168
slave_repl_offset:168
slave_priority:100
slave_read_only:1
replica_announced:1
connected_slaves:0
master_failover_state:no-failover
master_replid:8580184ed64d032578dbe46e8733621411a6799a
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:168
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:168
```

**复制原理**

master 传送整个文件到 slave,完成一次完全同步

**全量复制**：

**增量复制**：

只要重连 master ,一次全量复制将被自动执行


如果主机断开，群龙无首，`SLAVEOF NO ONE`手动设置自己为主机
```SQL 
127.0.0.1:6381> SLAVEOF NO ONE
OK
127.0.0.1:6381> INFO REPLICATION
# Replication
role:master
connected_slaves:0
master_failover_state:no-failover
master_replid:514e4a232ffe5feb8713008deee2c14234d56f24
master_replid2:8580184ed64d032578dbe46e8733621411a6799a
master_repl_offset:1425
second_repl_offset:1426
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:1425
```



```sql
127.0.0.1:6380> slaveof no one
OK
127.0.0.1:6380> info replication
# Replication
role:master
connected_slaves:1
slave0:ip=127.0.0.1,port=6381,state=online,offset=1607,lag=1
master_failover_state:no-failover
master_replid:fa39d1c1df4a3c77d68c4707d0b586022b7dd91c
master_replid2:8580184ed64d032578dbe46e8733621411a6799a
master_repl_offset:1607
second_repl_offset:1594
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:1607
```




## 第8章 Redis Sentinel

**技巧**
```sql

sed "s/7000/7001/g" redis-7000.conf > redis-7001.conf

echo "slaveof 127.0.0.1 7000" >> redis-7001.conf

cat redis-sentinel.conf | grep -v "#" | grep -v "^$"

cat redis-sentinel.conf | grep -v "#" | grep -v "^$" > redis-sentinel-26379.conf

sed "s/26379/26380/g" redis-sentinel-26379.conf > redis-sentinel-26380.conf
```







```sql
 x@192  /Users/X/config  redis-server redis-7000.conf
 x@192  /Users/X/config  redis-cli -p 7000 ping
PONG
 x@192  /Users/X/config  redis-server redis-7001.conf
 x@192  /Users/X/config  redis-server redis-7002.conf
 x@192  /Users/X/config  ps -ef | grep redis-server | grep 700
  501 65114     1   0 10:48下午 ??         0:00.40 redis-server 127.0.0.1:7000
  501 65290     1   0 10:49下午 ??         0:00.18 redis-server 127.0.0.1:7001
  501 65314     1   0 10:49下午 ??         0:00.16 redis-server 127.0.0.1:7002
 x@192  /Users/X/config  redis-cli -p 7000 info replication
# Replication
role:master
connected_slaves:2
slave0:ip=127.0.0.1,port=7001,state=online,offset=182,lag=1
slave1:ip=127.0.0.1,port=7002,state=online,offset=182,lag=1
master_failover_state:no-failover
master_replid:774cd605ab429270af9c644a6c72fb949a2507c1
master_replid2:0000000000000000000000000000000000000000
master_repl_offset:182
second_repl_offset:-1
repl_backlog_active:1
repl_backlog_size:1048576
repl_backlog_first_byte_offset:1
repl_backlog_histlen:182

 x@192  /Users/X/config  cat redis-sentinel.conf | grep -v "#" | grep -v "^$"

port 26379
daemonize no
pidfile /var/run/redis-sentinel.pid
logfile ""
dir /tmp
sentinel monitor mymaster 127.0.0.1 6379 2
sentinel down-after-milliseconds mymaster 30000
acllog-max-len 128
sentinel parallel-syncs mymaster 1
sentinel failover-timeout mymaster 180000
sentinel deny-scripts-reconfig yes
SENTINEL resolve-hostnames no
SENTINEL announce-hostnames no


 x@192  /Users/X/config  cat redis-sentinel.conf | grep -v "#" | grep -v "^$" > redis-sentinel-26379.conf
 x@192  /Users/X/config  ls
7000.log                  data                      redis-7002.conf           sentinel.conf
7001.log                  redis-7000.conf           redis-sentinel-26379.conf

```



```sql
code . 
vim redis-sentinel-26379.conf

port 26379
daemonize yes
pidfile "/var/run/redis-sentinel.pid"
logfile "26379.log"
dir "/private/tmp"
sentinel monitor mymaster 127.0.0.1 7002 2

```


```sql
 x@192  /Users/X/config  sed "s/26379/26380/g" redis-sentinel-26379.conf > redis-sentinel-26380.conf
 x@192  /Users/X/config  sed "s/26379/26381/g" redis-sentinel-26379.conf > redis-sentinel-26381.conf
 x@192  /Users/X/config  redis-sentinel redis-sentinel-26379.conf
 x@192  /Users/X/config  redis-sentinel redis-sentinel-26380.conf
 x@192  /Users/X/config  redis-sentinel redis-sentinel-26381.conf
 x@192  /Users/X/config  ps -ef | grep redis-sentinel
  501 12003     1   0  9:02上午 ??         0:02.19 redis-sentinel *:26379 [sentinel]
  501 13135     1   0  9:08上午 ??         0:00.10 redis-sentinel *:26380 [sentinel]
  501 13158     1   0  9:08上午 ??         0:00.05 redis-sentinel *:26381 [sentinel]
  501 13186 68861   0  9:08上午 ttys003    0:00.00 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn --exclude-dir=.idea --exclude-dir=.tox redis-sentinel
```

注意：**先拷贝，再执行** 

redis-sentinel redis-sentinel-26379.conf 

**否则拷贝 sentinel myid **
```sql
不能拷贝：sentinel myid 6b49391b3bfba5fdf2020826897cd988aeb846c4  

否则导致：sentinels=1
master0:name=mymaster,status=ok,address=127.0.0.1:7000,slaves=2,sentinels=1
```

```sql
 x@192  /Users/X/config  redis-cli -p 26379
127.0.0.1:26379> info
# Server
redis_version:6.2.6
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:c6f3693d1aced7d9
redis_mode:sentinel
os:Darwin 20.6.0 x86_64
arch_bits:64
multiplexing_api:kqueue
atomicvar_api:c11-builtin
gcc_version:4.2.1
process_id:14392
process_supervised:no
run_id:5addbfe11092c7933ad354d82bc5cf9d82a36a88
tcp_port:26379
server_time_usec:1634519695812819
uptime_in_seconds:27
uptime_in_days:0
hz:15
configured_hz:10
lru_clock:7129743
executable:/Users/x/config/redis-sentinel
config_file:/Users/x/config/redis-sentinel-26379.conf
io_threads_active:0

# Clients
connected_clients:3
cluster_connections:0
maxclients:10000
client_recent_max_input_buffer:48
client_recent_max_output_buffer:0
blocked_clients:0
tracking_clients:0
clients_in_timeout_table:0

# CPU
used_cpu_sys:0.149000
used_cpu_user:0.052356
used_cpu_sys_children:0.000000
used_cpu_user_children:0.000000

# Stats
total_connections_received:3
total_commands_processed:50
instantaneous_ops_per_sec:3
total_net_input_bytes:2588
total_net_output_bytes:1458
instantaneous_input_kbps:0.28
instantaneous_output_kbps:0.02
rejected_connections:0
sync_full:0
sync_partial_ok:0
sync_partial_err:0
expired_keys:0
expired_stale_perc:0.00
expired_time_cap_reached_count:0
expire_cycle_cpu_milliseconds:0
evicted_keys:0
keyspace_hits:0
keyspace_misses:0
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:0
total_forks:0
migrate_cached_sockets:0
slave_expires_tracked_keys:0
active_defrag_hits:0
active_defrag_misses:0
active_defrag_key_hits:0
active_defrag_key_misses:0
tracking_total_keys:0
tracking_total_items:0
tracking_total_prefixes:0
unexpected_error_replies:0
total_error_replies:0
dump_payload_sanitizations:0
total_reads_processed:49
total_writes_processed:48
io_threaded_reads_processed:0
io_threaded_writes_processed:0

# Sentinel
sentinel_masters:1
sentinel_tilt:0
sentinel_running_scripts:0
sentinel_scripts_queue_length:0
sentinel_simulate_failure_flags:0
master0:name=mymaster,status=ok,address=127.0.0.1:7000,slaves=2,sentinels=3
```



**连接 golang 客户端**

Go语言：go-redis客户端对sentinel模式下（非集群cluster）redis-server主从切换的支持

 通过调用NewFailoverClient函数可以创建一个能支持redis-server主从切换(sentinel模式下)的client, 基本用法如下：

redis主从节点

主127.0.0.1:7000
从127.0.0.1:7001
从127.0.0.1:7002
假如有3个sentinel实例依次为：127.0.0.1:26379，127.0.0.1:26380，127.0.0.1:26381；sentinel的配置如下，

redis-sentinel-26379.conf
```sql
port 26379
daemonize yes
pidfile "/var/run/redis-sentinel.pid"
logfile "26379.log"
dir "/private/tmp"
sentinel monitor mymaster 127.0.0.1 7000 2
```



```go
package main
 
import (
        "fmt"
        "github.com/go-redis/redis"
        "time"
)
 
func main() {
        client := redis.NewFailoverClient(&redis.FailoverOptions{
                MasterName:    "mymaster",
                SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"},
                Password:      "",
                DB:            0,
        })
 
        for {
                reply, err := client.Incr("pvcount").Result()
                fmt.Printf("reply=%v err=%v\n", reply, err)
                time.Sleep(1 * time.Second)
        }
 
}

```

**实验1：redis节点主从切换 **
故障转移

运行redissentinel.go，并在执行过程中，shutdown redis主节点127.0.0.1:7000，可以看到在reply等于132和133之间，客户端监测到了主从切换，并重新连接到新的主节点，这段时间大致等于sentinel配置down-after-milliseconds的时长。



```go
redis: 2021/10/18 09:15:53 sentinel.go:379: sentinel: discovered new sentinel="c15d5f8249d4c33a860cf6e2c80ff3aa680a1e59" for master="mymaster"
redis: 2021/10/18 09:15:53 sentinel.go:379: sentinel: discovered new sentinel="2f89646d5118ede604c27ba065b97b8821613837" for master="mymaster"
redis: 2021/10/18 09:15:53 sentinel.go:332: sentinel: new master="mymaster" addr="127.0.0.1:7000"
reply=109 err=<nil>
reply=110 err=<nil>
reply=130 err=<nil>
reply=131 err=<nil>
reply=132 err=<nil>
reply=0 err=EOF
reply=0 err=dial tcp 127.0.0.1:7000: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:7000: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:7000: connect: connection refused
redis: 2021/10/18 09:16:53 sentinel.go:332: sentinel: new master="mymaster" addr="127.0.0.1:7002"
reply=133 err=<nil>

```


**实验2：sentinel实例全部挂掉后，redis的读写操作**

如果把3个sentinel实例全部 kill掉，则go-redis会记录一条日志，而对redis的读写操作仍然正常。

```bash
 x@192  ~/config  ps -ef | grep redis-sentinel
  501 43868     1   0 11:26上午 ??         0:04.80 redis-sentinel *:26379 [sentinel]
  501 43890     1   0 11:26上午 ??         0:04.67 redis-sentinel *:26380 [sentinel]
  501 43914     1   0 11:26上午 ??         0:04.65 redis-sentinel *:26381 [sentinel]
  501 46306 39657   0 11:37上午 ttys005    0:00.01 grep --color=auto --exclude-dir=.bzr --exclude-dir=CVS --exclude-dir=.git --exclude-dir=.hg --exclude-dir=.svn --exclude-dir=.idea --exclude-dir=.tox redis-sentinel
 x@192  ~/config  kill 43868
 x@192  ~/config  kill 43890 
 x@192  ~/config  kill 43914
```

```sql

redis: 2021/10/18 11:37:14 pubsub.go:159: redis: discarding bad PubSub connection: EOF

```


**实验3：sentinel实例部分挂掉后（剩余实例数目多于配置的quorum），redis实例的主从切换**




如果只kill掉1个sentinel，则剩余两个sentinel还能正常监测redis主从切换。例如下面是kill掉127.0.0.1:26379之后，再kill redis主节点127.0.0.1:7000的实验：

```sql
reply=322 err=<nil>
reply=323 err=<nil>
reply=324 err=<nil>
reply=325 err=<nil>
reply=326 err=<nil>
reply=327 err=<nil>
reply=328 err=<nil>
reply=329 err=<nil>
reply=330 err=<nil>
reply=0 err=EOF
redis: 2020/10/30 22:08:19 sentinel.go:313: sentinel: GetMasterAddrByName name="mymaster" failed: EOF
redis: 2020/10/30 22:08:19 sentinel.go:313: sentinel: GetMasterAddrByName name="mymaster" failed: dial tcp 127.0.0.1:26379: connect: connection refused
redis: 2020/10/30 22:08:19 sentinel.go:287: sentinel: GetMasterAddrByName master="mymaster" failed: dial tcp 127.0.0.1:26379: connect: connection refused
redis: 2020/10/30 22:08:19 sentinel.go:379: sentinel: discovered new sentinel="9a023490096f5f87db1c7f445d883f56e75275db" for master="mymaster"
reply=0 err=dial tcp 127.0.0.1:6388: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6388: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6388: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6388: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6388: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6388: connect: connection refused
redis: 2020/10/30 22:08:24 sentinel.go:332: sentinel: new master="mymaster" addr="127.0.0.1:6398"
reply=331 err=<nil>
reply=332 err=<nil>
reply=333 err=<nil>
reply=334 err=<nil>


```

**实验4：sentinel实例都正常的情况下，进行redis实例迁移**

当3个sentinel都正常运行的情况下，动态增减从节点对客户端没有影响。例如增加一个从节点127.0.0.1:6378再把原从节点127.0.0.1:6388 shutdown，客户端没有影响。进而把主节点127.0.0.1:6398 shutdown，迫使主从切换，结果证明客户端能够切换到新的主节点127.0.0.1:6378上。


```sql
reply=525 err=<nil>
reply=526 err=<nil>
reply=527 err=<nil>
reply=528 err=<nil>
reply=529 err=<nil>
reply=530 err=<nil>
reply=531 err=<nil>
reply=0 err=EOF
reply=0 err=dial tcp 127.0.0.1:6398: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6398: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6398: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6398: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6398: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6398: connect: connection refused
reply=0 err=dial tcp 127.0.0.1:6398: connect: connection refused
redis: 2020/10/30 22:43:19 sentinel.go:332: sentinel: new master="mymaster" addr="127.0.0.1:6378"
reply=532 err=<nil>
reply=533 err=<nil>
reply=534 err=<nil>
reply=535 err=<nil>
reply=536 err=<nil>
reply=537 err=<nil>


```

小结：

通过上面的实验证明，**go-redis 的 NewFailoverClient 对 sentinel 模式下的 redis 应用是稳定可靠的**。



**主观下线/客观下线**

1. 主观下线：每个sentinel节点对redis节点失败的偏见
2. 客观下线：所有sentinel节点对redis节点失败“达成共识"(超过quorum个 统一下线)


**Sentinel领领导者选举**

Redis使用了Raft算法实 现领导者选举









## 第9章 初识Redis Cluster 

Redis Cluster是Redis 3提供的分布式解决方案，有效解决了Redis分布式方面的需求，同时它也是学习分布式存储的绝佳案例。本章将针对Redis Cluster的数据分布，搭建集群进行分析说明。

```sql
➜  ~ redis-cli -v
redis-cli 6.2.6

➜  ~ redis-cli --version
redis-cli 6.2.6
```


**说明：redis-cli --cluster help**
```sql
➜  ~ redis-cli --cluster help

Cluster Manager Commands:
  create         host1:port1 ... hostN:portN   #创建集群
                 --cluster-replicas <arg>      #从节点个数
  check          host:port                     #检查集群
                 --cluster-search-multiple-owners #检查是否有槽同时被分配给了多个节点
  info           host:port                     #查看集群状态
  fix            host:port                     #修复集群
                 --cluster-search-multiple-owners #修复槽的重复分配问题
                 --cluster-fix-with-unreachable-masters
  reshard        host:port                     #指定集群的任意一节点进行迁移slot，重新分slots
                 --cluster-from <arg>          #需要从哪些源节点上迁移slot，可从多个源节点完成迁移，以逗号隔开，传递的是节点的node id，还可以直接传递--from all，这样源节点就是集群的所有节点，不传递该参数的话，则会在迁移过程中提示用户输入
                 --cluster-to <arg>            #slot需要迁移的目的节点的node id，目的节点只能填写一个，不传递该参数的话，则会在迁移过程中提示用户输入
                 --cluster-slots <arg>         #需要迁移的slot数量，不传递该参数的话，则会在迁移过程中提示用户输入。
                 --cluster-yes                 #指定迁移时的确认输入
                 --cluster-timeout <arg>       #设置migrate命令的超时时间
                 --cluster-pipeline <arg>      #定义cluster getkeysinslot命令一次取出的key数量，不传的话使用默认值为10
                 --cluster-replace             #是否直接replace到目标节点
  rebalance      host:port                                      #指定集群的任意一节点进行平衡集群节点slot数量 
                 --cluster-weight <node1=w1...nodeN=wN>         #指定集群节点的权重
                 --cluster-use-empty-masters                    #设置可以让没有分配slot的主节点参与，默认不允许
                 --cluster-timeout <arg>                        #设置migrate命令的超时时间
                 --cluster-simulate                             #模拟rebalance操作，不会真正执行迁移操作
                 --cluster-pipeline <arg>                       #定义cluster getkeysinslot命令一次取出的key数量，默认值为10
                 --cluster-threshold <arg>                      #迁移的slot阈值超过threshold，执行rebalance操作
                 --cluster-replace                              #是否直接replace到目标节点
  add-node       new_host:new_port existing_host:existing_port  #添加节点，把新节点加入到指定的集群，默认添加主节点
                 --cluster-slave                                #新节点作为从节点，默认随机一个主节点
                 --cluster-master-id <arg>                      #给新节点指定主节点
  del-node       host:port node_id                              #删除给定的一个节点，成功后关闭该节点服务
  call           host:port command arg arg .. arg               #在集群的所有节点执行相关命令
                 --cluster-only-masters
                 --cluster-only-replicas
  set-timeout    host:port milliseconds                         #设置cluster-node-timeout
  import         host:port                                      #将外部redis数据导入集群
                 --cluster-from <arg>                           #将指定实例的数据导入到集群
                 --cluster-copy                                 #migrate时指定copy
                 --cluster-replace                              #migrate时指定replace
  backup         host:port backup_directory
  help           

For check, fix, reshard, del-node, set-timeout you can specify the host and port of any working node in the cluster.

Cluster Manager Options:
  --cluster-yes  Automatic yes to cluster commands prompts
```
注意：Redis Cluster最低要求是3个主节点，如果需要集群需要认证，则在最后加入 -a xx 即可。


```sql
➜  config cat redis-7000.conf
port 7000
daemonize yes
dir "/Users/X/config/data"
logfile "7000.log"
dbfilename "dump-7000.rdb"
cluster-enabled yes
cluster-config-file nodes-7000.conf
cluster-require-full-coverage no
➜  config redis-server redis-7000.conf
➜  config redis-server redis-7001.conf
➜  config redis-server redis-7002.conf
➜  config redis-server redis-7003.conf
➜  config redis-server redis-7004.conf
➜  config redis-server redis-7005.conf

➜  config ps -ef | grep redis-server
  501 77992     1   0  2:22下午 ??         0:01.01 redis-server *:7000 [cluster]
  501 78011     1   0  2:22下午 ??         0:01.00 redis-server *:7001 [cluster]
  501 78039     1   0  2:22下午 ??         0:00.96 redis-server *:7002 [cluster]
  501 78063     1   0  2:22下午 ??         0:00.93 redis-server *:7003 [cluster]
  501 78113     1   0  2:22下午 ??         0:00.86 redis-server *:7004 [cluster]
  501 78129     1   0  2:22下午 ??         0:00.85 redis-server *:7005 [cluster]


➜  config redis-cli -p 7000
127.0.0.1:7000> cluster nodes
922c869a43ca12bbf725e7051ae933acef977e9b :7000@17000 myself,master - 0 0 0 connected
127.0.0.1:7000> exit
```

**1. 创建集群主节点**
`redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002`

```sql
➜  config redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002

>>> Performing hash slots allocation on 3 nodes...
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
M: 922c869a43ca12bbf725e7051ae933acef977e9b 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
M: 15eb2b06f7b0951e6a2a88c79c727516fd651760 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
M: bd4d81e92785fa415c807342fba5d941e52c1d43 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
Can I set the above configuration? (type 'yes' to accept): yes
>>> Nodes configuration updated
>>> Assign a different config epoch to each node
>>> Sending CLUSTER MEET messages to join the cluster
Waiting for the cluster to join
..
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: 922c869a43ca12bbf725e7051ae933acef977e9b 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
M: 15eb2b06f7b0951e6a2a88c79c727516fd651760 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
M: bd4d81e92785fa415c807342fba5d941e52c1d43 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
```

**2. 创建集群主从节点**

`redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 --cluster-replicas 1`
说明：--cluster-replicas 参数为数字，1表示每个主节点需要1个从节点。

```sql
➜  config redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 --cluster-replicas 1
>>> Performing hash slots allocation on 6 nodes...
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
Adding replica 127.0.0.1:7004 to 127.0.0.1:7000
Adding replica 127.0.0.1:7005 to 127.0.0.1:7001
Adding replica 127.0.0.1:7003 to 127.0.0.1:7002
>>> Trying to optimize slaves allocation for anti-affinity
[WARNING] Some slaves are in the same host as their master
M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
M: 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
S: 72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003
   replicates 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
S: a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004
   replicates a293e248c9d0c282c2a0542392b0392d7e43c74f
S: 147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005
   replicates 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
Can I set the above configuration? (type 'yes' to accept): yes
>>> Nodes configuration updated
>>> Assign a different config epoch to each node
>>> Sending CLUSTER MEET messages to join the cluster
Waiting for the cluster to join

>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004
   slots: (0 slots) slave
   replicates a293e248c9d0c282c2a0542392b0392d7e43c74f
S: 147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
M: 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
S: 72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
```

通过该方式创建的带有从节点的机器不能够自己手动指定主节点，所以如果需要指定的话，需要自己手动指定，先使用①或③创建好主节点后，再通过④来处理。

```sql
➜  config redis-cli -p 7000
127.0.0.1:7000> cluster nodes
a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004@17004 slave a293e248c9d0c282c2a0542392b0392d7e43c74f 0 1634629362000 1 connected
a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000@17000 myself,master - 0 1634629362000 1 connected 0-5460
147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005@17005 slave 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 0 1634629364248 2 connected
8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001@17001 master - 0 1634629365259 2 connected 5461-10922
7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002@17002 master - 0 1634629362196 3 connected 10923-16383
72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003@17003 slave 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 0 1634629363229 3 connected
127.0.0.1:7000> cluster info
cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:6
cluster_size:3
cluster_current_epoch:6
cluster_my_epoch:1
cluster_stats_messages_ping_sent:1519
cluster_stats_messages_pong_sent:1534
cluster_stats_messages_sent:3053
cluster_stats_messages_ping_received:1529
cluster_stats_messages_pong_received:1519
cluster_stats_messages_meet_received:5
cluster_stats_messages_received:3053
```

**3. 添加集群主节点**

```bash
➜  config cat redis-7006.conf

port 7006
daemonize yes
dir "/Users/X/config/data"
logfile "7006.log"
dbfilename "dump-7006.rdb"
cluster-enabled yes
cluster-config-file nodes-7006.conf
cluster-require-full-coverage no

➜  config sed 's/7000/7007/g' redis-7000.conf > redis-7007.conf

➜  config redis-server redis-7006.conf
➜  config redis-server redis-7007.conf

➜  config ps -ef | grep redis-server
  501  7162     1   0  4:47下午 ??         0:00.60 redis-server *:7007 [cluster]
  501  7271     1   0  4:48下午 ??         0:00.45 redis-server *:7006 [cluster]
  501 88878     1   0  3:17下午 ??         0:25.86 redis-server *:7000 [cluster]
  501 88902     1   0  3:17下午 ??         0:25.86 redis-server *:7001 [cluster]
  501 88919     1   0  3:17下午 ??         0:25.78 redis-server *:7002 [cluster]
  501 88937     1   0  3:17下午 ??         0:25.66 redis-server *:7003 [cluster]
  501 88951     1   0  3:17下午 ??         0:25.78 redis-server *:7004 [cluster]
  501 88970     1   0  3:17下午 ??         0:25.73 redis-server *:7005 [cluster]
```


`config redis-cli --cluster add-node 127.0.0.1:7006 127.0.0.1:7000`

```sql
➜  config redis-cli --cluster add-node 127.0.0.1:7006 127.0.0.1:7000

>>> Adding node 127.0.0.1:7006 to cluster 127.0.0.1:7000
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004
   slots: (0 slots) slave
   replicates a293e248c9d0c282c2a0542392b0392d7e43c74f
S: 147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
M: 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
S: 72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
>>> Send CLUSTER MEET to node 127.0.0.1:7006 to make it join the cluster.
[OK] New node added correctly.
```
说明：为一个指定集群添加节点，需要先连到该集群的任意一个节点IP（127.0.0.1:7000），再把新节点加入。该2个参数的顺序有要求：新加入的节点放前


```sql
➜  config redis-cli -p 7006
127.0.0.1:7006> cluster info
cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:7
cluster_size:3
cluster_current_epoch:6
cluster_my_epoch:0
cluster_stats_messages_ping_sent:189
cluster_stats_messages_pong_sent:190
cluster_stats_messages_meet_sent:1
cluster_stats_messages_sent:380
cluster_stats_messages_ping_received:190
cluster_stats_messages_pong_received:190
cluster_stats_messages_received:380
127.0.0.1:7006> cluster nodes
147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005@17005 slave 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 0 1634634150934 2 connected
a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004@17004 slave a293e248c9d0c282c2a0542392b0392d7e43c74f 0 1634634150000 1 connected
465bd254b10b12945667df07e5ba211d57c8b978 127.0.0.1:7006@17006 myself,master - 0 1634634149000 0 connected
a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000@17000 master - 0 1634634150000 1 connected 0-5460
72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003@17003 slave 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 0 1634634151947 3 connected
7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002@17002 master - 0 1634634151000 3 connected 10923-16383
8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001@17001 master - 0 1634634148000 2 connected 5461-10922
```

**4. 添加集群从节点**

```sql
➜  config redis-cli --cluster add-node 127.0.0.1:7007 127.0.0.1:7000 --cluster-slave --cluster-master-id 465bd254b10b12945667df07e5ba211d57c8b978
>>> Adding node 127.0.0.1:7007 to cluster 127.0.0.1:7000
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004
   slots: (0 slots) slave
   replicates a293e248c9d0c282c2a0542392b0392d7e43c74f
S: 147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
M: 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
M: 465bd254b10b12945667df07e5ba211d57c8b978 127.0.0.1:7006
   slots: (0 slots) master
M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
S: 72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
>>> Send CLUSTER MEET to node 127.0.0.1:7007 to make it join the cluster.
Waiting for the cluster to join

>>> Configure node as replica of 127.0.0.1:7006.
[OK] New node added correctly.
```

说明：把 7007 节点加入到 7000 节点的集群中，并且当做node_id为 465bd254b10b12945667df07e5ba211d57c8b978 的从节点。如果不指定 --cluster-master-id 会随机分配到任意一个主节点。

```sql
➜  config redis-cli -p 7007 cluster nodes
7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002@17002 master - 0 1634634528000 3 connected 10923-16383
72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003@17003 slave 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 0 1634634531424 3 connected
8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001@17001 master - 0 1634634530410 2 connected 5461-10922
465bd254b10b12945667df07e5ba211d57c8b978 127.0.0.1:7006@17006 master - 0 1634634532438 0 connected
a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000@17000 master - 0 1634634529000 1 connected 0-5460
a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004@17004 slave a293e248c9d0c282c2a0542392b0392d7e43c74f 0 1634634530000 1 connected
92aa10c225bdd14f3034830d34873fd90969d50a 127.0.0.1:7007@17007 myself,slave 465bd254b10b12945667df07e5ba211d57c8b978 0 1634634530000 0 connected
147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005@17005 slave 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 0 1634634529000 2 connected
➜  config redis-cli -p 7007 cluster info
cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:8
cluster_size:3
cluster_current_epoch:6
cluster_my_epoch:0
cluster_stats_messages_ping_sent:110
cluster_stats_messages_pong_sent:112
cluster_stats_messages_meet_sent:1
cluster_stats_messages_sent:223
cluster_stats_messages_ping_received:112
cluster_stats_messages_pong_received:111
cluster_stats_messages_received:223
```

**5. 删除节点**


说明：指定IP、端口和node_id 来删除一个节点，从节点可以直接删除，有slot分配的主节点不能直接删除。删除之后，该节点会被shutdown。

```sql
➜  config redis-cli --cluster del-node 127.0.0.1:7007 92aa10c225bdd14f3034830d34873fd90969d50a
>>> Removing node 92aa10c225bdd14f3034830d34873fd90969d50a from cluster 127.0.0.1:7007
>>> Sending CLUSTER FORGET messages to the cluster...
>>> Sending CLUSTER RESET SOFT to the deleted node.

➜  config redis-cli --cluster del-node 127.0.0.1:7006 465bd254b10b12945667df07e5ba211d57c8b978
>>> Removing node 465bd254b10b12945667df07e5ba211d57c8b978 from cluster 127.0.0.1:7006
>>> Sending CLUSTER FORGET messages to the cluster...
>>> Sending CLUSTER RESET SOFT to the deleted node.

➜  config redis-cli --cluster del-node 127.0.0.1:7001 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
>>> Removing node 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 from cluster 127.0.0.1:7001
[ERR] Node 127.0.0.1:7001 is not empty! Reshard data away and try again.
```


注意：当被删除掉的节点重新起来之后不能自动加入集群，但其和主的复制不是正常的，可以通过该节点看到集群信息（通过其他正常节点已经看不到该被del-node节点的信息）。如果想要再次加入集群，则需要先在该节点执行cluster reset，再用add-node进行添加，进行增量同步复制。
到此，目前整个集群的状态如下：

```sql
➜  config redis-cli -p 7000
127.0.0.1:7000> cluster nodes

a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004@17004 slave a293e248c9d0c282c2a0542392b0392d7e43c74f 0 1634635406000 1 connected
a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000@17000 myself,master - 0 1634635407000 1 connected 0-5460
147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005@17005 slave 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 0 1634635404198 2 connected
8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001@17001 master - 0 1634635406224 2 connected 5461-10922
7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002@17002 master - 0 1634635407234 3 connected 10923-16383
72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003@17003 slave 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 0 1634635408246 3 connected

➜  config redis-cli -p 7006 cluster info
cluster_state:fail
cluster_slots_assigned:0
cluster_slots_ok:0
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:1
cluster_size:0
cluster_current_epoch:6
cluster_my_epoch:0
cluster_stats_messages_ping_sent:1227
cluster_stats_messages_pong_sent:1241
cluster_stats_messages_meet_sent:1
cluster_stats_messages_sent:2469
cluster_stats_messages_ping_received:1241
cluster_stats_messages_pong_received:1228
cluster_stats_messages_received:2469


➜  config redis-cli -p 7006 cluster nodes
465bd254b10b12945667df07e5ba211d57c8b978 127.0.0.1:7006@17006 myself,master - 0 1634635175000 0 connected
```


**6. 检查集群**

说明：任意连接一个集群节点，进行集群状态检查

```sql
➜  config redis-cli --cluster check 127.0.0.1:7006 --cluster-search-multiple-owners

127.0.0.1:7006 (465bd254...) -> 0 keys | 0 slots | 0 slaves.
[OK] 0 keys in 1 masters.
0.00 keys per slot on average.
>>> Performing Cluster Check (using node 127.0.0.1:7006)
M: 465bd254b10b12945667df07e5ba211d57c8b978 127.0.0.1:7006
   slots: (0 slots) master
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[ERR] Not all 16384 slots are covered by nodes.

>>> Check for multiple slot owners...
[OK] No multiple owners found.



➜  config redis-cli --cluster check 127.0.0.1:7000 --cluster-search-multiple-owners
127.0.0.1:7000 (a293e248...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:7001 (8add3cfb...) -> 0 keys | 5462 slots | 1 slaves.
127.0.0.1:7002 (7c71bccc...) -> 0 keys | 5461 slots | 1 slaves.
[OK] 0 keys in 3 masters.
0.00 keys per slot on average.
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004
   slots: (0 slots) slave
   replicates a293e248c9d0c282c2a0542392b0392d7e43c74f
S: 147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
M: 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
S: 72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
>>> Check for multiple slot owners...
[OK] No multiple owners found.
```

**7. 集群信息查看**

说明：检查key、slots、从节点个数的分配情况

```sql
➜  config redis-cli --cluster info 127.0.0.1:7005
127.0.0.1:7002 (7c71bccc...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:7000 (a293e248...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:7001 (8add3cfb...) -> 0 keys | 5462 slots | 1 slaves.
[OK] 0 keys in 3 masters.
0.00 keys per slot on average.
```

**8. 修复集群**

说明：修复集群和槽的重复分配问题

```sql
➜  config redis-cli --cluster fix 127.0.0.1:7005 --cluster-search-multiple-owners
127.0.0.1:7002 (7c71bccc...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:7000 (a293e248...) -> 0 keys | 5461 slots | 1 slaves.
127.0.0.1:7001 (8add3cfb...) -> 0 keys | 5462 slots | 1 slaves.
[OK] 0 keys in 3 masters.
0.00 keys per slot on average.
>>> Performing Cluster Check (using node 127.0.0.1:7005)
S: 147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: 72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
M: 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004
   slots: (0 slots) slave
   replicates a293e248c9d0c282c2a0542392b0392d7e43c74f
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
>>> Check for multiple slot owners...
[OK] No multiple owners found.
```

**9. 设置集群的超时时间**

说明：连接到集群的任意一节点来设置集群的超时时间参数cluster-node-timeout

```sql
➜  config redis-cli --cluster set-timeout 127.0.0.1:7005 1000
>>> Reconfiguring node timeout in every cluster node...
*** New timeout set for 127.0.0.1:7005
*** New timeout set for 127.0.0.1:7002
*** New timeout set for 127.0.0.1:7000
*** New timeout set for 127.0.0.1:7003
*** New timeout set for 127.0.0.1:7001
*** New timeout set for 127.0.0.1:7004
>>> New node timeout set. 6 OK, 0 ERR.
```

**10. 集群中执行相关命令**


```sql
redis-cli --cluster call 127.0.0.1:7005 config set requirepass cc
redis-cli -a cc --cluster call 127.0.0.1:7005 config set masterauth cc
redis-cli -a cc --cluster call 127.0.0.1:7005 config rewrite
```

说明：连接到集群的任意一节点来对整个集群的所有节点进行设置。

```sql
➜  config redis-cli --cluster call 127.0.0.1:7005 config set cluster-node-timeout 12000
>>> Calling config set cluster-node-timeout 12000
127.0.0.1:7005: OK
127.0.0.1:7002: OK
127.0.0.1:7000: OK
127.0.0.1:7003: OK
127.0.0.1:7001: OK
127.0.0.1:7004: OK
```

到此，相关集群的基本操作已经介绍完，现在说明集群迁移的相关操作。

Redis 6.0 新增了几个命令：

1. fix 的子命令：--cluster-fix-with-unreachable-masters

2. call的子命令：--cluster-only-masters、--cluster-only-replicas

3. 集群节点备份：backup

```sql
➜  config redis-cli --cluster backup 127.0.0.1:7005 /Users/X/config/backup
>>> Performing Cluster Check (using node 127.0.0.1:7005)
S: 147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: 72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
M: 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004
   slots: (0 slots) slave
   replicates a293e248c9d0c282c2a0542392b0392d7e43c74f
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
>>> Node 127.0.0.1:7002 -> Saving RDB...
SYNC sent to master, writing 176 bytes to '/Users/X/config/backup/redis-node-127.0.0.1-7002-7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f.rdb'
Transfer finished with success.
>>> Node 127.0.0.1:7000 -> Saving RDB...
SYNC sent to master, writing 176 bytes to '/Users/X/config/backup/redis-node-127.0.0.1-7000-a293e248c9d0c282c2a0542392b0392d7e43c74f.rdb'
Transfer finished with success.
>>> Node 127.0.0.1:7001 -> Saving RDB...
SYNC sent to master, writing 176 bytes to '/Users/X/config/backup/redis-node-127.0.0.1-7001-8add3cfbff39a147b31b975f4fe8eab1dafa24e9.rdb'
Transfer finished with success.
Saving cluster configuration to: /Users/X/config/backup/nodes.json
[OK] Backup created into: /Users/X/config/backup


➜  config cd backup
➜  backup ls
nodes.json
redis-node-127.0.0.1-7000-a293e248c9d0c282c2a0542392b0392d7e43c74f.rdb
redis-node-127.0.0.1-7001-8add3cfbff39a147b31b975f4fe8eab1dafa24e9.rdb
redis-node-127.0.0.1-7002-7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f.rdb
```






## 第10章 深入Redis Cluster 

本章将针对Redis Cluster的集群伸缩，请求路由，故障转移等方面进行分析说明。










**迁移相关**

**① 在线迁移slot ：在线把集群的一些slot从集群原来slot节点迁移到新的节点，即可以完成集群的在线横向扩容和缩容。有2种方式进行迁移**

一是根据提示来进行操作：

直接连接到集群的任意一节点
```sql
➜  config redis-cli --cluster reshard 127.0.0.1:7000
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
S: a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004
   slots: (0 slots) slave
   replicates a293e248c9d0c282c2a0542392b0392d7e43c74f
S: 147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
M: 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
S: 72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
How many slots do you want to move (from 1 to 16384)? 1
What is the receiving node ID? 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
Please enter all the source node IDs.
  Type 'all' to use all the nodes as source nodes for the hash slots.
  Type 'done' once you entered all the source nodes IDs.
Source node #1: a293e248c9d0c282c2a0542392b0392d7e43c74f
Source node #2: done

Ready to move 1 slots.
  Source nodes:
    M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
       slots:[0-5460] (5461 slots) master
       1 additional replica(s)
  Destination node:
    M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
       slots:[10923-16383] (5461 slots) master
       1 additional replica(s)
  Resharding plan:
    Moving slot 0 from a293e248c9d0c282c2a0542392b0392d7e43c74f
Do you want to proceed with the proposed reshard plan (yes/no)? yes
Moving slot 0 from 127.0.0.1:7000 to 127.0.0.1:7002:

```

二是根据参数进行操作：

```sql
redis-cli  --cluster reshard 127.0.0.1:7000 --cluster-from a293e248c9d0c282c2a0542392b0392d7e43c74f --cluster-to 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f --cluster-slots 10 --cluster-yes --cluster-timeout 5000 --cluster-pipeline 10 --cluster-replace
```

说明：连接到集群的任意一节点来对指定节点指定数量的slot进行迁移到指定的节点。 

```sql
➜  config redis-cli  --cluster reshard 127.0.0.1:7000 --cluster-from a293e248c9d0c282c2a0542392b0392d7e43c74f --cluster-to 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f --cluster-slots 10 --cluster-yes --cluster-timeout 5000 --cluster-pipeline 10 --cluster-replace
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
   slots:[1-5460] (5460 slots) master
   1 additional replica(s)
S: a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004
   slots: (0 slots) slave
   replicates a293e248c9d0c282c2a0542392b0392d7e43c74f
S: 147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
M: 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
   slots:[0],[10923-16383] (5462 slots) master
   1 additional replica(s)
S: 72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.

Ready to move 10 slots.
  Source nodes:
    M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
       slots:[1-5460] (5460 slots) master
       1 additional replica(s)
  Destination node:
    M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
       slots:[0],[10923-16383] (5462 slots) master
       1 additional replica(s)
  Resharding plan:
    Moving slot 1 from a293e248c9d0c282c2a0542392b0392d7e43c74f
    Moving slot 2 from a293e248c9d0c282c2a0542392b0392d7e43c74f
    Moving slot 3 from a293e248c9d0c282c2a0542392b0392d7e43c74f
    Moving slot 4 from a293e248c9d0c282c2a0542392b0392d7e43c74f
    Moving slot 5 from a293e248c9d0c282c2a0542392b0392d7e43c74f
    Moving slot 6 from a293e248c9d0c282c2a0542392b0392d7e43c74f
    Moving slot 7 from a293e248c9d0c282c2a0542392b0392d7e43c74f
    Moving slot 8 from a293e248c9d0c282c2a0542392b0392d7e43c74f
    Moving slot 9 from a293e248c9d0c282c2a0542392b0392d7e43c74f
    Moving slot 10 from a293e248c9d0c282c2a0542392b0392d7e43c74f
Moving slot 1 from 127.0.0.1:7000 to 127.0.0.1:7002:
Moving slot 2 from 127.0.0.1:7000 to 127.0.0.1:7002:
Moving slot 3 from 127.0.0.1:7000 to 127.0.0.1:7002:
Moving slot 4 from 127.0.0.1:7000 to 127.0.0.1:7002:
Moving slot 5 from 127.0.0.1:7000 to 127.0.0.1:7002:
Moving slot 6 from 127.0.0.1:7000 to 127.0.0.1:7002:
Moving slot 7 from 127.0.0.1:7000 to 127.0.0.1:7002:
Moving slot 8 from 127.0.0.1:7000 to 127.0.0.1:7002:
Moving slot 9 from 127.0.0.1:7000 to 127.0.0.1:7002:
Moving slot 10 from 127.0.0.1:7000 to 127.0.0.1:7002:
```


**② 平衡（rebalance）slot:**

1）平衡集群中各个节点的slot数量

```sql
➜  config redis-cli --cluster rebalance 127.0.0.1:7000
>>> Performing Cluster Check (using node 127.0.0.1:7000)
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
*** No rebalancing needed! All nodes are within the 2.00% threshold.
```


 2）根据集群中各个节点设置的权重等平衡slot数量（不执行，只模拟）

 ```sql
redis-cli  --cluster rebalance --cluster-weight a293e248c9d0c282c2a0542392b0392d7e43c74f=5 8add3cfbff39a147b31b975f4fe8eab1dafa24e9=4 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f=3 --cluster-simulate 127.0.0.1:7000
 ```


```sql
➜  config redis-cli  --cluster rebalance --cluster-weight a293e248c9d0c282c2a0542392b0392d7e43c74f=5 8add3cfbff39a147b31b975f4fe8eab1dafa24e9=4 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f=3 --cluster-simulate 127.0.0.1:7000
>>> Performing Cluster Check (using node 127.0.0.1:7000)
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
>>> Rebalancing across 3 nodes. Total weight = 12.00
Moving 1376 slots from 127.0.0.1:7002 to 127.0.0.1:7000
################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################################
Moving 1 slots from 127.0.0.1:7001 to 127.0.0.1:7000
#
```


**③ 导入集群**

```sql
➜  config vim redis-7009.conf
➜  config cat redis-7009.conf
port 7009
daemonize yes
dir "/Users/x/config/data"
logfile "7009.log"
dbfilename "dump-7009.rdb"
➜  config redis-server redis-7009.conf
➜  config ps -ef | grep redis-server | grep 7009
  501 42874     1   0  7:47下午 ??         0:00.05 redis-server *:7009
➜  config redis-cli --cluster import 127.0.0.1:7000 --cluster-from 127.0.0.1:7009 --cluster-replace
>>> Importing data from 127.0.0.1:7009 to cluster 127.0.0.1:7000
>>> Performing Cluster Check (using node 127.0.0.1:7000)
M: a293e248c9d0c282c2a0542392b0392d7e43c74f 127.0.0.1:7000
   slots:[11-5460] (5450 slots) master
   1 additional replica(s)
S: a7562ebcccbd8a5d177e02a7700dc0e65856dc06 127.0.0.1:7004
   slots: (0 slots) slave
   replicates a293e248c9d0c282c2a0542392b0392d7e43c74f
S: 147384aa85f3c9630b29ee55e2a426584d57d529 127.0.0.1:7005
   slots: (0 slots) slave
   replicates 8add3cfbff39a147b31b975f4fe8eab1dafa24e9
M: 8add3cfbff39a147b31b975f4fe8eab1dafa24e9 127.0.0.1:7001
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
M: 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f 127.0.0.1:7002
   slots:[0-10],[10923-16383] (5472 slots) master
   1 additional replica(s)
S: 72ce50cff5a2a4610ecdeba534aa0cf22541e47a 127.0.0.1:7003
   slots: (0 slots) slave
   replicates 7c71bccc1ae77038bbb43d8c2076c1c44e72ab4f
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
*** Importing 0 keys from DB 0
```

说明：外部Redis实例（9021）导入到集群中的任意一节点。



注意：测试下来发现参数--cluster-replace没有用，如果集群中已经包含了某个key，在导入的时候会失败，不会覆盖，只有清空集群key才能导入。








**哨兵模式（自动选举老大）**

后台监控主机，如果主机故障了根据票数**自动将从库转换为主库**


1. 配置哨兵配置文件
```sql
$ cd /usr/local/etc/config

vim sentinel.conf

sentinel monitor myredis 127.0.0.1 6379 1
```

2. 启动哨兵

```sql
redis-sentinel /usr/local/etc/config/sentinel.conf
```

```sql
8170:X 17 Oct 2021 17:56:14.051 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
8170:X 17 Oct 2021 17:56:14.051 # Redis version=6.2.6, bits=64, commit=00000000, modified=0, pid=8170, just started
8170:X 17 Oct 2021 17:56:14.051 # Configuration loaded
8170:X 17 Oct 2021 17:56:14.053 * monotonic clock: POSIX clock_gettime
                _._                                                  
           _.-``__ ''-._                                             
      _.-``    `.  `_.  ''-._           Redis 6.2.6 (00000000/0) 64 bit
  .-`` .-```.  ```\/    _.,_ ''-._                                  
 (    '      ,       .-`  | `,    )     Running in sentinel mode
 |`-._`-...-` __...-.``-._|'` _.-'|     Port: 26379
 |    `-._   `._    /     _.-'    |     PID: 8170
  `-._    `-._  `-./  _.-'    _.-'                                   
 |`-._`-._    `-.__.-'    _.-'_.-'|                                  
 |    `-._`-._        _.-'_.-'    |           https://redis.io       
  `-._    `-._`-.__.-'_.-'    _.-'                                   
 |`-._`-._    `-.__.-'    _.-'_.-'|                                  
 |    `-._`-._        _.-'_.-'    |                                  
  `-._    `-._`-.__.-'_.-'    _.-'                                   
      `-._    `-.__.-'    _.-'                                       
          `-._        _.-'                                           
              `-.__.-'                                               

8170:X 17 Oct 2021 17:56:14.066 # Sentinel ID is 0372f4e224f9d51efb90638ea333552a1998e46c
8170:X 17 Oct 2021 17:56:14.066 # +monitor master myredis 127.0.0.1 6379 quorum 1
8170:X 17 Oct 2021 17:56:14.093 * +slave slave 127.0.0.1:6381 127.0.0.1 6381 @ myredis 127.0.0.1 6379
8170:X 17 Oct 2021 17:56:14.095 * +slave slave 127.0.0.1:6380 127.0.0.1 6380 @ myredis 127.0.0.1 6379
```
**哨兵模式：**
主机宕机、自动选举主机、主机回来也只能归并到新主机下当做从机

优点：
  1. 哨兵集群，基于主从复制模式，所有的主从配置优点他都有
  1. 主从可以切换，故障可以转移，系统可用性就会更好
  1. 哨兵模式就是主从模式升级，手动到自动，更加健壮
  
缺点：
  1. redis 不好在线扩容
  1. 实现哨兵模式的配置其实很麻烦，里面有很多选择


## 第11章 缓存设计与优化

**缓存穿透（查不到）**
  用户想要查询一个数据，发现redis内存数据库没有（缓存没有命中），于是向持久层数据库查询，发现也没有，于是本次查询失败。当用户很多的时候，缓存都没有命中（秒杀场景），于是都去请求持久层数据库。这会给持久层数据库造成很大的压力，这时候就相当于出现了缓存穿透。


解决方案
**1. 布隆过滤器拦截**

布隆过滤器是一种数据结构，对所有可能查询的参数以 hash 形式存储，在控制层先进行校验，不符合则丢弃，从而避免了对底层存储系统的压力。

在访问缓存层和存储层之前，将存在的key用布隆过滤器提前保存起来，做第一层拦截。

**2. 缓存空对象**

存储层不命中后，仍然将空对象保留到缓存层 中，之后再访问这个数据将会从缓存中获取，这样就保护了后端数据源。

缓存空对象会有两个问题：

第一，空值做了缓存，意味着缓存层中存了 更多的键，需要更多的内存空间（如果是攻击，问题更严重），比较有效的 方法是针对这类数据设置一个较短的过期时间，让其自动剔除。

第二，缓存 层和存储层的数据会有一段时间窗口的不一致，可能会对业务有一定影响。




**缓存击穿（量太大，缓存过期）**

缓存击穿是指一个key非常热点，在不停扛着大并发，大并发集中对这一个点访问，当这个key在失效的瞬间，持续的大并发就击穿缓存，直接请求数据库，就像在一个屏障上凿开了一个洞。

当某个key过期的瞬间，有大量的请求并发访问，这类数据一般是热点数据，由于缓存过期，会同时访问数据库来查询最新数据，并且回写缓存，会导致数据库瞬间压力过大。

解决方案：

**1. 设置热点数据永不过期**

没有设置过期时间，所以不会产生热点key过期后产生的问题

**2. 加互斥锁**

分布式锁：加分布式锁，保证对于每个key同时只有一个线程去查询后端服务，其他线程没有获得分布式锁的权限，因此只需等待即可。这种方式将高并发压力转移到分布式锁，对分布式锁考验很大。


**缓存雪崩**

缓存雪崩：指某一个时间段，缓存集体过期失效，redis宕机。断电

由于缓存层承载着大量请求，有效地保护了存储层，但是如果缓存层由于某些原因不能提供服务，于是所有的请求都会达到存储层，存储层的调用量会暴增，造成存储层也会级联宕机的情况。

缓存雪崩的英文原意是stampeding herd（奔逃的野牛），指的是缓存层宕掉后，流量会像奔逃的野牛一样，打向后端存储。

解决方案：

1. redis高可用。
1. 限流降级。
1. 数据预热。




**补充：**


## redis cli命令

```sql
可执行文件	作用
redis-server 	启动redis
redis-cli	redis命令行工具
redis-benchmark	基准测试工具
redis-check-aof	AOF持久化文件检测工具和修复工具
redis-check-dump	RDB持久化文件检测工具和修复工具
redis-sentinel	启动redis-sentinel
本文重点介绍的redis-cli命令。

```

可以使用两种方式连接redis服务器。

第一种：交互式方式     
```sql
redis-cli -h {host} -p {port}方式连接，然后所有的操作都是在交互的方式实现，不需要再执行redis-cli了。

$redis-cli -h 127.0.0.1-p 6379

127.0.0.1：6379>set hello world

OK

127.0.0.1：6379>get hello

"world"
```
 

第二种方式：命令方式
```sql
redis-cli -h {host} -p {port} {command}直接得到命令的返回结果。

$redis-cli -h 127.0.0.1-p 6379 get hello

"world"
```
 

redis-cli包含很多参数，如-h，-p，要了解全部参数，可用redis-cli -help命令。

 

第一部分 命令方式

介绍一些重要参数以及使用场景。

**1. -r   代表将命令重复执行多次**

```sql
$redis-cli -r 3 ping

PONG

PONG

PONG
```
ping命令可用于检测redis实例是否存活，如果存活则显示PONG。

 

**2. -i**

每隔几秒(如果想用ms，如10ms则写0.01)执行一次命令，必须与-r一起使用。
```sql
$redis-cli -r 3 -i 1 ping

PONG

PONG

PONG
```
 
```sql
$redis-cli -r 10 -i 1 info|grep used_memory_human

used_memory_human:2.95G

.....................................

used_memory_human:2.95G
```
每隔1秒输出内存的使用量，一共输出10次。

```sql

$redis-cli -h ip -p port info server|grep process_id

process_id:999
```
获取redis的进程号999

 

**3. -x**

代表从标准输入读取数据作为该命令的最后一个参数。
```sql
$echo "world" |redis-cli -x set hello

Ok
```
 

**4. -c**

连接集群结点时使用，此选项可防止moved和ask异常。

**5. -a**

如配置了密码，可用a选项。

**6. --scan和--pattern**

用于扫描指定模式的键，相当于scan命令。

 

**7. --slave**

当当前客户端模拟成当前redis节点的从节点，可用来获取当前redis节点的更新操作。合理利用可用于记录当前连接redis节点的一些更新操作，这些更新可能是实开发业务时需要的数据。

**8. --rdb**

会请求redis实例生成并发送RDB持久化文件，保存在本地。可做定期备份。

**9. --pipe**

将命令封装成redis通信协议定义的数据格式，批量发送给redis执行。

**10. --bigkeys**

统计bigkey的分布，使用scan命令对redis的键进行采样，从中找到内存占用比较大的键，这些键可能是系统的瓶颈。

**11. --eval**

用于执行lua脚本

**12. --latency**

有三个选项，--latency、--latency-history、--latency-dist。它们可检测网络延迟，展现的形式不同。

**13. --stat**

可实时获取redis的重要统计信息。info命令虽然比较全，但这里可看到一些增加的数据，如requests（每秒请求数）

**14. --raw 和 --no-raw**

--no-raw 要求返回原始格式。--raw 显示格式化的效果。



