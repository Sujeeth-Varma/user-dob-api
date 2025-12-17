package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Sujeeth-Varma/user-dob-api/internal/handler"
)

func Register(app *fiber.App, UserHandler *handler.UserHandler) {
	users := app.Group("/users")

	users.Post("/", UserHandler.Create)
	users.Get("/", UserHandler.GetList)
	users.Get("/:id", UserHandler.GetById)
	users.Put("/:id", UserHandler.Update)
	users.Delete("/:id", UserHandler.Delete)
}
