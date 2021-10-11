
**insert语句可以一次插入多条记录吗？【掌握】**
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

