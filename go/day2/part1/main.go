package part1

import "aoc2023day2/games"

func Solution(input string) (int, error) {
	gameconf := games.GameConfig{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	games := games.GamesFromInput(input)
	gameValue := 0
	for _, game := range games {
		if game.IsPossibleGame(gameconf) {
			gameValue += game.GameNumber
		}
	}
	return gameValue, nil
}
