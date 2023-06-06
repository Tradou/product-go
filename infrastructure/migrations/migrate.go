package migrations

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main/infrastructure/models"
	"os"
)

func Migrate() {
	db, err := gorm.Open(mysql.Open(getDsn()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalln(err)
	}
}

func getDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))
}
