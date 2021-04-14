package models

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	Address  string     `json:"address"`
	Workers  []Worker   `json:"workers"`
	Products []*Product `gorm:"many2many:inventory;" json:"products"`
}
