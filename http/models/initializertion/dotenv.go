package initializertion

import (
	"log"

	"github.com/joho/godotenv"
)

func Dotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
