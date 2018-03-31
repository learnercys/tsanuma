package games

// AvailableGames TODO review this name
import (
	"encoding/json"
	"net/http"
)

// AvailableGames TODO review this name
type AvailableGames struct {
	Archives []string `json:"archives"`
}

// AvailableArchives monthly archive games by player
func AvailableArchives(username string) (AvailableGames, error) {
	var url = "https://api.chess.com/pub/player/" + username + "/games/archives"
	games := AvailableGames{}

	response, requestError := http.Get(url)

	if requestError != nil {
		return games, requestError
	}

	decodingError := json.NewDecoder(response.Body).Decode(&games)

	if decodingError != nil {
		return games, decodingError
	}

	return games, nil
}
