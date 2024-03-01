package RabbitMQ

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

const MQURL = "amqp://mojo:mojo@127.0.0.1:5672/x"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string // 队列名称
	Exchange  string // 交换机
	Key       string // key
	Mqurl     string // 连接信息
}

// 订阅模式：消息被路由投递给多个队列，一个消息被多个消费者获取

// 创建 RabbitMQ 结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	var err error
	// 创建 rabbitmq 连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取 channel 失败")
	return rabbitmq
}

// 订阅模式创建 RabbotMQ 实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	// 创建 RabbotMQ 实例
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
}
