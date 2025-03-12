package storage

import (
	"log"

	"github.com/connor-davis/zingreports-portal-go/internal/environment"
	"github.com/gofiber/fiber/v2/middleware/session"
	fiberPg "github.com/gofiber/storage/postgres/v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Storage struct {
	Postgres *gorm.DB
	Radius   *gorm.DB
	Ekasi    *gorm.DB
	Sessions *session.Store
}

func New() *Storage {
	sessions := session.New(session.Config{
		Storage: fiberPg.New(fiberPg.Config{
			Table:         "sessions",
			ConnectionURI: string(environment.POSTGRES_DSN),
		}),
		KeyLookup:         "cookie:zing_session",
		CookieDomain:      string(environment.COOKIE_DOMAIN),
		CookiePath:        "/",
		CookieSecure:      true,
		CookieSameSite:    "Strict",
		CookieSessionOnly: true,
		CookieHTTPOnly:    true,
	})

	return &Storage{
		Sessions: sessions,
	}
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
