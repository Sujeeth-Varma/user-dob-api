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

func (r *UserRepository) GetById(ctx context.Context, id int32) (db.User, error) {
	return r.q.GetUserById(ctx, id)
}

func (r *UserRepository) Update(ctx context.Context, id int32, name string, dob time.Time) (db.User, error) {
	return r.q.UpdateUser(ctx, db.UpdateUserParams{ID: id, Name: name, Dob: dob})
}

func (r *UserRepository) Delete(ctx context.Context, id int32) error {
	return r.q.DeleteUser(ctx, id)
}

func (r *UserRepository) List(ctx context.Context) ([]db.User, error) {
	return r.q.ListUsers(ctx)
}
