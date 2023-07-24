package controllers

import (
	"api.virak.me/modules/user/domain/dto"
	"api.virak.me/modules/user/domain/validators"
	"api.virak.me/modules/user/infrastructure/services"
	"api.virak.me/shared/models"
	shared_services "api.virak.me/shared/services"
	"api.virak.me/utils"
	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	userId := c.Query("userId")
	user := services.GetOneUser(userId)
	return shared_services.HandlerResponseMapper(c, user.StatusCode, user)
}

func Create(c *fiber.Ctx) error {
	user := new(dto.CreateUserDTO)

	err := c.BodyParser(user)

	if utils.IsNotNil(err) {
		return shared_services.HandlerResponseMapper(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	validatedErrors := validators.ValidateCreateUserDTO(*user)

	if validatedErrors != nil {
		return shared_services.HandlerResponseMapper(c, fiber.StatusInternalServerError, validatedErrors)
	}

	newUser := services.CreateUser(*user)

	return shared_services.HandlerResponseMapper(c, newUser.StatusCode, newUser)
}

func Update(c *fiber.Ctx) error {
	userId := c.Query(`userId`)
	updateUser := new(dto.UpdateUserDTO)

	err := c.BodyParser(updateUser)

	if utils.IsNotNil(err) {
		return shared_services.HandlerResponseMapper(c, fiber.StatusInternalServerError, models.Map{
			"message": err.Error(),
		})
	}

	result := services.UpdateUser(userId, *updateUser)

	return shared_services.HandlerResponseMapper(c, result.StatusCode, result)
}

func Delete(c *fiber.Ctx) error {
	result := services.DeleteUser(c.Query("userId"))

	return shared_services.HandlerResponseMapper(c, fiber.StatusOK, result)
}
