package main

import (
	"aoc2023day3/part1"
	"aoc2023day3/part2"
	"fmt"
	"os"
)

func readFileError(err error) {
	if err != nil {
		panic(fmt.Errorf("error on read %w", err))
	}
}
func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	input, err := os.ReadFile("./input.txt")
	readFileError(err)
	part1Solution, part1SolErr := part1.Solution(string(input))
	panicOnErr(part1SolErr)
	println("Part1 Solution", part1Solution)

	part2Solution, part2SolErr := part2.Solution(string(input))
	panicOnErr(part2SolErr)

	println("Part2 Solution", part2Solution)

}
