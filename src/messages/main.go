package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	go server()
	go client()
}

func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello from the other side -> %s"),
	}
	for {
		ch.Publish("", q.Name, false, false, msg)
	}
}

func client() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	failOnErr(err, "Consumer couldn't register")

	for msg := range msgs {
		log.Printf("Ah ha I have recived: %s", msg.Body)
	}
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnErr(err, "Failed to connect")

	ch, err := conn.Channel()
	failOnErr(err, "Fail to open a channel")

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil)
	failOnErr(err, "Fail to declare the Queue")

	return conn, ch, &q
}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}

}
