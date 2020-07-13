package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
)

var (
	Transactionrepo transactionrepo = transactionrepo{}
)

///curtesy to gorm
type transactionrepo struct{}

func (transactionRepo transactionrepo) Getconnected() (GormDB *gorm.DB, err *httperors.HttpError) {
	dbURI := "root@/store?charset=utf8&parseTime=True&loc=Local"
	GormDB, err1 := gorm.Open("mysql", dbURI)
	if err1 != nil {
		return nil, httperors.NewNotFoundError("No Mysql db connection")
	}
	GormDB.AutoMigrate(&model.Transaction{})
	return GormDB, nil
}
func (transactionRepo transactionrepo) DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (transactionRepo transactionrepo) Create(transaction *model.Transaction) (*model.Transaction, *httperors.HttpError) {
	GormDB, err1 := transactionRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Create(&transaction)
	transactionRepo.DbClose(GormDB)
	return transaction, nil
}
func (transactionRepo transactionrepo) GetOne(id int) (*model.Transaction, *httperors.HttpError) {
	ok := transactionRepo.transactionUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("transaction with that id does not exists!")
	}
	transaction := model.Transaction{}
	GormDB, err1 := transactionRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Model(&transaction).Where("id = ?", id).First(&transaction)
	transactionRepo.DbClose(GormDB)
	
	return &transaction, nil
}

func (transactionRepo transactionrepo) GetAll(transactions []model.Transaction) ([]model.Transaction, *httperors.HttpError) {
	GormDB, err1 := transactionRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	transaction := model.Transaction{}
	GormDB.Model(&transaction).Find(&transactions)
	
	transactionRepo.DbClose(GormDB)
	if len(transactions) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return transactions, nil
}

func (transactionRepo transactionrepo) Update(id int, transaction *model.Transaction) (*model.Transaction, *httperors.HttpError) {
	ok := transactionRepo.transactionUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("transaction with that id does not exists!")
	}
	
	GormDB, err1 := transactionRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	atransaction := model.Transaction{}
	
	GormDB.Model(&atransaction).Where("id = ?", id).First(&atransaction)
	// if transaction.Name  == "" {
	// 	transaction.Name = atransaction.Name
	// }
	// if transaction.Qty  == 0 {
	// 	transaction.Qty = atransaction.Qty
	// }
	// if transaction.Price  == 0 {
	// 	transaction.Price = atransaction.Price
	// }
	
	// if transaction.Discount  == 0 {
	// 	transaction.Discount = atransaction.Discount
	// }
	// if transaction.Tax  == 0 {
	// 	transaction.Tax = atransaction.Tax
	// }
	GormDB.Model(&transaction).Where("id = ?", id).First(&transaction).Update(&atransaction)
	
	transactionRepo.DbClose(GormDB)

	return transaction, nil
}
func (transactionRepo transactionrepo) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := transactionRepo.transactionUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("transaction with that id does not exists!")
	}
	transaction := model.Transaction{}
	GormDB, err1 := transactionRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&transaction).Where("id = ?", id).First(&transaction)
	GormDB.Delete(transaction)
	transactionRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (transactionRepo transactionrepo)transactionUserExistByid(id int) bool {
	transaction := model.Transaction{}
	GormDB, err1 := transactionRepo.Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&transaction, "id =?", id).RecordNotFound(){
	   return false
	}
	transactionRepo.DbClose(GormDB)
	return true
	
}