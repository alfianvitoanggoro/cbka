package kafka

import (
	"go-kafka/internal/config"
	"go-kafka/pkg/logger"
)

type KafkaClient struct {
	Brokers []string
	Topic   string
	GroupID string
}

func NewKafkaClient(ConfigKafka *config.Kafka) *KafkaClient {
	brokers := []string{ConfigKafka.Broker}

	logger.Infof("✅ Kafka client initialized. Brokers: %v, Topic: %s, GroupID: %s", brokers, ConfigKafka.Topic, ConfigKafka.GroupID)

	return &KafkaClient{
		Brokers: brokers,
		Topic:   ConfigKafka.Topic,
		GroupID: ConfigKafka.GroupID,
	}
}
