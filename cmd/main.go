package main

import (
	// "log"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/taofit/scrabble/database"
	"github.com/taofit/scrabble/handlers"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	var game = handlers.NewGame()
	setupRoutes(app, game)
	log.Printf("game: %v", game)

	app.Listen(":3000")
}
