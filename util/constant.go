package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// server
var PORT string

// mongo
var MONGO_URI string
var DATABASE string
var USER_COLLECTION string

// redis
var REDIS_URI string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// server
	PORT = os.Getenv("PORT")
	// mongo
	MONGO_URI = os.Getenv("MONGO_URI")
	DATABASE = os.Getenv("DATABASE")
	USER_COLLECTION = os.Getenv("USER_COLLECTION")
	// redis
	REDIS_URI = os.Getenv("REDIS_URI")
}
