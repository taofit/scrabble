package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/taofit/scrabble/database"
	"github.com/taofit/scrabble/models"
)

type Point struct {
	X int
	Y int
}
type Word struct {
	Word       string
	PointStart Point
	PointEnd   Point
}
type Move struct {
	Word   Word
	Player int
}

func (game *Game) Move(c *fiber.Ctx) error {
	move := new(Move)

	if err := c.BodyParser(move); err != nil {
		return err
	}
	// game.Bag = append(game.Bag, "4")
	fmt.Println("body: ", move)
	return c.Status(200).JSON(move)
}

func isValidWord(word Word) bool {
	wordLen := len(word.Word)
	if wordLen < 2 || wordLen > 7 {
		return false
	}
	for _, r := range word.Word {
		if (r < 'a' || r > 'z') {
			return false
		}
	}

	pointStart := word.PointStart
	pointEnd := word.PointEnd

	if pointStart.X == pointEnd.X && pointStart.Y - pointEnd.Y == wordLen {
		return true
	}
	if pointStart.Y == pointEnd.Y && pointStart.X - pointEnd.X == wordLen {
		return true
	}

	return false
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)
	return c.Status(200).JSON(fact)
}
