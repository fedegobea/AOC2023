package part2

import (
	"os"
	"testing"
)

func TestPart2(t *testing.T) {
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
	if solution != 281 {
		t.Log("solution did not match expected value")
		t.Log(solution)
		t.Fail()
	}
	testData2, err2 := os.ReadFile("testinput2.txt")
	if err2 != nil {
		t.Log("err2")
		t.Fail()
	}
	solution2, err3 := Solution(string(testData2))

	if err3 != nil {
		t.Log("solution errored")
		t.Log(err3)
		t.Fail()
	}
	res := 37 + 66 + 14 + 75 + 49 + 18 + 96 + 45 + 58 + 58 + 99 + 29 + 26 + 78 + 68 + 11 + 77 + 33 + 61 + 89 + 21
	if solution2 != res {
		t.Log("solution did not match expected value")
		t.Log("expected", res, "got", solution2)
		t.Fail()
	}
}
