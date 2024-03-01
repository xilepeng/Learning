package main

import "RabbitMQ/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "Simple")
	rabbitmq.ConsumeSimple()
}
