package schematicparts

import (
	"strconv"
	"strings"
	"unicode"
)

const NotSymbolChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789."

func IsSymbol(char rune) bool {
	return !strings.Contains(NotSymbolChars, string(char))
}

func GetDiagonalPos(startPos, finalPos, maxLen int) (int, int) {
	var x = startPos - 1
	var y = finalPos + 1
	if x < 0 {
		x = 0
	}
	if y > maxLen {
		y = maxLen
	}
	return x, y
}

func checkIfLineHasSymbol(line string) bool {
	for _, c := range line {
		if IsSymbol(c) {
			return true
		}
	}
	return false
}

func IsSymbolAdjacent(line, beforeLine, afterLine string, startPos, finalPos int) bool {
	start, end := GetDiagonalPos(startPos, finalPos, len(line))
	for _, checkedLine := range [3]string{line, beforeLine, afterLine} {
		if checkedLine == "" {
			continue
		}
		if checkIfLineHasSymbol(checkedLine[start:end]) {
			return true
		}
	}
	return false

}

func IsPart(line, beforeLine, afterLine string, startPos, finalPos int) bool {
	return IsSymbolAdjacent(line, beforeLine, afterLine, startPos, finalPos)
}

type EngineSchematicPart struct {
	Digit    int
	StartPos int
	FinalPos int
}

func GetNextNumber(line string, startPos int) (EngineSchematicPart, error) {
	var finalPos int
	newStartPos := startPos
	hasAssignedStartPos := false
	for pos, c := range line[startPos:] {
		if unicode.IsDigit(c) && !hasAssignedStartPos {
			newStartPos = startPos + pos
			hasAssignedStartPos = true
		}
		if !unicode.IsDigit(c) && hasAssignedStartPos {
			finalPos = startPos + pos
			break
		}
	}
	if finalPos < newStartPos {
		finalPos = len(line)
	}
	digit, err := strconv.Atoi(line[newStartPos:finalPos])
	return EngineSchematicPart{digit, newStartPos, finalPos}, err
}

func GetBeforeAndAfterLine(strs []string, currentLineIndex int) (string, string) {
	var beforeLine, afterLine string
	if currentLineIndex > 0 {
		beforeLine = strs[currentLineIndex-1]
	}
	if currentLineIndex < len(strs)-1 {
		afterLine = strs[currentLineIndex+1]
	}
	return beforeLine, afterLine
}

func GetGearPos(line string, gearPosStart int) (int, bool) {
	for i, c := range line[gearPosStart:] {
		if c == '*' {
			println("found gear at", i)
			return gearPosStart + i, true
		}

	}
	return 0, false
}
