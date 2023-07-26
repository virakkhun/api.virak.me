package routes

import (
	"api.virak.me/modules/user/infrastructure/controllers"
	sharedServices "api.virak.me/shared/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiPrefix := sharedServices.ApiPrefix("user")
	userRoutes := app.Group(apiPrefix)

	userRoutes.Get("/", controllers.Get)
	userRoutes.Post("/", controllers.Create)
	userRoutes.Patch("/", controllers.Update)
	userRoutes.Delete("/", controllers.Delete)
}
