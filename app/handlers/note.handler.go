package handlers

import (
	"github.com/ajikamaludin/go-fiber-rest/app/models"
	gormdb "github.com/ajikamaludin/go-fiber-rest/pkg/gorm.db"
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/validator"
	"github.com/gofiber/fiber/v2"
)

func GetAllNotes(c *fiber.Ctx) error {
	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	var notes []models.Note

	db.Find(&notes)

	return c.Status(fiber.StatusOK).JSON(notes)
}

func GetNoteById(c *fiber.Ctx) error {
	id := c.Params("id")

	db, err := gormdb.GetInstance()

	if err != nil {
		return err
	}

	note := models.Note{}
	err = db.First(&note, id).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "note not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(note)
}

func CreateNote(c *fiber.Ctx) error {
	noteRequest := new(models.Note)

	if err := c.BodyParser(&noteRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := validator.ValidateRequest(noteRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	var note = models.Note{
		User_id: "1",
		Title:   noteRequest.Note,
		Note:    noteRequest.Title,
	}

	db.Create(&note)

	return c.Status(fiber.StatusCreated).JSON(note)
}

func UpdateNote(c *fiber.Ctx) error {
	// validate request first
	noteRequest := new(models.Note)

	if err := c.BodyParser(&noteRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := validator.ValidateRequest(noteRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	// find records
	id := c.Params("id")

	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	note := models.Note{}
	err = db.First(&note, id).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "note not found",
		})
	}

	// Update
	db.Model(&note).Updates(noteRequest)

	return c.Status(fiber.StatusCreated).JSON(note)
}

func DeleteNote(c *fiber.Ctx) error {
	// find records
	id := c.Params("id")

	db, err := gormdb.GetInstance()
	if err != nil {
		return err
	}

	note := models.Note{}
	err = db.First(&note, id).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "note not found",
		})
	}

	db.Delete(&note)

	return c.SendStatus(fiber.StatusNoContent)
}
