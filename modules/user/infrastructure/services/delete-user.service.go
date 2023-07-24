package services

import (
	"api.virak.me/config/database"
	"api.virak.me/modules/user/domain/entities"
	"api.virak.me/shared/models"
	shared_services "api.virak.me/shared/services"
	"api.virak.me/utils"
	"github.com/gofiber/fiber/v2"
)

func DeleteUser(id string) models.BaseResponse {
	deleteResult := database.DB.Delete(&entities.UserEntity{}, id)

	if utils.IsNotNil(deleteResult.Error) {
		return shared_services.ServiceResponseMapper(
			nil,
			deleteResult.Error.Error(),
			fiber.StatusBadRequest,
		)
	}

	return shared_services.ServiceResponseMapper(
		nil,
		"deleted success",
		fiber.StatusOK,
	)
}
