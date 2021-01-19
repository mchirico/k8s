package nqueens

//import "fmt"

// Safe true is safe
func Safe(board [][]int, row int, col int) bool {
	var x, y int = 0, 0
	for x := 0; x < col; x++ {
		if board[row][x] == 1 {
			return false
		}
	}

	for x, y := row, col; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if board[x][y] == 1 {
			return false
		}
	}
	if board[x][y] == 1 {
		return false
	}

	for x, y = row, col; y >= 0 && x < 4; x, y = x+1, y-1 {
		if board[x][y] == 1 {
			return false
		}
	}
	return true

}


