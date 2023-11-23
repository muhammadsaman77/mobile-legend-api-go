package herorole

import (
	"context"
	"database/sql"
	"mobile-legend-api/helper"
	"mobile-legend-api/pkg/entities"
)

type HeroRoleService interface {
	GetAllHeroRole(ctx context.Context) (*[]entities.HeroRole,error)
	GetHeroRoleById(ctx context.Context, id int)(*entities.HeroRole,error)
	AddNewHeroRole(ctx context.Context, heroRole *entities.HeroRole) (*entities.HeroRole,error)
	UpdateHeroRole(ctx context.Context,heroRole *entities.HeroRole)(*entities.HeroRole,error)
	DeleteHeroRole(ctx context.Context, id int)(error)

}


type heroRoleService struct{
	repository HeroRoleRepository
	DB *sql.DB
}

func NewHeroRoleService(repository HeroRoleRepository, DB *sql.DB) HeroRoleService{
	return &heroRoleService{
		repository: repository,
		DB: DB,
	}
}

func (service *heroRoleService)	GetAllHeroRole(ctx context.Context) (*[]entities.HeroRole, error){
	tx,err:=  service.DB.Begin()
	if (err!= nil){
		return nil,err
	}
	defer tx.Commit()
	heroroles, err:= service.repository.FindAll(ctx,tx)
	if (err!= nil){
		return nil,err
	}
	return heroroles,nil
}
func (service *heroRoleService) GetHeroRoleById(ctx context.Context, id int)(*entities.HeroRole,error){
	tx,err:=  service.DB.Begin()
	if (err!= nil){
		return nil,err
	}
	defer tx.Commit()
	heroRole,err := service.repository.FindById(ctx,tx,id)
	if (err!= nil){
		return nil,err
	}
	return heroRole,nil
}


func (service *heroRoleService)	AddNewHeroRole(ctx context.Context, herorole *entities.HeroRole) (*entities.HeroRole,error){
	tx,err:= service.DB.Begin()
	helper.ErrorException(err)
	defer tx.Commit()
	newHeroRole,err:=  service.repository.AddNewData(ctx,tx,herorole)
	helper.ErrorException(err)
	return newHeroRole,nil
}


func (service *heroRoleService) 	UpdateHeroRole(ctx context.Context,herorole *entities.HeroRole)(*entities.HeroRole,error){
	tx,err:= service.DB.Begin()
	helper.ErrorException(err)
	defer tx.Commit()
	heroroledata,err := service.repository.FindById(ctx,tx,herorole.Id)
	helper.ErrorException(err)
	heroroledata.RoleId = herorole.RoleId
	updatedRole,err:= service.repository.UpdateData(ctx,tx,heroroledata)
	helper.ErrorException(err)
	return updatedRole,nil
}


func (service *heroRoleService) DeleteHeroRole(ctx context.Context, id int)(error){
	tx,err:= service.DB.Begin()
	if err != nil {
		return err
}
	defer tx.Commit()
	heroRoleData,err := service.repository.FindById(ctx,tx,id)
	if err != nil {
		return err
}
	err = service.repository.DeleteData(ctx,tx,heroRoleData.Id)
	if err != nil {
		return err
}
	return nil
}