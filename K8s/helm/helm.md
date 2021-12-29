
```go

ubuntu@master:~/helm$ cat config.yaml

mysqlUser: xUser
mysqlDatabase: xDB
persistence:
   enable: false
service:
  type: NodePort


ubuntu@master:~/helm$ helm install mysql stable/mysql

ubuntu@master:~/helm$ helm upgrade -f config.yaml mysql stable/mysql

Release "mysql" has been upgraded. Happy Helming!
NAME: mysql


```






           