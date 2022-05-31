package request

import "sudoku-assignment/generics"

type CheckSudokuRequest struct {
	Board              generics.SudokuBoard `json:"board"`
	HorizontalPosition int                  `json:"horizontal_position"`
	VerticalPosition   int                  `json:"vertical_position"`
	Value              int                  `json:"value"`
}

type SolveSudokuRequest struct {
	Board generics.SudokuBoard `json:"board"`
}
