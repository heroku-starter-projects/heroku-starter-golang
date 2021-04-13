package routes

import (
	fiber "github.com/gofiber/fiber/v2"
)

type Pizza struct {
	Name     string `json:"name"`
	Calories int32  `json:"calories"`
}

type PizzaResponse struct {
	Message string `json:"message"`
	Pizza   *Pizza `json:"pizza"`
}

func GetPizza(c *fiber.Ctx) error {
	Pizza := Pizza{"Cake", 3200}

	return c.JSON(Pizza)
}

func PostPizza(c *fiber.Ctx) error {
	pizza := new(Pizza)

	if err := c.BodyParser(pizza); err != nil {
		return err
	}

	response := PizzaResponse{"Success", pizza}

	return c.JSON(response)
}
