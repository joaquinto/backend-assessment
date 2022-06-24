package domain

import (
	"context"
	"time"
)

type IUserService interface {
	FindUserByEmail(ctx context.Context, email string) (User, error)
	FindUsers(ctx context.Context, limit, pageNumber int, from, to time.Time) ([]User, int, error)
}