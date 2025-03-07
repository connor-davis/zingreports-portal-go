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
	P *gorm.DB
	R *gorm.DB
	E *gorm.DB
}

func New() *Storage {
	d, err := gorm.Open(postgres.Open(string(environment.POSTGRES_DSN)))

	if err != nil {
		log.Printf("🔥 Error connecting to Postgres with Gorm:\n%v", err.Error())
	}

	e, err := gorm.Open(mysql.Open(string(environment.MYSQL_DSN)))

	if err != nil {
		log.Printf("🔥 Error connecting to MySQL with Gorm:\n%v", err.Error())
	}

	r, err := gorm.Open(sqlserver.Open(string(environment.SQLSERVER_DSN)))

	if err != nil {
		log.Printf("🔥 Error connecting to SQLServer with Gorm:\n%v", err.Error())
	}

	return &Storage{
		P: d,
		E: e,
		R: r,
	}
}
