package part1

import (
	"aoc2023day4/cards"
	"strings"
)

func Solution(input string) (int, error) {
	splittedInput := strings.Split(input, "\n")
	var points int
	for _, line := range splittedInput {
		card, _ := cards.ParseFromString(line)
		points += card.GetWinningNumberPoints()
	}
	return points, nil
}
