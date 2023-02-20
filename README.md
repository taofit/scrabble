# scrabble
scrabble in golang

## run program

Under root, enter command: `docker compose up`
in postman or insomnia enter url to make a post request: `http://localhost:3000/move`
choose JSON format and past json string as follows to make a move for a player

`{
	"Word": 
		{
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
}`

The word entered is `aim`, the start point's coordinate is {"x": 3, "y": 7} the end point's coordinate is {"x": 3, "y": 9}, so user can eliminate entering all the points in the middle of the word.