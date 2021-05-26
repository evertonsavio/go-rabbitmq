package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main()  {
	fmt.Println("Go with RabbitMQ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/");

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to rabbitMQ")

}