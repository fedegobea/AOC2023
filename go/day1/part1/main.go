package part1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
--- Day 1: Trebuchet?! ---
Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

*/

func GetFirstNumber(line string) (int, error) {
	for _, first := range line {
		if unicode.IsDigit(first) {
			return strconv.Atoi(string(first))
		}
	}
	return 0, fmt.Errorf("No digit found in %s", line)
}

func GetLastNumber(line string) (int, error) {

	for i := len(line) - 1; i >= 0; i-- {
		first := rune(line[i])
		if unicode.IsDigit(first) {
			return strconv.Atoi(string(first))
		}
	}
	return 0, nil
}
func GetFirstAndLastDigits(line string) (int, error) {
	first, firstErr := GetFirstNumber(line)
	if firstErr != nil {
		return 0, firstErr
	}
	last, lastErr := GetLastNumber(line)
	if lastErr != nil {
		return 0, lastErr
	}
	concatStr := fmt.Sprint(first) + fmt.Sprint(last)
	return strconv.Atoi(concatStr)
}

func Solution(input string) (int, error) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		res, error := GetFirstAndLastDigits(line)
		if error != nil {
			return 0, error
		}
		sum += res
	}
	return sum, nil
}
