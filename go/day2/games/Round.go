package games

import (
	"strconv"
	"strings"
)

type Round struct {
	Red   int
	Green int
	Blue  int
}

func RoundFromString(input string) Round {
	// 3 blue, 4 red
	splitted := strings.Split(input, ",")
	roundMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, s := range splitted {
		cube := strings.Split(strings.Trim(s, " "), " ")
		val, err := strconv.Atoi(cube[0])
		panicOnErr(err)
		roundMap[cube[1]] = val
	}

	return Round{
		Red:   roundMap["red"],
		Green: roundMap["green"],
		Blue:  roundMap["blue"],
	}
}
