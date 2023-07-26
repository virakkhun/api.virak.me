package dto

type RegisterDTO struct {
	Username string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required,min=6,max=12"`
}
