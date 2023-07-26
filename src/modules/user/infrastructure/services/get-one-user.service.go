package services

import (
	"api.virak.me/src/config/database"
	"api.virak.me/src/modules/user/domain/entities"
	"api.virak.me/src/shared/models"
	sharedServices "api.virak.me/src/shared/services"
	"api.virak.me/src/utils"
	"github.com/gofiber/fiber/v2"
)

func GetOneUser(id string) models.BaseResponse {
	var user entities.UserEntity

	result := database.DB.First(&user, id)

	err := result.Error

	if utils.IsNotNil(err) {
		return sharedServices.ServiceResponseMapper(nil, err.Error(), fiber.StatusNotFound)
	}

	userMap, _ := utils.StructToMap(user)
	modelMap, _ := utils.StructToMap(user.Model)
	mergedMap := utils.MergeMap(userMap, modelMap)
	utils.DeleteMultiKeys(&mergedMap, []string{"Password", "Model"})

	return sharedServices.ServiceResponseMapper(mergedMap, "success", fiber.StatusOK)
}

func GetOneUserByEmail(email string) (entities.UserEntity, error) {
	var user entities.UserEntity

	result := database.DB.Where("Email = ?", email).First(&user)

	err := result.Error

	if utils.IsNotNil(err) {
		return user, err
	}

	return user, nil
}

func GetOneUserByID(id string) (entities.UserEntity, error) {
	var user entities.UserEntity

	result := database.DB.Where("ID = ?", id).First(&user)

	err := result.Error

	if utils.IsNotNil(err) {
		return user, err
	}

	return user, nil
}
