package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
)

var (
	Majorcategoryrepo majorcategoryrepo = majorcategoryrepo{}
)

///curtesy to gorm
type majorcategoryrepo struct{}

func (majorcategoryRepo majorcategoryrepo) Getconnected() (GormDB *gorm.DB, err *httperors.HttpError) {
	dbURI := "root@/store?charset=utf8&parseTime=True&loc=Local"
	GormDB, err1 := gorm.Open("mysql", dbURI)
	if err1 != nil {
		return nil, httperors.NewNotFoundError("No Mysql db connection")
	}
	GormDB.AutoMigrate(&model.Majorcategory{})
	return GormDB, nil
}
func (majorcategoryRepo majorcategoryrepo) DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (majorcategoryRepo majorcategoryrepo) Create(majorcategory *model.Majorcategory) (*model.Majorcategory, *httperors.HttpError) {
	if err := majorcategory.Validate(); err != nil {
		return nil, err
	}
	GormDB, err1 := majorcategoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Create(&majorcategory)
	majorcategoryRepo.DbClose(GormDB)
	return majorcategory, nil
}
func (majorcategoryRepo majorcategoryrepo) GetOne(id int) (*model.Majorcategory, *httperors.HttpError) {
	ok := majorcategoryRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("majorcategory with that id does not exists!")
	}
	majorcategory := model.Majorcategory{}
	GormDB, err1 := majorcategoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Model(&majorcategory).Where("id = ?", id).First(&majorcategory)
	majorcategoryRepo.DbClose(GormDB)
	
	return &majorcategory, nil
}

func (majorcategoryRepo majorcategoryrepo) GetAll(majorcategorys []model.Majorcategory) ([]model.Majorcategory, *httperors.HttpError) {
	GormDB, err1 := majorcategoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	majorcategory := model.Majorcategory{}
	GormDB.Model(&majorcategory).Find(&majorcategorys)
	
	majorcategoryRepo.DbClose(GormDB)
	if len(majorcategorys) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return majorcategorys, nil
}

func (majorcategoryRepo majorcategoryrepo) Update(id int, majorcategory *model.Majorcategory) (*model.Majorcategory, *httperors.HttpError) {
	ok := majorcategoryRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("majorcategory with that id does not exists!")
	}
	
	GormDB, err1 := majorcategoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	amajorcategory := model.Majorcategory{}
	
	GormDB.Model(&amajorcategory).Where("id = ?", id).First(&amajorcategory)
	if majorcategory.Name  == "" {
		majorcategory.Name = amajorcategory.Name
	}
	if majorcategory.Title  == "" {
		majorcategory.Title = amajorcategory.Title
	}
	if majorcategory.Description  == "" {
		majorcategory.Description = amajorcategory.Description
	}
	GormDB.Model(&majorcategory).Where("id = ?", id).First(&majorcategory).Update(&amajorcategory)
	
	majorcategoryRepo.DbClose(GormDB)

	return majorcategory, nil
}
func (majorcategoryRepo majorcategoryrepo) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := majorcategoryRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("Product with that id does not exists!")
	}
	majorcategory := model.Majorcategory{}
	GormDB, err1 := majorcategoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&majorcategory).Where("id = ?", id).First(&majorcategory)
	GormDB.Delete(majorcategory)
	majorcategoryRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (majorcategoryRepo majorcategoryrepo)ProductUserExistByid(id int) bool {
	majorcategory := model.Majorcategory{}
	GormDB, err1 := majorcategoryRepo.Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&majorcategory, "id =?", id).RecordNotFound(){
	   return false
	}
	majorcategoryRepo.DbClose(GormDB)
	return true
	
}