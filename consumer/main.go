package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "my-topic",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("Received message: %s\n", m.Value)
	}
	r.Close()
}
