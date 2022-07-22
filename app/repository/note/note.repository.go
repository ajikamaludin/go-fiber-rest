package note

import (
	"time"

	"github.com/ajikamaludin/go-fiber-rest/app/models"
	gormdb "github.com/ajikamaludin/go-fiber-rest/pkg/gorm.db"
	redisclient "github.com/ajikamaludin/go-fiber-rest/pkg/redis.client"
)

func GetAllNotes(notes *[]models.Note) error {
	err := redisclient.Get("allnotes", &notes)
	if err != nil {
		db, err := gormdb.GetInstance()
		if err != nil {
			return err
		}
		db.Find(&notes)

		redisclient.Set("allnotes", &notes, 30*time.Second)
	}
	return nil
}

func GetNoteById(id string, note *models.Note) error {
	key := "note+" + id
	err := redisclient.Get(key, &note)
	if err != nil {
		db, err := gormdb.GetInstance()

		if err != nil {
			return err
		}

		err = db.First(&note, id).Error

		if err != nil {
			return err
		}

		redisclient.Set(key, &note, 30*time.Second)
	}
	return nil
}
