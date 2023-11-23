package role

import (
	"context"
	"database/sql"
	"mobile-legend-api/pkg/entities"
)

type RoleRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx,  )(*[]entities.Role,error)
	FindById(ctx context.Context, tx *sql.Tx, id int)(*entities.Role,error)
	AddNewRole(ctx context.Context, tx *sql.Tx,role *entities.Role)(*entities.Role,error)
	UpdateRole(ctx context.Context, tx *sql.Tx, role *entities.Role)(*entities.Role,error)
	DeleteRole(ctx context.Context, tx *sql.Tx, id int)(error)
}

type roleRepository struct{

}

func NewRoleRepository()RoleRepository{
	return &roleRepository{}
}

func (repository *roleRepository) 	FindAll(ctx context.Context, tx *sql.Tx,  )(*[]entities.Role,error){
	SQL := "CALL GetRoleViewData()"
	rows, err:= tx.QueryContext(ctx,SQL)
	if(err!=nil){
		return nil, err
	}
	defer rows.Close()
	var roles []entities.Role
	for rows.Next(){
		role:= entities.Role{}
		err:=rows.Scan(&role.Id,&role.Name)
		if(err!=nil){
			return nil, err
		}
		roles = append(roles, role)
	}
	return &roles,nil
} 

func (repository *roleRepository) 	FindById(ctx context.Context, tx *sql.Tx, id int)(*entities.Role,error){
	SQL := "CALL GetRoleViewById(?)	"
	row, err:= tx.QueryContext(ctx,SQL,id)
	if(err!=nil){
		return nil, err
	}
	defer row.Close()
	var role entities.Role
	if row.Next(){
		err:= row.Scan(&role.Id,&role.Name)
		if(err!=nil){
			return nil, err
		}
	}
	return &role,nil
}

func (repository *roleRepository) AddNewRole(ctx context.Context, tx *sql.Tx,role *entities.Role)(*entities.Role,error){
	SQl := "CALL InsertRole(?)	"
	result,err:= tx.ExecContext(ctx,SQl,&role.Name)
	if (err!=nil){
		return nil,err
	}
	id,err := result.LastInsertId()
	if err!=nil {
		return nil,err
	}
	role.Id = int(id)
	return role,nil
}

func (repository *roleRepository) UpdateRole(ctx context.Context, tx *sql.Tx, role *entities.Role)(*entities.Role,error){
	SQL := "CALL UpdateRoleName(?, ?)	"
	_,err:= tx.ExecContext(ctx,SQL,&role.Name,&role.Id)
	if err!=nil{
		return nil,err
	}
	return role,nil
}

func (repository *roleRepository) DeleteRole(ctx context.Context, tx *sql.Tx, id int)(error){
	SQL := "CALL DeleteRoleById(?)	"
	_,err :=tx.ExecContext(ctx,SQL,id)
	if (err!=nil){
		return err
	}
	return nil
}