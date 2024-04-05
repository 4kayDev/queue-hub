package main

import (
	"context"
	queuehub "github.com/4kayDev/queuehub/interface"
	"github.com/4kayDev/queuehub/pkg/nats"
	"log"
)

type MessageExample struct {
	Data string
}

func main() {

	cfg := nats.Config{
		Storage:            nats.NewInMem(),
		MaxRedeliveryCount: 5,
		BatchSize:          32,
		QueueName:          "test",
		ConnectionDSN:      "localhost:4222",
	}

	q, err := nats.New[MessageExample](cfg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Starting consumer...")
	err = q.Consume(context.Background(), func(ctx context.Context, msg MessageExample, meta *queuehub.Meta) (queuehub.Result, error) {
		log.Println("Received message:", msg, "Attempts:", meta.AttemptNumber)
		return queuehub.DEFER, nil
	})

	log.Fatalln(err)
}
