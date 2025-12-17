package models

import (
	"time"
)

type User struct {
	ID   int64     `json:"id"`
	Name string    `json:"name" validate:"required,min2"`
	DOB  time.Time `json:"dob" validate:"required"`
	Age  int       `json:"age",omitempty`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}
