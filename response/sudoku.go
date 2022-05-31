package response

type CheckSudokuResponse struct {
	ValidPosition bool `json:"valid_position"`
}

type SolveSudokuResponse struct {
	Message string `json:"message"`
}
