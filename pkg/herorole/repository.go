package herorole

import (
	"context"
	"database/sql"
	"mobile-legend-api/pkg/entities"
)

type HeroRoleRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) (*[]entities.HeroRole, error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (*entities.HeroRole, error)
	AddNewData(ctx context.Context, tx *sql.Tx, herorole *entities.HeroRole) (*entities.HeroRole, error)
	UpdateData(ctx context.Context, tx *sql.Tx, herorole *entities.HeroRole) (*entities.HeroRole, error)
	DeleteData(ctx context.Context, tx *sql.Tx, id int) error
}

type heroRoleRepository struct{

}

func NewHeroRoleRepository()HeroRoleRepository{
	return &heroRoleRepository{}
}

func (repository *heroRoleRepository) 	FindAll(ctx context.Context, tx *sql.Tx,  )(*[]entities.HeroRole,error){
	SQL := "CALL GetHeroRoleData()"
	rows, err:= tx.QueryContext(ctx,SQL)
	if(err!=nil){
		return nil, err
	}
	defer rows.Close()
	var heroRoles []entities.HeroRole
	for rows.Next(){
		heroRole:= entities.HeroRole{}
		err:=rows.Scan(&heroRole.Id ,&heroRole.HeroId,&heroRole.RoleId )
		if(err!=nil){
			return nil, err
		}
		heroRoles = append(heroRoles, heroRole)
	}
	return &heroRoles,nil
} 

func (repository *heroRoleRepository) 	FindById(ctx context.Context, tx *sql.Tx, id int)(*entities.HeroRole,error){
	SQL := "CALL GetHeroRoleById(?)	"
	row, err:= tx.QueryContext(ctx,SQL,id)
	if(err!=nil){
		return nil, err
	}
	defer row.Close()
	var heroRole entities.HeroRole
	if row.Next(){
		err:= row.Scan(&heroRole.Id,&heroRole.HeroId,&heroRole.RoleId)
		if(err!=nil){
			return nil, err
		}
	}
	return &heroRole,nil
}

func (repository *heroRoleRepository) AddNewData(ctx context.Context, tx *sql.Tx,heroRole *entities.HeroRole)(*entities.HeroRole,error){
	SQl := "CALL InsertHeroRole(?, ?)"
	result,err:= tx.ExecContext(ctx,SQl,&heroRole.HeroId,&heroRole.RoleId)
	if (err!=nil){
		return nil,err
	}
	id,err := result.LastInsertId()
	if err!=nil {
		return nil,err
	}
	heroRole.Id = int(id)
	return heroRole,nil
}

func (repository *heroRoleRepository) UpdateData(ctx context.Context, tx *sql.Tx, heroRole *entities.HeroRole)(*entities.HeroRole,error){
	SQL := "CALL UpdateHeroRole(?, ?)	"
	_,err:= tx.ExecContext(ctx,SQL,&heroRole.RoleId,&heroRole.Id)
	if err!=nil{
		return nil,err
	}
	return heroRole,nil
}

func (repository *heroRoleRepository) DeleteData(ctx context.Context, tx *sql.Tx, id int)(error){
	SQL := "CALL DeleteHeroRoleById(?)"
	_,err :=tx.ExecContext(ctx,SQL,id)
	if (err!=nil){
		return err
	}
	return nil
}