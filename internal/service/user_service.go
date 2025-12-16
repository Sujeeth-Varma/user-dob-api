package service

import (
	"context"
	"errors"

	"github.com/Sujeeth-Varma/user-dob-api/internal/models"
	"github.com/Sujeeth-Varma/user-dob-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{r}
}

func (s *UserService) Create(ctx context.Context, user models.User) (*models.User, error) {
	if user.DOB.IsZero() {
		return nil, errors.New("dob is required")
	}

	dbUser, err := s.repo.Create(ctx, user.Name, user.DOB)

	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int64(dbUser.ID),
		Name: dbUser.Name,
		DOB:  dbUser.Dob,
	}, nil
}
