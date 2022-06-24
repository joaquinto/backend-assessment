package repository

import (
	"user/domain"
	"user/prisma/db"
)

type Repository struct {
	db *db.PrismaClient
}

func NewRepository(client *db.PrismaClient) *Repository {
	return &Repository{db: client}
}

var _ domain.IUserRepository = &Repository{}
