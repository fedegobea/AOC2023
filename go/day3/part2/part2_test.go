package part2

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
	if solution != 467835 {
		t.Log("solution did not match expected value")
		t.Log(solution, "should have been 467835")
		t.Fail()
	}

}