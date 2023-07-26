package services

import (
	"api.virak.me/config/database"
	"api.virak.me/modules/user/domain/dto"
	"api.virak.me/modules/user/domain/entities"
	"api.virak.me/shared/models"
	sharedServices "api.virak.me/shared/services"
	"api.virak.me/utils"
	"github.com/gofiber/fiber/v2"
)

func UpdateUser(id string, dto dto.UpdateUserDTO) models.BaseResponse {
	user, err := GetOneUserByID(id)

	if utils.IsNotNil(err) {
		return sharedServices.ServiceResponseMapper(user, err.Error(), fiber.StatusNotFound)
	}

	resultUpdated := database.DB.Model(&user).Updates(entities.UserEntity{
		Email:    dto.Email,
		Password: sharedServices.HashPassword(dto.Password),
		Username: dto.Username,
		Role:     dto.Role,
		Photo:    dto.Photo,
	})

	updateResultError := resultUpdated.Error

	if utils.IsNotNil(updateResultError) {
		return sharedServices.ServiceResponseMapper(user, updateResultError.Error(), fiber.StatusBadRequest)
	}

	userMap, _ := utils.StructToMap(user)
	modelMap, _ := utils.StructToMap(user.Model)
	mergedMap := utils.MergeMap(userMap, modelMap)

	delete(mergedMap, "Password")
	delete(mergedMap, "Model")

	return sharedServices.ServiceResponseMapper(mergedMap, "success updated", fiber.StatusOK)
}
