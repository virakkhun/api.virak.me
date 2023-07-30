package controllers

import (
	"api.virak.me/src/modules/blog/domain/dto"
	"api.virak.me/src/modules/blog/domain/entities"
	blogService "api.virak.me/src/modules/blog/infrastructure/services"
	sharedServices "api.virak.me/src/shared/services"
	sharedValidators "api.virak.me/src/shared/validators"
	"api.virak.me/src/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllBlog(c *fiber.Ctx) error {
	allBlogs := blogService.GetBlogs()

	return sharedServices.HandlerResponseMapper(c, allBlogs.StatusCode, allBlogs)
}

func GetOneBlog(c *fiber.Ctx) error {
	oneBlog := blogService.GetOneBlog(c.Query("blogId"))
	return sharedServices.HandlerResponseMapper(c, oneBlog.StatusCode, oneBlog)
}

func CreateOneBlog(c *fiber.Ctx) error {
	blog := new(dto.BlogDTO)

	err := c.BodyParser(blog)

	if utils.IsNotNil(err) {
		return sharedServices.HandlerResponseMapper(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	validatedErrors := sharedValidators.Validate[dto.BlogDTO](*blog)

	if validatedErrors != nil {
		return sharedServices.HandlerResponseMapper(c, fiber.StatusInternalServerError, validatedErrors)
	}

	newBlog := blogService.CreateOneBlog(entities.BlogEntity{
		Title:       blog.Title,
		Description: blog.Description,
		Content:     blog.Content,
		Viewer:      int16(blog.Viewer),
	})

	return sharedServices.HandlerResponseMapper(c, newBlog.StatusCode, newBlog)
}

func UpdateOneBlog(c *fiber.Ctx) error {
	blogId := c.Query("blogId")

	blog := new(dto.UpdateBlogDTO)

	errors := c.BodyParser(blog)

	if utils.IsNotNil(errors) {
		return sharedServices.HandlerResponseMapper(c, fiber.StatusInternalServerError, fiber.Map{
			"message": errors.Error(),
		})
	}

	res := blogService.UpdateBlog(blogId, *blog)

	return sharedServices.HandlerResponseMapper(c, res.StatusCode, res)
}

func DeleteOneBlog(c *fiber.Ctx) error {
	blogId := c.Query("blogId")

	deleteBlog := blogService.DeleteBlog(blogId)

	return sharedServices.HandlerResponseMapper(c, deleteBlog.StatusCode, deleteBlog)
}
