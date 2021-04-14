package models

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	Name        string
	WarehouseID uint
}
