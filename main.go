package main

import (
	"log"
	"time"

	"github.com/georgekakarlis/Go-React.git/server/database"
	"github.com/georgekakarlis/Go-React.git/server/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/utils"
)

func main() {
	app := fiber.New()

	app.Use(helmet.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
	}))

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
