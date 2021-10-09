1. [把查询结果去除重复记录 [distinct]](#把查询结果去除重复记录-distinct)
1. [连接查询](#连接查询)
1. [子查询？](#子查询)
1. [union合并查询结果集](#union合并查询结果集)

## 把查询结果去除重复记录 [distinct]
	注意：原表数据不会被修改，只是查询结果去重。
	去重需要使用一个关键字：distinct

```SQL
mysql> SELECT DISTINCT JOB FROM EMP;
+-----------+
| JOB       |
+-----------+
| CLERK     |
| SALESMAN  |
| MANAGER   |
| ANALYST   |
| PRESIDENT |
+-----------+
5 rows in set (0.00 sec)
```

distinct出现在job,deptno两个字段之前，表示两个字段联合起来去重。
```SQL
mysql> SELECT DISTINCT JOB,DEPTNO FROM EMP;
+-----------+--------+
| JOB       | DEPTNO |
+-----------+--------+
| CLERK     |     20 |
| SALESMAN  |     30 |
| MANAGER   |     20 |
| MANAGER   |     30 |
| MANAGER   |     10 |
| ANALYST   |     20 |
| PRESIDENT |     10 |
| CLERK     |     30 |
| CLERK     |     10 |
+-----------+--------+
9 rows in set (0.00 sec)
```

统计一下工作岗位的数量？
```SQL
mysql> SELECT COUNT(DISTINCT JOB) FROM EMP;
+---------------------+
| COUNT(DISTINCT JOB) |
+---------------------+
|                   5 |
+---------------------+
1 row in set (0.00 sec)
```

## 连接查询

- 什么是连接查询？
	从一张表中单独查询，称为单表查询。
	emp表和dept表联合起来查询数据，从emp表中取员工名字，从dept表中取部门名字。
	这种跨表查询，多张表联合起来查询数据，被称为连接查询。

- 连接查询的分类？

	根据语法的年代分类：
		SQL92：1992年的时候出现的语法
		SQL99：1999年的时候出现的语法
		我们这里重点学习SQL99.(这个过程中简单演示一个SQL92的例子)
	
	根据表连接的方式分类：
		内连接：
			等值连接
			非等值连接
			自连接

		外连接：
			左外连接（左连接）
			右外连接（右连接）

		全连接（不讲）



- 当两张表进行连接查询时，没有任何条件的限制会发生什么现象？

	案例：查询每个员工所在部门名称？

```SQL
mysql> SELECT ENAME,DEPTNO FROM EMP;
+--------+--------+
| ENAME  | DEPTNO |
+--------+--------+
| SMITH  |     20 |
| ALLEN  |     30 |
| WARD   |     30 |
| JONES  |     20 |
| MARTIN |     30 |
| BLAKE  |     30 |
| CLARK  |     10 |
| SCOTT  |     20 |
| KING   |     10 |
| TURNER |     30 |
| ADAMS  |     20 |
| JAMES  |     30 |
| FORD   |     20 |
| MILLER |     10 |
+--------+--------+
14 rows in set (0.00 sec)

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
```


两张表连接没有任何条件限制：
```SQL
mysql> SELECT ENAME,LOC FROM EMP,DEPT;
+--------+----------+
| ENAME  | LOC      |
+--------+----------+
| SMITH  | BOSTON   |
| SMITH  | CHICAGO  |
| SMITH  | DALLAS   |
| SMITH  | NEW YORK |
| ALLEN  | BOSTON   |
| ALLEN  | CHICAGO  |
| ALLEN  | DALLAS   |
| ALLEN  | NEW YORK |
| WARD   | BOSTON   |
| WARD   | CHICAGO  |
| WARD   | DALLAS   |
| WARD   | NEW YORK |
| JONES  | BOSTON   |
| JONES  | CHICAGO  |
| JONES  | DALLAS   |
| JONES  | NEW YORK |
| MARTIN | BOSTON   |
| MARTIN | CHICAGO  |
| MARTIN | DALLAS   |
| MARTIN | NEW YORK |
| BLAKE  | BOSTON   |
| BLAKE  | CHICAGO  |
| BLAKE  | DALLAS   |
| BLAKE  | NEW YORK |
| CLARK  | BOSTON   |
| CLARK  | CHICAGO  |
| CLARK  | DALLAS   |
| CLARK  | NEW YORK |
| SCOTT  | BOSTON   |
| SCOTT  | CHICAGO  |
| SCOTT  | DALLAS   |
| SCOTT  | NEW YORK |
| KING   | BOSTON   |
| KING   | CHICAGO  |
| KING   | DALLAS   |
| KING   | NEW YORK |
| TURNER | BOSTON   |
| TURNER | CHICAGO  |
| TURNER | DALLAS   |
| TURNER | NEW YORK |
| ADAMS  | BOSTON   |
| ADAMS  | CHICAGO  |
| ADAMS  | DALLAS   |
| ADAMS  | NEW YORK |
| JAMES  | BOSTON   |
| JAMES  | CHICAGO  |
| JAMES  | DALLAS   |
| JAMES  | NEW YORK |
| FORD   | BOSTON   |
| FORD   | CHICAGO  |
| FORD   | DALLAS   |
| FORD   | NEW YORK |
| MILLER | BOSTON   |
| MILLER | CHICAGO  |
| MILLER | DALLAS   |
| MILLER | NEW YORK |
+--------+----------+
56 rows in set (0.01 sec)
```

14 * 4 = 56

当两张表进行连接查询，没有任何条件限制的时候，最终查询结果条数，是两张表条数的乘积，这种现象被称为：**笛卡尔积现象。**（笛卡尔发现的，这是一个数学现象。）


- 怎么避免笛卡尔积现象？
	**连接时加条件**，满足这个条件的记录被筛选出来！

```SQL
mysql> SELECT ENAME,DNAME FROM EMP,DEPT WHERE EMP.DEPTNO = DEPT.DEPTNO;
+--------+------------+
| ENAME  | DNAME      |
+--------+------------+
| SMITH  | RESEARCH   |
| ALLEN  | SALES      |
| WARD   | SALES      |
| JONES  | RESEARCH   |
| MARTIN | SALES      |
| BLAKE  | SALES      |
| CLARK  | ACCOUNTING |
| SCOTT  | RESEARCH   |
| KING   | ACCOUNTING |
| TURNER | SALES      |
| ADAMS  | RESEARCH   |
| JAMES  | SALES      |
| FORD   | RESEARCH   |
| MILLER | ACCOUNTING |
+--------+------------+
14 rows in set (0.00 sec)

mysql> SELECT EMP.ENAME,DEPT.DNAME FROM EMP,DEPT WHERE EMP.DEPTNO = DEPT.
DEPTNO;
+--------+------------+
| ENAME  | DNAME      |
+--------+------------+
| SMITH  | RESEARCH   |
| ALLEN  | SALES      |
| WARD   | SALES      |
| JONES  | RESEARCH   |
| MARTIN | SALES      |
| BLAKE  | SALES      |
| CLARK  | ACCOUNTING |
| SCOTT  | RESEARCH   |
| KING   | ACCOUNTING |
| TURNER | SALES      |
| ADAMS  | RESEARCH   |
| JAMES  | SALES      |
| FORD   | RESEARCH   |
| MILLER | ACCOUNTING |
+--------+------------+
14 rows in set (0.00 sec)
```

思考：最终查询的结果条数是14条，但是匹配的过程中，匹配的次数减少了吗？
		还是56次，只不过进行了四选一。次数没有减少。
	
	注意：通过笛卡尔积现象得出，表的连接次数越多效率越低，尽量避免表的
	连接次数。


**表起别名。很重要。效率问题。**

SQL92语法

```SQL
mysql> SELECT E.ENAME,D.DNAME FROM EMP E,DEPT D WHERE E.DEPTNO = D.DEPTNO;
+--------+------------+
| ENAME  | DNAME      |
+--------+------------+
| SMITH  | RESEARCH   |
| ALLEN  | SALES      |
| WARD   | SALES      |
| JONES  | RESEARCH   |
| MARTIN | SALES      |
| BLAKE  | SALES      |
| CLARK  | ACCOUNTING |
| SCOTT  | RESEARCH   |
| KING   | ACCOUNTING |
| TURNER | SALES      |
| ADAMS  | RESEARCH   |
| JAMES  | SALES      |
| FORD   | RESEARCH   |
| MILLER | ACCOUNTING |
+--------+------------+
14 rows in set (0.00 sec)
```

- 内连接之等值连接。

案例：查询每个员工所在部门名称，显示员工名和部门名？


SQL92语法

```SQL
mysql> SELECT E.ENAME,D.DNAME FROM EMP E,DEPT D WHERE E.DEPTNO = D.DEPTNO;
+--------+------------+
| ENAME  | DNAME      |
+--------+------------+
| SMITH  | RESEARCH   |
| ALLEN  | SALES      |
| WARD   | SALES      |
| JONES  | RESEARCH   |
| MARTIN | SALES      |
| BLAKE  | SALES      |
| CLARK  | ACCOUNTING |
| SCOTT  | RESEARCH   |
| KING   | ACCOUNTING |
| TURNER | SALES      |
| ADAMS  | RESEARCH   |
| JAMES  | SALES      |
| FORD   | RESEARCH   |
| MILLER | ACCOUNTING |
+--------+------------+
14 rows in set (0.00 sec)
```
sql92的缺点：结构不清晰，表的连接条件，和后期进一步筛选的条件，都放到了where后面。


SQL99语法：

```SQL
mysql> SELECT E.ENAME,D.DNAME FROM EMP E JOIN DEPT D ON E.DEPTNO = D.DEPTNO;
+--------+------------+
| ENAME  | DNAME      |
+--------+------------+
| SMITH  | RESEARCH   |
| ALLEN  | SALES      |
| WARD   | SALES      |
| JONES  | RESEARCH   |
| MARTIN | SALES      |
| BLAKE  | SALES      |
| CLARK  | ACCOUNTING |
| SCOTT  | RESEARCH   |
| KING   | ACCOUNTING |
| TURNER | SALES      |
| ADAMS  | RESEARCH   |
| JAMES  | SALES      |
| FORD   | RESEARCH   |
| MILLER | ACCOUNTING |
+--------+------------+
14 rows in set (0.02 sec)

mysql> SELECT E.ENAME,D.DNAME FROM EMP E INNER JOIN DEPT D ON E.DEPTNO =
D.DEPTNO;
+--------+------------+
| ENAME  | DNAME      |
+--------+------------+
| SMITH  | RESEARCH   |
| ALLEN  | SALES      |
| WARD   | SALES      |
| JONES  | RESEARCH   |
| MARTIN | SALES      |
| BLAKE  | SALES      |
| CLARK  | ACCOUNTING |
| SCOTT  | RESEARCH   |
| KING   | ACCOUNTING |
| TURNER | SALES      |
| ADAMS  | RESEARCH   |
| JAMES  | SALES      |
| FORD   | RESEARCH   |
| MILLER | ACCOUNTING |
+--------+------------+
14 rows in set (0.00 sec)
```

sql99优点：表连接的条件是独立的，连接之后，如果还需要进一步筛选，再往后继续添加where

```SQL
	SQL99语法：
		select 
			...
		from
			a
		join
			b
		on
			a和b的连接条件
		where
			筛选条件
```



- 内连接之非等值连接

案例：找出每个员工的薪资等级，要求显示员工名、薪资、薪资等级？

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
5 rows in set (0.02 sec)
```


```SQL
mysql> SELECT E.ENAME,E.SAL,S.GRADE FROM EMP E INNER JOIN SALGRADE S ON E.SAL BETWEEN S.LOSAL AND S.HISAL;
+--------+---------+-------+
| ENAME  | SAL     | GRADE |
+--------+---------+-------+
| SMITH  |  800.00 |     1 |
| ALLEN  | 1600.00 |     3 |
| WARD   | 1250.00 |     2 |
| JONES  | 2975.00 |     4 |
| MARTIN | 1250.00 |     2 |
| BLAKE  | 2850.00 |     4 |
| CLARK  | 2450.00 |     4 |
| SCOTT  | 3000.00 |     4 |
| KING   | 5000.00 |     5 |
| TURNER | 1500.00 |     3 |
| ADAMS  | 1100.00 |     1 |
| JAMES  |  950.00 |     1 |
| FORD   | 3000.00 |     4 |
| MILLER | 1300.00 |     2 |
+--------+---------+-------+
14 rows in set (0.00 sec)


SELECT
	E.ENAME,
	E.SAL,
	S.GRADE 
FROM
	EMP E
	INNER JOIN SALGRADE S ON E.SAL BETWEEN S.LOSAL 
	AND S.HISAL;
+--------+---------+-------+
| ENAME  | SAL     | GRADE |
+--------+---------+-------+
| SMITH  |  800.00 |     1 |
| ALLEN  | 1600.00 |     3 |
| WARD   | 1250.00 |     2 |
| JONES  | 2975.00 |     4 |
| MARTIN | 1250.00 |     2 |
| BLAKE  | 2850.00 |     4 |
| CLARK  | 2450.00 |     4 |
| SCOTT  | 3000.00 |     4 |
| KING   | 5000.00 |     5 |
| TURNER | 1500.00 |     3 |
| ADAMS  | 1100.00 |     1 |
| JAMES  |  950.00 |     1 |
| FORD   | 3000.00 |     4 |
| MILLER | 1300.00 |     2 |
+--------+---------+-------+
14 rows in set (0.00 sec)
```

- 内连接之自连接
案例：查询员工的上级领导，要求显示员工名和对应的领导名？

技巧：一张表看成两张表。

emp a 员工表
```SQL
+-------+--------+------+
| EMPNO | ENAME  | MGR  |
+-------+--------+------+
|  7369 | SMITH  | 7902 |
|  7499 | ALLEN  | 7698 |
|  7521 | WARD   | 7698 |
|  7566 | JONES  | 7839 |
|  7654 | MARTIN | 7698 |
|  7698 | BLAKE  | 7839 |
|  7782 | CLARK  | 7839 |
|  7788 | SCOTT  | 7566 |
|  7839 | KING   | NULL |
|  7844 | TURNER | 7698 |
|  7876 | ADAMS  | 7788 |
|  7900 | JAMES  | 7698 |
|  7902 | FORD   | 7566 |
|  7934 | MILLER | 7782 |
+-------+--------+------+
```

EMP B 领导表
```SQL
+-------+--------+------+
| EMPNO | ENAME  | MGR  |
+-------+--------+------+
|  7369 | SMITH  | 7902 |
|  7499 | ALLEN  | 7698 |
|  7521 | WARD   | 7698 |
|  7566 | JONES  | 7839 |
|  7654 | MARTIN | 7698 |
|  7698 | BLAKE  | 7839 |
|  7782 | CLARK  | 7839 |
|  7788 | SCOTT  | 7566 |
|  7839 | KING   | NULL |
|  7844 | TURNER | 7698 |
|  7876 | ADAMS  | 7788 |
|  7900 | JAMES  | 7698 |
|  7902 | FORD   | 7566 |
|  7934 | MILLER | 7782 |
+-------+--------+------+
```

```SQL
SELECT
	A.ENAME AS '员工名',
	B.ENAME AS '领导名' 
FROM
	EMP A
	INNER JOIN EMP B ON A.EMPNO = B.MGR;

+-------+--------+
| 员工名 | 领导名  |
+-------+--------+
| FORD  | SMITH  |
| BLAKE | ALLEN  |
| BLAKE | WARD   |
| KING  | JONES  |
| BLAKE | MARTIN |
| KING  | BLAKE  |
| KING  | CLARK  |
| JONES | SCOTT  |
| BLAKE | TURNER |
| SCOTT | ADAMS  |
| BLAKE | JAMES  |
| JONES | FORD   |
| CLARK | MILLER |
+-------+--------+
13 rows in set (0.01 sec)  
```



**外连接**


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


**内连接**
内连接：（A和B连接，AB两张表没有主次关系。平等的。）

内连接的特点：完成能够匹配上这个条件的数据查询出来。

```SQL
SELECT
	E.ENAME,
	D.DNAME 
FROM
	EMP E
	INNER JOIN DEPT D ON E.DEPTNO = D.DEPTNO;

+--------+------------+
| ENAME  | DNAME      |
+--------+------------+
| SMITH  | RESEARCH   |
| ALLEN  | SALES      |
| WARD   | SALES      |
| JONES  | RESEARCH   |
| MARTIN | SALES      |
| BLAKE  | SALES      |
| CLARK  | ACCOUNTING |
| SCOTT  | RESEARCH   |
| KING   | ACCOUNTING |
| TURNER | SALES      |
| ADAMS  | RESEARCH   |
| JAMES  | SALES      |
| FORD   | RESEARCH   |
| MILLER | ACCOUNTING |
+--------+------------+
14 rows in set (0.01 sec)
```


**外连接（右外连接）**

```SQL
SELECT
	E.ENAME,
	D.DNAME 
FROM
	EMP E
	RIGHT OUTER JOIN DEPT D ON E.DEPTNO = D.DEPTNO;

+--------+------------+
| ENAME  | DNAME      |
+--------+------------+
| MILLER | ACCOUNTING |
| KING   | ACCOUNTING |
| CLARK  | ACCOUNTING |
| FORD   | RESEARCH   |
| ADAMS  | RESEARCH   |
| SCOTT  | RESEARCH   |
| JONES  | RESEARCH   |
| SMITH  | RESEARCH   |
| JAMES  | SALES      |
| TURNER | SALES      |
| BLAKE  | SALES      |
| MARTIN | SALES      |
| WARD   | SALES      |
| ALLEN  | SALES      |
| NULL   | OPERATIONS |
+--------+------------+
15 rows in set (0.00 sec)
```

**right代表什么：表示将join关键字右边的这张表看成主表，主要是为了将这张表的数据全部查询出来，捎带着关联查询左边的表。**在外连接当中，两张表连接，产生了主次关系。


**外连接（左外连接）**

```SQL
SELECT
	E.ENAME,
	D.DNAME 
FROM
	DEPT D
	LEFT OUTER JOIN EMP E ON E.DEPTNO = D.DEPTNO;

+--------+------------+
| ENAME  | DNAME      |
+--------+------------+
| MILLER | ACCOUNTING |
| KING   | ACCOUNTING |
| CLARK  | ACCOUNTING |
| FORD   | RESEARCH   |
| ADAMS  | RESEARCH   |
| SCOTT  | RESEARCH   |
| JONES  | RESEARCH   |
| SMITH  | RESEARCH   |
| JAMES  | SALES      |
| TURNER | SALES      |
| BLAKE  | SALES      |
| MARTIN | SALES      |
| WARD   | SALES      |
| ALLEN  | SALES      |
| NULL   | OPERATIONS |
+--------+------------+
15 rows in set (0.00 sec)  
```

带有right的是右外连接，又叫做右连接。
带有left的是左外连接，又叫做左连接。
任何一个右连接都有左连接的写法。
任何一个左连接都有右连接的写法。

思考：外连接的查询结果条数一定是 >= 内连接的查询结果条数？

	**正确。**

```SQL
SELECT
	A.ENAME AS '员工名',
	B.ENAME AS '领导名'
FROM
	EMP A
	LEFT JOIN EMP B ON A.MGR = B.EMPNO;

+--------+-------+
| 员工名  | 领导名 |
+--------+-------+
| SMITH  | FORD  |
| ALLEN  | BLAKE |
| WARD   | BLAKE |
| JONES  | KING  |
| MARTIN | BLAKE |
| BLAKE  | KING  |
| CLARK  | KING  |
| SCOTT  | JONES |
| KING   | NULL  |
| TURNER | BLAKE |
| ADAMS  | SCOTT |
| JAMES  | BLAKE |
| FORD   | JONES |
| MILLER | CLARK |
+--------+-------+
14 rows in set (0.00 sec)
```



**三张表，四张表怎么连接？**
```SQL
	语法：
		select 
			...
		from
			a
		join
			b
		on
			a和b的连接条件
		join
			c
		on
			a和c的连接条件
		right join
			d
		on
			a和d的连接条件
```	
一条SQL中内连接和外连接可以混合。都可以出现！

**案例：找出每个员工的部门名称以及工资等级，**
	要求显示员工名、部门名、薪资、薪资等级？

```SQL
SELECT
	E.ENAME,
	D.DNAME,
	E.SAL,
	S.GRADE 
FROM
	EMP E
	JOIN DEPT D ON E.DEPTNO = D.DEPTNO
	JOIN SALGRADE S ON E.SAL BETWEEN S.LOSAL 
	AND S.HISAL;


+--------+------------+---------+-------+
| ENAME  | DNAME      | SAL     | GRADE |
+--------+------------+---------+-------+
| SMITH  | RESEARCH   |  800.00 |     1 |
| ALLEN  | SALES      | 1600.00 |     3 |
| WARD   | SALES      | 1250.00 |     2 |
| JONES  | RESEARCH   | 2975.00 |     4 |
| MARTIN | SALES      | 1250.00 |     2 |
| BLAKE  | SALES      | 2850.00 |     4 |
| CLARK  | ACCOUNTING | 2450.00 |     4 |
| SCOTT  | RESEARCH   | 3000.00 |     4 |
| KING   | ACCOUNTING | 5000.00 |     5 |
| TURNER | SALES      | 1500.00 |     3 |
| ADAMS  | RESEARCH   | 1100.00 |     1 |
| JAMES  | SALES      |  950.00 |     1 |
| FORD   | RESEARCH   | 3000.00 |     4 |
| MILLER | ACCOUNTING | 1300.00 |     2 |
+--------+------------+---------+-------+
14 rows in set (0.01 sec)
```


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
5 rows in set (0.01 sec)
```


**案例：找出每个员工的部门名称以及工资等级，还有上级领导，
	要求显示员工名、领导名、部门名、薪资、薪资等级？**

```SQL
SELECT
	E.ENAME,
	L.ENAME AS 'LEADER',
	D.DNAME,
	E.SAL,
	S.GRADE 
FROM
	EMP E
	JOIN DEPT D ON E.DEPTNO = D.DEPTNO
	JOIN SALGRADE S ON E.SAL BETWEEN S.LOSAL 
	AND S.HISAL
	LEFT JOIN EMP L ON E.MGR = L.EMPNO;

+--------+--------+------------+---------+-------+
| ENAME  | LEADER | DNAME      | SAL     | GRADE |
+--------+--------+------------+---------+-------+
| SMITH  | FORD   | RESEARCH   |  800.00 |     1 |
| ALLEN  | BLAKE  | SALES      | 1600.00 |     3 |
| WARD   | BLAKE  | SALES      | 1250.00 |     2 |
| JONES  | KING   | RESEARCH   | 2975.00 |     4 |
| MARTIN | BLAKE  | SALES      | 1250.00 |     2 |
| BLAKE  | KING   | SALES      | 2850.00 |     4 |
| CLARK  | KING   | ACCOUNTING | 2450.00 |     4 |
| SCOTT  | JONES  | RESEARCH   | 3000.00 |     4 |
| KING   | NULL   | ACCOUNTING | 5000.00 |     5 |
| TURNER | BLAKE  | SALES      | 1500.00 |     3 |
| ADAMS  | SCOTT  | RESEARCH   | 1100.00 |     1 |
| JAMES  | BLAKE  | SALES      |  950.00 |     1 |
| FORD   | JONES  | RESEARCH   | 3000.00 |     4 |
| MILLER | CLARK  | ACCOUNTING | 1300.00 |     2 |
+--------+--------+------------+---------+-------+
14 rows in set (0.00 sec)
```



## 子查询？

- 什么是子查询？
	select语句中嵌套select语句，被嵌套的select语句称为子查询。

- 子查询都可以出现在哪里呢？
```SQL
	select
		..(select).
	from
		..(select).
	where
		..(select).
```

- where子句中的子查询

**案例：找出比最低工资高的员工姓名和工资？**

```SQL
mysql> SELECT ENAME,SAL FROM EMP WHERE MIN(SAL);
ERROR 1111 (HY000): Invalid use of group function


mysql> SELECT MIN(SAL) FROM EMP;
+----------+
| MIN(SAL) |
+----------+
|   800.00 |
+----------+
1 row in set (0.00 sec)

mysql> SELECT ENAME,SAL FROM EMP WHERE SAL>800;
+--------+---------+
| ENAME  | SAL     |
+--------+---------+
| ALLEN  | 1600.00 |
| WARD   | 1250.00 |
| JONES  | 2975.00 |
| MARTIN | 1250.00 |
| BLAKE  | 2850.00 |
| CLARK  | 2450.00 |
| SCOTT  | 3000.00 |
| KING   | 5000.00 |
| TURNER | 1500.00 |
| ADAMS  | 1100.00 |
| JAMES  |  950.00 |
| FORD   | 3000.00 |
| MILLER | 1300.00 |
+--------+---------+
13 rows in set (0.00 sec)

mysql> SELECT ENAME,SAL FROM EMP WHERE SAL>(SELECT MIN(SAL) FROM EMP);
+--------+---------+
| ENAME  | SAL     |
+--------+---------+
| ALLEN  | 1600.00 |
| WARD   | 1250.00 |
| JONES  | 2975.00 |
| MARTIN | 1250.00 |
| BLAKE  | 2850.00 |
| CLARK  | 2450.00 |
| SCOTT  | 3000.00 |
| KING   | 5000.00 |
| TURNER | 1500.00 |
| ADAMS  | 1100.00 |
| JAMES  |  950.00 |
| FORD   | 3000.00 |
| MILLER | 1300.00 |
+--------+---------+
13 rows in set (0.02 sec)
```


**from子句中的子查询
	注意：from后面的子查询，可以将子查询的查询结果当做一张临时表。（技巧）**

- 案例：找出每个岗位的平均工资的薪资等级。

```SQL
SELECT T.*,S.GRADE FROM (
SELECT JOB,AVG(SAL) AS AVGSAL FROM EMP GROUP BY JOB) T JOIN SALGRADE S ON T.AVGSAL BETWEEN S.LOSAL AND S.HISAL;

+-----------+-------------+-------+
| JOB       | AVGSAL      | GRADE |
+-----------+-------------+-------+
| CLERK     | 1037.500000 |     1 |
| SALESMAN  | 1400.000000 |     2 |
| MANAGER   | 2758.333333 |     4 |
| ANALYST   | 3000.000000 |     4 |
| PRESIDENT | 5000.000000 |     5 |
+-----------+-------------+-------+
5 rows in set (0.00 sec)
```


select后面出现的子查询（这个内容不需要掌握，了解即可！！！）

案例：找出每个员工的部门名称，要求显示员工名，部门名？

```SQL
mysql> SELECT E.ENAME,(SELECT D.DNAME FROM DEPT D WHERE E.DEPTNO=D.DEPTNO) AS
DNAME FROM EMP E;
+--------+------------+
| ENAME  | DNAME      |
+--------+------------+
| SMITH  | RESEARCH   |
| ALLEN  | SALES      |
| WARD   | SALES      |
| JONES  | RESEARCH   |
| MARTIN | SALES      |
| BLAKE  | SALES      |
| CLARK  | ACCOUNTING |
| SCOTT  | RESEARCH   |
| KING   | ACCOUNTING |
| TURNER | SALES      |
| ADAMS  | RESEARCH   |
| JAMES  | SALES      |
| FORD   | RESEARCH   |
| MILLER | ACCOUNTING |
+--------+------------+
14 rows in set (0.00 sec)
```

```SQL
mysql> SELECT E.ENAME,(SELECT D.DNAME FROM DEPT D ) AS DNAME FROM EMP E;
ERROR 1242 (21000): Subquery returns more than 1 row
```

注意：对于select后面的子查询来说，这个子查询只能一次返回1条结果，
	多于1条，就报错了。！


## union合并查询结果集

- 案例：查询工作岗位是MANAGER和SALESMAN的员工？

```SQL
mysql> SELECT ENAME,JOB FROM EMP WHERE JOB='MANAGER' OR JOB='SALESMAN';
+--------+----------+
| ENAME  | JOB      |
+--------+----------+
| ALLEN  | SALESMAN |
| WARD   | SALESMAN |
| JONES  | MANAGER  |
| MARTIN | SALESMAN |
| BLAKE  | MANAGER  |
| CLARK  | MANAGER  |
| TURNER | SALESMAN |
+--------+----------+
7 rows in set (0.01 sec)

mysql> SELECT ENAME,JOB FROM EMP WHERE JOB IN('MANAGER','SALESMAN');
+--------+----------+
| ENAME  | JOB      |
+--------+----------+
| ALLEN  | SALESMAN |
| WARD   | SALESMAN |
| JONES  | MANAGER  |
| MARTIN | SALESMAN |
| BLAKE  | MANAGER  |
| CLARK  | MANAGER  |
| TURNER | SALESMAN |
+--------+----------+
7 rows in set (0.00 sec)
```


```SQL
SELECT ENAME,JOB FROM EMP WHERE JOB='MANAGER' UNION 
SELECT ENAME,JOB FROM EMP WHERE JOB='SALESMAN';

+--------+----------+
| ENAME  | JOB      |
+--------+----------+
| JONES  | MANAGER  |
| BLAKE  | MANAGER  |
| CLARK  | MANAGER  |
| ALLEN  | SALESMAN |
| WARD   | SALESMAN |
| MARTIN | SALESMAN |
| TURNER | SALESMAN |
+--------+----------+
7 rows in set (0.00 sec)
```


union的效率要高一些。对于表连接来说，每连接一次新表，
	则匹配的次数满足笛卡尔积，成倍的翻。。。
	但是union可以减少匹配的次数。在减少匹配次数的情况下，
	还可以完成两个结果集的拼接。

```SQL
	a 连接 b 连接 c
	a 10条记录
	b 10条记录
	c 10条记录
	匹配次数是：1000

	a 连接 b一个结果：10 * 10 --> 100次
	a 连接 c一个结果：10 * 10 --> 100次
	使用union的话是：100次 + 100次 = 200次。（**union把乘法变成了加法运算**）
```

union在使用的时候有注意事项吗？

	//错误的：union在进行结果集合并的时候，要求两个结果集的列数相同。


- **limit（非常重要）**
	
limit作用：将查询结果集的一部分取出来。通常使用在分页查询当中。
百度默认：一页显示10条记录。
分页的作用是为了提高用户的体验，因为一次全部都查出来，用户体验差。
可以一页一页翻页看。

limit怎么用呢？

	完整用法：limit startIndex, length
		startIndex是起始下标，length是长度。
		起始下标从0开始。

	缺省用法：limit 5; 这是取前5.

- 按照薪资降序，取出排名在前5名的员工？

```SQL
mysql> SELECT ENAME,SAL FROM EMP ORDER BY SAL DESC LIMIT 5;
+-------+---------+
| ENAME | SAL     |
+-------+---------+
| KING  | 5000.00 |
| FORD  | 3000.00 |
| SCOTT | 3000.00 |
| JONES | 2975.00 |
| BLAKE | 2850.00 |
+-------+---------+
5 rows in set (0.00 sec)

mysql> SELECT ENAME,SAL FROM EMP ORDER BY SAL DESC LIMIT 0,5;
+-------+---------+
| ENAME | SAL     |
+-------+---------+
| KING  | 5000.00 |
| FORD  | 3000.00 |
| SCOTT | 3000.00 |
| JONES | 2975.00 |
| BLAKE | 2850.00 |
+-------+---------+
5 rows in set (0.00 sec)
```

注意：mysql当中limit在order by之后执行！

- 取出工资排名在[3-5]名的员工？

```SQL
mysql> SELECT ENAME,SAL FROM EMP ORDER BY SAL DESC LIMIT 2,3;
+-------+---------+
| ENAME | SAL     |
+-------+---------+
| SCOTT | 3000.00 |
| JONES | 2975.00 |
| BLAKE | 2850.00 |
+-------+---------+
3 rows in set (0.00 sec)
```

- 取出工资排名在[5-9]名的员工？

```SQL
mysql> SELECT ENAME,SAL FROM EMP ORDER BY SAL DESC LIMIT 4,5;
+--------+---------+
| ENAME  | SAL     |
+--------+---------+
| BLAKE  | 2850.00 |
| CLARK  | 2450.00 |
| ALLEN  | 1600.00 |
| TURNER | 1500.00 |
| MILLER | 1300.00 |
+--------+---------+
5 rows in set (0.01 sec)
```


**分页**

每页显示3条记录
	第1页：limit 0,3		[0 1 2]
	第2页：limit 3,3		[3 4 5]
	第3页：limit 6,3		[6 7 8]
	第4页：limit 9,3		[9 10 11]

每页显示pageSize条记录
	第pageNo页：limit (pageNo - 1) * pageSize  , pageSize

	public static void main(String[] args){
		// 用户提交过来一个页码，以及每页显示的记录条数
		int pageNo = 5; //第5页
		int pageSize = 10; //每页显示10条

		int startIndex = (pageNo - 1) * pageSize;
		String sql = "select ...limit " + startIndex + ", " + pageSize;
	}

记公式：
	**limit (pageNo-1)*pageSize , pageSize**


**关于DQL语句的大总结：**

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
	limit
		...
	
	执行顺序？
		1.from
		2.where
		3.group by
		4.having
		5.select
		6.order by
		7.limit..
```


```sQL
创建一个学生表？
	学号、姓名、年龄、性别、邮箱地址
	create table t_student(
		no int,
		name varchar(32),
		sex char(1),
		age int(3),
		email varchar(255)
	);

	删除表：
		drop table t_student; // 当这张表不存在的时候会报错！

		// 如果这张表存在的话，删除
		drop table if exists t_student;

```




- 插入数据insert （DML）
	
	语法格式：
		insert into 表名(字段名1,字段名2,字段名3...) values(值1,值2,值3);

		注意：字段名和值要一一对应。什么是一一对应？
			数量要对应。数据类型要对应。

```SQL
mysql> INSERT INTO STUDENT (NO,NAME,SEX,EMAIL) VALUES (29,'MOJO','M','X@EMAIL.COM');
Query OK, 1 row affected (0.01 sec)
```

```SQL
DROP TABLE IF EXISTS STUDENT;
CREATE TABLE STUDENT (NO INT, SEX CHAR DEFAULT 'M',EMAIL CHAR(255));

mysql> DESC STUDENT;
+-------+-----------+------+-----+---------+-------+
| Field | Type      | Null | Key | Default | Extra |
+-------+-----------+------+-----+---------+-------+
| NO    | int       | YES  |     | NULL    |       |
| SEX   | char(1)   | YES  |     | M       |       |
| EMAIL | char(255) | YES  |     | NULL    |       |
+-------+-----------+------+-----+---------+-------+
3 rows in set (0.00 sec)

mysql> INSERT INTO STUDENT VALUES(2,'M','X@EMAIL.COM');
Query OK, 1 row affected (0.01 sec)

mysql> SELECT * FROM STUDENT;
+------+------+-------------+
| NO   | SEX  | EMAIL       |
+------+------+-------------+
|    2 | M    | X@EMAIL.COM |
+------+------+-------------+
1 row in set (0.00 sec)
```

- insert插入日期

	数字格式化：format

格式化数字：format(数字, '格式')

```SQL
mysql> SELECT ENAME,SAL FROM EMP;
+--------+---------+
| ENAME  | SAL     |
+--------+---------+
| SMITH  |  800.00 |
| ALLEN  | 1600.00 |
| WARD   | 1250.00 |
| JONES  | 2975.00 |
| MARTIN | 1250.00 |
| BLAKE  | 2850.00 |
| CLARK  | 2450.00 |
| SCOTT  | 3000.00 |
| KING   | 5000.00 |
| TURNER | 1500.00 |
| ADAMS  | 1100.00 |
| JAMES  |  950.00 |
| FORD   | 3000.00 |
| MILLER | 1300.00 |
+--------+---------+
14 rows in set (0.00 sec)

mysql> SELECT ENAME,FORMAT(SAL,'$999,999')AS SAL FROM EMP;
+--------+-------+
| ENAME  | SAL   |
+--------+-------+
| SMITH  | 800   |
| ALLEN  | 1,600 |
| WARD   | 1,250 |
| JONES  | 2,975 |
| MARTIN | 1,250 |
| BLAKE  | 2,850 |
| CLARK  | 2,450 |
| SCOTT  | 3,000 |
| KING   | 5,000 |
| TURNER | 1,500 |
| ADAMS  | 1,100 |
| JAMES  | 950   |
| FORD   | 3,000 |
| MILLER | 1,300 |
+--------+-------+
14 rows in set, 14 warnings (0.00 sec)
```

str_to_date：将字符串varchar类型转换成date类型

date_format：将date类型转换成具有一定格式的varchar字符串类型。

```SQL
drop table if exists t_user;
	create table t_user(
		id int,
		name varchar(32),
		birth date // 生日也可以使用date日期类型
	);

	create table t_user(
		id int,
		name varchar(32),
		birth char(10) // 生日可以使用字符串，没问题。
	);

```

	生日：1990-10-11 （10个字符）

	注意：数据库中的有一条命名规范：
		所有的标识符都是全部小写，单词和单词之间使用下划线进行衔接。


插入数据？
		insert into t_user(id,name,birth) values(1, 'zhangsan', '01-10-1990'); // 1990年10月1日
		出问题了：原因是类型不匹配。数据库birth是date类型，这里给了一个字符串varchar。

		怎么办？可以使用str_to_date函数进行类型转换。
```sql
		str_to_date函数可以将字符串转换成日期类型date？
		语法格式：
			str_to_date('字符串日期', '日期格式')

		mysql的日期格式：
			%Y	年
			%m 月
			%d 日
			%h	时
			%i	分
			%s	秒
		
		insert into t_user(id,name,birth) values(1, 'zhangsan', str_to_date('01-10-1990','%d-%m-%Y'));
```
		str_to_date函数可以把字符串varchar转换成日期date类型数据，
		通常使用在插入insert方面，因为插入的时候需要一个日期类型的数据，
		需要通过该函数将字符串转换成date。

	好消息？
		如果你提供的日期字符串是这个格式，str_to_date函数就不需要了！！！
			%Y-%m-%d
		insert into t_user(id,name,birth) values(2, 'lisi', '1990-10-01');
	
	查询的时候可以以某个特定的日期格式展示吗？
		date_format
		这个函数可以将日期类型转换成特定格式的字符串。
```sql
		select id,name,date_format(birth, '%m/%d/%Y') as birth from t_user;
		+------+----------+------------+
		| id   | name     | birth      |
		+------+----------+------------+
		|    1 | zhangsan | 10/01/1990 |
		|    2 | lisi     | 10/01/1990 |
		+------+----------+------------+
```
		date_format函数怎么用？
			date_format(日期类型数据, '日期格式')

```sql
		mysql> select id,name,birth from t_user;
		+------+----------+------------+
		| id   | name     | birth      |
		+------+----------+------------+
		|    1 | zhangsan | 1990-10-01 |
		|    2 | lisi     | 1990-10-01 |
		+------+----------+------------+
```

		以上的SQL语句实际上是进行了默认的日期格式化，
		自动将数据库中的date类型转换成varchar类型。
		并且采用的格式是mysql默认的日期格式：'%Y-%m-%d'
```sql
		select id,name,date_format(birth,'%Y/%m/%d') as birth from t_user;
		
```


- date和datetime两个类型的区别？
  
	date是短日期：只包括年月日信息。
	datetime是长日期：包括年月日时分秒信息。
```sql
	drop table if exists t_user;
	create table t_user(
		id int,
		name varchar(32),
		birth date,
		create_time datetime
	);
```

	id是整数
	name是字符串
	birth是短日期
	create_time是这条记录的创建时间：长日期类型

	mysql短日期默认格式：%Y-%m-%d
	mysql长日期默认格式：%Y-%m-%d %h:%i:%s
```sql
	insert into t_user(id,name,birth,create_time) values(1,'zhangsan','1990-10-01','2020-03-18 15:49:50');
```
	在mysql当中怎么获取系统当前时间？
		now() 函数，并且获取的时间带有：时分秒信息！！！！是datetime类型的。
```sql	
		insert into t_user(id,name,birth,create_time) values(2,'lisi','1991-10-01',now());
```


- 修改update（DML）

语法格式：
	update 表名 set 字段名1=值1,字段名2=值2,字段名3=值3... where 条件;

	注意：没有条件限制会导致所有数据全部更新。
```sql
	update t_user set name = 'jack', birth = '2000-10-11' where id = 2;
	+------+----------+------------+---------------------+
	| id   | name     | birth      | create_time         |
	+------+----------+------------+---------------------+
	|    1 | zhangsan | 1990-10-01 | 2020-03-18 15:49:50 |
	|    2 | jack     | 2000-10-11 | 2020-03-18 15:51:23 |
	+------+----------+------------+---------------------+


mysql> update t_user set name='mojo',birth='2000-10-01' where id=2;
Query OK, 1 row affected (0.01 sec)
Rows matched: 1  Changed: 1  Warnings: 0

mysql> select * from t_user;
+------+----------+------------+---------------------+
| id   | name     | birth      | create_time         |
+------+----------+------------+---------------------+
|    1 | zhangsan | 1990-10-01 | 2020-03-18 15:49:50 |
|    2 | mojo     | 2000-10-01 | 2021-10-09 10:36:50 |
+------+----------+------------+---------------------+
2 rows in set (0.00 sec)

```

```sql
	update t_user set name = 'jack', birth = '2000-10-11', create_time = now() where id = 2;

	更新所有？
		update t_user set name = 'abc';
```
- 删除数据 delete （DML）

```sql
	语法格式？
		delete from 表名 where 条件;

	注意：没有条件，整张表的数据会全部删除！

	delete from t_user where id = 2;

	insert into t_user(id) values(2);

	delete from t_user; // 删除所有！
```



