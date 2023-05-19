package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvByName(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(name)
}
