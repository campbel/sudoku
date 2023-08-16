package sudoku

func IsValid(puzzle []int) bool {
	// Check rows
	for i := 0; i < 9; i++ {
		if !isValidSet(puzzle[i*9 : i*9+9]) {
			return false
		}
	}

	// Check columns
	for i := 0; i < 9; i++ {
		set := make([]int, 9)
		for j := 0; j < 9; j++ {
			set[j] = puzzle[i+j*9]
		}
		if !isValidSet(set) {
			return false
		}
	}

	// Check blocks
	for i := 0; i < 9; i++ {
		set := make([]int, 9)
		for j := 0; j < 9; j++ {
			idx := (i/3)*27 + (i%3)*3 + j%3 + j/3*9
			set[j] = puzzle[idx]
		}
		if !isValidSet(set) {
			return false
		}
	}

	return true
}

func isValidSet(set []int) bool {
	if len(set) != 9 {
		return false
	}
	vals := make([]bool, 9)
	for _, v := range set {
		if v == 0 {
			continue
		}
		if vals[v-1] {
			return false
		}
		vals[v-1] = true
	}
	return true
}
