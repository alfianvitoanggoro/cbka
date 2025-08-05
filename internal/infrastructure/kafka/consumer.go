package kafka

import (
	"context"
	"fmt"
	"go-kafka/pkg/logger"

	"github.com/hamba/avro"
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

// ConsumeAvro reads and processes Kafka messages using Avro decoding
func (k *KafkaClient) ConsumeAvro() error {
	reader := k.NewReader()
	defer reader.Close()

	ctx := context.Background()
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			return fmt.Errorf("error reading message: %w", err)
		}

		var ur UserReconcile
		err = avro.Unmarshal(SchemaStr, msg.Value, &ur)
		if err != nil {
			logger.Errorf("❌ Failed to decode Avro message: %v", err)
			continue
		}

		// Call internal handler
		handleUserReconcile(ur)
	}
}

// handleUserReconcile is the internal handler for decoded messages
func handleUserReconcile(data UserReconcile) {
	logger.Infof("📥 [Kafka] Received user reconcile: %+v", data)
}
