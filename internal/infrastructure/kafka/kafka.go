package kafka

import (
	"fmt"
	"go-kafka/internal/config"
	"go-kafka/pkg/logger"
)

type KafkaClient struct {
	Brokers []string
	Topic   string
	GroupID string
}

func NewKafkaClient(ConfigKafka *config.Kafka) *KafkaClient {
	brokers := []string{fmt.Sprintf("%s:%d", ConfigKafka.Host, ConfigKafka.Port)}

	logger.Info("✅ Kafka client initialized Successfully")

	return &KafkaClient{
		Brokers: brokers,
		Topic:   ConfigKafka.Topic,
		GroupID: ConfigKafka.GroupID,
	}
}
