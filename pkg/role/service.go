package role

import (
	"context"
	"database/sql"
	"mobile-legend-api/helper"
	"mobile-legend-api/pkg/entities"
)

type RoleService interface {
	GetAllRole(ctx context.Context) (*[]entities.Role, error)
	AddNewRole(ctx context.Context, role *entities.Role) (*entities.Role,error)
	UpdateRole(ctx context.Context,role *entities.Role)(*entities.Role,error)
	DeleteRole(ctx context.Context, id int)(error)
}

type roleService struct{
	repository RoleRepository
	DB *sql.DB
}

func NewRoleService(repository RoleRepository, DB *sql.DB) RoleService{
	return &roleService{
		repository: repository,
		DB: DB,
	}
}

func (service *roleService)	GetAllRole(ctx context.Context) (*[]entities.Role, error){
	tx,err:=  service.DB.Begin()
	if (err!= nil){
		return nil,err
	}
	defer tx.Commit()
	roles, err:= service.repository.FindAll(ctx,tx)
	if (err!= nil){
		return nil,err
	}
	return roles,nil
}

func (service *roleService)	AddNewRole(ctx context.Context, role *entities.Role) (*entities.Role,error){
	tx,err:= service.DB.Begin()
	helper.ErrorException(err)
	defer tx.Commit()
	newRole,err:=  service.repository.AddNewRole(ctx,tx,role)
	helper.ErrorException(err)
	return newRole,nil
}


func (service *roleService) 	UpdateRole(ctx context.Context,role *entities.Role)(*entities.Role,error){
	tx,err:= service.DB.Begin()
	helper.ErrorException(err)
	defer tx.Commit()
	roleData,err := service.repository.FindById(ctx,tx,role.Id)
	helper.ErrorException(err)
	roleData.Name = role.Name
	updatedRole,err:= service.repository.UpdateRole(ctx,tx,roleData)
	helper.ErrorException(err)
	return updatedRole,nil
}


func (service *roleService) 	DeleteRole(ctx context.Context, id int)(error){
	tx,err:= service.DB.Begin()
	if err != nil {
		return err
}
	defer tx.Commit()
	roleData,err := service.repository.FindById(ctx,tx,id)
	if err != nil {
		return err
}
	err = service.repository.DeleteRole(ctx,tx,roleData.Id)
	if err != nil {
		return err
}
	return nil
}