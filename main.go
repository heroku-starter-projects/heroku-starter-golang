package main

import (
	"fmt"
	"os"

	models "github.com/heroku-starter-projects/heroku-starter-golang/models"
	routes "github.com/heroku-starter-projects/heroku-starter-golang/routes"

	fiber "github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") == "development" {
		godotenv.Load()
	}

	app := fiber.New()

	routes.RegisterRoutes(app)

	fmt.Println("Running Migrations")
	models.RunMigrations()
	fmt.Println("Migrations Ran Successfully")

	app.Listen(":5001")
}
