

## 1. 把查询结果去除重复记录【distinct】
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

## 2. 连接查询

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