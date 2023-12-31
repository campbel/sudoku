package sudoku

func Solve(puzzle []int) [][]int {
	// Solve via backtracking
	return solve(puzzle)
}

func solve(in []int) [][]int {
	var puzzle []int = make([]int, len(in))
	copy(puzzle, in)

	i := 0
	for ; i < 81; i++ {
		if puzzle[i] == 0 {
			break
		}
	}
	if i == 81 {
		return [][]int{puzzle}
	}

	// Find the possible values for the cell
	possibleValues := getPossibleValues(puzzle, i)
	if len(possibleValues) == 0 {
		return nil
	}

	// Try each possible value
	var solutions [][]int
	for _, v := range possibleValues {
		puzzle[i] = v
		if solved := solve(puzzle); solved != nil {
			solutions = append(solutions, solved...)
		}
	}

	// No solution found
	return solutions
}

func getPossibleValues(puzzle []int, i int) []int {
	exists := make([]bool, 9)

	// Check row
	row := i / 9
	for j := 0; j < 9; j++ {
		if puzzle[row*9+j] != 0 {
			exists[puzzle[row*9+j]-1] = true
		}
	}

	// Check column
	col := i % 9
	for j := 0; j < 9; j++ {
		if puzzle[col+j*9] != 0 {
			exists[puzzle[col+j*9]-1] = true
		}
	}

	// Check block
	blockRow := row / 3
	blockCol := col / 3
	for j := 0; j < 9; j++ {
		if puzzle[blockRow*27+blockCol*3+j%3+j/3*9] != 0 {
			exists[puzzle[blockRow*27+blockCol*3+j%3+j/3*9]-1] = true
		}
	}

	// Convert to slice of ints
	possibleValues := make([]int, 0, 9)
	for j := 0; j < 9; j++ {
		if !exists[j] {
			possibleValues = append(possibleValues, j+1)
		}
	}

	return possibleValues
}
