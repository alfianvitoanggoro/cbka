package kafka

import "github.com/segmentio/kafka-go"

type Producer interface {
	NewWriter() *kafka.Writer
	PublishMessage(key, value []byte) error
	Produce(key, value string) error
	ProduceAvro(key string, payload UserReconcile) error
}

type Consumer interface {
	NewReader() *kafka.Reader
	Consume(handler func(msg kafka.Message)) error
}

type UserReconcile struct {
	UserID    string `avro:"user_id"`
	Action    string `avro:"action"`
	Timestamp string `avro:"timestamp"`
}
