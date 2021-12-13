package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) (v string) {
	v = os.Getenv("PORT")
	if v != "" {
		return
	}

	v = dotEnvVariable(key)
	return
}

func dotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
