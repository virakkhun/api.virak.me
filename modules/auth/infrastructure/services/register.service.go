package services

import (
	"api.virak.me/modules/auth/domain/dto"
	createUserDTO "api.virak.me/modules/user/domain/dto"
	"api.virak.me/modules/user/infrastructure/services"
	"api.virak.me/shared/models"
	sharedServices "api.virak.me/shared/services"
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
