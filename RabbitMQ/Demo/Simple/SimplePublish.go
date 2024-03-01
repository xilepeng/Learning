package main

import (
	"RabbitMQ/RabbitMQ"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("xSimple")
	rabbitmq.PublishSimple("Hello x")
	fmt.Println("发送成功！")
}
