// TODO stub requests
package games

import (
	"testing"

	"github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

var tests = []struct {
	input  string
	output []Games
}{
	{
		input: "learnercys",
		output: []Games{
			{
				Player: "learnercys",
				Year:   "2015",
				Month:  "12",
			},
			{
				Player: "learnercys",
				Year:   "2018",
				Month:  "01",
			},
			{
				Player: "learnercys",
				Year:   "2018",
				Month:  "02",
			},
			{
				Player: "learnercys",
				Year:   "2018",
				Month:  "03",
			},
			{
				Player: "learnercys",
				Year:   "2018",
				Month:  "04",
			},
		},
	},
}

func TestAvailableArchivesOK(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.chess.com/pub/player/learnercys/games/archives",
		httpmock.NewStringResponder(200, `{
			"archives": [
				"https://api.chess.com/pub/player/learnercys/games/2015/12",
				"https://api.chess.com/pub/player/learnercys/games/2018/01",
				"https://api.chess.com/pub/player/learnercys/games/2018/02",
				"https://api.chess.com/pub/player/learnercys/games/2018/03",
				"https://api.chess.com/pub/player/learnercys/games/2018/04"
			]
		}`))

	for _, test := range tests {
		games, err := AvailableArchives(test.input)

		if err != nil {
			t.Fatal("couldn't get archived games")
		}

		assert.NotEmpty(t, games)

		// TODO This will fail in some cases, to me is good enough
		//      output should contain games
		//      games should contain output
		for _, game := range games {
			assert.Contains(t, test.output, game)
		}
	}
}
