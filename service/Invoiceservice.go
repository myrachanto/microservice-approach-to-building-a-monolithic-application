package service

import (
	// "fmt"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
	r "github.com/myrachanto/astore/repository"
)

var (
	Invoiceservice invoiceservice = invoiceservice{}

) 
type invoiceservice struct {
	
}

func (service invoiceservice) Create(invoice *model.Invoice) (*model.Invoice, *httperors.HttpError) {
	invoice, err1 := r.Invoicerepo.Create(invoice)
	if err1 != nil {
		return nil, err1
	}
	 return invoice, nil

}
func (service invoiceservice) GetOne(id int) (*model.Invoice, *httperors.HttpError) {
	invoice, err1 := r.Invoicerepo.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return invoice, nil
}

func (service invoiceservice) GetAll(invoices []model.Invoice) ([]model.Invoice, *httperors.HttpError) {
	invoices, err := r.Invoicerepo.GetAll(invoices)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (service invoiceservice) Update(id int, invoice *model.Invoice) (*model.Invoice, *httperors.HttpError) {
	invoice, err1 := r.Invoicerepo.Update(id, invoice)
	if err1 != nil {
		return nil, err1
	}
	
	return invoice, nil
}
func (service invoiceservice) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	
		success, failure := r.Invoicerepo.Delete(id)
		return success, failure
}
