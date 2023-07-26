package routes

import (
	"api.virak.me/src/modules/auth/infrastructure/controllers"
	sharedServices "api.virak.me/src/shared/services"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	apiPrefix := sharedServices.ApiPrefix("auth")
	authRoutes := app.Group(apiPrefix)

	authRoutes.Post("/login", controllers.Login)
	authRoutes.Post("/register", controllers.Register)
}
