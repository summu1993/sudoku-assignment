package request

type CheckSudokuRequest struct {
	Board              string `form:"board"`
	Dimension          int    `form:"dimension"`
	HorizontalPosition int    `form:"horizontalPosition"`
	VerticalPosition   int    `form:"verticalPosition"`
	Value              int    `form:"value"`
}

type SolveSudokuRequest struct {
	Board     string `form:"board"`
	Dimension int    `form:"dimension"`
}
