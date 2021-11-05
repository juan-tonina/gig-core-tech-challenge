package main

// I have never used golang in my life, so I hope the code below this line is not terrible. I followed a tutorial
// so some functions might be very similar to functions found on the internet

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	//The credentials, url and port are hardcoded. I would 100% move that to env variables if I had the time. Note that
	// for my current job I am using windows, so I have to connect using this url instead of the commented one.

	// Using docker, rabbitmq, node or go on windows is frustrating :)... actually, programming in windows is frustrating
	conn, err := amqp.Dial("amqp://guest:guest@host.docker.internal:5672/")
	//conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		failOnError(err, "Oh no :(")
	}(conn)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		failOnError(err, "'Oh no' but the face is upside down ):")
	}(ch)

	q, err := ch.QueueDeclare(
		"gig-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// TODO push message via websocket
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
