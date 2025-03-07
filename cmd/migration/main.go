package main

import (
	"log"

	"github.com/connor-davis/zingreports-portal-go/internal/models/postgres"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

func main() {
	storage := storage.New()

	log.Printf("🔃 Running Postgres Migrations...")

	storage.P.AutoMigrate(&postgres.User{})

	log.Printf("✅ Finished running Postgres Migrations...")
}
