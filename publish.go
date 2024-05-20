package main

import "GinProject/rabbitMQ"

func main() {
	rabbit := rabbitMQ.NewRabbitMQ("go_exchange", "go_route", "go_queue")
	defer rabbit.Close()
	rabbit.SendMessage(rabbitMQ.Message{Body: "这是一条普通消息"})
	rabbit.SendDelayMessage(rabbitMQ.Message{Body: "这是一条延时5秒的消息", DelayTime: 5})
}
