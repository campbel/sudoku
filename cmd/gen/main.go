package main

import (
	"fmt"

	"github.com/campbel/sudoku"
	"github.com/campbel/yoshi"
)

func main() {
	yoshi.New("gen").Run(func() {
		for {
			puzzle := sudoku.Generate(30)
			solutions := sudoku.Solve(puzzle)
			if len(solutions) == 1 {
				fmt.Println("Puzzle")
				fmt.Println(sudoku.Print(puzzle))
				fmt.Println("Solution")
				fmt.Println(sudoku.Print(solutions[0]))
				break
			}
		}
	})
}
