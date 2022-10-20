package bootstrappers

import (
	"log"

	"github.com/joho/godotenv"
)

func BootstrapEnvironment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
