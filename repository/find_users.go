package repository

import (
	"context"
	"encoding/json"
	"time"
	"user/domain"
	"user/prisma/db"
)

func (r *Repository) FindUsers(ctx context.Context, limit, pageNumber int, from, to time.Time) ([]domain.User, int, error) {
	var u []db.UserModel
	var err error
	var pn int
	var lmt int

	u, err = r.db.User.FindMany().Exec(ctx)
	if err != nil {
		return []domain.User{}, 0, err
	}

	totalUser := len(u)

	dd, _ := time.Parse("2006-01-02", "")

	if from != dd && to != dd {
		if limit != 0 {
			lmt = limit
			if pageNumber != 0 {
				pn = pageNumber * lmt
			} else {
				pn = 0
			}
			u, err = r.db.User.FindMany(
				db.User.CreatedAt.Gte(from),
				db.User.CreatedAt.Lte(to),
			).Take(lmt).Skip(pn).Exec(ctx)
		} else {
			lmt = 10
			pn = 0
			u, err = r.db.User.FindMany(
				db.User.CreatedAt.Gte(from),
				db.User.CreatedAt.Lte(to),
			).Take(lmt).Skip(pn).Exec(ctx)
		}
	} else {
		if limit != 0 {
			lmt = limit
			if pageNumber != 0 {
				pn = pageNumber * lmt
			} else {
				pn = 0
			}
			u, err = r.db.User.FindMany().Take(lmt).Skip(pn).Exec(ctx)
		} else {
			lmt = 10
			pn = 0
			u, err = r.db.User.FindMany().Take(lmt).Skip(pn).Exec(ctx)
		}
	}

	if err != nil {
		return []domain.User{}, 0, err
	}

	var users []domain.User
	uByte, _ := json.Marshal(u)
	_ = json.Unmarshal(uByte, &users)

	return users, totalUser, nil
}
