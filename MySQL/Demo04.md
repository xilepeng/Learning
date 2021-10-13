
1. [事务（重点：五颗星*****，必须理解，必须掌握）](#事务重点五颗星必须理解必须掌握)
1. [9.7、重点研究一下事务的隔离性！！！](#97重点研究一下事务的隔离性)
1. [索引（index）](#索引index)
1. [4.2、数据库设计范式共有？ 3个。](#42数据库设计范式共有-3个)

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








## 索引（index）

1.1、什么是索引？
	索引是在数据库表的字段上添加的，是为了提高查询效率存在的一种机制。
	一张表的一个字段可以添加一个索引，当然，多个字段联合起来也可以添加索引。
	索引相当于一本书的目录，是为了缩小扫描范围而存在的一种机制。

	对于一本字典来说，查找某个汉字有两种方式：
		第一种方式：一页一页挨着找，直到找到为止，这种查找方式属于全字典扫描。
		效率比较低。
		第二种方式：先通过目录（索引）去定位一个大概的位置，然后直接定位到这个
		位置，做局域性扫描，缩小扫描的范围，快速的查找。这种查找方式属于通过
		索引检索，效率较高。
	
	t_user
```sql
	id(idIndex)	name(nameIndex)	email(emailIndex)		address  (emailAddressIndex)
	----------------------------------------------------------------------------------
	1				zhangsan...
	2				lisi
	3				wangwu
	4				zhaoliu
	5				hanmeimei
	6				jack

	select * from t_user where name = 'jack';
```
	以上的这条SQL语句会去name字段上扫描，为什么？
		因为查询条件是：name='jack'
	
	如果name字段上没有添加索引（目录），或者说没有给name字段创建索引，
	MySQL会进行全扫描，会将name字段上的每一个值都比对一遍。效率比较低。

	MySQL在查询方面主要就是两种方式：
		第一种方式：全表扫描
		第二种方式：根据索引检索。
	
	注意：
		在实际中，汉语字典前面的目录是排序的，按照a b c d e f....排序，
		为什么排序呢？因为只有排序了才会有区间查找这一说！（缩小扫描范围
		其实就是扫描某个区间罢了！）

		在mysql数据库当中索引也是需要排序的，并且这个所以的排序和TreeSet
		数据结构相同。TreeSet（TreeMap）底层是一个自平衡的二叉树！在mysql
		当中索引是一个B-Tree数据结构。

		遵循左小又大原则存放。采用中序遍历方式遍历取数据。

1.2、索引的实现原理？

	假设有一张用户表：t_user
```sql
	id(PK)					name						每一行记录在硬盘上都有物理存储编号
	----------------------------------------------------------------------------------
	100						zhangsan					0x1111
	120						lisi						0x2222
	99							wangwu					0x8888
	88							zhaoliu					0x9999
	101						jack						0x6666
	55							lucy				    0x5555
	130						tom						    0x7777
```
	提醒1：在任何数据库当中主键上都会自动添加索引对象，id字段上自动有索引，
	因为id是PK。另外在mysql当中，一个字段上如果有unique约束的话，也会自动
	创建索引对象。

	提醒2：在任何数据库当中，任何一张表的任何一条记录在硬盘存储上都有
	一个硬盘的物理存储编号。

	提醒3：在mysql当中，索引是一个单独的对象，不同的存储引擎以不同的形式
	存在，在MyISAM存储引擎中，索引存储在一个.MYI文件中。在InnoDB存储引擎中
	索引存储在一个逻辑名称叫做tablespace的当中。在MEMORY存储引擎当中索引
	被存储在内存当中。不管索引存储在哪里，索引在mysql当中都是一个树的形式
	存在。（自平衡二叉树：B-Tree）

1.3、在mysql当中，主键上，以及unique字段上都会自动添加索引的！！！！
什么条件下，我们会考虑给字段添加索引呢？
	条件1：数据量庞大（到底有多么庞大算庞大，这个需要测试，因为每一个硬件环境不同）
	条件2：该字段经常出现在where的后面，以条件的形式存在，也就是说这个字段总是被扫描。
	条件3：该字段很少的DML(insert delete update)操作。（因为DML之后，索引需要重新排序。）

	建议不要随意添加索引，因为索引也是需要维护的，太多的话反而会降低系统的性能。
	建议通过主键查询，建议通过unique约束的字段进行查询，效率是比较高的。


1.4、索引怎么创建？怎么删除？语法是什么？

	创建索引：
		mysql> create index emp_ename_index on emp(ename);
		给emp表的ename字段添加索引，起名：emp_ename_index
	
	删除索引：
		mysql> drop index emp_ename_index on emp;
		将emp表上的emp_ename_index索引对象删除。


```SQL
mysql> create index emp_ename_index on EMP(ENAME);
Query OK, 0 rows affected (0.06 sec)
Records: 0  Duplicates: 0  Warnings: 0

mysql> DROP INDEX emp_ename_index ON EMP;
Query OK, 0 rows affected (0.04 sec)
Records: 0  Duplicates: 0  Warnings: 0
```


1.5、在mysql当中，怎么查看一个SQL语句是否使用了索引进行检索？


```sql

mysql> explain select * from EMP where ENAME='KING';
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
| id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra       |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
|  1 | SIMPLE      | EMP   | NULL       | ALL  | NULL          | NULL | NULL    | NULL |   14 |    10.00 | Using where |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
1 row in set, 1 warning (0.00 sec)
```
扫描14条记录：说明没有使用索引。type=ALL


```sql
mysql> create index emp_ename_index on EMP(ENAME);
Query OK, 0 rows affected (0.06 sec)
Records: 0  Duplicates: 0  Warnings: 0

mysql> explain select * from EMP where ENAME='KING';
+----+-------------+-------+------------+------+-----------------+-----------------+---------+-------+------+----------+-------+
| id | select_type | table | partitions | type | possible_keys   | key             | key_len | ref   | rows | filtered | Extra |
+----+-------------+-------+------------+------+-----------------+-----------------+---------+-------+------+----------+-------+
|  1 | SIMPLE      | EMP   | NULL       | ref  | emp_ename_index | emp_ename_index | 43      | const |    1 |   100.00 | NULL  |
+----+-------------+-------+------------+------+-----------------+-----------------+---------+-------+------+----------+-------+
1 row in set, 1 warning (0.00 sec)
```


1.6、索引有失效的时候，什么时候索引失效呢？
	
	失效的第1种情况：
		select * from emp where ename like '%T';

		ename上即使添加了索引，也不会走索引，为什么？
			原因是因为模糊匹配当中以“%”开头了！
			尽量避免模糊查询的时候以“%”开始。
			这是一种优化的手段/策略。
```sql
mysql> explain select * from EMP where ename like '%T';
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
| id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra       |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
|  1 | SIMPLE      | EMP   | NULL       | ALL  | NULL          | NULL | NULL    | NULL |   14 |    11.11 | Using where |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
1 row in set, 1 warning (0.00 sec)
```

失效的第2种情况：
		使用or的时候会失效，如果使用or那么要求or两边的条件字段都要有
		索引，才会走索引，如果其中一边有一个字段没有索引，那么另一个
		字段上的索引也会实现。所以这就是为什么不建议使用or的原因。
	

```sql
mysql> explain select * from EMP where ename = 'KING' or job = 'MANAGER';
+----+-------------+-------+------------+------+-----------------+------+---------+------+------+----------+-------------+
| id | select_type | table | partitions | type | possible_keys   | key  | key_len | ref  | rows | filtered | Extra       |
+----+-------------+-------+------------+------+-----------------+------+---------+------+------+----------+-------------+
|  1 | SIMPLE      | EMP   | NULL       | ALL  | emp_ename_index | NULL | NULL    | NULL |   14 |    16.43 | Using where |
+----+-------------+-------+------------+------+-----------------+------+---------+------+------+----------+-------------+
1 row in set, 1 warning (0.00 sec)
```


失效的第3种情况：
		使用复合索引的时候，没有使用左侧的列查找，索引失效
		什么是复合索引？
			两个字段，或者更多的字段联合起来添加一个索引，叫做复合索引。


```sql
mysql> create index emp_job_sal on EMP(job,sal);
Query OK, 0 rows affected (0.06 sec)
Records: 0  Duplicates: 0  Warnings: 0

mysql> EXPLAIN SELECT * FROM EMP WHERE JOB = 'MANAGER';
+----+-------------+-------+------------+------+---------------+-------------+---------+-------+------+----------+-------+
| id | select_type | table | partitions | type | possible_keys | key         | key_len | ref   | rows | filtered | Extra |
+----+-------------+-------+------------+------+---------------+-------------+---------+-------+------+----------+-------+
|  1 | SIMPLE      | EMP   | NULL       | ref  | emp_job_sal   | emp_job_sal | 39      | const |    3 |   100.00 | NULL  |
+----+-------------+-------+------------+------+---------------+-------------+---------+-------+------+----------+-------+
1 row in set, 1 warning (0.00 sec)

mysql> EXPLAIN SELECT * FROM EMP WHERE SAL > 2500;
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
| id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra       |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
|  1 | SIMPLE      | EMP   | NULL       | ALL  | NULL          | NULL | NULL    | NULL |   14 |    33.33 | Using where |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
1 row in set, 1 warning (0.00 sec)
```


失效的第4种情况：
		在where当中索引列参加了运算，索引失效。

```sql
mysql> create index emp_sal_index on EMP(SAL);
Query OK, 0 rows affected (0.07 sec)
Records: 0  Duplicates: 0  Warnings: 0

mysql> explain select * from EMP WHERE SAL='800';
+----+-------------+-------+------------+------+---------------+---------------+---------+-------+------+----------+-------+
| id | select_type | table | partitions | type | possible_keys | key           | key_len | ref   | rows | filtered | Extra |
+----+-------------+-------+------------+------+---------------+---------------+---------+-------+------+----------+-------+
|  1 | SIMPLE      | EMP   | NULL       | ref  | emp_sal_index | emp_sal_index | 9       | const |    1 |   100.00 | NULL  |
+----+-------------+-------+------------+------+---------------+---------------+---------+-------+------+----------+-------+
1 row in set, 1 warning (0.00 sec)

mysql> explain select * from EMP WHERE SAL+1='800';
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
| id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra       |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
|  1 | SIMPLE      | EMP   | NULL       | ALL  | NULL          | NULL | NULL    | NULL |   14 |   100.00 | Using where |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
1 row in set, 1 warning (0.00 sec)
```


失效的第5种情况：
		在where当中索引列使用了函数

```sql
mysql> EXPLAIN SELECT * FROM EMP WHERE LOWER(ENAME)='smith';
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
| id | select_type | table | partitions | type | possible_keys | key  | key_len | ref  | rows | filtered | Extra       |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
|  1 | SIMPLE      | EMP   | NULL       | ALL  | NULL          | NULL | NULL    | NULL |   14 |   100.00 | Using where |
+----+-------------+-------+------------+------+---------------+------+---------+------+------+----------+-------------+
1 row in set, 1 warning (0.00 sec)
```


1.7、索引是各种数据库进行优化的重要手段。优化的时候优先考虑的因素就是索引。
索引在数据库当中分了很多类？
	单一索引：一个字段上添加索引。
	复合索引：两个字段或者更多的字段上添加索引。

	主键索引：主键上添加索引。
	唯一性索引：具有unique约束的字段上添加索引。
	.....

	注意：唯一性比较弱的字段上添加索引用处不大。




**2、视图(view)**

2.1、什么是视图？
	view:站在不同的角度去看待同一份数据。

2.2、怎么创建视图对象？怎么删除视图对象？

表复制：
```SQL
mysql> CREATE TABLE DEPT2 AS SELECT * FROM DEPT;
Query OK, 4 rows affected (0.06 sec)
Records: 4  Duplicates: 0  Warnings: 0
```
dept2表中的数据：
```SQL
mysql> SELECT * FROM DEPT2;
+--------+------------+----------+
| DEPTNO | DNAME      | LOC      |
+--------+------------+----------+
|     10 | ACCOUNTING | NEW YORK |
|     20 | RESEARCH   | DALLAS   |
|     30 | SALES      | CHICAGO  |
|     40 | OPERATIONS | BOSTON   |
+--------+------------+----------+
4 rows in set (0.00 sec)
```

	创建视图对象：
		create view dept2_view as select * from dept2;
	
	删除视图对象：
		drop view dept2_view;

```SQL
mysql> CREATE VIEW DEPT2_VIEW AS SELECT * FROM DEPT2;
Query OK, 0 rows affected (0.02 sec)

mysql> DROP VIEW DEPT2_VIEW;
Query OK, 0 rows affected (0.02 sec)
```
	
	注意：只有DQL语句才能以view的形式创建。
		create view view_name as 这里的语句必须是DQL语句;

2.3、用视图做什么？

	我们可以面向视图对象进行增删改查，对视图对象的增删改查，会导致
	原表被操作！（视图的特点：通过对视图的操作，会影响到原表数据。）

	//面向视图查询
	select * from dept2_view; 

	// 面向视图插入
	insert into dept2_view(deptno,dname,loc) values(60,'SALES', 'BEIJING');

	// 查询原表数据
	mysql> select * from dept2;
	+--------+------------+----------+
	| DEPTNO | DNAME      | LOC      |
	+--------+------------+----------+
	|     10 | ACCOUNTING | NEW YORK |
	|     20 | RESEARCH   | DALLAS   |
	|     30 | SALES      | CHICAGO  |
	|     40 | OPERATIONS | BOSTON   |
	|     60 | SALES      | BEIJING  |
	+--------+------------+----------+

	// 面向视图删除
	mysql> delete from dept2_view;

	// 查询原表数据
	mysql> select * from dept2;
	Empty set (0.00 sec)
	

	// 创建视图对象
	create view 
		emp_dept_view
	as
		select 
			e.ename,e.sal,d.dname
		from
			emp e
		join
			dept d
		on
			e.deptno = d.deptno;

	// 查询视图对象
	mysql> select * from emp_dept_view;
	+--------+---------+------------+
	| ename  | sal     | dname      |
	+--------+---------+------------+
	| CLARK  | 2450.00 | ACCOUNTING |
	| KING   | 5000.00 | ACCOUNTING |
	| MILLER | 1300.00 | ACCOUNTING |
	| SMITH  |  800.00 | RESEARCH   |
	| JONES  | 2975.00 | RESEARCH   |
	| SCOTT  | 3000.00 | RESEARCH   |
	| ADAMS  | 1100.00 | RESEARCH   |
	| FORD   | 3000.00 | RESEARCH   |
	| ALLEN  | 1600.00 | SALES      |
	| WARD   | 1250.00 | SALES      |
	| MARTIN | 1250.00 | SALES      |
	| BLAKE  | 2850.00 | SALES      |
	| TURNER | 1500.00 | SALES      |
	| JAMES  |  950.00 | SALES      |
	+--------+---------+------------+

	// 面向视图更新
	update emp_dept_view set sal = 1000 where dname = 'ACCOUNTING';

	// 原表数据被更新
	mysql> select * from emp;
	+-------+--------+-----------+------+------------+---------+---------+--------+
	| EMPNO | ENAME  | JOB       | MGR  | HIREDATE   | SAL     | COMM    | DEPTNO |
	+-------+--------+-----------+------+------------+---------+---------+--------+
	|  7369 | SMITH  | CLERK     | 7902 | 1980-12-17 |  800.00 |    NULL |     20 |
	|  7499 | ALLEN  | SALESMAN  | 7698 | 1981-02-20 | 1600.00 |  300.00 |     30 |
	|  7521 | WARD   | SALESMAN  | 7698 | 1981-02-22 | 1250.00 |  500.00 |     30 |
	|  7566 | JONES  | MANAGER   | 7839 | 1981-04-02 | 2975.00 |    NULL |     20 |
	|  7654 | MARTIN | SALESMAN  | 7698 | 1981-09-28 | 1250.00 | 1400.00 |     30 |
	|  7698 | BLAKE  | MANAGER   | 7839 | 1981-05-01 | 2850.00 |    NULL |     30 |
	|  7782 | CLARK  | MANAGER   | 7839 | 1981-06-09 | 1000.00 |    NULL |     10 |
	|  7788 | SCOTT  | ANALYST   | 7566 | 1987-04-19 | 3000.00 |    NULL |     20 |
	|  7839 | KING   | PRESIDENT | NULL | 1981-11-17 | 1000.00 |    NULL |     10 |
	|  7844 | TURNER | SALESMAN  | 7698 | 1981-09-08 | 1500.00 |    0.00 |     30 |
	|  7876 | ADAMS  | CLERK     | 7788 | 1987-05-23 | 1100.00 |    NULL |     20 |
	|  7900 | JAMES  | CLERK     | 7698 | 1981-12-03 |  950.00 |    NULL |     30 |
	|  7902 | FORD   | ANALYST   | 7566 | 1981-12-03 | 3000.00 |    NULL |     20 |
	|  7934 | MILLER | CLERK     | 7782 | 1982-01-23 | 1000.00 |    NULL |     10 |
	+-------+--------+-----------+------+------------+---------+---------+--------+

2.4、视图对象在实际开发中到底有什么用？《方便，简化开发，利于维护》

		create view 
			emp_dept_view
		as
			select 
				e.ename,e.sal,d.dname
			from
				emp e
			join
				dept d
			on
				e.deptno = d.deptno;
		
		
		假设有一条非常复杂的SQL语句，而这条SQL语句需要在不同的位置上反复使用。
		每一次使用这个sql语句的时候都需要重新编写，很长，很麻烦，怎么办？
			可以把这条复杂的SQL语句以视图对象的形式新建。
			在需要编写这条SQL语句的位置直接使用视图对象，可以大大简化开发。
			并且利于后期的维护，因为修改的时候也只需要修改一个位置就行，只需要
			修改视图对象所映射的SQL语句。
		
		我们以后面向视图开发的时候，使用视图的时候可以像使用table一样。
		可以对视图进行增删改查等操作。视图不是在内存当中，视图对象也是
		存储在硬盘上的，不会消失。

		再提醒一下：
			视图对应的语句只能是DQL语句。
			但是视图对象创建完成之后，可以对视图进行增删改查等操作。

		小插曲：
			增删改查，又叫做：CRUD。
			CRUD是在公司中程序员之间沟通的术语。一般我们很少说增删改查。
			一般都说CRUD。

			C:Create（增）
			R:Retrive（查：检索）
			U:Update（改）
			D:Delete（删）

3、DBA常用命令？

	重点掌握：
		数据的导入和导出（数据的备份）
		其它命令了解一下即可。（这个培训日志文档留着，以后忘了，可以打开文档复制粘贴。）
	
	数据导出？


注意：在MacOs的iTerm命令窗口中：
```sql
mysqldump X>/Users/x/Desktop/X.sql -uroot -p123456
```
		
可以导出指定的表吗？
```SQL
mysqldump X EMP>/Users/x/Desktop/EMP.sql -uroot -p123456
```

	数据导入？
		注意：需要先登录到mysql数据库服务器上。
		然后创建数据库：create database X;
		使用数据库：use X;
		然后初始化数据库：source '/Users/x/Desktop/X.sql'

4、数据库设计三范式

4.1、什么是数据库设计范式？
	数据库表的设计依据。教你怎么进行数据库表的设计。

## 4.2、数据库设计范式共有？ 3个。

	**第一范式：要求任何一张表必须有主键，每一个字段原子性不可再分。**

	**第二范式：建立在第一范式的基础之上，要求所有非主键字段完全依赖主键，**
	**不要产生部分依赖。**

	**第三范式：建立在第二范式的基础之上，要求所有非主键字段直接依赖主键，**
	**不要产生传递依赖。**

	声明：三范式是面试官经常问的，所以一定要熟记在心！

	设计数据库表的时候，按照以上的范式进行，可以避免表中数据的冗余，空间的浪费。

4.3、第一范式
	最核心，最重要的范式，所有表的设计都需要满足。
	必须有主键，并且每一个字段都是原子性不可再分。

	学生编号 学生姓名 联系方式
	------------------------------------------
	1001		张三		zs@gmail.com,1359999999
	1002		李四		ls@gmail.com,13699999999
	1001		王五		ww@163.net,13488888888

	以上是学生表，满足第一范式吗？
		不满足，第一：没有主键。第二：联系方式可以分为邮箱地址和电话
	
	学生编号(pk) 学生姓名	邮箱地址			联系电话
	----------------------------------------------------
	1001				张三		zs@gmail.com	1359999999
	1002				李四		ls@gmail.com	13699999999
	1003				王五		ww@163.net		13488888888

4.4、第二范式：
	建立在第一范式的基础之上，
	要求所有非主键字段必须完全依赖主键，不要产生部分依赖。

	学生编号 学生姓名 教师编号 教师姓名
	----------------------------------------------------
	1001			张三		001		王老师
	1002			李四		002		赵老师
	1003			王五		001		王老师
	1001			张三		002		赵老师

	这张表描述了学生和老师的关系：（1个学生可能有多个老师，1个老师有多个学生）
	这是非常典型的：多对多关系！

	分析以上的表是否满足第一范式？
		不满足第一范式。
	
	怎么满足第一范式呢？修改

	学生编号+教师编号(pk)		学生姓名  教师姓名
	----------------------------------------------------
	1001			001				张三			王老师
	1002			002				李四			赵老师
	1003			001				王五			王老师
	1001			002				张三			赵老师

	学生编号 教师编号，两个字段联合做主键，复合主键（PK: 学生编号+教师编号）
	经过修改之后，以上的表满足了第一范式。但是满足第二范式吗？
		不满足，“张三”依赖1001，“王老师”依赖001，显然产生了部分依赖。
		产生部分依赖有什么缺点？
			数据冗余了。空间浪费了。“张三”重复了，“王老师”重复了。
	
	为了让以上的表满足第二范式，你需要这样设计：
		使用三张表来表示多对多的关系！！！！
		学生表
		学生编号(pk)		学生名字
		------------------------------------
		1001					张三
		1002					李四
		1003					王五
		
		教师表
		教师编号(pk)		教师姓名
		--------------------------------------
		001					王老师
		002					赵老师

		学生教师关系表
		id(pk)			学生编号(fk)			教师编号(fk)
		------------------------------------------------------
		1						1001						001
		2						1002						002
		3						1003						001
		4						1001						002
	

	背口诀：
		多对多怎么设计？
			**多对多，三张表，关系表两个外键**


4.5、第三范式
	第三范式建立在第二范式的基础之上
	要求所有非主键字典必须直接依赖主键，不要产生传递依赖。

	学生编号（PK） 学生姓名 班级编号  班级名称
	---------------------------------------------------------
		1001				张三		01			一年一班
		1002				李四		02			一年二班
		1003				王五		03			一年三班
		1004				赵六		03			一年三班
	
	以上表的设计是描述：班级和学生的关系。很显然是1对多关系！
	一个教室中有多个学生。

	分析以上表是否满足第一范式？
		满足第一范式，有主键。
	
	分析以上表是否满足第二范式？
		满足第二范式，因为主键不是复合主键，没有产生部分依赖。主键是单一主键。
	
	分析以上表是否满足第三范式？
		第三范式要求：不要产生传递依赖！
		一年一班依赖01，01依赖1001，产生了传递依赖。
		不符合第三范式的要求。产生了数据的冗余。
	
	那么应该怎么设计一对多呢？

		班级表：一
		班级编号(pk)				班级名称
		----------------------------------------
		01								一年一班
		02								一年二班
		03								一年三班

		学生表：多

		学生编号（PK） 学生姓名 班级编号(fk)
		-------------------------------------------
		1001				张三			01			
		1002				李四			02			
		1003				王五			03			
		1004				赵六			03		
		
		背口诀：
			**一对多，两张表，多的表加外键**

4.6、总结表的设计？

	一对多：
		**一对多，两张表，多的表加外键**

	多对多：
		**多对多，三张表，关系表两个外键**

	一对一：
		一对一放到一张表中不就行了吗？为啥还要拆分表？
		
		在实际的开发中，可能存在一张表字段太多，太庞大。这个时候要拆分表。
		一对一怎么设计？
			没有拆分表之前：一张表
				t_user
				id		login_name		login_pwd		real_name		email				address........
				---------------------------------------------------------------------------
				1			zhangsan		123				张三				zhangsan@xxx
				2			lisi			123				李四				lisi@xxx
				...
			
			这种庞大的表建议拆分为两张：
				t_login 登录信息表
				id(pk)		login_name		login_pwd	
				---------------------------------
				1				zhangsan		123			
				2				lisi			123			

				t_user 用户详细信息表
				id(pk)		real_name		email				address........	login_id(fk+unique)
				-----------------------------------------------------------------------------------------
				100			张三				zhangsan@xxx								1
				200			李四				lisi@xxx										2


				口诀：一对一，外键唯一！！！！！！！！！！

4.7、嘱咐一句话：

	数据库设计三范式是理论上的。

	实践和理论有的时候有偏差。

	最终的目的都是为了满足客户的需求，有的时候会拿冗余换执行速度。

	因为在sql当中，表和表之间连接次数越多，效率越低。（笛卡尔积）

	有的时候可能会存在冗余，但是为了减少表的连接次数，这样做也是合理的，
	并且对于开发人员来说，sql语句的编写难度也会降低。

	面试的时候把这句话说上：他就不会认为你是初级程序员了！