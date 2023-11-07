package main

import (
	"GoDynamoApiTemplate/src/database/migrations"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load("migration.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	err := migrations.CreateColorsTable()
	if err != nil {
		fmt.Print(err)
		return
	}
}
