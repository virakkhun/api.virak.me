package services

import (
	"api.virak.me/src/config/database"
	"api.virak.me/src/modules/user/domain/entities"
	"api.virak.me/src/shared/models"
	sharedServices "api.virak.me/src/shared/services"
	"api.virak.me/src/utils"
	"github.com/gofiber/fiber/v2"
)

func DeleteUser(id string) models.BaseResponse {
	deleteResult := database.DB.Delete(&entities.UserEntity{}, id)

	if utils.IsNotNil(deleteResult.Error) {
		return sharedServices.ServiceResponseMapper(
			nil,
			deleteResult.Error.Error(),
			fiber.StatusBadRequest,
		)
	}

	return sharedServices.ServiceResponseMapper(
		nil,
		"deleted success",
		fiber.StatusOK,
	)
}
