package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/Sujeeth-Varma/user-dob-api/internal/models"
	"github.com/Sujeeth-Varma/user-dob-api/internal/service"
)

type UserHandler struct {
	service   *service.UserService
	validator *validator.Validate
	logger    *zap.Logger
}

func NewUserHandler(s *service.UserService, l *zap.Logger) *UserHandler {
	return &UserHandler{s, validator.New(), l}
}

func (handler *UserHandler) Create(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := handler.validator.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "dob must be in YYYY-MM-DD format")
	}

	u := models.User{Name: req.Name, DOB: dob}

	res, err := handler.service.Create(c.Context(), u)
	if err != nil {
		handler.logger.Error("user creation failed", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	handler.logger.Info("User has been successfully created")

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":   res.ID,
		"name": res.Name,
		"dob":  res.DOB.Format("2006-01-02"),
	})
}

func (handler *UserHandler) GetById(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 32)

	user, err := handler.service.GetById(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.DOB.Format("2006-01-02"),
		"age":  handler.service.GetAge(user.DOB),
	})
}

func (handler *UserHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 32)

	_, err := handler.service.GetById(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}
	err = handler.service.Delete(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (handler *UserHandler) GetList(c *fiber.Ctx) error {
	users, err := handler.service.GetList(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	response := make([]fiber.Map, 0)
	for _, user := range users {
		response = append(response, fiber.Map{
			"id":   user.ID,
			"name": user.Name,
			"dob":  user.Dob.Format("2006-01-02"),
			"age":  handler.service.GetAge(user.Dob),
		})
	}
	return c.JSON(response)
}

func (handler *UserHandler) Update(c *fiber.Ctx) error {
	// parse the reqBody
	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	id, _ := strconv.ParseInt(c.Params("id"), 10, 32)

	// validate the reqBody
	if err := handler.validator.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "dob must be in YYYY-MM-DD format")
	}

	// check if the user exists to update
	_, err = handler.service.GetById(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}

	res, err := handler.service.Update(c.Context(), int32(id), models.User{
		Name: req.Name,
		DOB:  dob,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":   res.ID,
		"name": res.Name,
		"dob":  res.DOB.Format("2006-01-02"),
	})
}
