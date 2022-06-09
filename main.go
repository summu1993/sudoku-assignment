package main

import (
	"os"
	"sudoku-assignment/api/controllers"
	"sudoku-assignment/config"
)

/*

	Aim: Create a sudoku game
	Assumptions:
	      1. Taking standard sudoku of 9x9 board
		  2. For now we will be passing a 9x9 2-D matrix as game input
		  3. Main.go (entry point) will take a 9x9 2-D matrix as game input and give me the 1 possible solution
		  4. Although each utility function for solving the sudoku will be exposed a REST API as well so that
		     frontend service can consume it
			 i.e, while placing a number on empty spot, frontend can call our backend exposed API to see if it added number is valid or not


	library used:
		1. gin gonic http framewrok

*/

func main() {
	//load configurations
	config.GetEnvConfig()

	//create API entry points
	controllers.CreateUrlMappings()

	// for heroku deployment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" // Default port if not specified
	}

	//set application port
	controllers.Router.Run(":" + port)

}
