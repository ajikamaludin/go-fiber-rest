package main

import (
	"log"

	"github.com/ajikamaludin/go-fiber-rest/app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	app.Use(recover.New())

	// route here
	app.Get("/", handlers.Home)

	// route Note
	app.Get("/notes", handlers.GetAllNotes)
	app.Post("/notes", handlers.CreateNote)
	app.Get("/notes/:id", handlers.GetNoteById)
	app.Put("/notes/:id", handlers.UpdateNote)
	app.Delete("/notes/:id", handlers.DeleteNote)

	log.Fatal(app.Listen(":3000"))
}
