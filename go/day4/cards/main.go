package cards

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	CardNumber     int
	WinningNumbers []int
	MyNumbers      []int
}

func (card *Card) GetMyCardWinningNumbers() []int {
	var winningNumbers []int
	for _, myNumber := range card.MyNumbers {
		for _, winningNumber := range card.WinningNumbers {
			if myNumber == winningNumber {
				winningNumbers = append(winningNumbers, myNumber)
			}
		}
	}
	return winningNumbers
}

func (card *Card) GetWinningNumberPoints() int {
	var points int
	for range card.GetMyCardWinningNumbers() {
		if points == 0 {
			points = 1
		} else {

			points *= 2
		}
	}

	return points
}
func valueToNumber(value []string) []int {
	var numbers []int
	for _, v := range value {
		number, err := strconv.Atoi(strings.Trim(v, " "))
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func ParseFromString(line string) (Card, error) {
	cardSplit := strings.Split(line, ":")
	splittedCardNum := strings.Split(cardSplit[0], " ")
	cardNumStr := splittedCardNum[len(splittedCardNum)-1]
	cardNumer := strings.Trim(cardNumStr, " ")
	cardNumber, err := strconv.Atoi(cardNumer)
	if err != nil {
		fmt.Println("error parsing card number", splittedCardNum, len(splittedCardNum), cardNumStr, cardNumer)
		panic(err)
	}

	cardSplit = strings.Split(cardSplit[1], "|")

	return Card{cardNumber, valueToNumber(strings.Split(cardSplit[0], " ")), valueToNumber(strings.Split(cardSplit[1], " "))}, nil
}
