package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
)

var (
	Subcategoryrepo subcategoryrepo = subcategoryrepo{}
)

///curtesy to gorm
type subcategoryrepo struct{}

func (subcategoryRepo subcategoryrepo) Getconnected() (GormDB *gorm.DB, err *httperors.HttpError) {
	dbURI := "root@/store?charset=utf8&parseTime=True&loc=Local"
	GormDB, err1 := gorm.Open("mysql", dbURI)
	if err1 != nil {
		return nil, httperors.NewNotFoundError("No Mysql db connection")
	}
	GormDB.AutoMigrate(&model.Subcategory{})
	return GormDB, nil
}
func (subcategoryRepo subcategoryrepo) DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (subcategoryRepo subcategoryrepo) Create(subcategory *model.Subcategory) (*model.Subcategory, *httperors.HttpError) {
	if err := subcategory.Validate(); err != nil {
		return nil, err
	}
	GormDB, err1 := subcategoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Create(&subcategory)
	subcategoryRepo.DbClose(GormDB)
	return subcategory, nil
}
func (subcategoryRepo subcategoryrepo) GetOne(id int) (*model.Subcategory, *httperors.HttpError) {
	ok := subcategoryRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("subcategory with that id does not exists!")
	}
	subcategory := model.Subcategory{}
	GormDB, err1 := subcategoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Model(&subcategory).Where("id = ?", id).First(&subcategory)
	subcategoryRepo.DbClose(GormDB)
	
	return &subcategory, nil
}

func (subcategoryRepo subcategoryrepo) GetAll(subcategorys []model.Subcategory) ([]model.Subcategory, *httperors.HttpError) {
	GormDB, err1 := subcategoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	subcategory := model.Subcategory{}
	GormDB.Model(&subcategory).Find(&subcategorys)
	
	subcategoryRepo.DbClose(GormDB)
	if len(subcategorys) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return subcategorys, nil
}

func (subcategoryRepo subcategoryrepo) Update(id int, subcategory *model.Subcategory) (*model.Subcategory, *httperors.HttpError) {
	ok := subcategoryRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("subcategory with that id does not exists!")
	}
	
	GormDB, err1 := subcategoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	asubcategory := model.Subcategory{}
	
	GormDB.Model(&asubcategory).Where("id = ?", id).First(&asubcategory)
	if subcategory.Name  == "" {
		subcategory.Name = asubcategory.Name
	}
	if subcategory.Title  == "" {
		subcategory.Title = asubcategory.Title
	}
	if subcategory.Description  == "" {
		subcategory.Description = asubcategory.Description
	}
	GormDB.Model(&subcategory).Where("id = ?", id).First(&subcategory).Update(&asubcategory)
	
	subcategoryRepo.DbClose(GormDB)

	return subcategory, nil
}
func (subcategoryRepo subcategoryrepo) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := subcategoryRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("Product with that id does not exists!")
	}
	subcategory := model.Subcategory{}
	GormDB, err1 := subcategoryRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&subcategory).Where("id = ?", id).First(&subcategory)
	GormDB.Delete(subcategory)
	subcategoryRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (subcategoryRepo subcategoryrepo)ProductUserExistByid(id int) bool {
	subcategory := model.Subcategory{}
	GormDB, err1 := subcategoryRepo.Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&subcategory, "id =?", id).RecordNotFound(){
	   return false
	}
	subcategoryRepo.DbClose(GormDB)
	return true
	
}