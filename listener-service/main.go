package main

import (
	"fmt"
	event "listener/events"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Connect RabbitMQ
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// Waiting for message from broker
	log.Println("Listening for and consuming RabbitMQ message")

	// Create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}
	// Watch queue and consume event
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var count int64
	var undo = 1 * time.Second
	var connection *amqp.Connection

	// Start only if Rabbitmq would be ready
	for {
		conn, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("Waiting for rMQ")
			count++
		} else {
			log.Println("Exelent, u logged to rMQ")
			connection = conn
			break
		}

		if count > 4 {
			fmt.Println(err)
			return nil, err
		}

		undo = time.Duration(math.Pow(float64(count), 2)) * time.Second
		log.Println("Undo...")
		time.Sleep(undo)
		continue
	}

	return connection, nil
}
