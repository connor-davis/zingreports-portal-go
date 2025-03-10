package main

import (
	"log"

	"github.com/connor-davis/zingreports-portal-go/internal/models/postgres"
	"github.com/connor-davis/zingreports-portal-go/internal/services"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
)

func main() {
	storage := storage.New()
	storage.ConnectPostgres()

	log.Printf("🔃 Running Postgres Migrations...")

	storage.Postgres.AutoMigrate(
		&postgres.User{},
		&postgres.Poi{},
		&postgres.Report{},
		&postgres.ReportTable{},
		&postgres.ReportTableColumn{},
		&postgres.ReportTableReference{},
		&postgres.ReportColumn{},
		&postgres.ReportFilter{},
	)

	userService := services.NewUserService(storage)

	user, err := userService.FindUserById("test")

	if err != nil {
		log.Printf("🔥 Failed to find user:\n%s", err.Error())
	} else {
		log.Printf("🔎 Found user:\n%v", user)
	}

	log.Printf("✅ Finished running Postgres Migrations...")
}
