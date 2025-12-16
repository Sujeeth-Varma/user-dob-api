package models

import (
	"time"
)

type User struct {
	ID   int64     `json:"id"`
	Name string    `json:"name" validate:"required,min2"`
	DOB  time.Time `json:"dob" validate:"required"`
}
