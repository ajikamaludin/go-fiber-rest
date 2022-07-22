package apiv1

import (
	"github.com/ajikamaludin/go-fiber-rest/app/configs"
	authController "github.com/ajikamaludin/go-fiber-rest/app/controllers/auth"
	"github.com/ajikamaludin/go-fiber-rest/app/controllers/exception"
	noteController "github.com/ajikamaludin/go-fiber-rest/app/controllers/note"
	userController "github.com/ajikamaludin/go-fiber-rest/app/controllers/user"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func ApiRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Post("/auth/login", authController.Login)
	route.Post("/auth/register", authController.Register)

	route = route.Group("/", jwtware.New(jwtware.Config{
		SigningKey:   []byte(configs.GetInstance().Jwtconfig.Secret),
		ErrorHandler: exception.ExceptionNotFound,
	}))

	route.Get("/notes", noteController.GetAllNotes)
	route.Post("/notes", noteController.CreateNote)
	route.Get("/notes/:id", noteController.GetNoteById)
	route.Put("/notes/:id", noteController.UpdateNote)
	route.Delete("/notes/:id", noteController.DeleteNote)

	route.Post("/users", userController.CreateUser)
	route.Get("/users", userController.GetAllUsers)
}
