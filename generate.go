package sudoku

import "math/rand"

func Generate(prefilled int) []int {
	puzzle := make([]int, 81)

	cells := randomNums1to81(prefilled)
	for _, cell := range cells {
		values := getPossibleValues(puzzle, cell)
		if len(values) == 0 {
			return Generate(prefilled)
		}
		puzzle[cell] = values[rand.Int()%len(values)]
	}

	return puzzle
}

func randomNums1to81(count int) []int {
	var result []int

	for len(result) < count {
		rand := rand.Int() % 81
		if !contains(result, rand) {
			result = append(result, rand)
		}
	}

	return result
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}
