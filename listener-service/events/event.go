package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declateExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		// Name
		"logs_topic",
		// Type
		"topic",
		// Driable
		true,
		// Need to be auto delated?
		false,
		// Internatl?
		false,
		// Some wait?
		false,
		// Arguments
		nil,
	)
}

func declateRandomQue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		// Name
		"",
		// Durable, Yes or No
		false,
		// When its unused need to be removed
		false,
		// Exclusieve
		true,
		// Need to wait a little
		false,
		// Arguments
		nil,
	)
}
