package user

import (
	"context"
	"database/sql"
	"mobile-legend-api/pkg/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context,user *entities.User) (string,error)
	Register(ctx context.Context,user *entities.User) (*entities.User, error)
	SearchByName(ctx context.Context, name string)(*[]entities.User,error)
	FindByEmail(ctx context.Context, user *entities.User) (*entities.User,error)
	DeleteUser(ctx context.Context,id int)(*entities.User,error)
}

type userService struct {
	repository UserRepository
	DB *sql.DB
}
func NewUserService(repository UserRepository,DB *sql.DB) UserService{
	return &userService{
		repository: repository,
		DB: DB,
	}
}

func (service *userService) Login(ctx context.Context,user *entities.User) (string,error){
	tx,err := service.DB.Begin()
	if (err!= nil){
		return "",err
	}
	defer tx.Commit()
	userData,err:= service.repository.FindByEmail(ctx,tx,user.Email)
	 if err != nil {
		return "",err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return "",err
	}
	claims := jwt.MapClaims{
		"id":  userData.Id,
		"name": userData.Name,
		"email": userData.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
t, _ := token .SignedString([]byte("hja83jza"))

	return t,nil

}

func (service *userService) Register(ctx context.Context, user *entities.User)(*entities.User, error){
	tx,err := service.DB.Begin()
	if (err!= nil){
		return nil,err
	}
	defer tx.Commit()
	password := []byte(user.Password)
	hash,err:=  bcrypt.GenerateFromPassword(password,bcrypt.DefaultCost)
	if (err!= nil){
		return nil,err
	}
	user.Password = string(hash)
	newUser,err:= service.repository.AddNewUser(ctx,tx,user)
	if err!=nil {
		return nil,err
	}
	return  newUser,nil
}

func (service *userService) SearchByName(ctx context.Context, name string)(*[]entities.User,error){
	tx, err:=  service.DB.Begin()
	if err!=nil{
		return nil,err
	}
	defer tx.Commit()
	users,err:= service.repository.FindByName(ctx,tx,name)
	if err!=nil{
		return nil,err
	}
	return users,nil
}

func (service *userService) FindByEmail(ctx context.Context, user *entities.User)(*entities.User,error){
	tx, err:=  service.DB.Begin()
	if err!=nil{
		return nil,err
	}
	defer tx.Commit()
	userData, err := service.repository.FindByEmail(ctx,tx,user.Email)
	if err!=nil{
		return nil,err
	}
	return userData,nil
}

func (service *userService) DeleteUser(ctx context.Context, id int) (*entities.User,error){
	tx,errDb:= service.DB.Begin()
	if errDb!=nil{
		return nil,errDb
	}
	defer tx.Commit()
	user, errId:= service.repository.FindById(ctx,tx,id)
	if errId!=nil{
		return nil,errId
	}
	errDel := service.repository.DeleteById(ctx,tx,user.Id)
	if(errDel!=nil){
		return nil,errDel
	}
	return user, nil
}

