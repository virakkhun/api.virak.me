package controllers

import (
	"time"

	"api.virak.me/modules/auth/domain/dto"
	"api.virak.me/modules/auth/domain/validators"
	"api.virak.me/modules/auth/infrastructure/services"
	shared_services "api.virak.me/shared/services"
	"api.virak.me/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	loginDto := new(dto.LoginDTO)

	err := c.BodyParser(loginDto)

	if utils.IsNotNil(err) {
		return shared_services.HandlerResponseMapper(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	validatedErrors := validators.ValidateLoginDTO(*loginDto)

	if validatedErrors != nil {
		return shared_services.HandlerResponseMapper(c, fiber.StatusInternalServerError, validatedErrors)
	}

	loginResult := services.Login(*loginDto)

	if loginResult.Data == nil {
		return shared_services.HandlerResponseMapper(c, loginResult.StatusCode, loginResult)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    loginResult.Data.(string),
		Path:     "/",
		Domain:   "",
		MaxAge:   time.Now().Hour() * 24 * 7,
		Secure:   true,
		HTTPOnly: true,
	})

	return nil
}

func Register(c *fiber.Ctx) error {
	registerDto := new(dto.RegisterDTO)

	err := c.BodyParser(registerDto)

	if utils.IsNotNil(err) {
		return shared_services.HandlerResponseMapper(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	validatedErrors := validators.ValidateRegisterDTO(*registerDto)

	if validatedErrors != nil {
		return shared_services.HandlerResponseMapper(c, fiber.StatusInternalServerError, validatedErrors)
	}

	registerResult := services.Register(*registerDto)

	return shared_services.HandlerResponseMapper(c, registerResult.StatusCode, registerResult)
}
