package services

import (
	"user/domain"
	"user/repository"
)

type Services struct {
	userRepo *repository.Repository
}

func NewService(userRepo *repository.Repository) *Services {
	return &Services{userRepo: userRepo}
}

var _ domain.IUserService = &Services{}
