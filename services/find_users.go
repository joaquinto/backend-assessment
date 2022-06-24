package services

import (
	"context"
	"time"
	"user/domain"
)

func (s *Services) FindUsers(ctx context.Context, limit, pageNumber int, from, to time.Time) ([]domain.User, int, error) {
	return s.userRepo.FindUsers(ctx, limit, pageNumber, from, to)
}
