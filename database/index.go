package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var sqlDB *sql.DB

func setupDbPool() {
	// https://gorm.io/docs/generic_interface.html

	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	SSL_MODE := os.Getenv("SSL_MODE")

	// https://stackoverflow.com/a/30787336/1217998
	// dsn := "postgres://root:password@0.0.0.0:5432/postgres?sslmode=disable"
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_NAME, SSL_MODE)

	if sqlDB == nil {
		var err error
		sqlDB, err = sql.Open("postgres", dsn)
		if err != nil {
			panic(err)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

}

func CreateConnection() *gorm.DB {
	// Create a SQL connection if it does not exist
	setupDbPool()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func CloseConnection() error {
	err := sqlDB.Close()
	sqlDB = nil

	return err
}
