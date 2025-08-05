package kafka

import (
	"context"
	"go-kafka/pkg/logger"
	"time"

	"github.com/hamba/avro"
	"github.com/segmentio/kafka-go"
)

func (k *KafkaClient) NewWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(k.Brokers...),
		Topic:    k.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func (k *KafkaClient) PublishMessage(key, value []byte) error {
	writer := k.NewWriter()
	defer writer.Close()

	msg := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	return writer.WriteMessages(context.Background(), msg)
}

func (k *KafkaClient) Produce(key, value string) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  k.Brokers,
		Topic:    k.Topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	}

	err := w.WriteMessages(context.Background(), msg)
	if err != nil {
		logger.Errorf("❌ Failed to send message to Kafka: %v", err)
		return err
	}

	logger.Infof("✅ Kafka message sent. Key: %s, Value: %s", key, value)
	return nil
}

func (k *KafkaClient) ProduceAvro(key string, payload UserReconcile) error {
	value, err := avro.Marshal(SchemaStr, payload)
	if err != nil {
		logger.Errorf("❌ Failed to encode Avro: %v", err)
		return err
	}

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  k.Brokers,
		Topic:    k.Topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	msg := kafka.Message{
		Key:   []byte(key),
		Value: value,
		Time:  time.Now(),
	}

	err = w.WriteMessages(context.Background(), msg)
	if err != nil {
		logger.Errorf("❌ Kafka write failed: %v", err)
		return err
	}

	logger.Infof("✅ Kafka Avro message sent. Key: %s", key)
	return nil
}
