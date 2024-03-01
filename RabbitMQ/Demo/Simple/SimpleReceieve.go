package main

import (
	"RabbitMQ/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("xSimple")
	rabbitmq.ConsumeSimple()
}
