package app

import (
	auth "api.virak.me/src/modules/auth/infrastructure/routes"
	blog "api.virak.me/src/modules/blog/infrastructure/routes"
	user "api.virak.me/src/modules/user/infrastructure/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	blog.SetupRoutes(app)
	user.SetupRoutes(app)
	auth.SetupAuthRoutes(app)
}
