package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserId uuid.UUID
	Title  string `validate:"required,min=3"`
	Note   string `validate:"required,min=3"`
	Tag    string
}
