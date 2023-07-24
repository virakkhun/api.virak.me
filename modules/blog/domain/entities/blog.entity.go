package entities

import "gorm.io/gorm"

type BlogEntity struct {
	gorm.Model
	Title       string
	Description string
	Content     string
	Viewer      int16
}
