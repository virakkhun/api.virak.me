package services

import (
	"api.virak.me/src/config/database"
	"api.virak.me/src/modules/blog/domain/dto"
	"api.virak.me/src/modules/blog/domain/entities"
	"api.virak.me/src/shared/models"
	sharedServices "api.virak.me/src/shared/services"
	"api.virak.me/src/utils"
	"github.com/gofiber/fiber/v2"
)

func GetBlogs() models.BaseResponse {
	var blogs []entities.BlogEntity

	result := database.DB.Find(&blogs)

	if utils.IsNotNil(result.Error) {
		return sharedServices.ServiceResponseMapper(nil, result.Error.Error(), fiber.StatusBadRequest)
	}

	return sharedServices.ServiceResponseMapper(blogs, "get success", fiber.StatusOK)
}

func GetOneBlog(id string) models.BaseResponse {
	var blog entities.BlogEntity

	result := database.DB.First(&blog, id)

	if utils.IsNotNil(result.Error) {
		return sharedServices.ServiceResponseMapper(nil, result.Error.Error(), fiber.StatusNotFound)
	}

	return sharedServices.ServiceResponseMapper(blog, "get success", fiber.StatusOK)
}

func CreateOneBlog(blog entities.BlogEntity) models.BaseResponse {
	result := database.DB.Create(&blog)

	if utils.IsNotNil(result.Error) {
		return sharedServices.ServiceResponseMapper(nil, result.Error.Error(), fiber.StatusBadRequest)
	}

	return sharedServices.ServiceResponseMapper(blog, "created success", fiber.StatusOK)
}

func UpdateBlog(id string, blog dto.UpdateBlogDTO) models.BaseResponse {
	var _blog entities.BlogEntity

	database.DB.First(&_blog, id)

	result := database.DB.Model(&_blog).Updates(entities.BlogEntity{
		Title:       blog.Title,
		Description: blog.Description,
		Content:     blog.Content,
		Viewer:      int16(blog.Viewer),
	})

	if utils.IsNotNil(result.Error) {
		return sharedServices.ServiceResponseMapper(nil, "failed to updated", fiber.StatusBadRequest)
	}

	return sharedServices.ServiceResponseMapper(_blog, "updated success", fiber.StatusOK)
}

func DeleteBlog(id string) models.BaseResponse {
	var blog entities.BlogEntity

	findBlog := database.DB.First(&blog, id)

	if findBlog.Error.Error() != "" {
		return sharedServices.ServiceResponseMapper(nil, findBlog.Error.Error(), fiber.StatusNotFound)
	}

	result := database.DB.Delete(&blog)

	if result.Error != nil {
		return sharedServices.ServiceResponseMapper(nil, result.Error.Error(), fiber.StatusNotFound)
	}

	return sharedServices.ServiceResponseMapper(blog, "deleted successfully", fiber.StatusOK)
}
