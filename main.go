package main

import (
	routes "github.com/heroku-starter-projects/heroku-starter-golang/routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.RegisterRoutes(app)

	app.Listen(":5001")
}
