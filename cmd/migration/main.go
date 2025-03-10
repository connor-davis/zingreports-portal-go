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

	storage.P.AutoMigrate(
		&postgres.User{},
		&postgres.Poi{},
		&postgres.Report{},
		&postgres.ReportTable{},
		&postgres.ReportTableColumn{},
		&postgres.ReportTableReference{},
		&postgres.ReportColumn{},
		&postgres.ReportFilter{},
	)

	log.Printf("✅ Finished running Postgres Migrations...")
}
