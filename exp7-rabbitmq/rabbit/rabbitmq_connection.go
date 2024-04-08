package rabbit

import (
	"context"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQ, error) {
	conn, err := amqp.Dial("amqp://renan:secret@localhost:5672/")
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn, channel}, nil
}

func (r *RabbitMQ) DeclareQueues() error {
	storeQueue, err := r.channel.QueueDeclare(
		"store_orders",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	shippingQueue, err := r.channel.QueueDeclare(
		"shipping_orders",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Printf("Declared queues: %s, %s", storeQueue.Name, shippingQueue.Name)
	return nil
}

func (r *RabbitMQ) ConsumeMessage(ctx context.Context, queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := r.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	return msgs, err
}
