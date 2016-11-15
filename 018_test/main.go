package main

import (
	"log"

	"github.com/streadway/amqp"
)

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

var processBuff = make(chan bool, 4)

func publish() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:15672/")
	checkErr(err)

	ch, err := conn.Channel()
	checkErr(err)

	err = ch.Publish(
		"LiveFaresLog", // exchange
		"Error",        // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("test_body"),
		})
}

func main() {
	publish()
}
