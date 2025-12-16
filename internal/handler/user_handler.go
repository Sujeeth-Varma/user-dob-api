package handler

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/Sujeeth-Varma/user-dob-api/internal/models"
	"github.com/Sujeeth-Varma/user-dob-api/internal/service"
)

type UserHandler struct {
	svc *service.UserService
	val *validator.Validate
	log *zap.Logger
}

// createUserRequest is used only for request binding & validation.
// It accepts a date-only format for dob: YYYY-MM-DD
type createUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

func NewUserHandler(s *service.UserService, l *zap.Logger) *UserHandler {
	return &UserHandler{s, validator.New(), l}
}

func (handler *UserHandler) Create(c *fiber.Ctx) error {
	var req createUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := handler.val.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "dob must be in YYYY-MM-DD format")
	}

	u := models.User{Name: req.Name, DOB: dob}

	res, err := handler.svc.Create(c.Context(), u)
	if err != nil {
		handler.log.Error("user creation failed", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	handler.log.Info("User has been successfully created")

	return c.Status(fiber.StatusCreated).JSON(res)
}
