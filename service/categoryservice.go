package service

import (
	// "fmt"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
	r "github.com/myrachanto/astore/repository"
)

var (
	CategoryService categoryService = categoryService{}

) 
type redirectCategroy interface{
	Create(category *model.Category) (*model.Category, *httperors.HttpError)
	GetOne(id int) (*model.Category, *httperors.HttpError)
	GetAll(users []model.Category) ([]model.Category, *httperors.HttpError)
	Update(id int, category *model.Category) (*model.Category, *httperors.HttpError)
	Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError)
}


type categoryService struct {
	
}

func (service categoryService) Create(category *model.Category) (*model.Category, *httperors.HttpError) {
	if err := category.Validate(); err != nil {
		return nil, err
	}	
	category, err1 := r.Categoryrepo.Create(category)
	if err1 != nil {
		return nil, err1
	}
	 return category, nil

}
func (service categoryService) GetOne(id int) (*model.Category, *httperors.HttpError) {
	category, err1 := r.Categoryrepo.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return category, nil
}

func (service categoryService) GetAll(categorys []model.Category) ([]model.Category, *httperors.HttpError) {
	categorys, err := r.Categoryrepo.GetAll(categorys)
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (service categoryService) Update(id int, category *model.Category) (*model.Category, *httperors.HttpError) {
	category, err1 := r.Categoryrepo.Update(id, category)
	if err1 != nil {
		return nil, err1
	}
	
	return category, nil
}
func (service categoryService) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	
		success, failure := r.Categoryrepo.Delete(id)
		return success, failure
}
