package dto

type BlogDTO struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Content     string `validate:"required"`
	Viewer      int    `validate:"required"`
}
