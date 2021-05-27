package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

type MessageBody struct {
	Mac string `json mac`
	Log string `json log`
}

func main() {

	RABBIT_HOST := os.Getenv("RABBIT_HOST")
	fmt.Println("Rabbit host is: " + RABBIT_HOST)

	//RABBIT CONNECTION/////////////////////////////////////////////////
	conn, err := amqp.Dial("amqp://guest:guest@" + RABBIT_HOST + ":5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Successfully connected to RabbitMQ")

	//RABBIT CHANNEL//////////////
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer channel.Close()

	//CONSUMING QUEUE//////////////////
	msgs, err := channel.Consume(
		"log.firmware.queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//GO CHANNEL TO RUN FOREVER
	forever := make(chan bool)

	//GO ROUTINE -> LETS GO!!!
	go func() {
		for d := range msgs {
			fmt.Printf("Received msg: %s\n", d.Body)

			var msgBody MessageBody
			json.Unmarshal(d.Body, &msgBody)

			msgBody.Mac = strings.ReplaceAll(msgBody.Mac, ":", "")

			//CREATE DIR IF IF IT DOES NOT EXISTS///////////////
			if _, err := os.Stat("./logs"); os.IsNotExist(err) {
				os.Mkdir("./logs", 0700)
			}

			//WRITING TO FILE///////////////////////////////////////////////////////////////////////////
			f, err := os.OpenFile("./logs/"+msgBody.Mac+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			f.WriteString(msgBody.Log + "\n")

			defer f.Close()

			log.SetOutput(f)
		}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println("[*] - waiting for messages")

	<-forever

}
