package internals

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
