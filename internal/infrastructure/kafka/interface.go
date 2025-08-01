package kafka

import "github.com/segmentio/kafka-go"

type Producer interface {
	NewWriter() *kafka.Writer
	PublishMessage(key, value []byte) error
	Produce(key, value string) error
}

type Consumer interface {
	NewReader() *kafka.Reader
	Consume(handler func(msg kafka.Message)) error
}
