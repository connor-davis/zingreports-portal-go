package main

import (
	"fmt"
	"log"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/connor-davis/zingreports-portal-go/cmd/api/http"
	"github.com/connor-davis/zingreports-portal-go/internal/environment"
	"github.com/connor-davis/zingreports-portal-go/internal/services"
	"github.com/connor-davis/zingreports-portal-go/internal/storage"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title           ZingFibre Reports Portal API
// @version         1.0.0
// @description     This is the ZingFibre Reports Portal API built with Fiber

// @contact.name   Connor Davis
// @contact.url    https://thusa.co.za
// @contact.email  connor.davis@thusa.co.za

// @host      localhost:6173
// @BasePath  /api
func main() {
	app := fiber.New(fiber.Config{
		AppName:     "ZingFibre Reports Portal API",
		Prefork:     true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173,https://reports.core.zingfibre.co.za",
		AllowMethods:     "GET,POST,PATCH,PUT,DELETE",
		AllowCredentials: true,
	}))

	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} - ${latency} ${method} ${url}\n",
	}))

	api := app.Group("/api")

	api.Get("/api-spec", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	api.Get("/api-doc", func(c *fiber.Ctx) error {
		html, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL:       fmt.Sprintf("http://localhost:%s/api/api-spec", environment.PORT),
			Theme:         scalar.ThemeDefault,
			Layout:        scalar.LayoutClassic,
			BaseServerURL: fmt.Sprintf("http://localhost:%s", environment.PORT),
			DarkMode:      true,
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.Type("html").SendString(html)
	})

	storage := storage.New()

	userService := services.NewUserService(storage)
	poiService := services.NewPoiService(storage)

	http := http.NewHttpRouter(storage, userService, poiService)

	api.Route("/", http.LoadRoutes)

	if err := app.Listen(fmt.Sprintf(":%s", environment.PORT)); err != nil {
		log.Printf("🔥 An error occured that caused the API to shutdown:\n%s", err.Error())
	}
}
