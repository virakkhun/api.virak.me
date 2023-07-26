package sharedServices

import "api.virak.me/src/shared/models"

func ServiceResponseMapper(data models.AnyMap, message string, statusCode int) models.BaseResponse {
	return models.BaseResponse{
		Data:       data,
		Message:    message,
		StatusCode: statusCode,
	}
}