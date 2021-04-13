package routes

import fiber "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Get("/pizza", GetPizza)
	app.Post("/pizza", PostPizza)

	app.Get("/dessert", GetDessert)
	app.Post("/dessert", PostDessert)

	app.Get("/health", GetHealth)
}
