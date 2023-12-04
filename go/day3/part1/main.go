package part1

import (
	"aoc2023day3/schematicparts"
	"strings"
)

func GetNextPartNumber(beforeLine, line, afterLine string) int {
	startPos := 0
	partNumberSum := 0
	for nextNumber, err := schematicparts.GetNextNumber(line, startPos); err == nil; nextNumber, err = schematicparts.GetNextNumber(line, startPos) {
		if schematicparts.IsPart(line, beforeLine, afterLine, nextNumber.StartPos, nextNumber.FinalPos) {
			partNumberSum += nextNumber.Digit
		}
		startPos = nextNumber.FinalPos

	}
	return partNumberSum

}

func Solution(input string) (int, error) {
	strs := strings.Split(input, "\n")
	var sum int
	for i, line := range strs {
		beforeLine, afterLine := schematicparts.GetBeforeAndAfterLine(strs, i)
		sum += GetNextPartNumber(beforeLine, line, afterLine)
	}
	return sum, nil
}
