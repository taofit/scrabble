package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/taofit/scrabble/internal"
)

func main() {
	app := fiber.New()
	var game = internal.NewGame()
	internal.SetupRoutes(app, game)
	log.Printf("game: %v", game)

	app.Listen(":3000")
}
