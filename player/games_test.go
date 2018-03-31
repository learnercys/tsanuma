// TODO stub requests
package games

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvailableArchivesOK(t *testing.T) {
	games, err := AvailableArchives("learnercys")

	if err != nil {
		t.Fatal("couldn't get archive games")
	}

	assert.NotEmpty(t, games.Archives)
	assert.Contains(t, games.Archives, "https://api.chess.com/pub/player/learnercys/games/2015/12")
	assert.Contains(t, games.Archives, "https://api.chess.com/pub/player/learnercys/games/2018/01")
	assert.Contains(t, games.Archives, "https://api.chess.com/pub/player/learnercys/games/2018/02")
	assert.Contains(t, games.Archives, "https://api.chess.com/pub/player/learnercys/games/2018/03")

}
