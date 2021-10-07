

1. [数据处理函数](#数据处理函数)

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





