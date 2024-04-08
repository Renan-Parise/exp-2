package main

import (
	"context"
	"log"
	"rabbit/rabbit"

	"go.uber.org/fx"
)

func Register(lifeCycle fx.Lifecycle, rabbit *rabbit.RabbitMQ) {
	storeMsgs, err := rabbit.ConsumeMessage(context.Background(), "store_orders")
	if err != nil {
		log.Fatalf("Failed to consume messages from store_orders: %v", err)
	}

	shippingMsgs, err := rabbit.ConsumeMessage(context.Background(), "shipping_orders")
	if err != nil {
		log.Fatalf("Failed to consume messages from shipping_orders: %v", err)
	}

	lifeCycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				for message := range storeMsgs {
					log.Printf("Received message from store_orders: %s", string(message.Body))
				}
			}()

			go func() {
				for message := range shippingMsgs {
					log.Printf("Received message from shipping_orders: %s", string(message.Body))
				}
			}()
			return nil
		},
		OnStop: nil,
	})
}

func main() {
	app := fx.New(
		fx.Provide(rabbit.NewRabbitMQ),
		fx.Invoke(Register),
	)
	app.Run()
}
