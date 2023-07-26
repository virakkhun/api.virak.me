package services

import (
	"api.virak.me/src/modules/auth/domain/dto"
	"api.virak.me/src/modules/user/infrastructure/services"
	"api.virak.me/src/shared/models"
	sharedServices "api.virak.me/src/shared/services"
	"api.virak.me/src/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(dto dto.LoginDTO) models.BaseResponse {
	user, err := services.GetOneUserByEmail(dto.Email)

	if utils.IsNotNil(err) {
		return sharedServices.ServiceResponseMapper(nil, err.Error(), fiber.StatusForbidden)
	}

	compareHash := sharedServices.DeHashPassword(user.Password, dto.Password)

	if !compareHash {
		return sharedServices.ServiceResponseMapper(nil, "password is invalid", fiber.StatusForbidden)
	}

	token, _error := sharedServices.SignNewToken(models.Map{
		"ID": user.ID,
	})

	if utils.IsNotNil(_error) {
		return sharedServices.ServiceResponseMapper(nil, _error.Error(), fiber.StatusBadRequest)
	}

	return sharedServices.ServiceResponseMapper(token, "success login", fiber.StatusOK)
}
