package main

import (
	"RabbitMQ/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "Simple")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishSimple("hello work" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

// 工作模式：起到负载均衡的作用
