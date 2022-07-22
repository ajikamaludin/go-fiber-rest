package apiv1

import (
	authController "github.com/ajikamaludin/go-fiber-rest/app/controllers/auth"
	noteController "github.com/ajikamaludin/go-fiber-rest/app/controllers/note"
	userController "github.com/ajikamaludin/go-fiber-rest/app/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/notes", noteController.GetAllNotes)
	route.Post("/notes", noteController.CreateNote)
	route.Get("/notes/:id", noteController.GetNoteById)
	route.Put("/notes/:id", noteController.UpdateNote)
	route.Delete("/notes/:id", noteController.DeleteNote)

	route.Post("/users", userController.CreateUser)
	route.Get("/users", userController.GetAllUsers)

	route.Post("/auth/login", authController.Login)
}
