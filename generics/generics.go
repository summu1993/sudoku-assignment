package generics

import (
	"errors"
	"strconv"
	"strings"
)

type SudokuBoard [][]int

func MarshalSudoku(genericBoard string, dimension int) ([][]int, error) {

	matrix := make([][]int, dimension) // dimension*dimension matrix
	sudokuBoard := strings.Split(genericBoard, ",")
	if len(sudokuBoard) != dimension*dimension {
		return matrix, errors.New("dimention and grid values doesnt match")
	}

	for i := 0; i < dimension; i++ {
		matrix[i] = make([]int, 0, dimension)
		row := make([]int, dimension)
		for j := 0; j < dimension; j++ {
			value, _ := strconv.Atoi(sudokuBoard[i*dimension+j])
			row[j] = value
			matrix[i] = append(matrix[i], row[j])
		}
	}
	return matrix, nil
}
