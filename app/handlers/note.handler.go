package handlers

import (
	"github.com/ajikamaludin/go-fiber-rest/app/models"
	gormdb "github.com/ajikamaludin/go-fiber-rest/pkg/gorm.db"
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/constants"
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "Ok",
		"data":    notes,
	})
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
			"status":  constants.STATUS_FAIL,
			"message": "note not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  constants.STATUS_OK,
		"message": "note found",
		"data":    note,
	})
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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  constants.STATUS_OK,
		"message": "note created",
		"data":    note,
	})
}

func UpdateNote(c *fiber.Ctx) error {
	// validate request first
	noteRequest := new(models.Note)

	if err := c.BodyParser(&noteRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": err.Error(),
		})
	}

	errors := validator.ValidateRequest(noteRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "error",
			"errors":  errors,
		})
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
			"status":  constants.STATUS_FAIL,
			"message": "note not found",
		})
	}

	// Update
	db.Model(&note).Updates(noteRequest)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  constants.STATUS_OK,
		"message": "note updated",
		"data":    note,
	})
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
			"status":  constants.STATUS_FAIL,
			"message": "note not found",
		})
	}

	db.Delete(&note)

	return c.SendStatus(fiber.StatusNoContent)
}