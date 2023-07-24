package validators

import (
	"api.virak.me/modules/auth/domain/dto"
	"api.virak.me/shared/models"
	"api.virak.me/utils"
	"github.com/go-playground/validator/v10"
)

var registerValidate = validator.New()

func ValidateRegisterDTO(dto dto.RegisterDTO) []*models.ErrorResponse {
	var errors []*models.ErrorResponse

	err := registerValidate.Struct(dto)

	if utils.IsNotNil(err) {
		for _, e := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse

			element.FailedField = e.StructNamespace()
			element.Tag = e.Tag()
			element.Value = e.Param()

			errors = append(errors, &element)
		}
	}

	return errors
}
