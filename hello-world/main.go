package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	reader := bufio.NewReader(os.Stdin)
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

	go SetupConsumer(ch, &q)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		// TODO: produce message on user input
		Send(message)
	}
}
