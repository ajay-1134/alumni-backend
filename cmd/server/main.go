package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ajay-1134/alumni-backend/internal/app"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, relying on system env variables")
	}

	builder := app.NewBuilder()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	router := builder.Build(dsn)
	router.Run(":8080")
}
