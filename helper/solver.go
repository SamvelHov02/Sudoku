package helper

import "math/rand/v2"

type Game [9][9]int

func Solve(g Game) []Game {
	Changed := true
	for Changed {
		Changed = false
		if !g.isConsistent() {
			return nil
		}
		if g.NakedSingles() {
			Changed = true
			continue
		}

		if g.HiddenSingles() {
			Changed = true
			continue
		}

	}

	if g.isSolved() {
		return []Game{g}
	}

	return backtrack(g)
}

// Method finds and marks values for a Naked single,
// A naked single is a cell where only one possible value exists
// returns true if it made any changes
func (g *Game) NakedSingles() bool {
	Changed := false
	for true {
	outer:
		for i, row := range g {
			for j, val := range row {
				if val == 0 {
					// Check possible values
					poss := g.PossibleValues(i, j)
					if len(poss) == 1 {
						g[i][j] = poss[0]
						Changed = true
						break outer
					}
				}
			}
		}
		break
	}

	return Changed
}

// Method that finds and marks values hidden singles
func (g *Game) HiddenSingles() bool {
	Changed := false
	// Row by row find hidden singles
	for i, row := range g {
		RowNumbersLeft := g.PossibleRowValues(i)
		for j := range row {
			possible := g.PossibleValues(i, j)
			commonlity := Intersection(RowNumbersLeft, possible)

			// Found a hidden single
			if len(commonlity) == 1 {
				g[i][j] = commonlity[0]
				Changed = true
			}
		}
	}

	// Column by column find hidden singles

	return Changed
}

func (g *Game) PossibleValues(row int, col int) []int {
	not_possible := make([]bool, 9)

	// Row check
	for _, rowVal := range g[row] {
		if rowVal != 0 {
			not_possible[rowVal] = true
		}
	}

	// Column check
	for _, rows := range g {
		colVal := rows[col]
		if colVal != 0 {
			not_possible[colVal] = true
		}
	}

	// Sqaure check
	y := row / 3
	x := col / 3

	for _, rowVal := range g[x : x*3] {
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

func (g *Game) PossibleRowValues(row int) []int {
	taken := make([]bool, 10)
	for _, val := range g[row] {
		if val != 0 {
			taken[val] = true
		}
	}

	left := make([]int, 9)
	for num, used := range taken {
		if !used {
			left = append(left, num+1)
		}
	}

	return left
}

func (g *Game) PossibleColumnValues(col int) []int {
	taken := make([]bool, 10)
	column := make([]int, 9)

	for _, row := range g {
		column = append(column, row[col])

		if row[col] != 0 {
			taken[row[col]] = true
		}
	}

	left := make([]int, 9)

	for num, used := range taken {
		if !used {
			left = append(left, num+1)
		}
	}

	return left
}

func (g *Game) PossibleBlockValues(rowIdx int, colIdx int) []int {
	taken := make([]bool, 10)

	for _, row := range g[rowIdx*3 : (rowIdx+1)*3] {
		for _, val := range row[colIdx*3 : (colIdx+1)*3] {
			taken[val] = true
		}
	}

	left := make([]int, 9)

	for i, used := range taken {
		if !used {
			left = append(left, i+1)
		}
	}

	return left
}

func (g Game) FancyPossibleValues(row int, col int) []int {
	rowBlock := row / 3
	colBlock := col / 3

	rowPossible := g.PossibleRowValues(row)
	colPossible := g.PossibleColumnValues(col)
	tempPossible := Intersection(rowPossible, colPossible)

	return Intersection(tempPossible, g.PossibleBlockValues(rowBlock, colBlock))
}

// Checks if any cells has no possible value
func (g *Game) isConsistent() bool {
	for i, row := range g {
		for j := range row {
			possibilites := g.PossibleValues(i, j)
			if len(possibilites) == 0 {
				return false
			}
		}
	}
	return true
}

func (g *Game) isSolved() bool {
	consistent := g.isConsistent()

	if consistent {
		for _, row := range g {
			for _, val := range row {
				if val == 0 {
					return false
				}
			}
		}
	} else {
		return false
	}
	return true
}

// A Backtracking algorithm that tries to solve a sudoku by trying one of the possible values
func backtrack(g Game) []Game {

	for i, row := range g {
		for j, val := range row {
			if val != 0 {
				poss := g.PossibleValues(i, j)
				g[i][j] = poss[rand.IntN(len(poss))]
				if g.isSolved() {
					return []Game{g}
				}
				backtrack(g)
			}
		}
	}

	return []Game{g}
}

func Intersection(a []int, b []int) []int {
	returnSlice := make([]int, 9)

	for _, aVal := range a {
		AinB := false
		for _, bVal := range b {
			AinB = aVal == bVal
		}

		if AinB {
			returnSlice = append(returnSlice, aVal)
		}
	}

	return returnSlice
}
