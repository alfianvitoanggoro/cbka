package factory

import (
	"go-kafka/internal/handler"
	"go-kafka/internal/repository"
	"go-kafka/internal/usecase"

	"gorm.io/gorm"
)

func InitAuthFactory(db *gorm.DB) *handler.AuthHandler {
	userRepo := repository.NewUserRepository(db)
	authUC := usecase.NewAuthUsecase(userRepo)
	authHandler := handler.NewAuthHandler(authUC)

	return authHandler
}
