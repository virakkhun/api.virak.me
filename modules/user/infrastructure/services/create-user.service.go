package services

import (
	"api.virak.me/config/database"
	"api.virak.me/modules/user/domain/dto"
	"api.virak.me/modules/user/domain/entities"
	"api.virak.me/shared/models"
	shared_services "api.virak.me/shared/services"
	"api.virak.me/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(dto dto.CreateUserDTO) models.BaseResponse {
	_, err := GetOneUserByEmail(dto.Email)

	if !utils.IsNotNil(err) {
		return shared_services.ServiceResponseMapper(nil, "email is already existed", fiber.StatusBadRequest)
	}

	result := database.DB.Create(&entities.UserEntity{
		Username: dto.Username,
		Email:    dto.Email,
		Password: shared_services.HashPassword(dto.Password),
		Role:     dto.Role,
		Photo:    dto.Photo,
	})

	errResult := result.Error

	if utils.IsNotNil(errResult) {
		return shared_services.ServiceResponseMapper(nil, errResult.Error(), fiber.StatusBadRequest)
	}

	return shared_services.ServiceResponseMapper(nil, "user created", fiber.StatusOK)
}
