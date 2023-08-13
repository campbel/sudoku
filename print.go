package sudoku

import "fmt"

func Print(puzzle []int) string {
	result := ""
	result += "┏━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━┓\n"
	for i := 1; i <= 9; i++ {
		result += "┃"
		for j := 1; j <= 9; j++ {
			result += fmt.Sprintf(" %d ", puzzle[i*j-1])
			if j == 3 || j == 6 {
				result += "┃"
			}
		}
		result += "┃\n"
		if i == 3 || i == 6 {
			result += "┣━━━━━━━━━╋━━━━━━━━━╋━━━━━━━━━┫\n"
		}
	}
	result += "┗━━━━━━━━━┻━━━━━━━━━┻━━━━━━━━━┛"
	return result
}
