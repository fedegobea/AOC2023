package cards

import (
	"testing"
)

func TestWinningNumbersGet(t *testing.T) {
	card := Card{WinningNumbers: []int{41, 48, 83, 86, 17}, MyNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53}}
	winnings := card.GetMyCardWinningNumbers()
	expectedWinnings := [4]int{48, 83, 17, 86}
	for _, winning := range winnings {
		var isWon bool
		for _, expectedWinning := range expectedWinnings {
			if winning == expectedWinning {
				isWon = true
				break
			}
		}
		if !isWon {
			t.Log("winning number", winning, "was not in expected winnings")
			t.Fail()
		}
	}

}

func TestWinningNumberPointsGet(t *testing.T) {
	card := Card{WinningNumbers: []int{41, 48, 83, 86, 17}, MyNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53}}
	points := card.GetWinningNumberPoints()
	if points != 8 {
		t.Log("points were not 8")
		t.Fail()
	}
	card = Card{WinningNumbers: []int{87, 83, 26, 28, 32}, MyNumbers: []int{88, 30, 70, 12, 93, 22, 82, 36}}
	points = card.GetWinningNumberPoints()
	if points != 0 {
		t.Log("points were not 0")
		t.Fail()
	}

}

func TestParseFromString(t *testing.T) {
	card, err := ParseFromString("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	if err != nil {
		t.Log("error parsing card")
		t.Fail()
	}
	if len(card.WinningNumbers) != 5 {
		t.Log("winning numbers were not 5", card.WinningNumbers)
		t.Fail()
	}
	if len(card.MyNumbers) != 8 {
		t.Log("my numbers were not 8", card.MyNumbers)
		t.Fail()
	}
	t.Log(card.GetMyCardWinningNumbers())
	points := card.GetWinningNumberPoints()
	if points != 8 {
		t.Log("points were not 8")
		t.Fail()
	}
	lastInput := "Card 193: 53 40  5 39 13 12 27 57 68 45 | 67 10 87 64 22  6 77 17 20 24 78 52 19 18 99 88 66 31 65 47 11 61 90  9 92"
	card, err = ParseFromString(lastInput)
	if err != nil {
		t.Log("error parsing card")
		t.Fail()
	}
	t.Log("asdas", card.GetMyCardWinningNumbers())
}
