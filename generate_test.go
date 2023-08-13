package sudoku

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "try",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Generate()
			fmt.Println(Print(got))
			if !IsValid(got) {
				t.Errorf("Generate() = %v", got)
			}
		})
	}
}
