package part2

import "aoc2023day2/games"

func Solution(input string) (int, error) {
	games := games.GamesFromInput(input)
	gamesValue := 0
	for _, game := range games {
		gamesValue += game.GetLeastPossibleCombinationMultiplied()
	}
	return gamesValue, nil
}
