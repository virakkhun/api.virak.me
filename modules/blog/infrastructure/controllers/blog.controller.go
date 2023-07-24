package controllers

import (
	"fmt"

	"api.virak.me/modules/blog/domain/dto"
	"api.virak.me/modules/blog/domain/entities"
	"api.virak.me/modules/blog/domain/validators"
	blogService "api.virak.me/modules/blog/infrastructure/services"
	shared_services "api.virak.me/shared/services"
	"api.virak.me/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllBlog(c *fiber.Ctx) error {
	allBlogs := blogService.GetBlogs()

	return shared_services.HandlerResponseMapper(c, allBlogs.StatusCode, allBlogs)
}

func GetOneBlog(c *fiber.Ctx) error {
	oneBlog := blogService.GetOneBlog(c.Query("blogId"))
	return shared_services.HandlerResponseMapper(c, oneBlog.StatusCode, oneBlog)
}

func CreateOneBlog(c *fiber.Ctx) error {
	blog := new(dto.BlogDTO)

	err := c.BodyParser(blog)

	if utils.IsNotNil(err) {
		return shared_services.HandlerResponseMapper(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	validatedErrors := validators.ValidateCreateBlogSchema(*blog)

	if validatedErrors != nil {
		return shared_services.HandlerResponseMapper(c, fiber.StatusInternalServerError, validatedErrors)
	}

	newBlog := blogService.CreateOneBlog(entities.BlogEntity{
		Title:       blog.Title,
		Description: blog.Description,
		Content:     blog.Content,
		Viewer:      int16(blog.Viewer),
	})

	return shared_services.HandlerResponseMapper(c, newBlog.StatusCode, newBlog)
}

func UpdateOneBlog(c *fiber.Ctx) error {
	blogId := c.Query("blogId")

	blog := new(dto.UpdateBlogDTO)

	errors := c.BodyParser(blog)

	if utils.IsNotNil(errors) {
		return shared_services.HandlerResponseMapper(c, fiber.StatusInternalServerError, fiber.Map{
			"message": errors.Error(),
		})
	}

	res := blogService.UpdateBlog(blogId, *blog)

	return shared_services.HandlerResponseMapper(c, res.StatusCode, res)
}

func DeleteOneBlog(c *fiber.Ctx) error {
	blogId := c.Query("blogId")

	deleteBlog := blogService.DeleteBlog(blogId)

	fmt.Println(deleteBlog)

	return shared_services.HandlerResponseMapper(c, deleteBlog.StatusCode, deleteBlog)
}
