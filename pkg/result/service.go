package result

import (
	"context"
	"database/sql"
	"mobile-legend-api/api/presenter"
	"mobile-legend-api/helper"
	"mobile-legend-api/pkg/entities"
)

type ResultService interface {
	GetAllResult(ctx context.Context)(*[]entities.Result,error)
	GetResultByUserId(ctx context.Context,  id int)(*[]presenter.ResultByUserId,error)
	GetResultById(ctx context.Context,  id int) (*[]presenter.ResultByUserId, error)
	UpdateResult(ctx context.Context,  result *entities.Result) (*entities.Result, error)
}

type resultService struct{
	repository ResultRepository
	DB *sql.DB
}

func NewDetailResultService(repository ResultRepository, DB *sql.DB) ResultService{
	return &resultService{
		repository: repository,
		DB: DB,
	}
}

func (service * resultService) 	GetAllResult(ctx context.Context)(*[]entities.Result,error){
	tx,err:=  service.DB.Begin()
	helper.ErrorException(err)
	defer tx.Commit()
	results,err:= service.repository.FindAll(ctx,tx)
	helper.ErrorException(err)
	return results,nil

}

func(service * resultService) GetResultByUserId(ctx context.Context,  id int)(*[]presenter.ResultByUserId,error){
	tx,err:=  service.DB.Begin()
	helper.ErrorException(err)
	defer tx.Commit()
	results,err:= service.repository.FindResultByUserId(ctx,tx,id)
	helper.ErrorException(err)
	return results,nil

}

func(service * resultService) 	GetResultById(ctx context.Context,  id int) (*[]presenter.ResultByUserId, error){
	tx,err:=  service.DB.Begin()
	if (err!= nil){
		return nil,err
	}
	defer tx.Commit()
	result,err := service.repository.FindById(ctx,tx,id)
	if (err!= nil){
		return nil,err
	}
	return result,nil
}

func (service * resultService)  UpdateResult(ctx context.Context,  result *entities.Result) (*entities.Result, error){
	tx,err:= service.DB.Begin()
	helper.ErrorException(err)
	defer tx.Commit()
	_,err = service.repository.FindById(ctx,tx,result.Id)
	helper.ErrorException(err)
	UpdatedDetail,err:= service.repository.UpdateData(ctx,tx,result)
	helper.ErrorException(err)
	return UpdatedDetail,nil
}
