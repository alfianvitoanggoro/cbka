package usecase

import (
	"encoding/json"
	"fmt"
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
	payload := map[string]string{
		"user_id":   userId,
		"action":    "user.reconcile",
		"timestamp": fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	jsonVal, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return u.kafka.PublishMessage([]byte(userId), jsonVal)
}
