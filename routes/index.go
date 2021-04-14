package routes

import fiber "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Get("/pizza", GetPizza)
	app.Post("/pizza", PostPizza)

	app.Get("/warehouse/:id", GetWarehouse)
	app.Post("/warehouse", PostWarehouse)

	app.Get("/health", GetHealth)
}
