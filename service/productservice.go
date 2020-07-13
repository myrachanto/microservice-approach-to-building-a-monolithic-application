package service

import (
	// "fmt"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
	r "github.com/myrachanto/astore/repository"
)

var (
	Productservice productservice = productservice{}

) 
type productservice struct {
	
}

func (service productservice) Create(product *model.Product) (*model.Product, *httperors.HttpError) {
	if err := product.Validate(); err != nil {
		return nil, err
	}	
	product, err1 := r.Productrepo.Create(product)
	if err1 != nil {
		return nil, err1
	}
	 return product, nil

}
func (service productservice) GetOne(id int) (*model.Product, *httperors.HttpError) {
	product, err1 := r.Productrepo.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return product, nil
}

func (service productservice) GetAll(products []model.Product) ([]model.Product, *httperors.HttpError) {
	products, err := r.Productrepo.GetAll(products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (service productservice) Update(id int, product *model.Product) (*model.Product, *httperors.HttpError) {
	product, err1 := r.Productrepo.Update(id, product)
	if err1 != nil {
		return nil, err1
	}
	
	return product, nil
}
func (service productservice) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	
		success, failure := r.Productrepo.Delete(id)
		return success, failure
}
