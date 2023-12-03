package part1

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	testData, err := os.ReadFile("testinput.txt")
	if err != nil {
		t.Log("solution errored")
		t.Fail()
	}
	solution, err := Solution(string(testData))
	if err != nil {
		t.Log("solution errored")
		t.Log(err)
		t.Fail()
	}
	if solution != 142 {
		t.Log("solution did not match expected value")
		t.Log(solution)
		t.Fail()
	}

}
