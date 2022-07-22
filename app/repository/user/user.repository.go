package user

import (
	"github.com/ajikamaludin/go-fiber-rest/app/models"
	gormdb "github.com/ajikamaludin/go-fiber-rest/pkg/gorm.db"
)

func GetUserByEmail(email string, user *models.User) (err error) {
	db, _ := gormdb.GetInstance()
	err = db.Where("email = ?", email).First(&user).Error
	return
}

func CreateUser(user *models.User) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	err = db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
