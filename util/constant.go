package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var PORT string

// mongo
var MONGO_URI string
var DATABASE string
var USER_COLLECTION string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	// mongo
	MONGO_URI = os.Getenv("MONGO_URI")
	DATABASE = os.Getenv("DATABASE")
	USER_COLLECTION = os.Getenv("USER_COLLECTION")
}
