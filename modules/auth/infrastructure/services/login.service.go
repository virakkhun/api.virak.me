package services

import (
	"api.virak.me/modules/auth/domain/dto"
	"api.virak.me/modules/user/infrastructure/services"
	"api.virak.me/shared/models"
	shared_services "api.virak.me/shared/services"
	"api.virak.me/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(dto dto.LoginDTO) models.BaseResponse {
	user, err := services.GetOneUserByEmail(dto.Email)

	if utils.IsNotNil(err) {
		return shared_services.ServiceResponseMapper(nil, err.Error(), fiber.StatusForbidden)
	}

	compareHash := shared_services.DeHashPassword(user.Password, dto.Password)

	if !compareHash {
		return shared_services.ServiceResponseMapper(nil, "password is invalid", fiber.StatusForbidden)
	}

	token, _error := shared_services.SignNewToken(models.Map{
		"ID": user.ID,
	})

	if utils.IsNotNil(_error) {
		return shared_services.ServiceResponseMapper(nil, _error.Error(), fiber.StatusBadRequest)
	}

	return shared_services.ServiceResponseMapper(token, "success login", fiber.StatusOK)
}
