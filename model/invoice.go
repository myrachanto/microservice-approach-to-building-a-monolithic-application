package model

import (
	"github.com/jinzhu/gorm"
)
type Invoice struct {
	Customer string 
	Description string
	InSubtotal float64 
	InDiscount float64 
	InTax float64 
	InTotal float64 
	InPaidStatus bool 
	InAmountPaid float64 
	InBalance float64 
	OutSubtotal float64 
	OutDiscount float64 
	OutTax float64 
	OutTotal float64 
	OutCanceled bool 
	OutAmountPaid float64 
	OutBalance float64 
	gorm.Model
}