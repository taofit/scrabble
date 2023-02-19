package handlers

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
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
	player := move.Player

	if len(game.Bag) == 0 && (len(game.Deck1) == 0 || len(game.Deck2) == 0) {
		return c.Status(200).JSON("game over.")
	}
	if player != 1 && player != 2 {
		return c.Status(400).JSON("invalid player value, should be either 1 or 2")
	}
	if player == 1 && len(game.Bag) == 0 {
		remainLetterCount := len(game.Deck1) - len(move.Word.Word)
		if remainLetterCount == 1 {
			return c.Status(400).JSON("choose another word with different length, as the last word's length will be 1")
		}
	}
	if player == 2 && len(game.Bag) == 0 {
		remainLetterCount := len(game.Deck2) - len(move.Word.Word)
		if remainLetterCount == 1 {
			return c.Status(400).JSON("choose another word with different length, as the last word's length will be 1")
		}
	}
	valid, err := game.isValidWord(move.Word, player)
	if valid {
		inputWordPoints := buildWordPoints(move.Word)
		fmt.Println(inputWordPoints)
		if game.isBoardEmpty() || game.isWordInCorrectPosition(inputWordPoints) {
			for _, letterPoint := range inputWordPoints {
				game.mx.Lock()
				game.Board[letterPoint.Point.X][letterPoint.Point.Y] = letterPoint.Letter
				game.mx.Unlock()
			}
			game.calculateScore(inputWordPoints, player)
			game.buildDeck(move.Word.Word, player)

			if len(game.Bag) == 0 && (len(game.Deck1) == 0 || len(game.Deck2) == 0) {
				return c.Status(200).JSON("game over.")
			}

			return c.Status(200).JSON(move)
		}
		return c.Status(400).JSON("word is in wrong position, either overlap with current word or not adjacent to current word on board.")
	}

	return c.Status(400).JSON(err.Error())
}

func (game *Game) isBoardEmpty() bool {
	for _, row := range game.Board {
		for _, val := range row {
			if val != "" {
				return false
			}
		}
	}
	return true
}

func (game *Game) buildDeck(word string, player int) {
	if player == 1 {
		removeWordFromDeck(word, &game.Deck1)
		if len(game.Bag) > 0 {
			fillDeck(&game.Deck1, &game.Bag)
		}
	}
	if player == 2 {
		removeWordFromDeck(word, &game.Deck2)
		if len(game.Bag) > 0 {
			fillDeck(&game.Deck2, &game.Bag)
		}
	}
}

func removeWordFromDeck(word string, deck *[]string) {
	for _, inputLetter := range word {
		for j, deckLetter := range *deck {
			if string(inputLetter) == deckLetter {
				*deck = append((*deck)[:j], (*deck)[j+1:]...)
				break
			}
		}
	}
}

func fillDeck(deck *[]string, bag *[]string) {
	rand.Seed(time.Now().Unix())
	curDeckSize := len(*deck)
	for i := 0; i < deckSize-curDeckSize; i++ {
		randomIdx := rand.Intn(len(*bag))
		*deck = append(*deck, (*bag)[randomIdx])
		removeElementByIdx(bag, randomIdx)
	}
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

func (game *Game) isValidWord(word Word, player int) (bool, error) {
	wordLen := len(word.Word)
	if wordLen < 2 || wordLen > 7 {
		return false, errors.New("length of word must be between 2 and 7 inclusive")
	}
	inputWordArr := strings.Split(word.Word, "")
	if player == 1 && !areAllLettersInDeck(inputWordArr, game.Deck1) {
		return false, errors.New("not all letter of word are in deck, make sure to fetch the letter from deck")
	}

	if player == 2 && !areAllLettersInDeck(inputWordArr, game.Deck2) {
		return false, errors.New("not all letter of word are in deck, make sure to fetch the letter from deck")
	}

	pointStart := word.PointStart
	pointEnd := word.PointEnd

	if areAllLettersInBoard(pointStart, pointEnd) {
		if pointStart.X == pointEnd.X && abs(pointStart.Y, pointEnd.Y)+1 == wordLen {
			return true, nil
		}
		if pointStart.Y == pointEnd.Y && abs(pointStart.X, pointEnd.X)+1 == wordLen {
			return true, nil
		}
	}

	return false, errors.New("not all letters are in board or word not in one line")
}

func areAllLettersInDeck(word []string, deck []string) bool {
	set := make(map[string]int)
	for _, value := range deck {
		set[value] += 1
	}
	for _, value := range word {
		if count, ok := set[value]; !ok {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}
	return true
}

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
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
				letter := string(word.Word[i])
				inputWordPoints = append(inputWordPoints, LetterPoint{Letter: letter, Point: Point{x, y}})
				y += 1
			}
		}
		if pointEnd.Y < pointStart.Y {
			for i := 0; i < wordLen; i++ {
				letter := string(word.Word[i])
				inputWordPoints = append(inputWordPoints, LetterPoint{Letter: letter, Point: Point{x, y}})
				y -= 1
			}
		}
	}

	if pointStart.Y == pointEnd.Y {
		x := pointStart.X
		y := pointStart.Y
		if pointEnd.X > pointStart.X {
			for i := 0; i < wordLen; i++ {
				letter := string(word.Word[i])
				inputWordPoints = append(inputWordPoints, LetterPoint{Letter: letter, Point: Point{x, y}})
				x += 1
			}
		}
		if pointEnd.X < pointStart.X {
			for i := 0; i < wordLen; i++ {
				letter := string(word.Word[i])
				inputWordPoints = append(inputWordPoints, LetterPoint{Letter: letter, Point: Point{x, y}})
				x -= 1
			}
		}
	}
	return inputWordPoints
}

func (game *Game) isWordInCorrectPosition(inputWordPoints []LetterPoint) bool {
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

func (game *Game) Status(c *fiber.Ctx) error {
	return c.Status(200).JSON(game)
}
