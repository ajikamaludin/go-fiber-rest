package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	User_id string
	Title   string `validate:"required,min=3"`
	Note    string `validate:"required,min=3"`
	Tag     string
}
