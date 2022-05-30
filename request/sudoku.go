package request

type SolveSudokuRequest struct {
	Board              [][]int `json:"board"`
	HorizontalPosition string  `json:"horizontal_position"`
	VerticalPosition   string  `json:"vertical_position"`
	Value              string  `json:"value"`
}
