package services

import (
	"api.virak.me/src/modules/auth/domain/dto"
	createUserDTO "api.virak.me/src/modules/user/domain/dto"
	"api.virak.me/src/modules/user/infrastructure/services"
	"api.virak.me/src/shared/models"
	sharedServices "api.virak.me/src/shared/services"
)

func Register(dto dto.RegisterDTO) models.BaseResponse {
	newUser := services.CreateUser(createUserDTO.CreateUserDTO{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
		Role:     "",
		Photo:    "",
	})

	if newUser.Data == nil {
		return sharedServices.ServiceResponseMapper(nil, newUser.Message.(string), newUser.StatusCode)
	}

	return sharedServices.ServiceResponseMapper(nil, "success register", newUser.StatusCode)
}
