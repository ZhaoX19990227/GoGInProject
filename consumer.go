package main

import (
	"GinProject/rabbitMQ"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	// 新建连接
	rabbit := rabbitMQ.NewRabbitMQ("go_exchange", "go_route", "go_queue")
	// 一般来说消费者不关闭，常驻进程进行消息消费处理
	// defer rabbit.Close()

	// 执行消费
	rabbit.Consume(func(d amqp.Delivery) {
		fmt.Printf("Received a message: %s\n", d.Body)
	})
}
