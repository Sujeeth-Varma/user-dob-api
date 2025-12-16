package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Sujeeth-Varma/user-dob-api/internal/handler"
)

func Register(app *fiber.App, UserHandler *handler.UserHandler) {
	app.Post("/users", UserHandler.Create)
}
