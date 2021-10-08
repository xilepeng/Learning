

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

