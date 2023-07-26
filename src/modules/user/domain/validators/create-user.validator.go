package validators

import (
	"api.virak.me/src/modules/user/domain/dto"
	"api.virak.me/src/shared/models"
	"api.virak.me/src/utils"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateCreateUserDTO(createUserDTO dto.CreateUserDTO) []*models.ErrorResponse {
	var errors []*models.ErrorResponse

	err := validate.Struct(createUserDTO)

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
