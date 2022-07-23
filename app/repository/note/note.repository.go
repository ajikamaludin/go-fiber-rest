package note

import (
	"strconv"
	"time"

	"github.com/ajikamaludin/go-fiber-rest/app/models"
	gormdb "github.com/ajikamaludin/go-fiber-rest/pkg/gorm.db"
	"github.com/ajikamaludin/go-fiber-rest/pkg/jwtmanager"
	redisclient "github.com/ajikamaludin/go-fiber-rest/pkg/redis.client"
	"github.com/gofiber/fiber/v2"
)

func GetAllNotes(c *fiber.Ctx, notes *[]models.Note) error {
	err := redisclient.Get("allnotes", &notes)
	if err != nil {
		db, err := gormdb.GetInstance()
		if err != nil {
			return err
		}

		userId := jwtmanager.GetUserId(c)
		db.Where("user_id = ?", userId).Find(&notes)

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

func CreateNote(c *fiber.Ctx, noteRequest *models.NoteReq) (*models.Note, error) {
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}

	userId := jwtmanager.GetUserId(c)

	note := &models.Note{
		UserId: userId,
		Title:  noteRequest.Note,
		Note:   noteRequest.Title,
	}

	db.Create(&note)
	return note, nil
}

func UpdateNote(note *models.Note, noteRequest *models.NoteReq) (*models.NoteRes, error) {
	db, err := gormdb.GetInstance()
	if err != nil {
		return nil, err
	}

	db.Model(&note).Updates(&models.Note{
		Title: noteRequest.Title,
		Note:  noteRequest.Note,
	})
	key := "note+" + strconv.Itoa(int(note.ID))
	redisclient.Remove(key)

	return &models.NoteRes{
		UserId: note.UserId,
		Title:  note.Title,
		Note:   note.Note,
	}, err
}

func DeleteNote(note *models.Note) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	db.Delete(&note)
	key := "note+" + strconv.Itoa(int(note.ID))
	redisclient.Remove(key)

	return nil
}
