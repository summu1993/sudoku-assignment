package config

import "github.com/spf13/viper"

var Vars Configurations

// Configurations exported
type Configurations struct {
	SERVER Server
	SUDOKU Sudoku
}

type Server struct {
	ENV      string
	APP_PORT string
}

type Sudoku struct {
	SOLVE_SUDOKU_PROXY_URL string
}

func GetEnvConfig() {
	/*
		keeping config file at the root level although config can be added to a custom path(config dir)
		and its reference can be passed from main.go

		configPath := "config/" + env

	*/

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file at root level
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Couldn't load environment configuration, cannot start. Terminating. Error: " + err.Error())
	}

	//to unmarshal values into targer global config object
	err = viper.Unmarshal(&Vars)

	if err != nil {
		panic("Couldn't load environment configuration, cannot start. Terminating. Error: " + err.Error())
	}
}
