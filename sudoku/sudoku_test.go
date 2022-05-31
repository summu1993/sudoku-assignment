package sudoku

import (
	"sudoku-assignment/generics"
	"testing"
)

/*
 	command to run nested test cases files are
	go test ./...
*/

// 9x9 board
var board = generics.SudokuBoard{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

type addTest struct {
	arg1     generics.SudokuBoard
	arg2     int
	arg3     int
	arg4     int
	expected bool
}

var addTests = []addTest{
	{board, 0, 2, 3, false},
	{board, 1, 2, 2, true},
}

func TestSolveSudoku(t *testing.T) {

	for _, test := range addTests {
		if output := IsValidPosition(test.arg1, test.arg2, test.arg3, test.arg4); output != test.expected {
			t.Errorf("Output not equal to expected ")
		} else {
			t.Errorf("Test cases passed")
		}
	}
}
