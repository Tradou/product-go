package migrations

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"main/infrastructure/database"
	"main/infrastructure/models"
)

func Migrate() {
	db, err := database.GetDBConnection()

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(
		&models.Product{},
		&models.Attribute{},
		&models.AttributeProduct{},
	); err != nil {
		log.Fatalln(err)
	}
}
