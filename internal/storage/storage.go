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
	Postgres *gorm.DB
	Radius   *gorm.DB
	Ekasi    *gorm.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) ConnectPostgres() {
	d, err := gorm.Open(postgres.Open(string(environment.POSTGRES_DSN)))

	if err != nil {
		log.Printf("🔥 Error connecting to Postgres with Gorm:\n%v", err.Error())
	} else {
		log.Printf("✅ Connected to Postgres with Gorm.")
	}

	s.Postgres = d
}

func (s *Storage) ConnectMySQL() {
	d, err := gorm.Open(mysql.Open(string(environment.MYSQL_DSN)))

	if err != nil {
		log.Printf("🔥 Error connecting to MySQL with Gorm:\n%v", err.Error())
	} else {
		log.Printf("✅ Connected to MySQL with Gorm.")
	}

	s.Radius = d
}

func (s *Storage) ConnectSQLServer() {
	d, err := gorm.Open(sqlserver.Open(string(environment.SQLSERVER_DSN)))

	if err != nil {
		log.Printf("🔥 Error connecting to SQLServer with Gorm:\n%v", err.Error())
	} else {
		log.Printf("✅ Connected to SQLServer with Gorm.")
	}

	s.Ekasi = d
}
