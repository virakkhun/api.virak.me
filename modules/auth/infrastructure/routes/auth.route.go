package routes

import (
	"api.virak.me/modules/auth/infrastructure/controllers"
	shared_services "api.virak.me/shared/services"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	apiPrefix := shared_services.ApiPrefix("auth")
	authRoutes := app.Group(apiPrefix)

	authRoutes.Post("/login", controllers.Login)
	authRoutes.Post("/register", controllers.Register)
}
