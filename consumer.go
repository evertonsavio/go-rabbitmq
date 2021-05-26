package main


import (
	"fmt"
	"github.com/streadway/amqp"
)

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