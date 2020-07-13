package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
)

var (
	Invoicerepo invoicerepo = invoicerepo{}
)

///curtesy to gorm
type invoicerepo struct{}

func (invoiceRepo invoicerepo) Getconnected() (GormDB *gorm.DB, err *httperors.HttpError) {
	dbURI := "root@/store?charset=utf8&parseTime=True&loc=Local"
	GormDB, err1 := gorm.Open("mysql", dbURI)
	if err1 != nil {
		return nil, httperors.NewNotFoundError("No Mysql db connection")
	}
	GormDB.AutoMigrate(&model.Invoice{})
	return GormDB, nil
}
func (invoiceRepo invoicerepo) DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (invoiceRepo invoicerepo) Create(invoice *model.Invoice) (*model.Invoice, *httperors.HttpError) {
	GormDB, err1 := invoiceRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Create(&invoice)
	invoiceRepo.DbClose(GormDB)
	return invoice, nil
}
func (invoiceRepo invoicerepo) GetOne(id int) (*model.Invoice, *httperors.HttpError) {
	ok := invoiceRepo.invoiceUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("invoice with that id does not exists!")
	}
	invoice := model.Invoice{}
	GormDB, err1 := invoiceRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Model(&invoice).Where("id = ?", id).First(&invoice)
	invoiceRepo.DbClose(GormDB)
	
	return &invoice, nil
}

func (invoiceRepo invoicerepo) GetAll(invoices []model.Invoice) ([]model.Invoice, *httperors.HttpError) {
	GormDB, err1 := invoiceRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	invoice := model.Invoice{}
	GormDB.Model(&invoice).Find(&invoices)
	
	invoiceRepo.DbClose(GormDB)
	if len(invoices) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return invoices, nil
}

func (invoiceRepo invoicerepo) Update(id int, invoice *model.Invoice) (*model.Invoice, *httperors.HttpError) {
	ok := invoiceRepo.invoiceUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("invoice with that id does not exists!")
	}
	
	GormDB, err1 := invoiceRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	ainvoice := model.Invoice{}
	
	GormDB.Model(&ainvoice).Where("id = ?", id).First(&ainvoice)
	// if invoice.Customer  == "" {
	// 	invoice.Customer = ainvoice.Customer
	// }
	// if invoice.Description  == "" {
	// 	invoice.Description = ainvoice.Description
	// }
	// if invoice.Subtotal  == 0 {
	// 	invoice.Subtotal = ainvoice.Subtotal
	// }
	// if invoice.Discount  == 0 {
	// 	invoice.Discount = ainvoice.Discount
	// }	
	// if invoice.AmountPaid  == 0 {
	// 	invoice.AmountPaid = ainvoice.AmountPaid
	// }
	GormDB.Model(&invoice).Where("id = ?", id).First(&invoice).Update(&ainvoice)
	
	invoiceRepo.DbClose(GormDB)

	return invoice, nil
}
func (invoiceRepo invoicerepo) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := invoiceRepo.invoiceUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("invoice with that id does not exists!")
	}
	invoice := model.Invoice{}
	GormDB, err1 := invoiceRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&invoice).Where("id = ?", id).First(&invoice)
	GormDB.Delete(invoice)
	invoiceRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (invoiceRepo invoicerepo)invoiceUserExistByid(id int) bool {
	invoice := model.Invoice{}
	GormDB, err1 := invoiceRepo.Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&invoice, "id =?", id).RecordNotFound(){
	   return false
	}
	invoiceRepo.DbClose(GormDB)
	return true
	
}