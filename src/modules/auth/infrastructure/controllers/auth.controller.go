package controllers

import (
	"time"

	"api.virak.me/src/modules/auth/domain/dto"
	"api.virak.me/src/modules/auth/domain/validators"
	"api.virak.me/src/modules/auth/infrastructure/services"
	sharedServices "api.virak.me/src/shared/services"
	"api.virak.me/src/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	loginDto := new(dto.LoginDTO)

	err := c.BodyParser(loginDto)

	if utils.IsNotNil(err) {
		return sharedServices.HandlerResponseMapper(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	validatedErrors := validators.ValidateLoginDTO(*loginDto)

	if validatedErrors != nil {
		return sharedServices.HandlerResponseMapper(c, fiber.StatusInternalServerError, validatedErrors)
	}

	loginResult := services.Login(*loginDto)

	if loginResult.Data == nil {
		return sharedServices.HandlerResponseMapper(c, loginResult.StatusCode, loginResult)
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
		return sharedServices.HandlerResponseMapper(c, fiber.StatusInternalServerError, fiber.Map{
			"message": err.Error(),
		})
	}

	validatedErrors := validators.ValidateRegisterDTO(*registerDto)

	if validatedErrors != nil {
		return sharedServices.HandlerResponseMapper(c, fiber.StatusInternalServerError, validatedErrors)
	}

	registerResult := services.Register(*registerDto)

	return sharedServices.HandlerResponseMapper(c, registerResult.StatusCode, registerResult)
}
