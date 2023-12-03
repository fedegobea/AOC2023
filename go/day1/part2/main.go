package part2

import (
	"fmt"
	"strconv"
	"strings"
)

/*
--- Part Two ---
Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.
*/

func matchStrToInt(digit string) (int, error) {
	switch digit {
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	default:
		return strconv.Atoi(digit)

	}
}

var digits = [18]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func getStrDigitPos(line string, initialPos int, checkFn func(line string, strDigit string) int, comparisonFunc func(a, b int) bool) string {
	validPos := initialPos
	res := ""
	for _, strDigit := range digits {
		pos := checkFn(line, strDigit)
		if pos != -1 && comparisonFunc(pos, validPos) {
			validPos = pos
			res = strDigit
		}
	}
	return res
}

func GetFirstNumber(line string) string {
	return getStrDigitPos(line, len(line), strings.Index, lesserThanEq)

}

func GetLastNumber(line string) string {
	return getStrDigitPos(line, 0, strings.LastIndex, greaterThanEq)
}

func GetFirstAndLastDigits(line string) (int, error) {
	first, firstErr := matchStrToInt(GetFirstNumber(line))
	if firstErr != nil {
		return 0, firstErr
	}
	last, lastErr := matchStrToInt(GetLastNumber(line))
	if lastErr != nil {
		return 0, lastErr
	}

	concatStr := fmt.Sprint(first) + fmt.Sprint(last)
	return strconv.Atoi(concatStr)
}
func lesserThanEq(a, b int) bool {
	return a <= b
}

func greaterThanEq(a, b int) bool {
	return a >= b
}

func Solution(input string) (int, error) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {

		res, e := GetFirstAndLastDigits(line)
		if e != nil {
			return 0, e
		}
		sum += res
	}
	return sum, nil
}
