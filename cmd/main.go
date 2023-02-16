package main

import (
	"log"

	"gihub.com/taofit/scrabble/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	var Game = newGame()
	log.Printf("%v", Game)

	app.Listen(":3003")
}
