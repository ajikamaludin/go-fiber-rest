package note

import (
	"github.com/ajikamaludin/go-fiber-rest/app/models"
	noteRepository "github.com/ajikamaludin/go-fiber-rest/app/repository/note"
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/constants"
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/converter"
	"github.com/ajikamaludin/go-fiber-rest/pkg/utils/validator"
	"github.com/gofiber/fiber/v2"
)

func GetAllNotes(c *fiber.Ctx) error {
	var notes []models.Note

	err := noteRepository.GetAllNotes(c, &notes)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "Internal Service Error",
			"error":   err.Error(),
		})
	}

	notesRes := converter.MapNoteToNoteRes(notes)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "Ok",
		"data":    notesRes,
	})
}

func GetNoteById(c *fiber.Ctx) error {
	id := c.Params("id")
	note := models.Note{}

	err := noteRepository.GetNoteById(id, &note)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "note not found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "note found",
		"data":    note.ToNoteRes(),
	})
}

func CreateNote(c *fiber.Ctx) error {
	noteRequest := new(models.NoteReq)

	c.BodyParser(&noteRequest)

	errors := validator.ValidateRequest(noteRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	note, _ := noteRepository.CreateNote(c, noteRequest)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "note created",
		"data":    note.ToNoteRes(),
	})
}

func UpdateNote(c *fiber.Ctx) error {
	noteRequest := new(models.NoteReq)

	c.BodyParser(&noteRequest)

	errors := validator.ValidateRequest(noteRequest)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "error",
			"errors":  errors,
		})
	}

	id := c.Params("id")
	note := models.Note{}
	err := noteRepository.GetNoteById(id, &note)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "note not found",
		})
	}

	noteRepository.UpdateNote(&note, noteRequest)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  constants.STATUS_SUCCESS,
		"message": "note updated",
		"data":    note.ToNoteRes(),
	})
}

func DeleteNote(c *fiber.Ctx) error {
	id := c.Params("id")

	note := models.Note{}
	err := noteRepository.GetNoteById(id, &note)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  constants.STATUS_FAIL,
			"message": "note not found",
		})
	}

	noteRepository.DeleteNote(&note)

	return c.SendStatus(fiber.StatusNoContent)
}
