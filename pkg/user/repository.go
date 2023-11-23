package user

import (
	"context"
	"database/sql"
	"mobile-legend-api/pkg/entities"
)

type UserRepository interface {
	FindByEmail(ctx context.Context,tx *sql.Tx,email string) (*entities.User, error)
	AddNewUser(ctx context.Context, tx *sql.Tx,user *entities.User) (*entities.User,error)
	FindByName( ctx context.Context, tx *sql.Tx,name string) (*[]	entities.User,error)
	FindById( ctx context.Context, tx *sql.Tx,id int) (*entities.User,error)
	UpdateById( ctx context.Context, tx *sql.Tx,user *entities.User) (*entities.User, error) 
	DeleteById( ctx context.Context, tx *sql.Tx,id int) (error)

}

type userRepository struct{

}

func NewUserRepository () UserRepository{
	return &userRepository{}
}

func (repository *userRepository) FindByEmail(ctx context.Context, tx *sql.Tx,email string) (*entities.User,error){
	SQL:= "CALL GetUserByEmail(?)	"
	rows, err := tx.QueryContext(ctx, SQL,email)
	if err!= nil{
		return nil,err
	}
	defer rows.Close()  
	var user entities.User
	if rows.Next(){
		err= rows.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.Gender,&user.VictoryCount,&user.DefeatCount,&user.Winrate)
		if err!= nil{
			return nil,err
		}
	}
	return &user,nil
}

func (repository *userRepository) AddNewUser(ctx context.Context, tx *sql.Tx, user *entities.User) (*entities.User,error){
	SQL := "CALL InsertUserData(?,?,?,?)"
	result,err:= tx.ExecContext(ctx, SQL,&user.Name,&user.Email,&user.Password,&user.Gender)
	if (err!=nil){
		return nil,err
	}
	id,err := result.LastInsertId()
	if err!=nil {
		return nil,err
	}
	user.Id = int(id)
	return user,nil
}
func (repository *userRepository) FindByName( ctx context.Context, tx *sql.Tx,name string) (*[]	entities.User,error){
	SQL:= "CALL GetUserByName(?)"
	rows,err:= tx.QueryContext(ctx,SQL,"%"+name+"%")
	if err!=nil {
		return nil,err
	}
	defer rows.Close()
	var users []entities.User
	for rows.Next(){
		user:= entities.User{}
		err:= rows.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.Gender,&user.VictoryCount,&user.DefeatCount,&user.Winrate)
		if err!=nil {
			return nil,err
		}
		users= append(users, user)
	}

	return &users,err
}

func (repository *userRepository) FindById( ctx context.Context, tx *sql.Tx,id int) (*entities.User,error){
	SQL := "CALL GetUserById(?)"

	rows,err:= tx.QueryContext(ctx,SQL,id)
	rows.Close()
	if err!=nil {
		return nil,err
	}
	var user entities.User
	rows.Scan(&user.Id,&user.Name,&user.Email,&user.Password,&user.Gender,&user.VictoryCount,&user.DefeatCount,&user.Winrate)
	return &user, err
}

func (repository *userRepository) UpdateById( ctx context.Context, tx *sql.Tx,user *entities.User) (*entities.User, error)  {
	SQL := "CALL UpdateUserData(?,?,?,?,?,?,?,?)"
	_,err:= tx.ExecContext(ctx,SQL,&user.Name,&user.Email,&user.Password,&user.Gender,&user.VictoryCount,&user.DefeatCount,&user.Winrate)
	if err!=nil{
		return nil,err
	}
	return user,nil
}

func (repository *userRepository) DeleteById( ctx context.Context, tx *sql.Tx,id int) (error){
	SQL := "CALL DeleteUserById(?)"
	_,err :=tx.ExecContext(ctx,SQL,id)
	if (err!=nil){
		return err
	}
	return nil
}

