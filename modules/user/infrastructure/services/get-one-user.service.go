package services

import (
	"api.virak.me/config/database"
	"api.virak.me/modules/user/domain/entities"
	"api.virak.me/shared/models"
	shared_services "api.virak.me/shared/services"
	"api.virak.me/utils"
	"github.com/gofiber/fiber/v2"
)

func GetOneUser(id string) models.BaseResponse {
	var user entities.UserEntity

	result := database.DB.First(&user, id)

	err := result.Error

	if utils.IsNotNil(err) {
		return shared_services.ServiceResponseMapper(nil, err.Error(), fiber.StatusNotFound)
	}

	userMap := utils.MergeMap(user, user.Model)

	delete(userMap, "Password")
	delete(userMap, "Model")

	return shared_services.ServiceResponseMapper(userMap, "success", fiber.StatusOK)
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
