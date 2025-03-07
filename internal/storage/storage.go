package storage

import (
	"log"

	"github.com/connor-davis/zingreports-portal-go/internal/environment"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Storage struct {
	d *gorm.DB
	r *gorm.DB
	e *gorm.DB
}

func New() *Storage {
	d, err := gorm.Open(postgres.Open(string(environment.POSTGRES_DSN)))

	if err != nil {
		log.Fatalf("Error connecting to Postgres with Gorm:\n%v", err)
	}

	e, err := gorm.Open(mysql.Open(string(environment.MYSQL_DSN)))

	if err != nil {
		log.Fatalf("Error connecting to MySQL with Gorm:\n%v", err)
	}

	r, err := gorm.Open(sqlserver.Open(string(environment.SQLSERVER_DSN)))

	if err != nil {
		log.Fatalf("Error connecting to SQLServer with Gorm:\n%v", err)
	}

	return &Storage{
		d: d,
		e: e,
		r: r,
	}
}
