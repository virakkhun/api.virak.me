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

func UpdateUser(id string, dto dto.UpdateUserDTO) models.BaseResponse {
	user, err := GetOneUserByID(id)

	if utils.IsNotNil(err) {
		return shared_services.ServiceResponseMapper(user, err.Error(), fiber.StatusNotFound)
	}

	resultUpdated := database.DB.Model(&user).Updates(entities.UserEntity{
		Email:    dto.Email,
		Password: shared_services.HashPassword(dto.Password),
		Username: dto.Username,
		Role:     dto.Role,
		Photo:    dto.Photo,
	})

	updateResultError := resultUpdated.Error

	if utils.IsNotNil(updateResultError) {
		return shared_services.ServiceResponseMapper(user, updateResultError.Error(), fiber.StatusBadRequest)
	}

	userMap := utils.MergeMap(user, user.Model)

	delete(userMap, "Password")
	delete(userMap, "Model")

	return shared_services.ServiceResponseMapper(userMap, "success updated", fiber.StatusOK)
}
