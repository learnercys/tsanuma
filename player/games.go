package games

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// gameLinks TODO review this name
type gameLinks struct {
	Archives []string `json:"archives"`
}

// Games TODO review name, games by month
type Games struct {
	Player string
	Year   int
	Month  time.Month
}

// Date date
type Date struct {
	Year  int
	Month time.Month
}

// return player name
func buildPlayer(url string) string {
	// TODO review player's pattern rules
	playerMatchRegExp := regexp.MustCompile("player/[[:alpha:]]+/games")
	playerMatch := playerMatchRegExp.FindString(url)

	playerRegExp := regexp.MustCompile("/[[:alpha:]]/")
	player := playerRegExp.FindString(playerMatch)

	return player
}

func buildDate(url string) (Date, error) {
	yearAndMonthRegExp := regexp.MustCompile("games/[[:digit:]]{4}/[[:digit:]]{2}")
	yearAndMonth := yearAndMonthRegExp.FindString(url)

	yearRegExp := regexp.MustCompile("/[[:digit:]]{4}/")
	yearMatch := yearRegExp.FindString(yearAndMonth)
	year, errorMatchingYear := strconv.Atoi(yearMatch)

	if errorMatchingYear != nil {
		return Date{}, errorMatchingYear
	}

	monthRegExp := regexp.MustCompile("/[[:digit:]]{2}$")
	monthMatch := monthRegExp.FindString(yearAndMonth)
	monthValue, errorMatchingMonth := strconv.Atoi(monthMatch)

	if errorMatchingMonth != nil {
		return Date{}, errorMatchingMonth
	}

	month := time.Month(monthValue)

	return Date{Year: year, Month: month}, nil
}

// first I should get the player with something like this player/{player}/games
// and the the year and month with something like games/{YYYY}/{MM}$
func buildGames(url string) (Games, error) {
	player := buildPlayer(url)
	date, err := buildDate(url)

	if err != nil {
		return Games{}, err
	}

	return Games{Player: player, Year: date.Year, Month: date.Month}, nil
}

// AvailableArchives monthly archive games by player
func AvailableArchives(player string) ([]Games, error) {
	var url = "https://api.chess.com/pub/player/" + player + "/games/archives"
	var gamesByMonth []Games

	gameLinks := gameLinks{}

	response, requestError := http.Get(url)

	if requestError != nil {
		return gamesByMonth, requestError
	}

	// TODO maybe just for fun, if it's correct also, try to follow functional programming.
	// Basics statements:
	//
	// 1. No mutable data: what I'm not following here because decoding json from a `response.Body`
	//    by providing an argument that is going to be mutated, `gameLinks`.
	// 2. No state (no implicit, hidden state).
	//
	// ref: https://medium.com/@geisonfgfg/functional-go-bc116f4c96a4
	//
	decodingError := json.NewDecoder(response.Body).Decode(&gameLinks)

	if decodingError != nil {
		return gamesByMonth, decodingError
	}

	// here I'm converting a string with the pattern "https://api.chess.com/pub/player/{player}/games/{YYYY}/{MM}"
	// to a `Games` struct
	for _, link := range gameLinks.Archives {
		games, error := buildGames(link)
		if error != nil {

		}
		gamesByMonth = append(gamesByMonth, games)
	}

	return gamesByMonth, nil
}
