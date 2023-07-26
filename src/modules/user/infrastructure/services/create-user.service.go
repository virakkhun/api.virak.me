package services

import (
	"api.virak.me/src/config/database"
	"api.virak.me/src/modules/user/domain/dto"
	"api.virak.me/src/modules/user/domain/entities"
	"api.virak.me/src/shared/models"
	sharedServices "api.virak.me/src/shared/services"
	"api.virak.me/src/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(dto dto.CreateUserDTO) models.BaseResponse {
	_, err := GetOneUserByEmail(dto.Email)

	if !utils.IsNotNil(err) {
		return sharedServices.ServiceResponseMapper(nil, "email is already existed", fiber.StatusBadRequest)
	}

	result := database.DB.Create(&entities.UserEntity{
		Username: dto.Username,
		Email:    dto.Email,
		Password: sharedServices.HashPassword(dto.Password),
		Role:     dto.Role,
		Photo:    dto.Photo,
	})

	errResult := result.Error

	if utils.IsNotNil(errResult) {
		return sharedServices.ServiceResponseMapper(nil, errResult.Error(), fiber.StatusBadRequest)
	}

	return sharedServices.ServiceResponseMapper(nil, "user created", fiber.StatusOK)
}