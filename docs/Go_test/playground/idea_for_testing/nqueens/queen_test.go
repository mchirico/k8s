package nqueens

import "testing"

func Boards(i int) [][]int {

	board := [][][]int{
		{{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}},

		{{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0}},


		{{1, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0},
		{0, 1, 0, 0}},

	}

		return board[i]
}



func TestSafe(t *testing.T) {


	type args struct {
		board [][]int
		row   int
		col   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

		{name: "True", args: args{board: Boards(0),row: 1,col: 1,}, want: true},
		{name: "False", args: args{board: Boards(1),row: 1,col: 1,}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Safe(tt.args.board, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}
