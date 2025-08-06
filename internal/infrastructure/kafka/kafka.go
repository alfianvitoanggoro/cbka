package kafka

import (
	"go-kafka/internal/config"
	"go-kafka/pkg/logger"
)

type KafkaClient struct {
	Brokers  []string
	Topic    string
	GroupID  string
	Username string
	Password string
}

func NewKafkaClient(ConfigKafka *config.Kafka) *KafkaClient {

	logger.Info("✅ Kafka client initialized Successfully")

	return &KafkaClient{
		Brokers:  ConfigKafka.Brokers,
		Topic:    ConfigKafka.Topic,
		GroupID:  ConfigKafka.GroupID,
		Username: ConfigKafka.Username,
		Password: ConfigKafka.Password,
	}
}
