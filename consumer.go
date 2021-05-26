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
	Mac     string `json mac`
	Message string `json message`
}

func main(){
	fmt.Println("Consumer Application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func(){
		for d := range msgs {
			fmt.Printf("Received msg: %s\n", d.Body)

			var msgBody MessageBody
			json.Unmarshal(d.Body, &msgBody)

			//log.Println(msgBody.Mac)
			//log.Println(msgBody.Message)
			msgBody.Mac = strings.ReplaceAll(msgBody.Mac, ":", "") 
			fmt.Println(msgBody.Mac)

			//WRITING TO FILE
			f, err := os.OpenFile(msgBody.Mac, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			f.WriteString(string(d.Body) + "\n")

			defer f.Close()
		
			log.SetOutput(f)
		}
	}()

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println("[*] - waiting for messages")
	
	<-forever

}


//GO ROUTINE & CHANNELS EXAMPLE
//=========================================
// import (
// 	"fmt"
// 	"time"
// )

// func say(s string, done chan string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// 	done <- "Terminei"
	
// }

// func main() {
// 	done := make(chan string)
// 	go say("world", done)
// 	fmt.Println(<-done)
// }
//=========================================