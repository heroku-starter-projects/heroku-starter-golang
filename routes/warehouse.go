package routes

import (
	"strconv"

	fiber "github.com/gofiber/fiber/v2"
	database "github.com/heroku-starter-projects/heroku-starter-golang/database"
	models "github.com/heroku-starter-projects/heroku-starter-golang/models"
	"gorm.io/gorm/clause"
)

func GetWarehouse(c *fiber.Ctx) error {
	db := database.CreateConnection()
	defer database.CloseConnection()

	warehouseId, _ := strconv.Atoi(c.Params("id"))

	var warehouse models.Warehouse
	db.Preload(clause.Associations).First(&warehouse, warehouseId)

	return c.JSON(warehouse)
}

func PostWarehouse(c *fiber.Ctx) error {
	db := database.CreateConnection()
	defer database.CloseConnection()

	warehouse := new(models.Warehouse)

	if err := c.BodyParser(warehouse); err != nil {
		return err
	}

	db.Create(&warehouse)

	return c.JSON(warehouse)
}
