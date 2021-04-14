package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	Name        string `json:"name"`
	WarehouseID uint
}
