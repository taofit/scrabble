package internal

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, game *Game) {
	app.Post("/move", game.Move)
	app.Get("/Status", game.Status)
}
