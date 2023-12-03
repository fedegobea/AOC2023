package games

import (
	"strconv"
	"strings"
)

type Game struct {
	Rounds     []Round
	GameNumber int
}

func (g Game) IsPossibleGame(gameconf GameConfig) bool {
	for _, round := range g.Rounds {
		if !gameconf.IsPossibleRound(round) {
			return false
		}
	}
	return true
}

func (g Game) GetLeastPossibleCombinationMultiplied() int {
	red := 0
	green := 0
	blue := 0
	for _, round := range g.Rounds {
		if round.Red >= red {
			red = round.Red
		}
		if round.Green >= green {
			green = round.Green
		}
		if round.Blue >= blue {
			blue = round.Blue
		}
	}
	return red * green * blue
}

func getGameNumber(input string) int {
	splittedGame := strings.Split(input, " ")
	gameNumber, err := strconv.Atoi(splittedGame[1])
	panicOnErr(err)
	return gameNumber
}

func GameFromString(input string) Game {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	splitted := strings.Split(input, ":")
	gameNumber := getGameNumber(splitted[0])

	roundsStrArr := strings.Split(splitted[1], ";")

	rounds := make([]Round, len(roundsStrArr))
	for i, roundStr := range roundsStrArr {
		rounds[i] = RoundFromString(roundStr)
	}

	return Game{
		Rounds:     rounds,
		GameNumber: gameNumber,
	}
}

func GamesFromInput(input string) []Game {
	games := make([]Game, 0)
	for _, line := range strings.Split(input, "\n") {
		game := GameFromString(line)
		games = append(games, game)

	}
	return games
}
