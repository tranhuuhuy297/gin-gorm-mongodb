package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var PORT string

func Load_Env() {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
}
