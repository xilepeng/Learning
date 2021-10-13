
1. [事务（重点：五颗星*****，必须理解，必须掌握）](#事务重点五颗星必须理解必须掌握)
1. [9.7、重点研究一下事务的隔离性！！！](#97重点研究一下事务的隔离性)

**8. 存储引擎（了解内容）**

8.1 什么是存储引擎，有什么用呢？
	存储引擎是MySQL中特有的一个术语，其它数据库中没有。（Oracle中有，但是不叫这个名字）
	存储引擎这个名字高端大气上档次。
	实际上存储引擎是一个表存储/组织数据的方式。
	不同的存储引擎，表存储数据的方式不同。

8.2 怎么给表添加/指定“存储引擎”呢？

	show create table t_student;

	可以在建表的时候给表指定存储引擎。
```sql
	CREATE TABLE `t_student` (
	  `no` int(11) NOT NULL AUTO_INCREMENT,
	  `name` varchar(255) DEFAULT NULL,
	  `cno` int(11) DEFAULT NULL,
	  PRIMARY KEY (`no`),
	  KEY `cno` (`cno`),
	  CONSTRAINT `t_student_ibfk_1` FOREIGN KEY (`cno`) REFERENCES `t_class` (`classno`)
	) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8
```
	在建表的时候可以在最后小括号的")"的右边使用：
		ENGINE来指定存储引擎。
		CHARSET来指定这张表的字符编码方式。
	
		结论：
			mysql默认的存储引擎是：InnoDB
			mysql默认的字符编码方式是：utf8
	
	建表时指定存储引擎，以及字符编码方式。
```sql
	create table t_product(
		id int primary key,
		name varchar(255)
	)engine=InnoDB default charset=gbk;
```
8.3、怎么查看mysql支持哪些存储引擎呢？


```sql
mysql> show engines \G;
*************************** 1. row ***************************
      Engine: FEDERATED
     Support: NO
     Comment: Federated MySQL storage engine
Transactions: NULL
          XA: NULL
  Savepoints: NULL
*************************** 2. row ***************************
      Engine: MEMORY
     Support: YES
     Comment: Hash based, stored in memory, useful for temporary tables
Transactions: NO
          XA: NO
  Savepoints: NO
*************************** 3. row ***************************
      Engine: InnoDB
     Support: DEFAULT
     Comment: Supports transactions, row-level locking, and foreign keys
Transactions: YES
          XA: YES
  Savepoints: YES
*************************** 4. row ***************************
      Engine: PERFORMANCE_SCHEMA
     Support: YES
     Comment: Performance Schema
Transactions: NO
          XA: NO
  Savepoints: NO
*************************** 5. row ***************************
      Engine: MyISAM
     Support: YES
     Comment: MyISAM storage engine
Transactions: NO
          XA: NO
  Savepoints: NO
*************************** 6. row ***************************
      Engine: MRG_MYISAM
     Support: YES
     Comment: Collection of identical MyISAM tables
Transactions: NO
          XA: NO
  Savepoints: NO
*************************** 7. row ***************************
      Engine: BLACKHOLE
     Support: YES
     Comment: /dev/null storage engine (anything you write to it disappears)
Transactions: NO
          XA: NO
  Savepoints: NO
*************************** 8. row ***************************
      Engine: CSV
     Support: YES
     Comment: CSV storage engine
Transactions: NO
          XA: NO
  Savepoints: NO
*************************** 9. row ***************************
      Engine: ARCHIVE
     Support: YES
     Comment: Archive storage engine
Transactions: NO
          XA: NO
  Savepoints: NO
9 rows in set (0.00 sec)

ERROR: 
No query specified
```


8.4、关于mysql常用的存储引擎介绍一下

MyISAM存储引擎？
	它管理的表具有以下特征：
		使用三个文件表示每个表：
			格式文件 — 存储表结构的定义（mytable.frm）
			数据文件 — 存储表行的内容（mytable.MYD）
			索引文件 — 存储表上索引（mytable.MYI）：索引是一本书的目录，缩小扫描范围，提高查询效率的一种机制。
		可被转换为压缩、只读表来节省空间

		提示一下：
			对于一张表来说，只要是主键，
			或者加有unique约束的字段上会自动创建索引。

		MyISAM存储引擎特点：
			可被转换为压缩、只读表来节省空间
			这是这种存储引擎的优势！！！！
		
		MyISAM不支持事务机制，安全性低。

**InnoDB存储引擎？**
	这是mysql默认的存储引擎，同时也是一个重量级的存储引擎。
	InnoDB支持事务，支持数据库崩溃后自动恢复机制。
	InnoDB存储引擎最主要的特点是：非常安全。

	它管理的表具有下列主要特征：
		– 每个 InnoDB 表在数据库目录中以.frm 格式文件表示
		– InnoDB 表空间 tablespace 被用于存储表的内容（表空间是一个逻辑名称。表空间存储数据+索引。）

		– 提供一组用来记录事务性活动的日志文件
		– 用 COMMIT(提交)、SAVEPOINT 及ROLLBACK(回滚)支持事务处理
		– 提供全 ACID 兼容
		– 在 MySQL 服务器崩溃后提供自动恢复
		– 多版本（MVCC）和行级锁定
		– 支持外键及引用的完整性，包括级联删除和更新
	
	InnoDB最大的特点就是支持事务：
		**以保证数据的安全。效率不是很高，并且也不能压缩，不能转换为只读，
		不能很好的节省存储空间。**

MEMORY存储引擎？
	使用 MEMORY 存储引擎的表，其数据存储在内存中，且行的长度固定，
	这两个特点使得 MEMORY 存储引擎非常快。

	MEMORY 存储引擎管理的表具有下列特征：
		– 在数据库目录内，每个表均以.frm 格式的文件表示。
		– 表数据及索引被存储在内存中。（目的就是快，查询快！）
		– 表级锁机制。
		– 不能包含 TEXT 或 BLOB 字段。

	MEMORY 存储引擎以前被称为HEAP 引擎。

	MEMORY引擎优点：查询效率是最高的。不需要和硬盘交互。
	MEMORY引擎缺点：不安全，关机之后数据消失。因为数据和索引都是在内存当中。


## 事务（重点：五颗星*****，必须理解，必须掌握）

9.1、什么是事务？

	一个事务其实就是一个完整的业务逻辑。
	是一个最小的工作单元。不可再分。

	什么是一个完整的业务逻辑？
		假设转账，从A账户向B账户中转账10000.
		将A账户的钱减去10000（update语句）
		将B账户的钱加上10000（update语句）
		这就是一个完整的业务逻辑。

		以上的操作是一个最小的工作单元，要么同时成功，要么同时失败，不可再分。
		这两个update语句要求必须同时成功或者同时失败，这样才能保证钱是正确的。
	
9.2、只有DML语句才会有事务这一说，其它语句和事务无关！！！
	insert
	delete
	update
	只有以上的三个语句和事务有关系，其它都没有关系。

	因为 只有以上的三个语句是数据库表中数据进行增、删、改的。
	只要你的操作一旦涉及到数据的增、删、改，那么就一定要考虑安全问题。

	数据安全第一位！！！

9.3、假设所有的业务，只要一条DML语句就能完成，还有必要存在事务机制吗？
	正是因为做某件事的时候，需要多条DML语句共同联合起来才能完成，
	所以需要事务的存在。如果任何一件复杂的事儿都能一条DML语句搞定，
	那么事务则没有存在的价值了。

	到底什么是事务呢？
		说到底，说到本质上，一个事务其实就是多条DML语句同时成功，或者同时失败！
	
	事务：就是批量的DML语句同时成功，或者同时失败！

9.4、事务是怎么做到多条DML语句同时成功和同时失败的呢？

	InnoDB存储引擎：提供一组用来记录事务性活动的日志文件

	事务开启了：
	insert
	insert
	insert
	delete
	update
	update
	update
	事务结束了！

	在事务的执行过程中，每一条DML的操作都会记录到“事务性活动的日志文件”中。
	在事务的执行过程中，我们可以提交事务，也可以回滚事务。

	提交事务？
		清空事务性活动的日志文件，将数据全部彻底持久化到数据库表中。
		提交事务标志着，事务的结束。并且是一种全部成功的结束。

	回滚事务？
		将之前所有的DML操作全部撤销，并且清空事务性活动的日志文件
		回滚事务标志着，事务的结束。并且是一种全部失败的结束。

9.5、怎么提交事务，怎么回滚事务？
	提交事务：commit; 语句
	回滚事务：rollback; 语句（回滚永远都是只能回滚到上一次的提交点！）

	事务对应的英语单词是：transaction

	测试一下，在mysql当中默认的事务行为是怎样的？
		mysql默认情况下是支持自动提交事务的。（自动提交）
		什么是自动提交？
			每执行一条DML语句，则提交一次！

		这种自动提交实际上是不符合我们的开发习惯，因为一个业务
		通常是需要多条DML语句共同执行才能完成的，为了保证数据
		的安全，必须要求同时成功之后再提交，所以不能执行一条
		就提交一条。
	
	怎么将mysql的自动提交机制关闭掉呢？

先执行这个命令：
```sql
start transaction;

ROLLBACK;

COMMIT;
```

```SQL
mysql> SELECT * FROM DEPT_BAK;
Empty set (0.01 sec)

mysql> START TRANSACTION;
Query OK, 0 rows affected (0.00 sec)

mysql> INSERT INTO DEPT_BAK (DEPTNO,DNAME,LOC) VALUES (1,'MOJO','BEIJING');
Query OK, 1 row affected (0.00 sec)

mysql> SELECT * FROM DEPT_BAK;
+--------+-------+---------+
| DEPTNO | DNAME | LOC     |
+--------+-------+---------+
|      1 | MOJO  | BEIJING |
+--------+-------+---------+
1 row in set (0.00 sec)

mysql> ROLLBACK;
Query OK, 0 rows affected (0.01 sec)

mysql> SELECT * FROM DEPT_BAK;
Empty set (0.00 sec)

mysql> INSERT INTO DEPT_BAK (DEPTNO,DNAME,LOC) VALUES (1,'X','BEIJING');
Query OK, 1 row affected (0.01 sec)

mysql> COMMIT;
Query OK, 0 rows affected (0.00 sec)

mysql> SELECT * FROM DEPT_BAK;
+--------+-------+---------+
| DEPTNO | DNAME | LOC     |
+--------+-------+---------+
|      1 | X     | BEIJING |
+--------+-------+---------+
1 row in set (0.00 sec)
```



9.6、事务包括4个特性？

	A：原子性
		说明事务是最小的工作单元。不可再分。

	C：一致性
		所有事务要求，在同一个事务当中，所有操作必须同时成功，或者同时失败，
		以保证数据的一致性。

	**I：隔离性**
		A事务和B事务之间具有一定的隔离。
		教室A和教室B之间有一道墙，这道墙就是隔离性。
		A事务在操作一张表的时候，另一个事务B也操作这张表会那样？？？

	D：持久性
		事务最终结束的一个保障。事务提交，就相当于将没有保存到硬盘上的数据
		保存到硬盘上！

## 9.7、重点研究一下事务的隔离性！！！

	A教室和B教室中间有一道墙，这道墙可以很厚，也可以很薄。这就是事务的隔离级别。
	这道墙越厚，表示隔离级别就越高。

	事务和事务之间的隔离级别有哪些呢？4个级别

		读未提交：read uncommitted（最低的隔离级别）《没有提交就读到了》
			什么是读未提交？
				事务A可以读取到事务B未提交的数据。
			这种隔离级别存在的问题就是：
				脏读现象！(Dirty Read)
				我们称读到了脏数据。
			这种隔离级别一般都是理论上的，大多数的数据库隔离级别都是二档起步！

		读已提交：read committed《提交之后才能读到》
			什么是读已提交？
				事务A只能读取到事务B提交之后的数据。
			这种隔离级别解决了什么问题？
				解决了脏读的现象。
			这种隔离级别存在什么问题？
				不可重复读取数据。
				什么是不可重复读取数据呢？
					在事务开启之后，第一次读到的数据是3条，当前事务还没有
					结束，可能第二次再读取的时候，读到的数据是4条，3不等于4
					称为不可重复读取。

			这种隔离级别是比较真实的数据，每一次读到的数据是绝对的真实。
			oracle数据库默认的隔离级别是：read committed

		可重复读：repeatable read《提交之后也读不到，永远读取的都是刚开启事务时的数据》
			什么是可重复读取？
				事务A开启之后，不管是多久，每一次在事务A中读取到的数据
				都是一致的。即使事务B将数据已经修改，并且提交了，事务A
				读取到的数据还是没有发生改变，这就是可重复读。
			可重复读解决了什么问题？
				解决了不可重复读取数据。
			可重复读存在的问题是什么？
				可以会出现幻影读。
				每一次读取到的数据都是幻象。不够真实！
			
			早晨9点开始开启了事务，只要事务不结束，到晚上9点，读到的数据还是那样！
			读到的是假象。不够绝对的真实。

			mysql中默认的事务隔离级别就是这个！！！！！！！！！！！

		序列化/串行化：serializable（最高的隔离级别）
			这是最高隔离级别，效率最低。解决了所有的问题。
			这种隔离级别表示事务排队，不能并发！
			synchronized，线程同步（事务同步）
			每一次读取到的数据都是最真实的，并且效率是最低的。


**9.8、验证各种隔离级别**



```sql
**mysql默认的隔离级别**

mysql> SELECT @@global.transaction_isolation;
+--------------------------------+
| @@global.transaction_isolation |
+--------------------------------+
| REPEATABLE-READ                |
+--------------------------------+
1 row in set (0.00 sec)

mysql> SELECT @@transaction_isolation;
+-------------------------+
| @@transaction_isolation |
+-------------------------+
| REPEATABLE-READ         |
+-------------------------+
1 row in set (0.00 sec)
```

被测试的表t_user






**验证：read uncommited**

------

```sql
mysql> set global transaction isolation level read uncommitted;
Query OK, 0 rows affected (0.00 sec)
```

**事务A**

```sql
mysql> SELECT @@global.transaction_isolation;
+--------------------------------+
| @@global.transaction_isolation |
+--------------------------------+
| READ-UNCOMMITTED               |
+--------------------------------+
1 row in set (0.00 sec)

mysql> use X;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> start transaction;
Query OK, 0 rows affected (0.00 sec)

mysql> select * from t_user;
Empty set (0.00 sec)

mysql> select * from t_user;
+------+
| name |
+------+
| mojo |
+------+
1 row in set (0.00 sec)


mysql> select * from t_user;
Empty set (0.00 sec)
```

**事务B**
```sql
mysql> use X;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> SELECT @@global.transaction_isolation;
+--------------------------------+
| @@global.transaction_isolation |
+--------------------------------+
| READ-UNCOMMITTED               |
+--------------------------------+
1 row in set (0.00 sec)

mysql> start transaction;
Query OK, 0 rows affected (0.00 sec)

mysql> insert into t_user values ('mojo');
Query OK, 1 row affected (0.01 sec)

mysql> rollback;
Query OK, 0 rows affected (0.01 sec)
```


------



**验证：read commited**

------
**事务A**

```sql
mysql> use X;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> start transaction;
Query OK, 0 rows affected (0.00 sec)

mysql> select * from t_user;
Empty set (0.00 sec)

mysql> select * from t_user;
Empty set (0.00 sec)

mysql> select * from t_user;
+------+
| name |
+------+
| mojo |
+------+
1 row in set (0.00 sec)
```


**事务B**
```sql
mysql> use X;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> start transaction;
Query OK, 0 rows affected (0.01 sec)

mysql> insert into t_user values ('mojo');
Query OK, 1 row affected (0.01 sec)

mysql> commit;
Query OK, 0 rows affected (0.01 sec)
```
------


**验证：repeatable read**

------
**事务A**


```sql
mysql> start transaction;
Query OK, 0 rows affected (0.00 sec)

mysql> select * from t_user;
+------+
| name |
+------+
| mojo |
+------+
1 row in set (0.00 sec)

// commit 后依然读不到,读到的是幻像

mysql> select * from t_user;
+------+
| name |
+------+
| mojo |
+------+
1 row in set (0.00 sec)
```

**事务B**
```sql
mysql> start transaction;
Query OK, 0 rows affected (0.00 sec)

mysql> insert into t_user values ('x'),('mojoman');
Query OK, 2 rows affected (0.00 sec)
Records: 2  Duplicates: 0  Warnings: 0

mysql> select * from t_user;
+---------+
| name    |
+---------+
| mojo    |
| x       |
| mojoman |
+---------+
3 rows in set (0.00 sec)

mysql> commit;
Query OK, 0 rows affected (0.01 sec)
```

---




**验证：serializable**
mysql> set global transaction isolation level serializable;
```sql
事务A												事务B
--------------------------------------------------------------------------------
use X;
													use X;
start transaction;
													start transaction;
select * from t_user;
insert into t_user values('abc');
													select * from t_user;
```