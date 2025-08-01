package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func (k *KafkaClient) NewReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     k.Brokers,
		GroupID:     k.GroupID,
		Topic:       k.Topic,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset,
	})
}

func (k *KafkaClient) Consume(handler func(msg kafka.Message)) error {
	reader := k.NewReader()
	defer reader.Close()

	ctx := context.Background()
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			return fmt.Errorf("error reading message: %w", err)
		}
		handler(msg)
	}
}
