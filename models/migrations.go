package models

import (
	"fmt"

	database "github.com/heroku-starter-projects/heroku-starter-golang/database"
	"gorm.io/gorm"
)

func testConnection(db *gorm.DB) {
	fmt.Println(db.Migrator().CurrentDatabase())

	var result string
	err := db.Raw("SELECT 1+1 AS result").Scan(&result)

	if result != "2" {
		fmt.Println(result)
		panic(err)
	}
}

func RunMigrations() {
	db := database.CreateConnection()

	// Test DB connection
	testConnection(db)

	err := db.AutoMigrate(
		&Product{},
		&Warehouse{},
		&Worker{},
	)

	if err != nil {
		fmt.Println(err)
	}

}
