package repository

import (
	"context"
	"encoding/json"
	"user/domain"
	"user/prisma/db"
)

// FindUserByEmail This returns a valid user by passing a valid email parameter
func (r *Repository) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.db.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(ctx)

	if err != nil {
		return domain.User{}, err
	}

	var user domain.User

	uByte, _ := json.Marshal(u)
	_ = json.Unmarshal(uByte, &user)

	return user, nil
}
