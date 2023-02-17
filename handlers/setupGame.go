package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Letter struct {
	Points int `json:"points"`
	Tiles  int `json:"tiles"`
}

type LettersStruct struct {
	Letters map[string]Letter `json:"letters"`
}

type Game struct {
	Bag   []string
	Board [15][15]byte
	Deck1 [7]string
	Deck2 [7]string
}

func NewGame() *Game {
	bag := resetBag()
	board := resetBoard()

	return &Game{
		Bag:   bag,
		Board: board,
	}
}

func resetBoard() [15][15]byte {
	board := [15][15]byte{}
	return board
}

func resetBag() []string {
	var lettersMap = parseFile()
	var bag []string
	for key, letter := range lettersMap {
		for i := 0; i < letter.Tiles; i++ {
			bag = append(bag, key)
		}
	}

	return bag
}

func parseFile() map[string]Letter {
	content, err := ioutil.ReadFile("./letters.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var letters LettersStruct
	err = json.Unmarshal(content, &letters)
	if err != nil {
		log.Fatal("Error during Unmarshall():", err)
	}

	return letters.Letters
}
