package model

import (
	"github.com/jinzhu/gorm"
)
type Cart struct {
	Name string `gorm:"not null"`
	Qty uint16 `gorm:"not null"`
	Price float64 `gorm:"not null"`
	Subtotal float64 `gorm:"not null"`
	Discount float64 
	Tax float64 
	Total float64 
	Cartstatus bool 
	Picture string 
	gorm.Model
}