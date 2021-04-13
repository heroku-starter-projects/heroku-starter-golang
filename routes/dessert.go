package routes

import fiber "github.com/gofiber/fiber/v2"

type Dessert struct {
	Name     string `json:"name"`
	Diabetic bool   `json:"diabetic"`
}

type DessertResponse struct {
	Message string   `json:"message"`
	Dessert *Dessert `json:"dessert"`
}

func GetDessert(c *fiber.Ctx) error {
	dessert := Dessert{"Cake", false}

	return c.JSON(dessert)
}

func PostDessert(c *fiber.Ctx) error {
	dessert := new(Dessert)

	if err := c.BodyParser(dessert); err != nil {
		return err
	}

	response := DessertResponse{"Success", dessert}

	return c.JSON(response)
}
