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

type LetterPoint struct {
	Letter string
	Point  Point
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

	fmt.Println("body: ", move)
	if isValidWord(move.Word) {
		inputWordPoints := buildWordPoints(move.Word)
		if game.isWordInCorrectPosition(inputWordPoints) {
			for _, letterPoint := range inputWordPoints {
				game.Board[letterPoint.Point.X][letterPoint.Point.Y] = letterPoint.Letter
			}
			return c.Status(200).JSON(move)
		}
	}

	return c.Status(400).JSON("invalid word format")
}

func (game *Game) calculateScore(inputWordPoints []LetterPoint, player int) {
	letterMap := parseFile()
	if player == 1 {
		for _, letterPoint := range inputWordPoints {
			game.Score1 += letterMap[letterPoint.Letter].Points
		}
	}
	if player == 2 {
		for _, letterPoint := range inputWordPoints {
			game.Score2 += letterMap[letterPoint.Letter].Points
		}
	}
}

func isValidWord(word Word) bool {
	wordLen := len(word.Word)
	if wordLen < 2 || wordLen > 7 {
		return false
	}
	for _, r := range word.Word {
		if r < 'a' || r > 'z' {
			return false
		}
	}

	pointStart := word.PointStart
	pointEnd := word.PointEnd

	if areAllLettersInBoard(pointStart, pointEnd) {
		if pointStart.X == pointEnd.X && pointStart.Y-pointEnd.Y+1 == wordLen {
			return true
		}
		if pointStart.Y == pointEnd.Y && pointStart.X-pointEnd.X+1 == wordLen {
			return true
		}
	}

	return false
}

func areAllLettersInBoard(pointStart Point, pointEnd Point) bool {
	if pointStart.X < 0 || pointStart.X >= boardSize ||
		pointStart.Y < 0 || pointStart.Y >= boardSize ||
		pointEnd.X < 0 || pointEnd.X >= boardSize ||
		pointEnd.Y < 0 || pointEnd.Y >= boardSize {
		return false
	}
	return true
}

func buildWordPoints(word Word) []LetterPoint {
	var inputWordPoints []LetterPoint
	wordLen := len(word.Word)
	pointStart := word.PointStart
	pointEnd := word.PointEnd
	if pointStart.X == pointEnd.X {
		x := pointStart.X
		y := pointStart.Y
		if pointEnd.Y > pointStart.Y {
			for i := 0; i < wordLen; i++ {
				y += i
				letter := string(word.Word[i])
				inputWordPoints = append(inputWordPoints, LetterPoint{Letter: letter, Point: Point{x, y}})
			}
		}
		if pointEnd.Y < pointStart.Y {
			for i := 0; i < wordLen; i++ {
				y -= i
				letter := string(word.Word[i])
				inputWordPoints = append(inputWordPoints, LetterPoint{Letter: letter, Point: Point{x, y}})
			}
		}
	}

	if pointStart.Y == pointEnd.Y {
		x := pointStart.X
		y := pointStart.Y
		if pointEnd.X > pointStart.X {
			for i := 0; i < wordLen; i++ {
				x += i
				letter := string(word.Word[i])
				inputWordPoints = append(inputWordPoints, LetterPoint{Letter: letter, Point: Point{x, y}})
			}
		}
		if pointEnd.X < pointStart.X {
			for i := 0; i < wordLen; i++ {
				x -= i
				letter := string(word.Word[i])
				inputWordPoints = append(inputWordPoints, LetterPoint{Letter: letter, Point: Point{x, y}})
			}
		}
	}
	return inputWordPoints
}

func (game Game) isWordInCorrectPosition(inputWordPoints []LetterPoint) bool {
	for _, letterPoint := range inputWordPoints {
		if game.Board[letterPoint.Point.X][letterPoint.Point.Y] != "" {
			return false
		}
	}

	for _, letterPoint := range inputWordPoints {
		neighbourLeftPointX := letterPoint.Point.X - 1
		neighbourLeftPointY := letterPoint.Point.Y

		neighbourUpperPointX := letterPoint.Point.X
		neighbourUpperPointY := letterPoint.Point.Y + 1

		neighbourRightPointX := letterPoint.Point.X + 1
		neighbourRightPointY := letterPoint.Point.Y

		neighbourLowerPointX := letterPoint.Point.X
		neighbourLowerPointY := letterPoint.Point.Y - 1

		if game.isNeighbourPlacedOnBoard(neighbourLeftPointX, neighbourLeftPointY) {
			return true
		}
		if game.isNeighbourPlacedOnBoard(neighbourUpperPointX, neighbourUpperPointY) {
			return true
		}
		if game.isNeighbourPlacedOnBoard(neighbourRightPointX, neighbourRightPointY) {
			return true
		}
		if game.isNeighbourPlacedOnBoard(neighbourLowerPointX, neighbourLowerPointY) {
			return true
		}
	}
	return false
}

func (game *Game) isNeighbourPlacedOnBoard(neighbourX int, neighbourY int) bool {
	if neighbourX < 0 || neighbourX >= boardSize ||
		neighbourY < 0 || neighbourY >= boardSize {
		return false
	}
	return game.Board[neighbourX][neighbourY] != ""
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