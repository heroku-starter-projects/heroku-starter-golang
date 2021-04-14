package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code       string       `json:"code"`
	Price      uint         `json:"price"`
	Warehouses []*Warehouse `gorm:"many2many:inventory;" json:"warehouses"`
}
