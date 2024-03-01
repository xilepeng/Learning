package main

import (
	"RabbitMQ/RabbitMQ"
)

// 订阅消息
func main() {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduce")
	rabbitmq.ReceiverSub()
}
