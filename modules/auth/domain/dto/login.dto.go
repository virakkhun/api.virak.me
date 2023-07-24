package dto

type LoginDTO struct {
	Email    string `validate:"required"`
	Password string `validate:"required,min=6"`
}
