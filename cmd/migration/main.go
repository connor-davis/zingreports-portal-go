package main

import (
	"log"

	"github.com/connor-davis/zingreports-portal-go/internal/models/postgres"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

func main() {
	storage := storage.New()
	storage.ConnectPostgres()

	log.Printf("🔃 Running Postgres Migrations...")

	storage.P.AutoMigrate(&postgres.User{}, &postgres.Poi{})

	log.Printf("✅ Finished running Postgres Migrations...")
}
