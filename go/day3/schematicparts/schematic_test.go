package schematicparts

import testing "testing"

func TestGetDiagonalPos(t *testing.T) {
	x, y := GetDiagonalPos(0, 10, 9)
	if x != 0 {
		t.Log("x should have been 0")
		t.Fail()
	}
	if y != 9 {
		t.Log("y should have been 9")
		t.Fail()
	}
	x, y = GetDiagonalPos(1, 10, 10)
	if x != 0 {
		t.Log("x should have been 0")
		t.Fail()
	}
	if y != 10 {
		t.Log("y should have been 10", y)
		t.Fail()
	}

}
func TestIsSymbolAdjacent(t *testing.T) {
	var beforeLine = ""
	var line = "467..114.."
	var afterLine = "...*......"
	part := IsPart(line, beforeLine, afterLine, 0, 3)
	if part != true {
		t.Log("part 467 should have been true")
		t.Fail()
	}

	part2 := IsPart(line, beforeLine, afterLine, 5, 8)
	if part2 != false {
		t.Log("part 114 should have been false")
		t.Fail()
	}
}

func TestIsSymbolAdjacent2(t *testing.T) {

	var beforeLine = "617*......"
	var line = ".....+.58."
	var afterLine = "..592....."
	part := IsPart(line, beforeLine, afterLine, 7, 9)
	if part != false {
		t.Log("part 58 should have been false")
		t.Fail()
	}
}

func TestGetGear(t *testing.T) {
	var beforeLine = "617*......"
	var line = ".....+.58."
	var afterLine = "..592....."
	pos, ok := GetGearPos(beforeLine, 0)
	if !ok {
		t.Log(ok, "should have been nil")
		t.Fail()
	}
	if pos != 3 {
		t.Log("pos should have been 3")
		t.Fail()
	}
	_, ok = GetGearPos(line, 0)
	if ok {
		t.Log(ok, "should have been set")
		t.Fail()
	}
	_, ok = GetGearPos(afterLine, 0)
	if ok {
		t.Log(ok, "should have been set")
		t.Fail()
	}

}
