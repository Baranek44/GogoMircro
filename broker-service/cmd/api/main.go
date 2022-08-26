package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const webPort = "8080"

// Connect to rabbitmq
type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	// Connect RabbitMQ
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	//One root and responde json
	app := Config{
		Rabbit: rabbitConn,
	}

	log.Printf("Broker Service has started at port: %s!\n", webPort)

	// Defining http servers
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// Starting the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
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
