package main

import (
	"log"

	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

func main() {
	storage := storage.New()

	log.Printf("🔃 Running Postgres Migrations...")

	storage.P.AutoMigrate()

	log.Printf("✅ Finished running Postgres Migrations...")
}
