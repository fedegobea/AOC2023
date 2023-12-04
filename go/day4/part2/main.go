package part2

import (
	"aoc2023day4/cards"
	"strings"
)

func Solution(input string) (int, error) {
	// TODO: solution is pretty slow, should be optimized, and from the start I think I can get O(1) space complexity.
	// * If I have time later I will try to rewrite it and see what happens.
	splittedInput := strings.Split(input, "\n")
	var newInput []string = make([]string, len(splittedInput))
	copy(newInput[:], splittedInput[:])
	for i := 0; i < len(newInput); i++ {
		line := newInput[i]
		card, _ := cards.ParseFromString(line)
		copies := len(card.GetMyCardWinningNumbers())

		if copies == 0 {
			continue
		}
		cardsToAppend := splittedInput[card.CardNumber : card.CardNumber+copies]
		newInput = append(newInput, cardsToAppend...)
	}

	return len(newInput), nil
}
