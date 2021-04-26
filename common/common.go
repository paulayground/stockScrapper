package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// get environment value
func GetEnv(str string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv(str)
}
