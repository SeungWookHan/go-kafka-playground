package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "my-topic",
	}
	err := w.WriteMessages(
		context.Background(),
		kafka.Message{Value: []byte("Hello Kafka!")},
	)
	if err != nil {
		log.Fatal("Failed to send message", err)
	}
	w.Close()
}
