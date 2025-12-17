package service

import (
	"context"
	"errors"
	"time"

	db "github.com/Sujeeth-Varma/user-dob-api/db/sqlc"
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

func (s *UserService) GetAge(dob time.Time) int {
	return CalculateAge(dob)
}

func (s *UserService) GetById(ctx context.Context, id int32) (*models.User, error) {
	dbUser, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:   int64(dbUser.ID),
		Name: dbUser.Name,
		DOB:  dbUser.Dob,
	}, nil
}

func (s *UserService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) GetList(ctx context.Context) ([]db.User, error) {
	return s.repo.List(ctx)
}

func (s *UserService) Update(ctx context.Context, id int32, user models.User) (*models.User, error) {
	updatedUser, err := s.repo.Update(ctx, id, user.Name, user.DOB)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:   int64(updatedUser.ID),
		Name: updatedUser.Name,
		DOB:  updatedUser.Dob,
	}, nil
}
