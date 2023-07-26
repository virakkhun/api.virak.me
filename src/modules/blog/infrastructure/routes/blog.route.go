package blog

import (
	blogControllers "api.virak.me/src/modules/blog/infrastructure/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	blogRoutes := app.Group("/blogs")

	blogRoutes.Get("/", blogControllers.GetAllBlog)
	blogRoutes.Get("/one", blogControllers.GetOneBlog)
	blogRoutes.Post("/", blogControllers.CreateOneBlog)
	blogRoutes.Patch("/", blogControllers.UpdateOneBlog)
	blogRoutes.Delete("/", blogControllers.DeleteOneBlog)
}
