package apiv1

import (
	noteController "github.com/ajikamaludin/go-fiber-rest/app/controllers/note"
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/notes", noteController.GetAllNotes)
	route.Post("/notes", noteController.CreateNote)
	route.Get("/notes/:id", noteController.GetNoteById)
	route.Put("/notes/:id", noteController.UpdateNote)
	route.Delete("/notes/:id", noteController.DeleteNote)
}
