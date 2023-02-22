package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	deckSize  = 7
	boardSize = 15
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

type Letter struct {
	Points int `json:"points"`
	Tiles  int `json:"tiles"`
}

type LettersStruct struct {
	Letters map[string]Letter `json:"letters"`
}

type Game struct {
	mx     sync.Mutex
	Bag    []string
	Board  [boardSize][boardSize]string
	Deck1  []string
	Deck2  []string
	Score1 int
	Score2 int
}

func NewGame() *Game {
	bag := initBag()
	board := initBoard()
	deck1 := initDeck(&bag)
	deck2 := initDeck(&bag)

	return &Game{
		Bag:   bag,
		Board: board,
		Deck1: deck1,
		Deck2: deck2,
	}
}

func initBoard() [15][15]string {
	board := [15][15]string{}
	return board
}

func initBag() []string {
	var lettersMap = parseFile()
	var bag []string
	for key, letter := range lettersMap {
		if len(key) > 1 || key[0] < 'a' || key[0] > 'z' {
			log.Fatal("key must be one alphabetic in lower case")
		}
		for i := 0; i < letter.Tiles; i++ {
			bag = append(bag, key)
		}
	}

	return bag
}

func initDeck(bag *[]string) []string {
	var deck []string
	rand.Seed(time.Now().Unix())
	for i := 0; i < deckSize; i++ {
		randomIdx := rand.Intn(len(*bag))
		deck = append(deck, (*bag)[randomIdx])
		removeElementByIdx(bag, randomIdx)
	}
	return deck
}

func removeElementByIdx(bag *[]string, idx int) {
	bagLen := len(*bag)
	bagLastInx := bagLen - 1

	if idx != bagLastInx {
		(*bag)[idx] = (*bag)[bagLastInx]
	}
	*bag = (*bag)[:bagLastInx]
}

func parseFile() map[string]Letter {
	content, err := ioutil.ReadFile(basePath + "/../letters.json")
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
