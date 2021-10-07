
1. [简单查询](#简单查询)
1. [条件查询](#条件查询)
1. [排序](#排序)

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