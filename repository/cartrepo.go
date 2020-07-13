package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
)

var (
	Cartrepo cartrepo = cartrepo{}
)

///curtesy to gorm
type cartrepo struct{}

func (cartRepo cartrepo) Getconnected() (GormDB *gorm.DB, err *httperors.HttpError) {
	dbURI := "root@/store?charset=utf8&parseTime=True&loc=Local"
	GormDB, err1 := gorm.Open("mysql", dbURI)
	if err1 != nil {
		return nil, httperors.NewNotFoundError("No Mysql db connection")
	}
	GormDB.AutoMigrate(&model.Cart{})
	return GormDB, nil
}
func (cartRepo cartrepo) DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (cartRepo cartrepo) Create(cart *model.Cart) (*model.Cart, *httperors.HttpError) {
	GormDB, err1 := cartRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Create(&cart)
	cartRepo.DbClose(GormDB)
	return cart, nil
}
func (cartRepo cartrepo) GetOne(id int) (*model.Cart, *httperors.HttpError) {
	ok := cartRepo.cartUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("cart with that id does not exists!")
	}
	cart := model.Cart{}
	GormDB, err1 := cartRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Model(&cart).Where("id = ?", id).First(&cart)
	cartRepo.DbClose(GormDB)
	
	return &cart, nil
}

func (cartRepo cartrepo) GetAll(carts []model.Cart) ([]model.Cart, *httperors.HttpError) {
	GormDB, err1 := cartRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	cart := model.Cart{}
	GormDB.Model(&cart).Find(&carts)
	
	cartRepo.DbClose(GormDB)
	if len(carts) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return carts, nil
}

func (cartRepo cartrepo) Update(id int, cart *model.Cart) (*model.Cart, *httperors.HttpError) {
	ok := cartRepo.cartUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("cart with that id does not exists!")
	}
	
	GormDB, err1 := cartRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	acart := model.Cart{}
	
	GormDB.Model(&acart).Where("id = ?", id).First(&acart)
	if cart.Name  == "" {
		cart.Name = acart.Name
	}
	if cart.Qty  == 0 {
		cart.Qty = acart.Qty
	}
	if cart.Price  == 0 {
		cart.Price = acart.Price
	}
	
	if cart.Discount  == 0 {
		cart.Discount = acart.Discount
	}
	if cart.Tax  == 0 {
		cart.Tax = acart.Tax
	}
	GormDB.Model(&cart).Where("id = ?", id).First(&cart).Update(&acart)
	
	cartRepo.DbClose(GormDB)

	return cart, nil
}
func (cartRepo cartrepo) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := cartRepo.cartUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("cart with that id does not exists!")
	}
	cart := model.Cart{}
	GormDB, err1 := cartRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&cart).Where("id = ?", id).First(&cart)
	GormDB.Delete(cart)
	cartRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (cartRepo cartrepo)cartUserExistByid(id int) bool {
	cart := model.Cart{}
	GormDB, err1 := cartRepo.Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&cart, "id =?", id).RecordNotFound(){
	   return false
	}
	cartRepo.DbClose(GormDB)
	return true
	
}