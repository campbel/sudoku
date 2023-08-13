package sudoku

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		puzzle []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				puzzle: []int{
					5, 3, 4, 6, 7, 8, 9, 1, 2,
					6, 7, 2, 1, 9, 5, 3, 4, 8,
					1, 9, 8, 3, 4, 2, 5, 6, 7,
					8, 5, 9, 7, 6, 1, 4, 2, 3,
					4, 2, 6, 8, 5, 3, 7, 9, 1,
					7, 1, 3, 9, 2, 4, 8, 5, 6,
					9, 6, 1, 5, 3, 7, 2, 8, 4,
					2, 8, 7, 4, 1, 9, 6, 3, 5,
					3, 4, 5, 2, 8, 6, 1, 7, 9,
				},
			},
			want: true,
		},
		{
			name: "valid",
			args: args{
				puzzle: []int{
					5, 2, 4, 6, 7, 8, 9, 1, 2,
					6, 7, 2, 1, 9, 5, 3, 4, 8,
					1, 9, 8, 3, 4, 2, 5, 6, 7,
					8, 5, 9, 7, 6, 1, 4, 2, 3,
					4, 2, 6, 8, 5, 3, 7, 9, 1,
					7, 1, 3, 9, 2, 4, 8, 5, 6,
					9, 6, 1, 5, 3, 7, 2, 8, 4,
					2, 8, 7, 4, 1, 9, 6, 3, 5,
					3, 4, 5, 2, 8, 6, 1, 7, 9,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.puzzle); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
