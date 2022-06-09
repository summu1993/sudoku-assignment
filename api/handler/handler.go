package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sudoku-assignment/config"
	"sudoku-assignment/generics"
	"sudoku-assignment/request"
	"sudoku-assignment/response"
	"sudoku-assignment/sudoku"

	"github.com/gin-gonic/gin"
)

/*
    For Scalibilty -> When we have to break this service into multiple services and this entry point will act as reverse proxy

 	This reverse proxy will forward the request so the same network no external RPC calls
	UseCase:
	  1. If my service was deplpoyed somewhere else then this proxy architecture would be a scalable approach
	  2. right now reverse proxy will loop back to the same host and port but different end point
*/
func ReverseProxy(proxyUrl string) gin.HandlerFunc {

	remote, err := url.Parse(proxyUrl)
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		path := RedirectUrlPath(c.Request.Header.Get("Service"))
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.URL.Scheme = remote.Scheme
			req.Host = remote.Host
			req.URL.Host = remote.Host
			req.URL.Path = path
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

/*
 TODO: revisit this logic
       static logic for reverse proxy to redirect request based on Service Header
*/

func RedirectUrlPath(service string) string {
	var redirectedPath string
	if service == "solve" {
		redirectedPath = "/v1/solve/sudoku"
	} else if service == "check" {
		redirectedPath = "/v1/check/sudoku"
	}
	return redirectedPath
}

/*
	CHANGES: This API will now accept dimention from user and its comma separated values 
			 If no sudoku string is passed then it will default to a static test sudoko board
*/

func SolveSudoku() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var solveSudokuRequest request.SolveSudokuRequest
		var solveSudokuResponse response.SolveSudokuResponse
		var httpResponse response.HttpResponse

		bindError := ctx.ShouldBindQuery(&solveSudokuRequest)
		if bindError != nil {
			log.Println("solveSudokuRequest error " + bindError.Error())
			setError(httpResponse, ctx, bindError.Error())
			return
		}
		var board [][]int

		if solveSudokuRequest.Board == "" {

			// test static sudoku board if no sudoku board was passed from the query paramater

			board = generics.SudokuBoard{
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
		} else {
			board = generics.MarshalSudoku(solveSudokuRequest.Board, solveSudokuRequest.Dimension)
		}

		sudoku.SolveSudoku(board, 0, 0)
		solveSudokuResponse.Message = "Sudoku has been solved"
		validResponse, marshalError := json.Marshal(solveSudokuResponse)

		// handling error while marshalling the response object
		if marshalError != nil {
			log.Println("error in solveSudokuResponse marshal " + bindError.Error())
			setError(httpResponse, ctx, marshalError.Error())
		}
		httpResponse.RequestId = config.Vars.SERVER.REQUEST_ID
		httpResponse.Data = validResponse
		ctx.JSON(http.StatusOK, httpResponse)
	}
}

/*
    Sudoku board will be passed as comman sparated strings of values

    ASSUMPTION: Number of Values and dimention has to reciprocate

	CHANGES: This API will now accept dimention from user and its comma separated values
*/
func CheckSudokuValidity() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sudokuRequest request.CheckSudokuRequest
		var sudokuResponse response.CheckSudokuResponse
		var httpResponse response.HttpResponse

		bindError := ctx.ShouldBindQuery(&sudokuRequest)

		if bindError != nil {
			log.Println("sudokuRequest error " + bindError.Error())
			setError(httpResponse, ctx, bindError.Error())
			return
		}

		board := generics.MarshalSudoku(sudokuRequest.Board, sudokuRequest.Dimension)
		horizontalPosition := sudokuRequest.HorizontalPosition
		verticalPosition := sudokuRequest.VerticalPosition
		value := sudokuRequest.Value

		isValid := sudoku.IsValidPosition(board, horizontalPosition, verticalPosition, value)
		sudokuResponse.ValidPosition = isValid
		validResponse, marshalError := json.Marshal(sudokuResponse)

		// handling error while marshalling the response object
		if marshalError != nil {
			log.Println("error in sudokuRequest marshal " + marshalError.Error())
			setError(httpResponse, ctx, marshalError.Error())
		}

		httpResponse.RequestId = config.Vars.SERVER.REQUEST_ID
		httpResponse.Data = validResponse
		ctx.JSON(http.StatusOK, httpResponse)
	}
}

func setError(httpResponse response.HttpResponse, ctx *gin.Context, errorMsg string) {

	// Request id is uuid which is sent by frontend everytime it hits backend API , to track/debug the request
	httpResponse.RequestId = config.Vars.SERVER.REQUEST_ID
	httpResponse.ErrorResponse.Message = errorMsg
	httpResponse.ErrorResponse.Code = http.StatusBadRequest
	ctx.JSON(http.StatusBadRequest, httpResponse)
	ctx.Abort()
}
