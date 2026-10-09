package helper

import (
	"fmt"
	"testing"
)

func TestGameCorrectlyGenerated(t *testing.T) {
	game := Grid{}
	game.FillGrid()

	for _, row := range game.HiddenGrid {
		fmt.Println(row)
	}

	if !game.HiddenGrid.isConsistent() {
		t.Errorf("The game wasn't consistent with sudoku rules.")
	}
}

func TestPossibleValues(t *testing.T) {
	testGame := Grid{}
	var testHidden Game

	testHidden[0] = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	testGame.HiddenGrid = testHidden

	actual := PossibleValues(&testGame, 1, 0)
	expected := []int{4, 5, 6, 7, 8, 9}

	if len(actual) == len(expected) {
		for _, i := range expected {
			found := false
			for _, j := range actual {
				if i == j {
					found = true
				}
			}
			if !found {
				t.Errorf("Couldn't find expected value in actual returned")
			}
		}
	} else {
		t.Error("Actuall implementation finds more / less candidates than expected")
	}
}
