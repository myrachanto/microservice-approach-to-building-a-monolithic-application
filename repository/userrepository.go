package repository

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/myrachanto/astore/httperors"
	"github.com/myrachanto/astore/model"
)

var (
	Userrepo userrepo = userrepo{}
)

///curtesy to gorm
type userrepo struct{}
func (userRepo userrepo)  Getconnected() (GormDB *gorm.DB, err *httperors.HttpError) {
	dbURI := "root@/store?charset=utf8&parseTime=True&loc=Local"
	GormDB, err1 := gorm.Open("mysql", dbURI)
	if err1 != nil {
		return nil, httperors.NewNotFoundError("No Mysql db connection")
	}

	GormDB.AutoMigrate(&model.User{})
	GormDB.AutoMigrate(&model.Auth{})
	return GormDB, nil
}
func (userRepo userrepo) DbClose(GormDB *gorm.DB) {
	defer GormDB.Close()
}
func (userRepo userrepo) Create(user *model.User) (*model.User, *httperors.HttpError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	ok, err1 := user.ValidatePassword(user.Password)
	if !ok {
		return nil, err1
	}
	ok = user.ValidateEmail(user.Email)
	if !ok {
		return nil, httperors.NewNotFoundError("Your email format is wrong!")
	}
	ok = userRepo.UserExist(user.Email)
	if ok {
		return nil, httperors.NewNotFoundError("Your email already exists!")
	}
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return nil, err2
	}
	user.Password = hashpassword
	//  encyKey := Enkey()
	// // p := support.Encrypt([]byte(user.Password), encyKey)
	// // user.Password = string(p)
	// user.Password = support.Hash(encyKey,user.Password)
	GormDB, err1 := userRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Create(&user)
	userRepo.DbClose(GormDB)
	return user, nil
}
func (userRepo userrepo) Login(auser *model.LoginUser) (*model.Auth, *httperors.HttpError) {
	if err := auser.Validate(); err != nil {
		return nil, err
	}
	ok := userRepo.UserExist(auser.Email)
	if !ok {
		return nil, httperors.NewNotFoundError("Your email does not exists!")
	}
	GormDB, err1 := userRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	user := model.User{}
	GormDB.Model(&user).Where("email = ?", auser.Email).First(&user)
	ok = user.Compare(auser.Password, user.Password)
	if !ok {
		return nil, httperors.NewNotFoundError("wrong email password combo!")
	}
	tk := &model.Token{
		UserID: user.ID,
		UserName:   user.UName,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: model.ExpiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading key")
	}
	encyKey := os.Getenv("EncryptionKey")
	tokenString, error := token.SignedString([]byte(encyKey))
	if error != nil {
		fmt.Println(error)
	}
	
	auth := &model.Auth{UserID:user.ID, Token:tokenString}
	GormDB.Create(&auth)
	userRepo.DbClose(GormDB)
	
	return auth, nil
}
func (userRepo userrepo) Logout(token string) (*httperors.HttpError) {
	auth := model.Auth{}
	GormDB, err1 := userRepo.Getconnected()
	if err1 != nil {
		return err1
	}
	if GormDB.First(&auth, "token =?", token).RecordNotFound(){
		return httperors.NewNotFoundError("Something went wrong logging out!")
	 }
	
	GormDB.Model(&auth).Where("token =?", token).First(&auth)
	
	GormDB.Delete(auth)
	userRepo.DbClose(GormDB)
	
	return  nil
}
func (userRepo userrepo) GetOne(id int) (*model.User, *httperors.HttpError) {
	ok := userRepo.UserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("User with that id does not exists!")
	}
	user := model.User{}
	GormDB, err1 := userRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	
	GormDB.Model(&user).Where("id = ?", id).First(&user)
	userRepo.DbClose(GormDB)
	
	return &user, nil
}

func (userRepo userrepo) GetAll(users []model.User) ([]model.User, *httperors.HttpError) {
	GormDB, err1 := userRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	User := model.User{}
	GormDB.Model(&User).Find(&users)
	
	userRepo.DbClose(GormDB)
	if len(users) == 0 {
		return nil, httperors.NewNotFoundError("No results found!")
	}
	return users, nil
}

func (userRepo userrepo) Update(id int, user *model.User) (*model.User, *httperors.HttpError) {
	ok := userRepo.UserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("User with that id does not exists!")
	}
	
	
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return nil, err2
	}
	user.Password = hashpassword
	GormDB, err1 := userRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	User := model.User{}
	uuser := model.User{}
	
	GormDB.Model(&User).Where("id = ?", id).First(&uuser)
	if user.FName  == "" {
		user.FName = uuser.FName
	}
	if user.LName  == "" {
		user.LName = uuser.LName
	}
	if user.UName  == "" {
		user.UName = uuser.UName
	}
	if user.Phone  == "" {
		user.Phone = uuser.Phone
	}
	if user.Address  == "" {
		user.Address = uuser.Address
	}
	if user.Picture  == "" {
		user.Picture = uuser.Picture
	}
	if user.Email  == "" {
		user.Email = uuser.Email
	}
	if user.Password  == "" {
		user.Password = uuser.Password
	}
	GormDB.Model(&User).Where("id = ?", id).First(&User).Update(&user)
	
	userRepo.DbClose(GormDB)

	return user, nil
}
func (userRepo userrepo) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := userRepo.UserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("User with that id does not exists!")
	}
	user := model.User{}
	GormDB, err1 := userRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&user).Where("id = ?", id).First(&user)
	GormDB.Delete(user)
	userRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (userRepo userrepo)UserExist(email string) bool {
	user := model.User{}
	GormDB, err1 := userRepo.Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&user, "email =?", email).RecordNotFound(){
	   return false
	}
	userRepo.DbClose(GormDB)
	return true
	
}
func (userRepo userrepo)UserExistByid(id int) bool {
	user := model.User{}
	GormDB, err1 := userRepo.Getconnected()
	if err1 != nil {
		return false
	}
	if GormDB.First(&user, "id =?", id).RecordNotFound(){
	   return false
	}
	userRepo.DbClose(GormDB)
	return true
	
}
