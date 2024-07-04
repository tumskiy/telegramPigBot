package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("not find .env file")
	}
}

func ParseEnv(key string) string {
	return os.Getenv(key)
}
