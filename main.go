package main

import (
	_ "github.com/joho/godotenv/autoload"
	"main/endpoints"
	"main/infrastructure/migrations"
)

func main() {
	migrations.Migrate()
	endpoints.RegisterEndpoints()
}
