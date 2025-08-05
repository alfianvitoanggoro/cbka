package usecase

import (
	"go-kafka/internal/infrastructure/kafka"
	"go-kafka/internal/repository"
	"time"
)

type UserUsecase struct {
	repo  repository.UserRepository
	kafka kafka.Producer
}

func NewUserUsecase(repo repository.UserRepository, kafka kafka.Producer) *UserUsecase {
	return &UserUsecase{repo: repo, kafka: kafka}
}

func (u *UserUsecase) UserReconcile(userId string) error {
	payload := kafka.UserReconcile{
		UserID:    userId,
		Action:    "user.reconcile",
		Timestamp: time.Now().Format(time.RFC3339), // bisa juga pakai UnixMilli jika konsisten
	}

	return u.kafka.ProduceAvro(userId, payload)
}
