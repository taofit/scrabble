# scrabble

Scrabble game implementation in golang

## run program

Under root, enter command: `docker compose up` to start the game.
There are two end point:
The first is to place a word in POST request: `http://localhost:3000/move`
in postman or insomnia, choose JSON format and past json string as follows to make a move for a player

```
{
"Word": {
    "Word": "aim",
    "PointStart": {
        "x": 3,
        "y": 7
    },
    "PointEnd": {
        "x": 3,
        "y": 9
    }
},
"Player": 1
}
```

The word entered is `aim`, the start point's coordinate is {"x": 3, "y": 7}, the end point's coordinate is {"x": 3, "y": 9}, so user can eliminate entering all the points between the two points in the word.

The second end point is to fetch current game status, it is a GET request: `http://localhost:3000/status`

the result returned from the end point is as follows:

`{
	"Bag": [
		"y",
		"i",
		"i",
		"c",
		"d",
		"d",
		"d",
		"d",
		"g",
		"g",
		"g",
		"l",
		"i",
		"l",
		"l",
		"n",
		"n",
		"n",
		"n",
		"n",
		"n",
		"s",
		"s",
		"s",
		"i",
		"t",
		"x",
		"t",
		"t",
		"t",
		"t",
		"w",
		"e",
		"b",
		"b",
		"i",
		"f",
		"j",
		"k",
		"p",
		"p",
		"i",
		"o",
		"e",
		"o",
		"o",
		"o",
		"m",
		"o",
		"q",
		"e",
		"r",
		"h",
		"r",
		"e",
		"r",
		"r",
		"i",
		"u",
		"u",
		"e",
		"v",
		"v",
		"a",
		"a",
		"e",
		"a",
		"e",
		"m",
		"a",
		"a",
		"a",
		"e",
		"i",
		"e"
	],
	"Board": [
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"t",
			"a",
			"o",
			"u",
			"a",
			"e",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"o",
			"s",
			"e",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		],
		[
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			""
		]
	],
	"Deck1": [
		"c",
		"r",
		"u",
		"i",
		"z",
		"r",
		"e"
	],
	"Deck2": [
		"l",
		"y",
		"f",
		"a",
		"o",
		"w",
		"h"
	],
	"Score1": 6,
	"Score2": 3
}`

`Bag` contains the tiles left in the bag,
`Board` displays the tiles placed on the board.
`Deck1` shows player 1's tiles in his/her deck
`Deck2` shows player 2's tiles in his/her deck
`Score1` current score of player 1
`Score2` current score of player 2

## program description

The game status is saved in a `Game` struct, each time a player successfully places a word on the board, the Game struct will be updated, and its status can be fetched via the Get `status` request. Therefore no database is setup to save the game data.

## testing

run testing under internal folder by entering : `go test -v ./...`
