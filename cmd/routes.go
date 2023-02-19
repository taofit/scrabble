package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/taofit/scrabble/handlers"
)

func setupRoutes(app *fiber.App, game *handlers.Game) {
	app.Post("/move", game.Move)
	app.Get("/Status", game.Status)
}
