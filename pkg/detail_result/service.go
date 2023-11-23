package detailresult

import (
	"context"
	"database/sql"
	"mobile-legend-api/api/presenter"
	"mobile-legend-api/helper"
	"mobile-legend-api/pkg/entities"
)

type DetailResultService interface {
	GetAllDetailResult(ctx context.Context)(*[]presenter.AllDetailResult,error)
	GetDetailByResultId(ctx context.Context,  id int)(*[]presenter.DetailResult,error)
	GetDetailById(ctx context.Context,  id int) (*presenter.DetailResult, error)
	AddNewDetail(ctx context.Context,  detailresult *entities.DetailResult) (*entities.DetailResult, error)

	DeleteDetail(ctx context.Context,  id int) error
}

type detailResultService struct{
	repository DetailResultRepository
	DB *sql.DB
}

func NewDetailResultService(repository DetailResultRepository, DB *sql.DB) DetailResultService{
	return &detailResultService{
		repository: repository,
		DB: DB,
	}
}

func (service * detailResultService) 	GetAllDetailResult(ctx context.Context)(*[]presenter.AllDetailResult,error){
	tx,err:=  service.DB.Begin()
	if (err!= nil){
		return nil,err
	}
	defer tx.Commit()
	detailresult,err:= service.repository.FindAll(ctx,tx)
	if (err!= nil){
		return nil,err
	}
	return detailresult,nil

}

func(service * detailResultService) GetDetailByResultId(ctx context.Context,  id int)(*[]presenter.DetailResult,error){
	tx,err:=  service.DB.Begin()
	if (err!= nil){
		return nil,err
	}
	defer tx.Commit()
	detailresult,err:= service.repository.FindByResultId(ctx,tx,id)
	if (err!= nil){
		return nil,err
	}
	return detailresult,nil

}

func(service * detailResultService) 	GetDetailById(ctx context.Context,  id int) (*presenter.DetailResult, error){
	tx,err:=  service.DB.Begin()
	if (err!= nil){
		return nil,err
	}
	defer tx.Commit()
	detailresult,err := service.repository.FindById(ctx,tx,id)
	if (err!= nil){
		return nil,err
	}
	return detailresult,nil
}
func(service * detailResultService) 	AddNewDetail(ctx context.Context,  detailresult *entities.DetailResult) (*entities.DetailResult, error){
	tx,err:= service.DB.Begin()
	helper.ErrorException(err)
	defer tx.Commit()
	newDetail,err:=  service.repository.AddNewData(ctx,tx,detailresult)
	helper.ErrorException(err)
	return newDetail,nil
}

	
func (service * detailResultService)DeleteDetail(ctx context.Context,  id int) error{
	tx,err:= service.DB.Begin()
	if err != nil {
		return err
}
	defer tx.Commit()
	deletedDetail,err := service.repository.FindById(ctx,tx,id)
	if err != nil {
		return err
}
	err = service.repository.DeleteData(ctx,tx,deletedDetail.Id)
	if err != nil {
		return err
}
	return nil
}