
1. [简单查询](#简单查询)
1. [条件查询](#条件查询)
1. [排序](#排序)
1. [数据处理函数](#数据处理函数)
1. [分组函数（多行处理函数）](#分组函数多行处理函数)
1. [✅ 分组查询 （非常重要）](#-分组查询-非常重要)

什么是数据库？什么是数据库管理系统？什么是SQL？他们之间的关系是什么？

	数据库：
		英文单词DataBase，简称DB。按照一定格式存储数据的一些文件的组合。
		顾名思义：存储数据的仓库，实际上就是一堆文件。这些文件中存储了
		具有特定格式的数据。

	数据库管理系统：
		DataBaseManagement，简称DBMS。
		数据库管理系统是专门用来管理数据库中数据的，数据库管理系统可以
		对数据库当中的数据进行增删改查。

		常见的数据库管理系统：
			MySQL、Oracle、MS SqlServer、DB2、sybase等....
	
	SQL：结构化查询语言
		程序员需要学习SQL语句，程序员通过编写SQL语句，然后DBMS负责执行SQL
		语句，最终来完成数据库中数据的增删改查操作。

		SQL是一套标准，程序员主要学习的就是SQL语句，这个SQL在mysql中可以使用，
		同时在Oracle中也可以使用，在DB2中也可以使用。
	
	三者之间的关系？
		DBMS--执行--> SQL --操作--> DB
	
	先安装数据库管理系统MySQL，然后学习SQL语句怎么写，编写SQL语句之后，DBMS
	对SQL语句进行执行，最终来完成数据库的数据管理。



**关于SQL语句的分类？**

	SQL语句有很多，最好进行分门别类，这样更容易记忆。
		分为：
			**DQL：**
				数据查询语言（凡是带有select关键字的都是查询语句）
				select...

			**DML：**
				数据操作语言（凡是对表当中的数据进行增删改的都是DML）
				insert delete update
				insert 增
				delete 删
				update 改

				这个主要是操作表中的数据data。

			**DDL：**
				数据定义语言
				凡是带有create、drop、alter的都是DDL。
				DDL主要操作的是表的结构。不是表中的数据。
				create：新建，等同于增
				drop：删除
				alter：修改
				这个增删改和DML不同，这个主要是对表结构进行操作。

			**TCL：**
				不是王牌电视。
				是事务控制语言
				包括：
					事务提交：commit;
					事务回滚：rollback;

			**DCL：**
				是数据控制语言。
				例如：授权grant、撤销权限revoke....


``` sql
docker exec -it mysql bash

mysql -uroot -p123456
```

```SQL
x@n  docker exec -it mysql bash
root@f68a3d361cc9:/# mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 15
Server version: 8.0.26 MySQL Community Server - GPL

Copyright (c) 2000, 2021, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
```


```sql
mysql> SELECT VERSION();
+-----------+
| VERSION() |
+-----------+
| 8.0.26    |
+-----------+
1 row in set (0.00 sec)

mysql> CREATE DATABASE X;
Query OK, 1 row affected (0.01 sec)

mysql> SHOW DATABASES;
+--------------------+
| Database           |
+--------------------+
| X                  |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)

mysql> USE X;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed

mysql> SHOW TABLES;
+-------------+
| Tables_in_X |
+-------------+
| DEPT        |
| EMP         |
| SALGRADE    |
+-------------+
3 rows in set (0.00 sec)

mysql> SELECT DATABASE();
+------------+
| DATABASE() |
+------------+
| X          |
+------------+
1 row in set (0.00 sec)
```


```sql
mysql> SELECT * FROM DEPT;
+--------+------------+----------+
| DEPTNO | DNAME      | LOC      |
+--------+------------+----------+
|     10 | ACCOUNTING | NEW YORK |
|     20 | RESEARCH   | DALLAS   |
|     30 | SALES      | CHICAGO  |
|     40 | OPERATIONS | BOSTON   |
+--------+------------+----------+
4 rows in set (0.01 sec)

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

mysql> SELECT * FROM SALGRADE;
+-------+-------+-------+
| GRADE | LOSAL | HISAL |
+-------+-------+-------+
|     1 |   700 |  1200 |
|     2 |  1201 |  1400 |
|     3 |  1401 |  2000 |
|     4 |  2001 |  3000 |
|     5 |  3001 |  9999 |
+-------+-------+-------+
5 rows in set (0.00 sec)
```


```SQL
mysql> DESC DEPT;
+--------+-------------+------+-----+---------+-------+
| Field  | Type        | Null | Key | Default | Extra |
+--------+-------------+------+-----+---------+-------+
| DEPTNO | int         | NO   | PRI | NULL    |       |部门编号
| DNAME  | varchar(14) | YES  |     | NULL    |       |部门名称
| LOC    | varchar(13) | YES  |     | NULL    |       |地理位置
+--------+-------------+------+-----+---------+-------+
3 rows in set (0.00 sec)

mysql> DESC EMP;
+----------+-------------+------+-----+---------+-------+
| Field    | Type        | Null | Key | Default | Extra |
+----------+-------------+------+-----+---------+-------+
| EMPNO    | int         | NO   | PRI | NULL    |       |员工编号
| ENAME    | varchar(10) | YES  |     | NULL    |       |员工姓名
| JOB      | varchar(9)  | YES  |     | NULL    |       |工作岗位
| MGR      | int         | YES  |     | NULL    |       |上级编号
| HIREDATE | date        | YES  |     | NULL    |       |入职日期
| SAL      | double(7,2) | YES  |     | NULL    |       |工资
| COMM     | double(7,2) | YES  |     | NULL    |       |补助
| DEPTNO   | int         | YES  |     | NULL    |       |部门编号
+----------+-------------+------+-----+---------+-------+
8 rows in set (0.00 sec)

mysql> DESC SALGRADE;
+-------+------+------+-----+---------+-------+
| Field | Type | Null | Key | Default | Extra |
+-------+------+------+-----+---------+-------+
| GRADE | int  | YES  |     | NULL    |       |工资等级
| LOSAL | int  | YES  |     | NULL    |       |最低工资
| HISAL | int  | YES  |     | NULL    |       |最高工资
+-------+------+------+-----+---------+-------+
3 rows in set (0.00 sec)
```


终止一条命令的输入
```SQL
mysql> \c  
```

退出 MySQL
```SQL
mysql> exit;
Bye
root@f68a3d361cc9:/
```

## 简单查询

1. 查询一个字段：
   
查询部门名

```SQL
mysql> SELECT DNAME FROM DEPT;
+------------+
| DNAME      |
+------------+
| ACCOUNTING |
| RESEARCH   |
| SALES      |
| OPERATIONS |
+------------+
4 rows in set (0.00 sec)
```


2. 查询2个字段或多个字段? 使用逗号隔开“，”

查询部门编号和部门名

```sql
mysql> SELECT DEPTNO,DNAME FROM DEPT;
+--------+------------+
| DEPTNO | DNAME      |
+--------+------------+
|     10 | ACCOUNTING |
|     20 | RESEARCH   |
|     30 | SALES      |
|     40 | OPERATIONS |
+--------+------------+
4 rows in set (0.00 sec)
```

3. 查询所有
   

* 效率低、可读性差
```SQL
mysql> SELECT DEPTNO,DNAME,LOC FROM DEPT;
+--------+------------+----------+
| DEPTNO | DNAME      | LOC      |
+--------+------------+----------+
|     10 | ACCOUNTING | NEW YORK |
|     20 | RESEARCH   | DALLAS   |
|     30 | SALES      | CHICAGO  |
|     40 | OPERATIONS | BOSTON   |
+--------+------------+----------+
4 rows in set (0.00 sec)

mysql> SELECT * FROM DEPT;
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


4. 给查询的列起别名
使用 AS 关键字起别名，只是显示的列名为AS后的名字，SELECT 不修改数据
AS 可以省略，不能加逗号
```SQL
mysql> SELECT DNAME AS DEPTNAME FROM DEPT;
+------------+
| DEPTNAME   |
+------------+
| ACCOUNTING |
| RESEARCH   |
| SALES      |
| OPERATIONS |
+------------+
4 rows in set (0.00 sec)

mysql> SELECT DNAME DEPTNAME FROM DEPT;
+------------+
| DEPTNAME   |
+------------+
| ACCOUNTING |
| RESEARCH   |
| SALES      |
| OPERATIONS |
+------------+
4 rows in set (0.00 sec)

mysql> SELECT DNAME `DEPT NAME` FROM DEPT;
+------------+
| DEPT NAME  |
+------------+
| ACCOUNTING |
| RESEARCH   |
| SALES      |
| OPERATIONS |
+------------+
4 rows in set (0.00 sec)

```


5. 计算员工年薪

字段可以进行"加减乘除"运算
```SQL
mysql> SELECT ENAME,SAL*12 FROM EMP;
+--------+----------+
| ENAME  | SAL*12   |
+--------+----------+
| SMITH  |  9600.00 |
| ALLEN  | 19200.00 |
| WARD   | 15000.00 |
| JONES  | 35700.00 |
| MARTIN | 15000.00 |
| BLAKE  | 34200.00 |
| CLARK  | 29400.00 |
| SCOTT  | 36000.00 |
| KING   | 60000.00 |
| TURNER | 18000.00 |
| ADAMS  | 13200.00 |
| JAMES  | 11400.00 |
| FORD   | 36000.00 |
| MILLER | 15600.00 |
+--------+----------+
14 rows in set (0.00 sec)

mysql> SELECT ENAME,SAL*12 AS YEARSAL FROM EMP;
+--------+----------+
| ENAME  | YEARSAL  |
+--------+----------+
| SMITH  |  9600.00 |
| ALLEN  | 19200.00 |
| WARD   | 15000.00 |
| JONES  | 35700.00 |
| MARTIN | 15000.00 |
| BLAKE  | 34200.00 |
| CLARK  | 29400.00 |
| SCOTT  | 36000.00 |
| KING   | 60000.00 |
| TURNER | 18000.00 |
| ADAMS  | 13200.00 |
| JAMES  | 11400.00 |
| FORD   | 36000.00 |
| MILLER | 15600.00 |
+--------+----------+
14 rows in set (0.00 sec)

mysql> SELECT ENAME,SAL*12 AS `年薪` FROM EMP;
+--------+----------+
| ENAME  |    年薪   |
+--------+----------+
| SMITH  |  9600.00 |
| ALLEN  | 19200.00 |
| WARD   | 15000.00 |
| JONES  | 35700.00 |
| MARTIN | 15000.00 |
| BLAKE  | 34200.00 |
| CLARK  | 29400.00 |
| SCOTT  | 36000.00 |
| KING   | 60000.00 |
| TURNER | 18000.00 |
| ADAMS  | 13200.00 |
| JAMES  | 11400.00 |
| FORD   | 36000.00 |
| MILLER | 15600.00 |
+--------+----------+
14 rows in set (0.00 sec)

```


## 条件查询

1. 

```sql

mysql> SELECT EMPNO,ENAME,SAL FROM EMP WHERE SAL <> 800;
+-------+--------+---------+
| EMPNO | ENAME  | SAL     |
+-------+--------+---------+
|  7499 | ALLEN  | 1600.00 |
|  7521 | WARD   | 1250.00 |
|  7566 | JONES  | 2975.00 |
|  7654 | MARTIN | 1250.00 |
|  7698 | BLAKE  | 2850.00 |
|  7782 | CLARK  | 2450.00 |
|  7788 | SCOTT  | 3000.00 |
|  7839 | KING   | 5000.00 |
|  7844 | TURNER | 1500.00 |
|  7876 | ADAMS  | 1100.00 |
|  7900 | JAMES  |  950.00 |
|  7902 | FORD   | 3000.00 |
|  7934 | MILLER | 1300.00 |
+-------+--------+---------+
13 rows in set (0.01 sec)

mysql> SELECT EMPNO,ENAME,SAL FROM EMP WHERE SAL != 800;
+-------+--------+---------+
| EMPNO | ENAME  | SAL     |
+-------+--------+---------+
|  7499 | ALLEN  | 1600.00 |
|  7521 | WARD   | 1250.00 |
|  7566 | JONES  | 2975.00 |
|  7654 | MARTIN | 1250.00 |
|  7698 | BLAKE  | 2850.00 |
|  7782 | CLARK  | 2450.00 |
|  7788 | SCOTT  | 3000.00 |
|  7839 | KING   | 5000.00 |
|  7844 | TURNER | 1500.00 |
|  7876 | ADAMS  | 1100.00 |
|  7900 | JAMES  |  950.00 |
|  7902 | FORD   | 3000.00 |
|  7934 | MILLER | 1300.00 |
+-------+--------+---------+
13 rows in set (0.00 sec)

mysql> SELECT EMPNO,ENAME,SAL FROM EMP WHERE SAL < 2000;
+-------+--------+---------+
| EMPNO | ENAME  | SAL     |
+-------+--------+---------+
|  7369 | SMITH  |  800.00 |
|  7499 | ALLEN  | 1600.00 |
|  7521 | WARD   | 1250.00 |
|  7654 | MARTIN | 1250.00 |
|  7844 | TURNER | 1500.00 |
|  7876 | ADAMS  | 1100.00 |
|  7900 | JAMES  |  950.00 |
|  7934 | MILLER | 1300.00 |
+-------+--------+---------+
8 rows in set (0.00 sec)

mysql> SELECT EMPNO,ENAME,SAL FROM EMP WHERE SAL <= 3000;

+-------+--------+---------+
| EMPNO | ENAME  | SAL     |
+-------+--------+---------+
|  7369 | SMITH  |  800.00 |
|  7499 | ALLEN  | 1600.00 |
|  7521 | WARD   | 1250.00 |
|  7566 | JONES  | 2975.00 |
|  7654 | MARTIN | 1250.00 |
|  7698 | BLAKE  | 2850.00 |
|  7782 | CLARK  | 2450.00 |
|  7788 | SCOTT  | 3000.00 |
|  7844 | TURNER | 1500.00 |
|  7876 | ADAMS  | 1100.00 |
|  7900 | JAMES  |  950.00 |
|  7902 | FORD   | 3000.00 |
|  7934 | MILLER | 1300.00 |
+-------+--------+---------+
13 rows in set (0.00 sec)


mysql> SELECT EMPNO,ENAME,SAL FROM EMP WHERE SAL > 3000;+-------+-------+---------+
| EMPNO | ENAME | SAL     |
+-------+-------+---------+
|  7839 | KING  | 5000.00 |
+-------+-------+---------+
1 row in set (0.00 sec)

mysql> SELECT EMPNO,ENAME,SAL FROM EMP WHERE SAL >= 3000;

+-------+-------+---------+
| EMPNO | ENAME | SAL     |
+-------+-------+---------+
|  7788 | SCOTT | 3000.00 |
|  7839 | KING  | 5000.00 |
|  7902 | FORD  | 3000.00 |
+-------+-------+---------+
3 rows in set (0.00 sec)

mysql> SELECT EMPNO,ENAME,SAL FROM EMP WHERE ENAME = 'JONES';
+-------+-------+---------+
| EMPNO | ENAME | SAL     |
+-------+-------+---------+
|  7566 | JONES | 2975.00 |
+-------+-------+---------+
1 row in set (0.00 sec)

mysql> SELECT EMPNO,ENAME,SAL FROM EMP WHERE SAL >= 2450 && SAL <= 3000;
+-------+-------+---------+
| EMPNO | ENAME | SAL     |
+-------+-------+---------+
|  7566 | JONES | 2975.00 |
|  7698 | BLAKE | 2850.00 |
|  7782 | CLARK | 2450.00 |
|  7788 | SCOTT | 3000.00 |
|  7902 | FORD  | 3000.00 |
+-------+-------+---------+
5 rows in set, 1 warning (0.01 sec)

mysql> SELECT EMPNO,ENAME,SAL FROM EMP WHERE SAL BETWEEN 2450 AND 3000;
+-------+-------+---------+
| EMPNO | ENAME | SAL     |
+-------+-------+---------+
|  7566 | JONES | 2975.00 |
|  7698 | BLAKE | 2850.00 |
|  7782 | CLARK | 2450.00 |
|  7788 | SCOTT | 3000.00 |
|  7902 | FORD  | 3000.00 |
+-------+-------+---------+
5 rows in set (0.00 sec)
```

查询哪些员工的津贴/补助为null
```SQL
mysql> SELECT EMPNO,ENAME,COMM FROM EMP WHERE COMM IS NULL;
+-------+--------+------+
| EMPNO | ENAME  | COMM |
+-------+--------+------+
|  7369 | SMITH  | NULL |
|  7566 | JONES  | NULL |
|  7698 | BLAKE  | NULL |
|  7782 | CLARK  | NULL |
|  7788 | SCOTT  | NULL |
|  7839 | KING   | NULL |
|  7876 | ADAMS  | NULL |
|  7900 | JAMES  | NULL |
|  7902 | FORD   | NULL |
|  7934 | MILLER | NULL |
+-------+--------+------+
10 rows in set (0.00 sec)

mysql> SELECT EMPNO,ENAME,COMM FROM EMP WHERE COMM IS NOT NULL;
+-------+--------+---------+
| EMPNO | ENAME  | COMM    |
+-------+--------+---------+
|  7499 | ALLEN  |  300.00 |
|  7521 | WARD   |  500.00 |
|  7654 | MARTIN | 1400.00 |
|  7844 | TURNER |    0.00 |
+-------+--------+---------+
4 rows in set (0.00 sec)
```

查询 JOB = 'MANAGER' 并且 SAL > 2500 的员工信息

```sql
mysql> SELECT EMPNO,ENAME,JOB,SAL FROM EMP WHERE JOB = 'MANAGER' AND SAL > 2500; 
+-------+-------+---------+---------+
| EMPNO | ENAME | JOB     | SAL     |
+-------+-------+---------+---------+
|  7566 | JONES | MANAGER | 2975.00 |
|  7698 | BLAKE | MANAGER | 2850.00 |
+-------+-------+---------+---------+
2 rows in set (0.00 sec)
```

查询工作岗位是 MANAGER 和 SALESMAN 的员工


IN 相当于多个 OR, IN 后面不是一个区间，而是多个具体的值
```SQL
mysql> SELECT EMPNO,ENAME,JOB FROM EMP WHERE JOB = 'MANAGER' OR JOB = 'SALESMAN';
+-------+--------+----------+
| EMPNO | ENAME  | JOB      |
+-------+--------+----------+
|  7499 | ALLEN  | SALESMAN |
|  7521 | WARD   | SALESMAN |
|  7566 | JONES  | MANAGER  |
|  7654 | MARTIN | SALESMAN |
|  7698 | BLAKE  | MANAGER  |
|  7782 | CLARK  | MANAGER  |
|  7844 | TURNER | SALESMAN |
+-------+--------+----------+
7 rows in set (0.01 sec)


mysql> SELECT EMPNO,ENAME,JOB FROM EMP WHERE JOB IN('MANAGER','SALESMAN');
+-------+--------+----------+
| EMPNO | ENAME  | JOB      |
+-------+--------+----------+
|  7499 | ALLEN  | SALESMAN |
|  7521 | WARD   | SALESMAN |
|  7566 | JONES  | MANAGER  |
|  7654 | MARTIN | SALESMAN |
|  7698 | BLAKE  | MANAGER  |
|  7782 | CLARK  | MANAGER  |
|  7844 | TURNER | SALESMAN |
+-------+--------+----------+
7 rows in set (0.00 sec)
```




查找工资等于800 和 5000的员工

```SQL
mysql> SELECT EMPNO,ENAME,JOB,SAL FROM EMP WHERE SAL IN(800,5000);
+-------+-------+-----------+---------+
| EMPNO | ENAME | JOB       | SAL     |
+-------+-------+-----------+---------+
|  7369 | SMITH | CLERK     |  800.00 |
|  7839 | KING  | PRESIDENT | 5000.00 |
+-------+-------+-----------+---------+
2 rows in set (0.00 sec)

mysql> SELECT EMPNO,ENAME,JOB,SAL FROM EMP WHERE SAL = 800 OR SAL = 5000;
+-------+-------+-----------+---------+
| EMPNO | ENAME | JOB       | SAL     |
+-------+-------+-----------+---------+
|  7369 | SMITH | CLERK     |  800.00 |
|  7839 | KING  | PRESIDENT | 5000.00 |
+-------+-------+-----------+---------+
2 rows in set (0.00 sec)
```


```SQL
mysql> SELECT EMPNO,ENAME,JOB,SAL FROM EMP WHERE SAL NOT IN(800,3000,5000);
+-------+--------+----------+---------+
| EMPNO | ENAME  | JOB      | SAL     |
+-------+--------+----------+---------+
|  7499 | ALLEN  | SALESMAN | 1600.00 |
|  7521 | WARD   | SALESMAN | 1250.00 |
|  7566 | JONES  | MANAGER  | 2975.00 |
|  7654 | MARTIN | SALESMAN | 1250.00 |
|  7698 | BLAKE  | MANAGER  | 2850.00 |
|  7782 | CLARK  | MANAGER  | 2450.00 |
|  7844 | TURNER | SALESMAN | 1500.00 |
|  7876 | ADAMS  | CLERK    | 1100.00 |
|  7900 | JAMES  | CLERK    |  950.00 |
|  7934 | MILLER | CLERK    | 1300.00 |
+-------+--------+----------+---------+
10 rows in set (0.01 sec)
```


AND 优先级高于 OR

查询工资大于 2500，并且部门编号为10或20 的员工 ？
```SQL
mysql> SELECT * FROM EMP WHERE SAL > 2500  AND (DEPTNO = 10 OR DEPTNO = 20);
+-------+-------+-----------+------+------------+---------+------+--------+
| EMPNO | ENAME | JOB       | MGR  | HIREDATE   | SAL     | COMM | DEPTNO |
+-------+-------+-----------+------+------------+---------+------+--------+
|  7566 | JONES | MANAGER   | 7839 | 1981-04-02 | 2975.00 | NULL |     20 |
|  7788 | SCOTT | ANALYST   | 7566 | 1987-04-19 | 3000.00 | NULL |     20 |
|  7839 | KING  | PRESIDENT | NULL | 1981-11-17 | 5000.00 | NULL |     10 |
|  7902 | FORD  | ANALYST   | 7566 | 1981-12-03 | 3000.00 | NULL |     20 |
+-------+-------+-----------+------+------------+---------+------+--------+
4 rows in set (0.00 sec)
```



**模糊查询：like   %任意多个字符   _任意一个字符**

% _  代表特殊字符，使用\ 转义

查找名字中含有 o的员工名

```SQL
mysql> SELECT ENAME FROM EMP WHERE ENAME LIKE '%O%';
+-------+
| ENAME |
+-------+
| JONES |
| SCOTT |
| FORD  |
+-------+
3 rows in set (0.00 sec)
```

查找名字中以T结尾的员工名

```SQL
mysql> SELECT ENAME FROM EMP WHERE ENAME LIKE '%T';
+-------+
| ENAME |
+-------+
| SCOTT |
+-------+
1 row in set (0.00 sec)
```

```SQL
mysql> SELECT ENAME FROM EMP WHERE ENAME LIKE 'K%';
+-------+
| ENAME |
+-------+
| KING  |
+-------+
1 row in set (0.00 sec)
```

找出第二个字母是A的

```SQL
mysql> SELECT ENAME FROM EMP WHERE ENAME LIKE '_A%';
+--------+
| ENAME  |
+--------+
| WARD   |
| MARTIN |
| JAMES  |
+--------+
3 rows in set (0.00 sec)
```

找出第3个字母是R的

```SQL
mysql> SELECT ENAME FROM EMP WHERE ENAME LIKE '__R%';
+--------+
| ENAME  |
+--------+
| WARD   |
| MARTIN |
| TURNER |
| FORD   |
+--------+
4 rows in set (0.00 sec)
```


## 排序


查询所有员工薪资，排序？

默认是升序
```SQL
mysql> SELECT ENAME,SAL FROM EMP ORDER BY SAL;
+--------+---------+
| ENAME  | SAL     |
+--------+---------+
| SMITH  |  800.00 |
| JAMES  |  950.00 |
| ADAMS  | 1100.00 |
| WARD   | 1250.00 |
| MARTIN | 1250.00 |
| MILLER | 1300.00 |
| TURNER | 1500.00 |
| ALLEN  | 1600.00 |
| CLARK  | 2450.00 |
| BLAKE  | 2850.00 |
| JONES  | 2975.00 |
| SCOTT  | 3000.00 |
| FORD   | 3000.00 |
| KING   | 5000.00 |
+--------+---------+
14 rows in set (0.00 sec)
```

指定降序

```SQL
mysql> SELECT ENAME,SAL FROM EMP ORDER BY SAL DESC;
+--------+---------+
| ENAME  | SAL     |
+--------+---------+
| KING   | 5000.00 |
| SCOTT  | 3000.00 |
| FORD   | 3000.00 |
| JONES  | 2975.00 |
| BLAKE  | 2850.00 |
| CLARK  | 2450.00 |
| ALLEN  | 1600.00 |
| TURNER | 1500.00 |
| MILLER | 1300.00 |
| WARD   | 1250.00 |
| MARTIN | 1250.00 |
| ADAMS  | 1100.00 |
| JAMES  |  950.00 |
| SMITH  |  800.00 |
+--------+---------+
14 rows in set (0.00 sec)
```
指定升序 
```SQL
mysql> SELECT ENAME,SAL FROM EMP ORDER BY SAL ASC;
+--------+---------+
| ENAME  | SAL     |
+--------+---------+
| SMITH  |  800.00 |
| JAMES  |  950.00 |
| ADAMS  | 1100.00 |
| WARD   | 1250.00 |
| MARTIN | 1250.00 |
| MILLER | 1300.00 |
| TURNER | 1500.00 |
| ALLEN  | 1600.00 |
| CLARK  | 2450.00 |
| BLAKE  | 2850.00 |
| JONES  | 2975.00 |
| SCOTT  | 3000.00 |
| FORD   | 3000.00 |
| KING   | 5000.00 |
+--------+---------+
14 rows in set (0.00 sec)
```

查询员工名字和薪资，要求按照薪资升序，如果薪资一样的话，再按照名字升序排列。

```SQL
SELECT ENAME,SAL FROM EMP ORDER BY SAL ASC, ENAME ASC;
```
SAL 在前起主导，只有 SAL 相同，启用 ENAME 排序 

```SQL
mysql> SELECT ENAME,SAL FROM EMP ORDER BY SAL ASC, ENAME ASC;
+--------+---------+
| ENAME  | SAL     |
+--------+---------+
| SMITH  |  800.00 |
| JAMES  |  950.00 |
| ADAMS  | 1100.00 |
| MARTIN | 1250.00 |
| WARD   | 1250.00 |
| MILLER | 1300.00 |
| TURNER | 1500.00 |
| ALLEN  | 1600.00 |
| CLARK  | 2450.00 |
| BLAKE  | 2850.00 |
| JONES  | 2975.00 |
| FORD   | 3000.00 |
| SCOTT  | 3000.00 |
| KING   | 5000.00 |
+--------+---------+
14 rows in set (0.01 sec)
```


了解：按查询结果的第 2 列排序，不健壮。
```SQL
mysql> SELECT ENAME,SAL FROM EMP ORDER BY 2;
+--------+---------+
| ENAME  | SAL     |
+--------+---------+
| SMITH  |  800.00 |
| JAMES  |  950.00 |
| ADAMS  | 1100.00 |
| WARD   | 1250.00 |
| MARTIN | 1250.00 |
| MILLER | 1300.00 |
| TURNER | 1500.00 |
| ALLEN  | 1600.00 |
| CLARK  | 2450.00 |
| BLAKE  | 2850.00 |
| JONES  | 2975.00 |
| SCOTT  | 3000.00 |
| FORD   | 3000.00 |
| KING   | 5000.00 |
+--------+---------+
14 rows in set (0.00 sec)
```

综合一点的案例：
	找出工资在1250到3000之间的员工信息，要求按照薪资降序排列。
```SQL
mysql> SELECT ENAME,SAL FROM EMP WHERE SAL BETWEEN 1250 AND 3000 ORDER BY SAL DESC;
+--------+---------+
| ENAME  | SAL     |
+--------+---------+
| SCOTT  | 3000.00 |
| FORD   | 3000.00 |
| JONES  | 2975.00 |
| BLAKE  | 2850.00 |
| CLARK  | 2450.00 |
| ALLEN  | 1600.00 |
| TURNER | 1500.00 |
| MILLER | 1300.00 |
| WARD   | 1250.00 |
| MARTIN | 1250.00 |
+--------+---------+
10 rows in set (0.00 sec)
```



## 数据处理函数

数据处理函数又被称为单行处理函数

    - 单行处理函数的特点：一个输入对应一个输出。

	- 和单行处理函数相对的是：多行处理函数。（多行处理函数特点：多个输入，对应1个输出！）




```SQL
mysql> SELECT LOWER(ENAME) FROM EMP;
+--------------+
| LOWER(ENAME) |
+--------------+
| smith        |
| allen        |
| ward         |
| jones        |
| martin       |
| blake        |
| clark        |
| scott        |
| king         |
| turner       |
| adams        |
| james        |
| ford         |
| miller       |
+--------------+
14 rows in set (0.02 sec)

mysql> SELECT LOWER(ENAME) AS ename FROM EMP;
+--------+
| ename  |
+--------+
| smith  |
| allen  |
| ward   |
| jones  |
| martin |
| blake  |
| clark  |
| scott  |
| king   |
| turner |
| adams  |
| james  |
| ford   |
| miller |
+--------+
14 rows in set (0.00 sec)

mysql> SELECT UPPER(ENAME) FROM EMP;
+--------------+
| UPPER(ENAME) |
+--------------+
| SMITH        |
| ALLEN        |
| WARD         |
| JONES        |
| MARTIN       |
| BLAKE        |
| CLARK        |
| SCOTT        |
| KING         |
| TURNER       |
| ADAMS        |
| JAMES        |
| FORD         |
| MILLER       |
+--------------+
14 rows in set (0.00 sec)


```


substr 取子串（substr( 被截取的字符串, 起始下标,截取的长度)） 
	注意：起始下标从1开始，没有0.

```SQL
mysql> SELECT SUBSTR(ENAME,1,2) FROM EMP;
+-------------------+
| SUBSTR(ENAME,1,2) |
+-------------------+
| SM                |
| AL                |
| WA                |
| JO                |
| MA                |
| BL                |
| CL                |
| SC                |
| KI                |
| TU                |
| AD                |
| JA                |
| FO                |
| MI                |
+-------------------+
14 rows in set (0.00 sec)
```

		找出员工名字第一个字母是A的员工信息？
			第一种方式：模糊查询
			第二种方式：substr函数

```SQL
mysql> SELECT ENAME FROM EMP WHERE ENAME LIKE 'A%';
+-------+
| ENAME |
+-------+
| ALLEN |
| ADAMS |
+-------+
2 rows in set (0.00 sec)

mysql> SELECT ENAME FROM EMP WHERE SUBSTR(ENAME, 1, 1) = 'A';
+-------+
| ENAME |
+-------+
| ALLEN |
| ADAMS |
+-------+
2 rows in set (0.00 sec)
```


concat函数进行字符串的拼接

```SQL
mysql> SELECT CONCAT(EMPNO,ENAME) FROM EMP;
+---------------------+
| CONCAT(EMPNO,ENAME) |
+---------------------+
| 7369SMITH           |
| 7499ALLEN           |
| 7521WARD            |
| 7566JONES           |
| 7654MARTIN          |
| 7698BLAKE           |
| 7782CLARK           |
| 7788SCOTT           |
| 7839KING            |
| 7844TURNER          |
| 7876ADAMS           |
| 7900JAMES           |
| 7902FORD            |
| 7934MILLER          |
+---------------------+
14 rows in set (0.00 sec)
```



```SQL
mysql> SELECT LENGTH(ENAME) FROM EMP;
+---------------+
| LENGTH(ENAME) |
+---------------+
|             5 |
|             5 |
|             4 |
|             5 |
|             6 |
|             5 |
|             5 |
|             5 |
|             4 |
|             6 |
|             5 |
|             5 |
|             4 |
|             6 |
+---------------+
14 rows in set (0.00 sec)

mysql> SELECT * FROM EMP WHERE ENAME = TRIM('  KING');
+-------+-------+-----------+------+------------+---------+------+--------+
| EMPNO | ENAME | JOB       | MGR  | HIREDATE   | SAL     | COMM | DEPTNO |
+-------+-------+-----------+------+------------+---------+------+--------+
|  7839 | KING  | PRESIDENT | NULL | 1981-11-17 | 5000.00 | NULL |     10 |
+-------+-------+-----------+------+------------+---------+------+--------+
1 row in set (0.00 sec)


```




round 四舍五入

select后面直接跟“字面量/字面值”

```SQL
mysql> SELECT 'ABC' FROM EMP;
+-----+
| ABC |
+-----+
| ABC |
| ABC |
| ABC |
| ABC |
| ABC |
| ABC |
| ABC |
| ABC |
| ABC |
| ABC |
| ABC |
| ABC |
| ABC |
| ABC |
+-----+
14 rows in set (0.00 sec)

mysql> SELECT 1000 FROM EMP;
+------+
| 1000 |
+------+
| 1000 |
| 1000 |
| 1000 |
| 1000 |
| 1000 |
| 1000 |
| 1000 |
| 1000 |
| 1000 |
| 1000 |
| 1000 |
| 1000 |
| 1000 |
| 1000 |
+------+
14 rows in set (0.00 sec)
```
结论：select后面可以跟某个表的字段名（可以等同看做变量名），也可以跟字面量/字面值（数据）。



```SQL
mysql> SELECT ROUND(3.1415, 0) AS RESULT FROM EMP;
+--------+
| RESULT |
+--------+
|      3 |
|      3 |
|      3 |
|      3 |
|      3 |
|      3 |
|      3 |
|      3 |
|      3 |
|      3 |
|      3 |
|      3 |
|      3 |
|      3 |
+--------+
14 rows in set (0.01 sec)
```

rand() 生成随机数
```SQL
mysql> SELECT RAND() FROM EMP;
+--------------------+
| RAND()             |
+--------------------+
|  0.220449269023211 |
| 0.9055510674654982 |
| 0.8664075609915028 |
| 0.6153846332946649 |
| 0.4777005142324609 |
| 0.5423487020119547 |
| 0.2786419980960358 |
| 0.7661639151779599 |
|  0.994893612335337 |
| 0.6759763468764503 |
| 0.3952009281098851 |
|   0.94807563065372 |
| 0.5547753996725896 |
| 0.9296501371354332 |
+--------------------+
14 rows in set (0.00 sec)

mysql> SELECT ROUND(RAND()*100,0) FROM EMP;
+---------------------+
| ROUND(RAND()*100,0) |
+---------------------+
|                  62 |
|                  93 |
|                  79 |
|                  15 |
|                  39 |
|                  50 |
|                  34 |
|                  17 |
|                  86 |
|                  76 |
|                  23 |
|                  88 |
|                  70 |
|                  87 |
+---------------------+
14 rows in set (0.00 sec)
```

ifnull 可以将 null 转换成一个具体值
		ifnull是空处理函数。专门处理空的。
		在所有数据库当中，只要有NULL参与的数学运算，最终结果就是NULL。

```SQL
mysql> SELECT ENAME,(SAL+COMM)*12 AS YEARSAL FROM EMP;
+--------+----------+
| ENAME  | YEARSAL  |
+--------+----------+
| SMITH  |     NULL |
| ALLEN  | 22800.00 |
| WARD   | 21000.00 |
| JONES  |     NULL |
| MARTIN | 31800.00 |
| BLAKE  |     NULL |
| CLARK  |     NULL |
| SCOTT  |     NULL |
| KING   |     NULL |
| TURNER | 18000.00 |
| ADAMS  |     NULL |
| JAMES  |     NULL |
| FORD   |     NULL |
| MILLER |     NULL |
+--------+----------+
14 rows in set (0.00 sec)
```


注意：NULL只要参与运算，最终结果一定是NULL。为了避免这个现象，需要使用ifnull函数。
			ifnull函数用法：ifnull(数据, 被当做哪个值)
				如果“数据”为NULL的时候，把这个数据结构当做哪个值。
			
			补助为NULL的时候，将补助当做0

```SQL
mysql> SELECT ENAME,(SAL+IFNULL(COMM,0))*12 AS YEARSAL FROM EMP;
+--------+----------+
| ENAME  | YEARSAL  |
+--------+----------+
| SMITH  |  9600.00 |
| ALLEN  | 22800.00 |
| WARD   | 21000.00 |
| JONES  | 35700.00 |
| MARTIN | 31800.00 |
| BLAKE  | 34200.00 |
| CLARK  | 29400.00 |
| SCOTT  | 36000.00 |
| KING   | 60000.00 |
| TURNER | 18000.00 |
| ADAMS  | 13200.00 |
| JAMES  | 11400.00 |
| FORD   | 36000.00 |
| MILLER | 15600.00 |
+--------+----------+
14 rows in set (0.00 sec)
```

	case..when..then..when..then..else..end
		当员工的工作岗位是MANAGER的时候，工资上调10%，当工作岗位是SALESMAN的时候，工资上调50%,其它正常。
		（注意：不修改数据库，只是将查询结果显示为工资上调）

```SQL
mysql> SELECT ENAME,JOB,SAL AS OLDSAL,(CASE JOB WHEN 'MANALSE' THEN SAL*1.1 WHEN 'SALSMAN' THEN SAL*1.5 ELSE SAL END) AS NEWSAL FROM EMP;
+--------+-----------+---------+---------+
| ENAME  | JOB       | OLDSAL  | NEWSAL  |
+--------+-----------+---------+---------+
| SMITH  | CLERK     |  800.00 |  800.00 |
| ALLEN  | SALESMAN  | 1600.00 | 1600.00 |
| WARD   | SALESMAN  | 1250.00 | 1250.00 |
| JONES  | MANAGER   | 2975.00 | 2975.00 |
| MARTIN | SALESMAN  | 1250.00 | 1250.00 |
| BLAKE  | MANAGER   | 2850.00 | 2850.00 |
| CLARK  | MANAGER   | 2450.00 | 2450.00 |
| SCOTT  | ANALYST   | 3000.00 | 3000.00 |
| KING   | PRESIDENT | 5000.00 | 5000.00 |
| TURNER | SALESMAN  | 1500.00 | 1500.00 |
| ADAMS  | CLERK     | 1100.00 | 1100.00 |
| JAMES  | CLERK     |  950.00 |  950.00 |
| FORD   | ANALYST   | 3000.00 | 3000.00 |
| MILLER | CLERK     | 1300.00 | 1300.00 |
+--------+-----------+---------+---------+
14 rows in set (0.00 sec)
```

## 分组函数（多行处理函数）

	多行处理函数的特点：输入多行，最终输出一行。

	5个：
		count	计数
		sum	求和
		avg	平均值
		max	最大值
		min	最小值
	
	注意：
		分组函数在使用的时候必须先进行分组，然后才能用。
		如果你没有对数据进行分组，整张表默认为一组。


```SQL
mysql> SELECT MAX(SAL) FROM EMP;
+----------+
| MAX(SAL) |
+----------+
|  5000.00 |
+----------+
1 row in set (0.01 sec)

mysql> SELECT MIN(SAL) FROM EMP;
+----------+
| MIN(SAL) |
+----------+
|   800.00 |
+----------+
1 row in set (0.00 sec)

mysql> SELECT SUM(SAL) FROM EMP;
+----------+
| SUM(SAL) |
+----------+
| 29025.00 |
+----------+
1 row in set (0.00 sec)

mysql> SELECT AVG(SAL) FROM EMP;
+-------------+
| AVG(SAL)    |
+-------------+
| 2073.214286 |
+-------------+
1 row in set (0.01 sec)

mysql> SELECT COUNT(ENAME) FROM EMP;
+--------------+
| COUNT(ENAME) |
+--------------+
|           14 |
+--------------+
1 row in set (0.00 sec)
```

分组函数在使用的时候需要注意哪些？

第一点：分组函数自动忽略NULL，你不需要提前对NULL进行处理。

```SQL
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

mysql> SELECT SUM(COMM) FROM EMP;
+-----------+
| SUM(COMM) |
+-----------+
|   2200.00 |
+-----------+
1 row in set (0.00 sec)
```


第二点：分组函数中count(*)和count(具体字段)有什么区别？


```SQL
mysql> SELECT COUNT(COMM) FROM EMP;
+-------------+
| COUNT(COMM) |
+-------------+
|           4 |
+-------------+
1 row in set (0.00 sec)

mysql> SELECT COUNT(*) FROM EMP;
+----------+
| COUNT(*) |
+----------+
|       14 |
+----------+
1 row in set (0.01 sec)
```

count(具体字段)：表示统计该字段下所有不为NULL的元素的总数。

count(*)：统计表当中的总行数。（只要有一行数据count则++）

因为每一行记录不可能都为NULL，一行数据中有一列不为NULL，则这行数据就是有效的。



**第三点：分组函数不能够直接使用在where子句中。**
			
找出比最低工资高的员工信息。

```SQL
mysql> SELECT ENAME,SAL FROM EMP WHERE SAL > MIN(SAL);
ERROR 1111 (HY000): Invalid use of group function
```


第四点：所有的分组函数可以组合起来一起用。

```SQL
mysql> SELECT SUM(SAL),AVG(SAL),MIN(SAL),MAX(SAL),COUNT(*) FROM EMP;
+----------+-------------+----------+----------+----------+
| SUM(SAL) | AVG(SAL)    | MIN(SAL) | MAX(SAL) | COUNT(*) |
+----------+-------------+----------+----------+----------+
| 29025.00 | 2073.214286 |   800.00 |  5000.00 |       14 |
+----------+-------------+----------+----------+----------+
1 row in set (0.00 sec)
```






## ✅ 分组查询 （非常重要）

将之前的关键字全部组合在一起，来看一下他们的执行顺序？

```SQL
		select
			...
		from
			...
		where
			...
		group by
			...
		order by
			...

```
以上关键字的顺序不能颠倒，需要记忆。
		执行顺序是什么？
			1. from
			2. where
			3. group by
			4. select
			5. order by


为什么分组函数不能直接使用在where后面？

			select ename,sal from emp where sal > min(sal);//报错。

**因为分组函数在使用的时候必须先分组之后才能使用。**

where执行的时候，还没有分组。所以where后面不能出现分组函数。

			select sum(sal) from emp; 
			这个没有分组，为啥sum()函数可以用呢？
				**因为select在group by之后执行。**



**找出每个工作岗位的工资和？**
	
		实现思路：按照工作岗位分组，然后对工资求和。

```SQL
mysql> SELECT JOB,SUM(SAL) FROM EMP GROUP BY JOB;
+-----------+----------+
| JOB       | SUM(SAL) |
+-----------+----------+
| CLERK     |  4150.00 |
| SALESMAN  |  5600.00 |
| MANAGER   |  8275.00 |
| ANALYST   |  6000.00 |
| PRESIDENT |  5000.00 |
+-----------+----------+
5 rows in set (0.00 sec)
```

以上这个语句的执行顺序？

- 先从emp表中查询数据。
- 根据job字段进行分组。
- 然后对每一组的数据进行sum(sal)




```SQL
mysql> SELECT ENAME,JOB,SUM(SAL) FROM EMP GROUP BY JOB;
ERROR 1055 (42000): Expression #1 of SELECT list is not in GROUP BY clause and contains nonaggregated column 'X.EMP.ENAME' which is not functionally dependent on columns in GROUP BY clause; this is incompatible with sql_mode=only_full_group_by
mysql>
```

		以上语句在mysql中可以执行，但是毫无意义。
		以上语句在oracle中执行报错。
		oracle的语法比mysql的语法严格。（mysql的语法相对来说松散一些！）

		重点结论：
			在一条select语句当中，如果有group by语句的话，
			select后面只能跟：参加分组的字段，以及分组函数。
			其它的一律不能跟。



**找出每个部门的最高薪资**
		实现思路是什么？
			按照部门编号分组，求每一组的最大值。

```SQL
mysql> SELECT DEPTNO,MAX(SAL) FROM EMP GROUP BY DEPTNO;
+--------+----------+
| DEPTNO | MAX(SAL) |
+--------+----------+
|     20 |  3000.00 |
|     30 |  2850.00 |
|     10 |  5000.00 |
+--------+----------+
3 rows in set (0.01 sec)
```

**找出“每个部门，不同工作岗位”的最高薪资？**

技巧：两个字段联合成1个字段看。（两个字段联合分组）

```SQL
mysql> SELECT JOB,SAL,DEPTNO FROM EMP ORDER BY DEPTNO;
+-----------+---------+--------+
| JOB       | SAL     | DEPTNO |
+-----------+---------+--------+
| MANAGER   | 2450.00 |     10 |
| PRESIDENT | 5000.00 |     10 |
| CLERK     | 1300.00 |     10 |
| CLERK     |  800.00 |     20 |
| MANAGER   | 2975.00 |     20 |
| ANALYST   | 3000.00 |     20 |
| CLERK     | 1100.00 |     20 |
| ANALYST   | 3000.00 |     20 |
| SALESMAN  | 1600.00 |     30 |
| SALESMAN  | 1250.00 |     30 |
| SALESMAN  | 1250.00 |     30 |
| MANAGER   | 2850.00 |     30 |
| SALESMAN  | 1500.00 |     30 |
| CLERK     |  950.00 |     30 |
+-----------+---------+--------+
14 rows in set (0.00 sec)

mysql> SELECT DEPTNO,JOB,MAX(SAL) FROM EMP GROUP BY DEPTNO,JOB;
+--------+-----------+----------+
| DEPTNO | JOB       | MAX(SAL) |
+--------+-----------+----------+
|     20 | CLERK     |  1100.00 |
|     30 | SALESMAN  |  1600.00 |
|     20 | MANAGER   |  2975.00 |
|     30 | MANAGER   |  2850.00 |
|     10 | MANAGER   |  2450.00 |
|     20 | ANALYST   |  3000.00 |
|     10 | PRESIDENT |  5000.00 |
|     30 | CLERK     |   950.00 |
|     10 | CLERK     |  1300.00 |
+--------+-----------+----------+
9 rows in set (0.00 sec)
```


使用having可以对分完组之后的数据进一步过滤。
	having不能单独使用，having不能代替where，having必须
	和group by联合使用。

**找出每个部门最高薪资，要求显示最高薪资大于3000的？**

```SQL
mysql> SELECT DEPTNO,MAX(SAL) FROM EMP GROUP BY DEPTNO;
+--------+----------+
| DEPTNO | MAX(SAL) |
+--------+----------+
|     20 |  3000.00 |
|     30 |  2850.00 |
|     10 |  5000.00 |
+--------+----------+
3 rows in set (0.00 sec)

mysql> SELECT DEPTNO,MAX(SAL) FROM EMP GROUP BY DEPTNO HAVING MAX(SAL)>3000;
+--------+----------+
| DEPTNO | MAX(SAL) |
+--------+----------+
|     10 |  5000.00 |
+--------+----------+
1 row in set (0.00 sec)
```


思考一个问题：以上的sql语句执行效率是不是低？
			比较低，实际上可以这样考虑：先将大于3000的都找出来，然后再分组。

```SQL
mysql> SELECT DEPTNO,MAX(SAL) FROM EMP WHERE SAL>3000 GROUP BY DEPTNO;
+--------+----------+
| DEPTNO | MAX(SAL) |
+--------+----------+
|     10 |  5000.00 |
+--------+----------+
1 row in set (0.00 sec)
```

优化策略：

where和having，优先选择where，where实在完成不了了，再选择 having。

where没办法的?

**找出每个部门平均薪资，要求显示平均薪资高于2500的。**

第一步：找出每个部门平均薪资
```SQL
mysql> SELECT DEPTNO,AVG(SAL) FROM EMP GROUP BY DEPTNO;
+--------+-------------+
| DEPTNO | AVG(SAL)    |
+--------+-------------+
|     20 | 2175.000000 |
|     30 | 1566.666667 |
|     10 | 2916.666667 |
+--------+-------------+
3 rows in set (0.00 sec)
```

第二步：要求显示平均薪资高于2500的
```SQL
mysql> SELECT DEPTNO,AVG(SAL) FROM EMP GROUP BY DEPTNO HAVING AVG(SAL)>2500;
+--------+-------------+
| DEPTNO | AVG(SAL)    |
+--------+-------------+
|     10 | 2916.666667 |
+--------+-------------+
1 row in set (0.01 sec)
```



大总结（单表的查询学完了）

```SQL
	select 
		...
	from
		...
	where
		...
	group by
		...
	having
		...
	order by
		...
```
	以上关键字只能按照这个顺序来，不能颠倒。

	执行顺序？
		1. from
		2. where
		3. group by
		4. having
		5. select
		6. order by

	从某张表中查询数据，
	先经过where条件筛选出有价值的数据。
	对这些有价值的数据进行分组。
	分组之后可以使用having继续筛选。
	select查询出来。
	最后排序输出！

	找出每个岗位的平均薪资，要求显示平均薪资大于1500的，除MANAGER岗位之外，
	要求按照平均薪资降序排。

```SQL
mysql> SELECT JOB,AVG(SAL) FROM EMP WHERE JOB!='MANAGER' GROUP BY JOB HAVING AVG(SAL)>1500 ORDER BY AVG(SAL) DESC;
+-----------+-------------+
| JOB       | AVG(SAL)    |
+-----------+-------------+
| PRESIDENT | 5000.000000 |
| ANALYST   | 3000.000000 |
+-----------+-------------+
2 rows in set (0.00 sec)
```