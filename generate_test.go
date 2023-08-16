package sudoku

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name      string
		prefilled int
		want      []int
	}{
		{
			name:      "try",
			prefilled: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Generate(tt.prefilled)
			fmt.Println(Print(got))
			if !IsValid(got) {
				t.Errorf("Generate() = %v", got)
			}
			solutions := Solve(got)
			if len(solutions) != 1 {
				t.Errorf("Solve() = %d", len(solutions))
			}
		})
	}
}
