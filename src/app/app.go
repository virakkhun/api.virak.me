package app

import (
	"os"
	"strings"

	"api.virak.me/src/config/database"
	"github.com/gofiber/fiber/v2"
)

func Bootstrap() {
	database.DBInitializer()

	database.Migrator(database.DB)

	app := fiber.New()

	PORT := strings.Join([]string{"", os.Getenv("PORT")}, ":")

	Middleware(app)

	SetupRoutes(app)

	app.Listen(PORT)
}
