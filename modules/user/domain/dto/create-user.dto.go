package dto

type CreateUserDTO struct {
	Username string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required,min=6,max=12"`
	Role     string `validate:"required"`
	Photo    string `validate:"required"`
}
