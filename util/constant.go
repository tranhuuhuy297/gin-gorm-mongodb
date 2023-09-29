package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var PORT string
var MONGO_URI string

func Load_Env() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	MONGO_URI = os.Getenv("MONGO_URI")
}
