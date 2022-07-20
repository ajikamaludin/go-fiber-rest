package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	User_id string
	Title   string `validate:"required,min=3"`
	Note    string `validate:"required,min=3"`
	Tag     string
}

func (note Note) MarshalBinary() ([]byte, error) {
	return []byte(fmt.Sprintf("%v-%v", note.Title, note.Note)), nil
}
