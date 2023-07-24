package app

import (
	"os"

	"api.virak.me/config/database"
	"github.com/gofiber/fiber/v2"
)

func Bootstrap() {
	database.DBInitializer()

	database.Migrator(database.DB)

	app := fiber.New()

	PORT := os.Getenv("PORT")

	Middleware(app)

	SetupRoutes(app)

	app.Listen(PORT)
}
