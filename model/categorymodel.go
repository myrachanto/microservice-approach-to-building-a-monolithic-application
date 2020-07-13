package model

import (
	"github.com/jinzhu/gorm"
	"github.com/myrachanto/astore/httperors"
)

type Category struct {
	Name string `gorm:"not null"`
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
	Majorcategory Majorcategory `gorm:"foreignKey:UserID; not null"`
	MajorcategoryID uint `json:"userid"`
	Subcategory []Subcategory
	gorm.Model
}
func (category Category) Validate() *httperors.HttpError{ 
	if category.Name == "" && len(category.Name) < 3 {
		return httperors.NewNotFoundError("Invalid Name")
	}
	if category.Title == "" && len(category.Title) < 3 {
		return httperors.NewNotFoundError("Invalid Title")
	}
	
	if category.Description == "" && len(category.Description) < 10 {
		return httperors.NewNotFoundError("Invalid description")
	}
	return nil
}