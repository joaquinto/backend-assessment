package services

import (
	"context"
	"user/domain"
)

func (s *Services) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.userRepo.FindUserByEmail(ctx, email)
}
