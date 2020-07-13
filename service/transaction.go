package service

import (
	// "fmt"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
	r "github.com/myrachanto/astore/repository"
)

var (
	Transactionservice transactionservice = transactionservice{}

) 
type transactionservice struct {
	
}

func (service transactionservice) Create(transaction *model.Transaction) (*model.Transaction, *httperors.HttpError) {
	transaction, err1 := r.Transactionrepo.Create(transaction)
	if err1 != nil {
		return nil, err1
	}
	 return transaction, nil

}
func (service transactionservice) GetOne(id int) (*model.Transaction, *httperors.HttpError) {
	transaction, err1 := r.Transactionrepo.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return transaction, nil
}

func (service transactionservice) GetAll(transactions []model.Transaction) ([]model.Transaction, *httperors.HttpError) {
	transactions, err := r.Transactionrepo.GetAll(transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (service transactionservice) Update(id int, transaction *model.Transaction) (*model.Transaction, *httperors.HttpError) {
	transaction, err1 := r.Transactionrepo.Update(id, transaction)
	if err1 != nil {
		return nil, err1
	}
	
	return transaction, nil
}
func (service transactionservice) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	
		success, failure := r.Transactionrepo.Delete(id)
		return success, failure
}
