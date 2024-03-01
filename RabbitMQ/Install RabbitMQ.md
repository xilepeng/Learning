

Docker 安装  RabbitMQ

```shell
docker run -itd --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq

```

开启rabbit节点
```shell
# 开启RabbitMQ节点
rabbitmqctl start_app
```

开启RabbitMQ管理模块的插件
```shell
# 开启RabbitMQ管理模块的插件，并配置到RabbitMQ节点上
rabbitmq-plugins enable rabbitmq_management


# 查看插件
rabbitmq-plugins list
```


关闭rabbit节点
```shell
# 关闭RabbitMQ节点
rabbitmqctl stop
```

此时RabbitMQ的管理模块已经安装上去了，你可以回到浏览器登陆。

[http://localhost:15672](http://localhost:15672)


用户名：guest
密码：  guest





