package usecase

import (
	"errors"
	"go-kafka/internal/dto"
	"go-kafka/internal/model"
	"go-kafka/internal/repository"
	"go-kafka/pkg/utils"
)

type AuthUsecase interface {
	Create(req *dto.RequestRegistration) (*dto.ResponseRegistration, error)
	Login(req *dto.RequestLogin) (*dto.ResponseLogin, error)
}

type authUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) AuthUsecase {
	return &authUsecase{
		repo: repo,
	}
}

func (u *authUsecase) Create(req *dto.RequestRegistration) (*dto.ResponseRegistration, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	reqUser := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Role:     req.Role,
		Password: hashedPassword,
		IsActive: true,
	}

	data, err := u.repo.Create(reqUser)

	if err != nil {
		return nil, err
	}

	res := &dto.ResponseRegistration{
		UserID:    data.UserID,
		Username:  data.Username,
		Email:     data.Email,
		Role:      data.Role,
		IsActive:  data.IsActive,
		LastLogin: data.LastLogin,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return res, nil
}

func (u *authUsecase) Login(req *dto.RequestLogin) (*dto.ResponseLogin, error) {
	user, err := u.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.ComparePassword(user.Password, req.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.UserID, user.Email, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &dto.ResponseLogin{Token: token}, nil
}
