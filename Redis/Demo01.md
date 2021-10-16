
docker redis
```bash
docker exec -it redis /bin/bash
root@74db1bab4c23:/data# redis-cli
127.0.0.1:6379>
```


brew redis
```bash
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



```sql

```