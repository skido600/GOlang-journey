package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	// PGHOST=localhost
	// PGUSER=postgres
	// PGPORT=5432
	// PGPASSWORD=12345
	// PGDATABASE=demopost
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(("failed to connect"))
	}
}
