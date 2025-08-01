package factory

import (
	"go-kafka/internal/handler"
	"go-kafka/internal/infrastructure/kafka"

	"gorm.io/gorm"
)

type Factory struct {
	AuthHandler *handler.AuthHandler
	UserHandler *handler.UserHandler
}

func NewFactory(db *gorm.DB, kafkaProducer kafka.Producer) *Factory {

	return &Factory{
		AuthHandler: InitAuthFactory(db),
		UserHandler: InitUserFactory(db, kafkaProducer),
	}
}
