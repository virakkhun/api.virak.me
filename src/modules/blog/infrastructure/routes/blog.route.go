package blog

import (
	blogControllers "api.virak.me/src/modules/blog/infrastructure/controllers"
	sharedServices "api.virak.me/src/shared/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	apiPrefix := sharedServices.ApiPrefix("blogs")
	blogRoutes := app.Group(apiPrefix)

	blogRoutes.Get("/one", blogControllers.GetOneBlog)
	blogRoutes.Get("/", blogControllers.GetAllBlog)
	blogRoutes.Post("/", blogControllers.CreateOneBlog)
	blogRoutes.Patch("/", blogControllers.UpdateOneBlog)
	blogRoutes.Delete("/", blogControllers.DeleteOneBlog)
}
