package helper

import "math/rand/v2"

type Grid struct {
	HiddenGrid   [9][9]int
	RevealedGrid [9][9]int
	Mistakes     int
	Hints        int
}

// This method creates the underlying structure / grid
func (g *Grid) fillGrid() {
	for i := 0; i < len(g.HiddenGrid); i++ {
		row := g.HiddenGrid[i]
		tempRow := make([]int, 9)
		for j, val := range row {
			// Unassigned cell
			if val == 0 {
				possible := PossibleValues(g, i, j)

				if len(possible) > 0 {
					tempRow[j] = possible[rand.IntN(len(possible))]
				} else {
					// No possible value, some conflict that needs to be solved
					i--
					break
				}
			}
		}
	}

}

func (g *Grid) MarkCell(row int, col int, val int) {
	g.RevealedGrid[row][col] = val
	// TODO : Check if its correct and how to handle incorrect inputs
}

// This function returns the possible values for a cell if there according to sudoku rules
func PossibleValues(g *Grid, row int, col int) []int {
	not_possible := make([]bool, 9)

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
	y := row / 3
	x := col / 3

	for _, rowVal := range g.HiddenGrid[x : x*3] {
		for _, colVal := range rowVal[y : y*3] {
			if colVal != 0 {
				not_possible[colVal] = true
			}
		}
	}

	retSlice := make([]int, 0, 9)

	for i, not_pos := range not_possible {
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
