package sudoku

func Candidates(puzzle []int) [][]int {
	if !IsValid(puzzle) {
		return nil
	}
	result := make([][]int, 81)

	for i, cell := range puzzle {
		if cell != 0 {
			result[i] = []int{cell}
			continue
		}

		// Find the possible values for the cell
		result[i] = getPossibleValues(puzzle, i)
	}

	return result
}
