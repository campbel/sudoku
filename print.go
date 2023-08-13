package sudoku

import "fmt"

func Print(puzzle []int) string {
	result := ""
	result += "┏━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━┓\n"
	for i := 0; i < 9; i++ {
		result += "┃"
		for j := 0; j < 9; j++ {
			result += fmt.Sprintf(" %d ", puzzle[i*9+j])
			if j == 2 || j == 5 {
				result += "┃"
			}
		}
		result += "┃\n"
		if i == 2 || i == 5 {
			result += "┣━━━━━━━━━╋━━━━━━━━━╋━━━━━━━━━┫\n"
		}
	}
	result += "┗━━━━━━━━━┻━━━━━━━━━┻━━━━━━━━━┛"
	return result
}
