package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
)

var (
	Productrepo productrepo = productrepo{}
)

///curtesy to gorm
type productrepo struct{}

func (productRepo productrepo) Getconnected() (GormDB *gorm.DB, err *httperors.HttpError) {
	dbURI := "root@/store?charset=utf8&parseTime=True&loc=Local"
	GormDB, err1 := gorm.Open("mysql", dbURI)
	if err1 != nil {
		return nil, httperors.NewNotFoundError("No Mysql db connection")
	}
	GormDB.AutoMigrate(&model.Product{})
	return GormDB, nil
}
func (productRepo productrepo) DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (productRepo productrepo) Create(product *model.Product) (*model.Product, *httperors.HttpError) {
	if err := product.Validate(); err != nil {
		return nil, err
	}
	GormDB, err1 := productRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Create(&product)
	productRepo.DbClose(GormDB)
	return product, nil
}
func (productRepo productrepo) GetOne(id int) (*model.Product, *httperors.HttpError) {
	ok := productRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("product with that id does not exists!")
	}
	product := model.Product{}
	GormDB, err1 := productRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Model(&product).Where("id = ?", id).First(&product)
	productRepo.DbClose(GormDB)
	
	return &product, nil
}

func (productRepo productrepo) GetAll(products []model.Product) ([]model.Product, *httperors.HttpError) {
	GormDB, err1 := productRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	product := model.Product{}
	GormDB.Model(&product).Find(&products)
	
	productRepo.DbClose(GormDB)
	if len(products) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return products, nil
}

func (productRepo productrepo) Update(id int, product *model.Product) (*model.Product, *httperors.HttpError) {
	ok := productRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("product with that id does not exists!")
	}
	
	GormDB, err1 := productRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	aproduct := model.Product{}
	
	GormDB.Model(&aproduct).Where("id = ?", id).First(&aproduct)
	if product.Name  == "" {
		product.Name = aproduct.Name
	}
	if product.Title  == "" {
		product.Title = aproduct.Title
	}
	if product.Description  == "" {
		product.Description = aproduct.Description
	}
	GormDB.Model(&product).Where("id = ?", id).First(&product).Update(&aproduct)
	
	productRepo.DbClose(GormDB)

	return product, nil
}
func (productRepo productrepo) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := productRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("Product with that id does not exists!")
	}
	product := model.Product{}
	GormDB, err1 := productRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&product).Where("id = ?", id).First(&product)
	GormDB.Delete(product)
	productRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (productRepo productrepo)ProductUserExistByid(id int) bool {
	product := model.Product{}
	GormDB, err1 := productRepo.Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&product, "id =?", id).RecordNotFound(){
	   return false
	}
	productRepo.DbClose(GormDB)
	return true
	
}