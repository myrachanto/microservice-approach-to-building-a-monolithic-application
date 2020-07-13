package model

import (
	"github.com/jinzhu/gorm"
	"github.com/myrachanto/astore/httperors"
)


type Product struct {
	Name string `gorm:"not null"`
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
	Subcategory Subcategory `gorm:"foreignKey:UserID; not null"`
	SubcategoryID uint `gorm:"not null"`
	Picture string 
	gorm.Model
}
func (product Product) Validate() *httperors.HttpError{ 
	if product.Name == "" && len(product.Name) < 3 {
		return httperors.NewNotFoundError("Invalid Name")
	}
	if product.Title == "" && len(product.Title) < 3 {
		return httperors.NewNotFoundError("Invalid Title")
	}
	
	if product.Description == "" && len(product.Description) < 10 {
		return httperors.NewNotFoundError("Invalid description")
	}
	return nil
}