package bootstrappers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func BootstrapDatabases() {
	dsn := fmt.Sprintf(
		"host=%s user=%s  password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_USER"),
		os.Getenv("PSQL_PASSWORD"),
		os.Getenv("PSQL_DATABASE"),
	)

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal("Database bootstrapping failed in 'server'", error)
	}
	DB = db
}
