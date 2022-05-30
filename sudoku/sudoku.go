package sudoku

func IsValidPosition(board [][]int, rowPosition int, columnPosition int, value int) bool {

	// checking horizontal position
	for i := range board {
		if board[rowPosition][i] == value {
			return false
		}
	}

	// checking vertical column position
	for i := range board {
		if board[i][columnPosition] == value {
			return false
		}
	}

	// checking the small grid board (3x3)
	sx, sy := int(rowPosition/3)*3, int(columnPosition/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if board[sy+dy][sx+dx] == value {
				return false
			}
		}
	}
	return true

}
