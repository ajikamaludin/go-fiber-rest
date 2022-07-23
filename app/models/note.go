package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserId uuid.UUID `gorm:"type:uuid"`
	Title  string
	Note   string
	Tag    string
	User   User `json:",omitempty"`
}

func (n Note) ToNoteRes() *NoteRes {
	return &NoteRes{
		UserId: n.UserId,
		ID:     n.ID,
		Title:  n.Title,
		Note:   n.Note,
	}
}

func (n Note) ToNoteWithUserRes() *NoteWithUserRes {
	return &NoteWithUserRes{
		UserId: n.UserId,
		ID:     n.ID,
		Title:  n.Title,
		Note:   n.Note,
		User:   n.User,
	}
}

type NoteRes struct {
	UserId uuid.UUID
	ID     uint
	Title  string
	Note   string
}

type NoteWithUserRes struct {
	UserId uuid.UUID
	ID     uint
	Title  string
	Note   string
	User   User
}

type NoteReq struct {
	Title string `validate:"required,min=3"`
	Note  string `validate:"required,min=3"`
}
