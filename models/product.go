package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code       string
	Price      uint
	Warehouses []*Warehouse `gorm:"many2many:inventory;"`
}
