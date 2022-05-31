package sudoku

import (
	"fmt"
	"sudoku-assignment/generics"
)

func SolveSudoku(board generics.SudokuBoard, rowPosition int, columnPosition int) {
	if rowPosition == len(board) {
		display(board)
		return
	}
	var nRow int
	var nCol int

	if columnPosition == len(board[0])-1 {
		nRow = rowPosition + 1
		nCol = 0
	} else {
		nRow = rowPosition
		nCol = columnPosition + 1
	}

	if board[rowPosition][columnPosition] != 0 {
		SolveSudoku(board, nRow, nCol)
	} else {
		for index := range [9]int{} {
			var value = index + 1
			if IsValidPosition(board, rowPosition, columnPosition, value) {
				board[rowPosition][columnPosition] = value
				SolveSudoku(board, nRow, nCol)
				board[rowPosition][columnPosition] = 0
			}
		}
	}
}

func display(board generics.SudokuBoard) {
	for _, row := range board {
		fmt.Println(row)
	}
	fmt.Println()
}

func IsValidPosition(board generics.SudokuBoard, rowPosition int, columnPosition int, value int) bool {

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
	for dx := range [3]int{} {
		for dy := range [3]int{} {
			if board[sx+dx][sy+dy] == value {
				return false
			}
		}
	}
	return true

}
