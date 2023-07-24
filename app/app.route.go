package app

import (
	auth "api.virak.me/modules/auth/infrastructure/routes"
	blog "api.virak.me/modules/blog/infrastructure/routes"
	user "api.virak.me/modules/user/infrastructure/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	blog.SetupRoutes(app)
	user.SetupRoutes(app)
	auth.SetupAuthRoutes(app)
}
