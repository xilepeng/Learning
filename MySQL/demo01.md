
``` sql
docker exec -it mysql bash

mysql -u root -p

123456
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


```SQL

```