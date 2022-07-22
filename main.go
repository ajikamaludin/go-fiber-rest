package main

import (
	"fmt"
	"log"

	"github.com/ajikamaludin/go-fiber-rest/app/configs"
	apiRoute "github.com/ajikamaludin/go-fiber-rest/routers/api/v1"
	exceptionRoute "github.com/ajikamaludin/go-fiber-rest/routers/exception"
	homeRoute "github.com/ajikamaludin/go-fiber-rest/routers/home"
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

	// default here : /
	homeRoute.HomeRoutes(app)
	// api route : api/v1
	apiRoute.ApiRoutes(app)
	// handle 404
	exceptionRoute.Routes(app)

	config := configs.GetInstance().Appconfig
	listenPort := fmt.Sprintf(":%v", config.Port)
	log.Fatal(app.Listen(listenPort))
}
