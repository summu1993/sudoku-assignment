package controllers

import (
	"sudoku-assignment/api/handler"
	"sudoku-assignment/api/middleware"
	"sudoku-assignment/config"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {

	Router = gin.Default()

	Router.Use(middleware.Headers())
	Router.Use(middleware.RequestLogger())

	/*
			apiRouter can accept a chain of handlers/middlwares, we can even authenticate a user by passing jwt as well
			Cases
				1. Added a support of reverse proxy (in case current repo is my entry point and i will verify the request
				   and then pass on the request forward to the desired destination)
				2. but in my case everything is one repo, just for scaliblity purpose added reverse proxy support


		   APIs
		     	1. 	/sudoku - GET is the entrypoint
				2. 	/solve/sudoku - solves a unsolved 2-D matrix of sudoku board
				3. 	/check/sudoku - will check if a placed number is valid or not (this API is ideal for frontend)
				    Whenever a player places a number in the grid(frontend) he/she will call this api to validate if this is valid placment or not
					like real sudoku game
	*/

	sudokuUrl := config.Vars.SUDOKU.SOLVE_SUDOKU_PROXY_URL

	apiRouter := Router.Group("/v1")
	{
		apiRouter.GET("/sudoku", handler.ReverseProxy(sudokuUrl))
		apiRouter.GET("/solve/sudoku", handler.SolveSudoku())
		apiRouter.GET("/check/sudoku", handler.CheckSudokuValidity())
	}
}
