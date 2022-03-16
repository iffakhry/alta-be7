package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
}
