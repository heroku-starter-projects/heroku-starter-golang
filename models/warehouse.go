package models

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	Address  string
	Workers  []Worker
	Products []*Product `gorm:"many2many:inventory;"`
}
