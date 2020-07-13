package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
)

var (
	Categoryrepo categoryrepo = categoryrepo{}
)

///curtesy to gorm
type categoryrepo struct{}

func (categoryRepo categoryrepo) Getconnected() (GormDB *gorm.DB, err *httperors.HttpError) {
	dbURI := "root@/store?charset=utf8&parseTime=True&loc=Local"
	GormDB, err1 := gorm.Open("mysql", dbURI)
	if err1 != nil {
		return nil, httperors.NewNotFoundError("No Mysql db connection")
	}
	GormDB.AutoMigrate(&model.Category{})
	return GormDB, nil
}
func (categoryRepo categoryrepo) DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (categoryRepo categoryrepo) Create(category *model.Category) (*model.Category, *httperors.HttpError) {
	if err := category.Validate(); err != nil {
		return nil, err
	}
	GormDB, err1 := categoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Create(&category)
	categoryRepo.DbClose(GormDB)
	return category, nil
}
func (categoryRepo categoryrepo) GetOne(id int) (*model.Category, *httperors.HttpError) {
	ok := categoryRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("category with that id does not exists!")
	}
	category := model.Category{}
	GormDB, err1 := categoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Model(&category).Where("id = ?", id).First(&category)
	categoryRepo.DbClose(GormDB)
	
	return &category, nil
}

func (categoryRepo categoryrepo) GetAll(categorys []model.Category) ([]model.Category, *httperors.HttpError) {
	GormDB, err1 := categoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	Category := model.Category{}
	GormDB.Model(&Category).Find(&categorys)
	
	categoryRepo.DbClose(GormDB)
	if len(categorys) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return categorys, nil
}

func (categoryRepo categoryrepo) Update(id int, category *model.Category) (*model.Category, *httperors.HttpError) {
	ok := categoryRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("category with that id does not exists!")
	}
	
	GormDB, err1 := categoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	Category := model.Category{}
	acategory := model.Category{}
	
	GormDB.Model(&Category).Where("id = ?", id).First(&acategory)
	if category.Name  == "" {
		category.Name = acategory.Name
	}
	if category.Title  == "" {
		category.Title = acategory.Title
	}
	if category.Description  == "" {
		category.Description = acategory.Description
	}
	GormDB.Model(&Category).Where("id = ?", id).First(&Category).Update(&acategory)
	
	categoryRepo.DbClose(GormDB)

	return category, nil
}
func (categoryRepo categoryrepo) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := categoryRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("Product with that id does not exists!")
	}
	category := model.Category{}
	GormDB, err1 := categoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&category).Where("id = ?", id).First(&category)
	GormDB.Delete(category)
	categoryRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (categoryRepo categoryrepo)ProductUserExistByid(id int) bool {
	category := model.Category{}
	GormDB, err1 := categoryRepo.Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&category, "id =?", id).RecordNotFound(){
	   return false
	}
	categoryRepo.DbClose(GormDB)
	return true
	
}