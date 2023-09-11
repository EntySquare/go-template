package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load("../.env.dev")
	//err := godotenv.Load("go-template/.env.dev")
	if err != nil {
		err = godotenv.Load("../.env.dev")
		if err != nil {
			fmt.Print("Error loading .env file")
		}
	}
	return os.Getenv(key)
}
