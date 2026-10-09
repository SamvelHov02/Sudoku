package helper

import (
	"fmt"
	"math/rand/v2"
)

type Grid struct {
	HiddenGrid   Game
	RevealedGrid Game
	Mistakes     int
	Hints        int
}

// This method creates the underlying structure / grid
func (g *Grid) FillGrid() bool {
	return g.fill(0)
}

func (g *Grid) fill(pos int) bool {
	if pos == 81 {
		return true
	}

	row, col := pos/9, pos%9

	candidates := PossibleValues(g, row, col)
	rand.Shuffle(len(candidates), func(a, b int) {
		candidates[a], candidates[b] = candidates[b], candidates[a]
	})

	for _, v := range candidates {
		g.HiddenGrid[row][col] = v
		if g.fill(pos + 1) {
			return true
		}
	}

	g.HiddenGrid[row][col] = 0
	return false
}

func (g *Grid) MarkCell(row int, col int, val int) {
	g.RevealedGrid[row][col] = val
	// TODO : Check if its correct and how to handle incorrect inputs
}

// This function returns the possible values for a cell if there according to sudoku rules
func PossibleValues(g *Grid, row int, col int) []int {
	fmt.Printf("row : %d, col : %d", row, col)
	not_possible := make([]bool, 10)

	// Row check
	for _, rowVal := range g.HiddenGrid[row] {
		if rowVal != 0 {
			not_possible[rowVal] = true
		}
	}

	// Column check
	for _, rows := range g.HiddenGrid {
		colVal := rows[col]
		if colVal != 0 {
			not_possible[colVal] = true
		}
	}

	// Sqaure check
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3

	for _, rowVal := range g.HiddenGrid[boxRow : boxRow+3] {
		for _, colVal := range rowVal[boxCol : boxCol+3] {
			if colVal != 0 {
				not_possible[colVal] = true
			}
		}
	}

	retSlice := make([]int, 0, 9)

	for i, not_pos := range not_possible[1:] {
		if !not_pos {
			retSlice = append(retSlice, i+1)
		}
	}

	return retSlice
}

// TODO: method that removes cells and turns a board into a puzzle
// Determine the difficulty Number of hints, type of cells to fill ?
func (g *Grid) MakePuzzle() {
	CurrHints := 81
	g.RevealedGrid = g.HiddenGrid

	for CurrHints > g.Hints {
		tempGame := g.RevealedGrid
		rowIdx := rand.IntN(9)
		colIdx := rand.IntN(9)

		tempGame[rowIdx][colIdx] = 0
		posSolution := len(Solve(tempGame))

		if posSolution == 1 {
			g.RevealedGrid = tempGame
		} else {
			return
		}
	}

}
