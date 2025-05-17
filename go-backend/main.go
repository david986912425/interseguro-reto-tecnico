package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"go-backend/internal/handlers"
	"go-backend/internal/middlewares"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Post("/api/matrix", middlewares.ValidateJWT, handlers.PostMatrix)
	app.Post("/login", handlers.Login)

	app.Listen(":8080")
}
