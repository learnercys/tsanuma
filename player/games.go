package games

import (
	"encoding/json"
	"net/http"
	"regexp"
)

// gameUrls all archived urls
type gameUrls struct {
	Archives []string `json:"archives"`
}

// Games TODO review name, games by month
type Games struct {
	Player string
	Year   string
	Month  string
}

// Date date
type Date struct {
	Year  string
	Month string
}

// return player name
func buildPlayer(url string) string {
	// TODO review player's pattern rules
	playerMatchServiceRegExp := regexp.MustCompile("player/[[:alpha:]]+/games")
	playerMatchService := playerMatchServiceRegExp.FindString(url)

	playerMatchRegExp := regexp.MustCompile("/[[:alpha:]]+/")
	playerMatch := playerMatchRegExp.FindString(playerMatchService)

	playerRegExp := regexp.MustCompile("[[:alpha:]]+")
	player := playerRegExp.FindString(playerMatch)

	return player
}

func buildDate(url string) Date {
	yearAndMonthRegExp := regexp.MustCompile("games/[[:digit:]]{4}/[[:digit:]]{2}")
	yearAndMonth := yearAndMonthRegExp.FindString(url)

	yearRegExp := regexp.MustCompile("/[[:digit:]]{4}/")
	yearMatch := yearRegExp.FindString(yearAndMonth)
	year := yearMatch[1 : len(yearMatch)-1]

	monthRegExp := regexp.MustCompile("/[[:digit:]]{2}$")
	month := monthRegExp.FindString(yearAndMonth)[1:]
	date := Date{Year: year, Month: month}

	return date
}

// first I should get the player with something like this player/{player}/games
// and the the year and month with something like games/{YYYY}/{MM}$
func buildGames(url string) Games {
	player := buildPlayer(url)
	date := buildDate(url)

	return Games{Player: player, Year: date.Year, Month: date.Month}
}

// AvailableArchives monthly archive games by player
func AvailableArchives(player string) ([]Games, error) {
	var url = "https://api.chess.com/pub/player/" + player + "/games/archives"
	var gamesByMonth []Games

	urls := gameUrls{}

	response, requestError := http.Get(url)

	if requestError != nil {
		return gamesByMonth, requestError
	}

	decodingError := json.NewDecoder(response.Body).Decode(&urls)

	if decodingError != nil {
		return gamesByMonth, decodingError
	}

	// here I'm converting a string with the pattern "https://api.chess.com/pub/player/{player}/games/{YYYY}/{MM}"
	// to a `Games` struct
	for _, url := range urls.Archives {
		games := buildGames(url)
		gamesByMonth = append(gamesByMonth, games)
	}

	return gamesByMonth, nil
}
