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

	assert.NotEmpty(t, games)

}
