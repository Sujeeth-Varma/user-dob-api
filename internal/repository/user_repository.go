package repository

import (
	"context"
	"time"

	"github.com/Sujeeth-Varma/user-dob-api/db/sqlc"
)

type UserRepository struct {
	q *db.Queries
}

func (r *UserRepository) Create(ctx context.Context, name string, dob time.Time) (db.User, error) {
	return r.q.CreateUser(ctx, db.CreateUserParams{Name: name, Dob: dob})
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{q}
}
