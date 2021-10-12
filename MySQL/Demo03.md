
1. [insert语句可以一次插入多条记录吗？【掌握】](#insert语句可以一次插入多条记录吗掌握)
1. [快速创建表？【了解内容】](#快速创建表了解内容)
1. [将查询结果插入到一张表当中？insert相关的！！！【了解内容】](#将查询结果插入到一张表当中insert相关的了解内容)
1. [快速删除表中的数据？【truncate比较重要，必须掌握】](#快速删除表中的数据truncate比较重要必须掌握)
1. [约束（非常重要，五颗星*****）](#约束非常重要五颗星)
1. [主键约束（primary key，简称PK）非常重要五颗星](#主键约束primary-key简称pk非常重要五颗星)
1. [外键约束（foreign key，简称FK）非常重要五颗星](#外键约束foreign-key简称fk非常重要五颗星)

## insert语句可以一次插入多条记录吗？【掌握】
	可以的！

```sql
mysql> desc t_user;

+-------------+-------------+------+-----+---------+-------+
| Field       | Type        | Null | Key | Default | Extra |
+-------------+-------------+------+-----+---------+-------+
| id          | int         | YES  |     | NULL    |       |
| name        | varchar(32) | YES  |     | NULL    |       |
| birth       | date        | YES  |     | NULL    |       |
| create_time | datetime    | YES  |     | NULL    |       |
+-------------+-------------+------+-----+---------+-------+
4 rows in set (0.00 sec)

mysql> select * from t_user;
Empty set (0.00 sec)

mysql> insert into t_user values (23,'mojo','1995-09-01',now()),(2,'mojo','1995-09-01',now());
Query OK, 2 rows affected (0.01 sec)
Records: 2  Duplicates: 0  Warnings: 0

mysql> select * from t_user;

+------+------+------------+---------------------+
| id   | name | birth      | create_time         |
+------+------+------------+---------------------+
|   23 | mojo | 1995-09-01 | 2021-10-10 00:45:27 |
|    2 | mojo | 1995-09-01 | 2021-10-10 00:45:27 |
+------+------+------------+---------------------+
2 rows in set (0.00 sec)
```

## 快速创建表？【了解内容】

```SQL
mysql> CREATE TABLE EMP2 AS SELECT * FROM EMP;
Query OK, 14 rows affected, 2 warnings (0.04 sec)
Records: 14  Duplicates: 0  Warnings: 2

mysql> SELECT * FROM EMP;
+-------+--------+-----------+------+------------+---------+---------+--------+
| EMPNO | ENAME  | JOB       | MGR  | HIREDATE   | SAL     | COMM    | DEPTNO |
+-------+--------+-----------+------+------------+---------+---------+--------+
|  7369 | SMITH  | CLERK     | 7902 | 1980-12-17 |  800.00 |    NULL |     20 |
|  7499 | ALLEN  | SALESMAN  | 7698 | 1981-02-20 | 1600.00 |  300.00 |     30 |
|  7521 | WARD   | SALESMAN  | 7698 | 1981-02-22 | 1250.00 |  500.00 |     30 |
|  7566 | JONES  | MANAGER   | 7839 | 1981-04-02 | 2975.00 |    NULL |     20 |
|  7654 | MARTIN | SALESMAN  | 7698 | 1981-09-28 | 1250.00 | 1400.00 |     30 |
|  7698 | BLAKE  | MANAGER   | 7839 | 1981-05-01 | 2850.00 |    NULL |     30 |
|  7782 | CLARK  | MANAGER   | 7839 | 1981-06-09 | 2450.00 |    NULL |     10 |
|  7788 | SCOTT  | ANALYST   | 7566 | 1987-04-19 | 3000.00 |    NULL |     20 |
|  7839 | KING   | PRESIDENT | NULL | 1981-11-17 | 5000.00 |    NULL |     10 |
|  7844 | TURNER | SALESMAN  | 7698 | 1981-09-08 | 1500.00 |    0.00 |     30 |
|  7876 | ADAMS  | CLERK     | 7788 | 1987-05-23 | 1100.00 |    NULL |     20 |
|  7900 | JAMES  | CLERK     | 7698 | 1981-12-03 |  950.00 |    NULL |     30 |
|  7902 | FORD   | ANALYST   | 7566 | 1981-12-03 | 3000.00 |    NULL |     20 |
|  7934 | MILLER | CLERK     | 7782 | 1982-01-23 | 1300.00 |    NULL |     10 |
+-------+--------+-----------+------+------------+---------+---------+--------+
14 rows in set (0.00 sec)
```


原理：
		将一个查询结果当做一张表新建！！！！！
		这个可以完成表的快速复制！！！！
		表创建出来，同时表中的数据也存在了！！！

```SQL
mysql> CREATE TABLE MYTABLE AS SELECT EMPNO,ENAME FROM EMP WHERE JOB = 'MANAGER';
Query OK, 3 rows affected (0.03 sec)
Records: 3  Duplicates: 0  Warnings: 0

mysql> SELECT * FROM MYTABLE;
+-------+-------+
| EMPNO | ENAME |
+-------+-------+
|  7566 | JONES |
|  7698 | BLAKE |
|  7782 | CLARK |
+-------+-------+
3 rows in set (0.00 sec)
```




 
## 将查询结果插入到一张表当中？insert相关的！！！【了解内容】

很少用

```SQL
mysql> CREATE TABLE DEPT_BAK AS SELECT * FROM DEPT;
Query OK, 4 rows affected (0.04 sec)
Records: 4  Duplicates: 0  Warnings: 0

mysql> SELECT * FROM DEPT_BAK;
+--------+------------+----------+
| DEPTNO | DNAME      | LOC      |
+--------+------------+----------+
|     10 | ACCOUNTING | NEW YORK |
|     20 | RESEARCH   | DALLAS   |
|     30 | SALES      | CHICAGO  |
|     40 | OPERATIONS | BOSTON   |
+--------+------------+----------+
4 rows in set (0.00 sec)

mysql> INSERT INTO DEPT_BAK SELECT * FROM DEPT;
Query OK, 4 rows affected (0.01 sec)
Records: 4  Duplicates: 0  Warnings: 0

mysql> SELECT * FROM DEPT_BAK;
+--------+------------+----------+
| DEPTNO | DNAME      | LOC      |
+--------+------------+----------+
|     10 | ACCOUNTING | NEW YORK |
|     20 | RESEARCH   | DALLAS   |
|     30 | SALES      | CHICAGO  |
|     40 | OPERATIONS | BOSTON   |
|     10 | ACCOUNTING | NEW YORK |
|     20 | RESEARCH   | DALLAS   |
|     30 | SALES      | CHICAGO  |
|     40 | OPERATIONS | BOSTON   |
+--------+------------+----------+
8 rows in set (0.00 sec)
```


## 快速删除表中的数据？【truncate比较重要，必须掌握】

```SQL
mysql> DELETE FROM DEPT_BAK;
Query OK, 8 rows affected (0.03 sec)

mysql> SELECT * FROM DEPT_BAK;
Empty set (0.01 sec)
```

```sql
	//删除dept_bak表中的数据
	//这种删除数据的方式比较慢。

delete语句删除数据的原理？（delete属于DML语句！！！）
		表中的数据被删除了，但是这个数据在硬盘上的真实存储空间不会被释放！！！
		这种删除缺点是：删除效率比较低。
		这种删除优点是：支持回滚，后悔了可以再恢复数据！！！
	
	truncate语句删除数据的原理？
		这种删除效率比较高，表被一次截断，物理删除。
		这种删除缺点：不支持回滚。
		这种删除优点：快速。

	用法：truncate table dept_bak; （这种操作属于DDL操作。）

	大表非常大，上亿条记录？？？？
		删除的时候，使用delete，也许需要执行1个小时才能删除完！效率较低。
		可以选择使用truncate删除表中的数据。只需要不到1秒钟的时间就删除结束。效率较高。
		但是使用truncate之前，必须仔细询问客户是否真的要删除，并警告删除之后不可恢复！

		truncate是删除表中的数据，表还在！
	
	删除表操作？
		drop table 表名; // 这不是删除表中的数据，这是把表删除。

```


对表结构的增删改？
	
	什么是对表结构的修改？
		添加一个字段，删除一个字段，修改一个字段！！！
	
	对表结构的修改需要使用：alter
	属于DDL语句

	DDL包括：create drop alter
	
	第一：在实际的开发中，需求一旦确定之后，表一旦设计好之后，很少的
	进行表结构的修改。因为开发进行中的时候，修改表结构，成本比较高。
	修改表的结构，对应的java代码就需要进行大量的修改。成本是比较高的。
	这个责任应该由设计人员来承担！

	第二：由于修改表结构的操作很少，所以我们不需要掌握，如果有一天
	真的要修改表结构，你可以使用工具！！！！

	修改表结构的操作是不需要写到java程序中的。实际上也不是java程序员的范畴。


## 约束（非常重要，五颗星*****）

**什么是约束？**
	约束对应的英语单词：constraint
	在创建表的时候，我们可以给表中的字段加上一些约束，来保证这个表中数据的
	完整性、有效性！！！

	约束的作用就是为了保证：表中的数据有效！！

**约束包括哪些？**
```sql
	非空约束：not null
	唯一性约束: unique
	主键约束: primary key （简称PK）
	外键约束：foreign key（简称FK）
	检查约束：check（mysql不支持，oracle支持）

	我们这里重点学习四个约束：
		not null
		unique
		primary key
		foreign key
```


**非空约束：not null**

	非空约束not null约束的字段不能为NULL。
```sql
	drop table if exists t_vip;
	create table t_vip(
		id int,
		name varchar(255) not null  // not null只有列级约束，没有表级约束！
	);
	insert into t_vip(id,name) values(1,'zhangsan');
	insert into t_vip(id,name) values(2,'lisi');
```
	insert into t_vip(id) values(3);
	ERROR 1364 (HY000): Field 'name' doesn't have a default value

	小插曲：
		xxxx.sql这种文件被称为sql脚本文件。
		sql脚本文件中编写了大量的sql语句。
		我们执行sql脚本文件的时候，该文件中所有的sql语句会全部执行！
		批量的执行SQL语句，可以使用sql脚本文件。
		在mysql当中怎么执行sql脚本呢？
			mysql> source D:\course\03-MySQL\document\vip.sql
		
		你在实际的工作中，第一天到了公司，项目经理会给你一个xxx.sql文件，
		你执行这个脚本文件，你电脑上的数据库数据就有了！

**唯一性约束: unique**

	唯一性约束unique约束的字段不能重复，但是可以为NULL。

**列级约束**
```sql
DROP TABLE IF EXISTS t_vip; 
CREATE TABLE t_vip (id INT,NAME VARCHAR (255) UNIQUE,email CHAR (255));
INSERT INTO t_vip(id,name,email) VALUES (1,'mojo1','x@gmail.com');
INSERT INTO t_vip(id,name,email) VALUES (2,'mojo2','x@gmail.com');
```

```sql
mysql> INSERT INTO t_vip(id,name,email) VALUES (2,'mojo2','x@gmail.com');
ERROR 1062 (23000): Duplicate entry 'mojo2' for key 't_vip.NAME'
```

**新需求：name和email两个字段联合起来具有唯一性！！！！**

**表级约束**
```sql
DROP TABLE IF EXISTS t_vip; 
CREATE TABLE t_vip (id INT,NAME VARCHAR (255),email CHAR (255),UNIQUE (NAME,email)); 
INSERT INTO t_vip (id,NAME,email) VALUES (1,'mojo','x@gmail.com'); 
INSERT INTO t_vip (id,NAME,email) VALUES (2,'mojo','x2@gmail.com');
```


什么时候使用表级约束呢？
			需要给多个字段联合起来添加某一个约束的时候，需要使用表级约束。

unique 和not null可以联合吗？

```sql
DROP TABLE IF EXISTS t_vip; 
CREATE TABLE t_vip (id INT,NAME VARCHAR (255) NOT NULL UNIQUE,email CHAR (255));

mysql> DESC t_vip;;
+-------+--------------+------+-----+---------+-------+
| Field | Type         | Null | Key | Default | Extra |
+-------+--------------+------+-----+---------+-------+
| id    | int          | YES  |     | NULL    |       |
| NAME  | varchar(255) | NO   | PRI | NULL    |       |
| email | char(255)    | YES  |     | NULL    |       |
+-------+--------------+------+-----+---------+-------+
3 rows in set (0.00 sec)

ERROR: 
No query specified
```

在mysql当中，如果一个字段同时被not null和unique约束的话，
		该字段自动变成主键字段。（注意：oracle中不一样！）

## 主键约束（primary key，简称PK）非常重要五颗星 

	主键约束的相关术语？
		主键约束：就是一种约束。
		主键字段：该字段上添加了主键约束，这样的字段叫做：主键字段
		主键值：主键字段中的每一个值都叫做：主键值。
	
	什么是主键？有啥用？
		主键值是每一行记录的唯一标识。
		主键值是每一行记录的身份证号！！！
	
	记住：任何一张表都应该有主键，没有主键，表无效！！

	主键的特征：not null + unique（主键值不能是NULL，同时也不能重复！）

	怎么给一张表添加主键约束呢？

```sql
DROP TABLE IF EXISTS t_vip; 
CREATE TABLE t_vip (id INT PRIMARY KEY,NAME VARCHAR (255)); 
INSERT INTO t_vip (id,NAME) VALUES (1,'mojo'); 
```


```sql
mysql> INSERT INTO t_vip (id,NAME) VALUES (1,'mojo');
ERROR 1062 (23000): Duplicate entry '1' for key 't_vip.PRIMARY'
```

表级约束
```sql
DROP TABLE IF EXISTS t_vip; 
CREATE TABLE t_vip (id INT,NAME VARCHAR (255),PRIMARY KEY (id)); 
INSERT INTO t_vip (id,NAME) VALUES (1,'mojo'); 
```


表级约束主要是给多个字段联合起来添加约束？

// id和name联合起来做主键：复合主键！！！！

```sql
DROP TABLE IF EXISTS t_vip; 
CREATE TABLE t_vip (id INT,NAME VARCHAR (255),email VARCHAR (255),PRIMARY KEY (id,NAME)); 
INSERT INTO t_vip (id,NAME,email) VALUES (1,'zhangsan','zhangsan@123.com'); 
INSERT INTO t_vip (id,NAME,email) VALUES (1,'lisi','lisi@123.com');

mysql> select * from t_vip;
+----+----------+------------------+
| id | NAME     | email            |
+----+----------+------------------+
|  1 | lisi     | lisi@123.com     |
|  1 | zhangsan | zhangsan@123.com |
+----+----------+------------------+
2 rows in set (0.00 sec)
```


主键值建议使用：
		int
		bigint
		char
		等类型。

		不建议使用：varchar来做主键。主键值一般都是数字，一般都是定长的！

	主键除了：单一主键和复合主键之外，还可以这样进行分类？
		自然主键：主键值是一个自然数，和业务没关系。
		业务主键：主键值和业务紧密关联，例如拿银行卡账号做主键值。这就是业务主键！

		在实际开发中使用业务主键多，还是使用自然主键多一些？
			自然主键使用比较多，因为主键只要做到不重复就行，不需要有意义。
			业务主键不好，因为主键一旦和业务挂钩，那么当业务发生变动的时候，
			可能会影响到主键值，所以业务主键不建议使用。尽量使用自然主键。


在mysql当中，有一种机制，可以帮助我们自动维护一个主键值？


```sql
drop table if exists t_vip;
create table t_vip(
			id int primary key auto_increment, 
			name varchar(255)
		);
		insert into t_vip(name) values('zhangsan');
		insert into t_vip(name) values('zhangsan');
		insert into t_vip(name) values('zhangsan');
		insert into t_vip(name) values('zhangsan');
		insert into t_vip(name) values('zhangsan');
		insert into t_vip(name) values('zhangsan');
		insert into t_vip(name) values('zhangsan');
		insert into t_vip(name) values('zhangsan');


mysql> select * from t_vip;
+----+----------+
| id | name     |
+----+----------+
|  1 | zhangsan |
|  2 | zhangsan |
|  3 | zhangsan |
|  4 | zhangsan |
|  5 | zhangsan |
|  6 | zhangsan |
|  7 | zhangsan |
|  8 | zhangsan |
+----+----------+
8 rows in set (0.00 sec)
```


## 外键约束（foreign key，简称FK）非常重要五颗星

外键约束涉及到的相关术语：
		外键约束：一种约束（foreign key）
		外键字段：该字段上添加了外键约束
		外键值：外键字段当中的每一个值。


	t_class是父表
			t_student是子表

			删除表的顺序？
				先删子，再删父。

			创建表的顺序？
				先创建父，再创建子。

			删除数据的顺序？
				先删子，再删父。

			插入数据的顺序？
				先插入父，再插入子。

		思考：子表中的外键引用的父表中的某个字段，被引用的这个字段必须是主键吗？
			不一定是主键，但至少具有unique约束。

		测试：外键可以为NULL吗？
			外键值可以为NULL。

```sql
drop table if exists t_student;
drop table if exists t_class;
create table t_class(classno int primary key, classname varchar(255));
create table t_student(no int primary key auto_increment,name varchar(255),cno int, foreign key(cno) references t_class(classno));
```

