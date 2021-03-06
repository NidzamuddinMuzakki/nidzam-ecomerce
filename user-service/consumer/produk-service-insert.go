package consumer

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func ConsumerInsert() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	amqpServerURL := os.Getenv("RABBITMQ_URL")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// Subscribing to QueueService1 for getting messages.
	messages, err := channelRabbitMQ.Consume(
		"CobaService", // queue name
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no local
		false,         // no wait
		nil,           // arguments
	)
	if err != nil {
		log.Println(err)
	}

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.

			log.Printf(" > Received message: %s\n", message.Body)
		}
	}()
	<-forever

	// fmt.Println(c)

}
