package factory

import (
	"go-kafka/internal/handler"
	"go-kafka/internal/infrastructure/kafka"
	"go-kafka/internal/repository"
	"go-kafka/internal/usecase"

	"gorm.io/gorm"
)

func InitUserFactory(db *gorm.DB, producer kafka.Producer) *handler.UserHandler {
	userRepo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(userRepo, producer)
	return handler.NewUserHandler(usecase)
}
