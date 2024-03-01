package RabbitMQ

import (
	"context"
	"fmt"
	"log"
	"time"

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

// 断开 channel connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Panicf("%s:%s", message, err)
	}
}

// 简单模式：step1 创建简单模式下 RabbitMQ 实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// 简单模式：step2 简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	// 1.申请队列， 如果队列不存在会自动创建，如果存在则跳过创建

	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, // 是否持久化
		false, // 是否自动删除
		false, // 是否具有排他性
		false, // 是否阻塞
		nil,   // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 2. 发送消息到队列中
	err = r.channel.PublishWithContext(ctx,
		r.Exchange,
		r.QueueName,
		// 如果为 true, 根据 exchange 类型和 route key规则，
		// 如果无法找到符合条件的队列，那么会把发送的消息返还给发送者，
		false,
		// 如果为true,当 exchange 发送消息
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	r.failOnErr(err, "Failed to publish a message")

}

// 简单模式：step3 简单模式下消费代码
func (r *RabbitMQ) ConsumeSimple() {
	// 1.申请队列， 如果队列不存在会自动创建，如果存在则跳过创建

	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		// 是否自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 2. 接受消息到队列中
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",    // 用来区分多个消费者
		true,  // 是否自动应答
		false, // 是否具有排他性
		false, // 如果设置为 true, 表示不能将同一个 connection 中发送的消息传递给这个connection 中的消费者
		false, // 队列消费是否阻塞
		nil,   // 其他参数
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	// 3. 启用协程处理消息
	go func() {
		for d := range msgs {
			// 实现我们要处理的逻辑函数
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("[*] Waiting for messages, To exit press")
	<-forever
}
