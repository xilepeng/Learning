
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
1. [哨兵模式（自动选举老大）](#哨兵模式自动选举老大)
1. [第11章 缓存设计与优化](#第11章-缓存设计与优化)


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

## 哨兵模式（自动选举老大）

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