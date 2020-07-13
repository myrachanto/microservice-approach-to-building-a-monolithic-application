package model

import (
	"github.com/jinzhu/gorm"
)
type Transaction struct {
	Productname string 
	Productid uint
	InQty uint16 
	InPrice float64 
	InTax float64 
	InSubtotal float64 
	InDiscount float64 
	InTotal float64 
	InAmountPaid float64 
	InBalance float64 
	OutQty uint16 
	OutPrice float64 
	OutTax float64 
	OutSubtotal float64 
	OutDiscount float64 
	OutTotal float64 
	OutAmountPaid float64 
	OutBalance float64 
	gorm.Model
}