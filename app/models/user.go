package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primarykey;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string
	Password  string
	Notes     []Note
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()

	return
}

func (user User) ToUserRes() *UserRes {
	return &UserRes{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}

type UserRes struct {
	ID        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
}

type UserReq struct {
	Email    string `validate:"required,min=3,email"`
	Password string `validate:"required,min=3"`
}
