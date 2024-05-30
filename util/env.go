package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DbUrl  string
	Secret string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	DbUrl = os.Getenv("DB_URL")
	Secret = os.Getenv("JWT_SECRET")

}
