package part2

import (
	"aoc2023day3/schematicparts"
	"strings"
)

func GearRatio(part1Digit, part2Digit int) int {
	return part1Digit * part2Digit
}

func loopyloop(line string, gearStartPos, gearFinalPos, part1Digit int) (int, int) {
	var part2Digit int
	var startPos int
	for nextNumber, err := schematicparts.GetNextNumber(line, startPos); err == nil; nextNumber, err = schematicparts.GetNextNumber(line, startPos) {
		if (nextNumber.FinalPos >= gearStartPos && nextNumber.StartPos <= gearStartPos) || (nextNumber.StartPos <= gearFinalPos && nextNumber.FinalPos >= gearFinalPos) {
			// println("nextNumber.FinalPos", nextNumber.FinalPos, "nextNumber.StartPos", nextNumber.StartPos)
			// println(line, nextNumber.Digit, "nextNumber.FinalPos >= startPos || nextNumber.StartPos <= finalPos", nextNumber.FinalPos, ">=", gearStartPos, "||", nextNumber.StartPos, "<=", gearFinalPos)
			if part1Digit == 0 {
				part1Digit = nextNumber.Digit
			} else {
				part2Digit = nextNumber.Digit
				break
			}
		}
		startPos = nextNumber.FinalPos

	}

	return part1Digit, part2Digit
}

func GetAdjacentPartDigitsToGear(beforeLine, line, afterLine string, gearPos int) (int, int) {
	var part1Digit, part2Digit int
	gearStartPos, gearFinalPos := schematicparts.GetDiagonalPos(gearPos+1, gearPos, len(line))
	for _, currentLine := range [3]string{beforeLine, line, afterLine} {
		if part1Digit == 0 {

			part1Digit, part2Digit = loopyloop(currentLine, gearStartPos, gearFinalPos, part1Digit)
		} else {
			_, part2Digit = loopyloop(currentLine, gearStartPos, gearFinalPos, part1Digit)

		}
		if part2Digit != 0 {
			return part1Digit, part2Digit
		}
	}
	return 0, 0

}

func Solution(input string) (int, error) {
	var sum int
	strs := strings.Split(input, "\n")
	for i, line := range strs {
		var gearPos int
		for gearPos, ok := schematicparts.GetGearPos(line, gearPos); ok; gearPos, ok = schematicparts.GetGearPos(line, gearPos+1) {
			println("gearPos", line, gearPos, ok)
			beforeLine, afterLine := schematicparts.GetBeforeAndAfterLine(strs, i)
			part1Digit, part2Digit := GetAdjacentPartDigitsToGear(beforeLine, line, afterLine, gearPos)
			sum += GearRatio(part1Digit, part2Digit)
		}
	}
	return sum, nil
}
