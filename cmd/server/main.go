package main

import (
	"database/sql"
	"os"

	sqlc "github.com/Sujeeth-Varma/user-dob-api/db/sqlc"
	"github.com/Sujeeth-Varma/user-dob-api/internal/handler"
	"github.com/Sujeeth-Varma/user-dob-api/internal/logger"
	"github.com/Sujeeth-Varma/user-dob-api/internal/middleware"
	"github.com/Sujeeth-Varma/user-dob-api/internal/repository"
	"github.com/Sujeeth-Varma/user-dob-api/internal/routes"
	"github.com/Sujeeth-Varma/user-dob-api/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load() // loads .env file
	conn, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	queries := sqlc.New(conn)

	log := logger.New()
	repo := repository.NewUserRepository(queries)
	svc := service.NewUserService(repo)
	UserHandler := handler.NewUserHandler(svc, log)

	app := fiber.New()
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(log))
	routes.Register(app, UserHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	app.Listen(":" + port)
}
