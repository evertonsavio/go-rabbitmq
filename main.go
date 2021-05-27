package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go with RabbitMQ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to rabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"log.firmware.queue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"log.firmware.queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",                                                                      //"text/plain",
			Body:        []byte("{\"mac\":\"11:22:33:44:55:66\",\"log\":\"this is my log sent by iot device\"}"), //[]byte("Message send to a rabbitmq queue and received by other service"),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully Publish method to queue!")

}
