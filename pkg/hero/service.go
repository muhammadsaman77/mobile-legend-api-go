package hero

import (
	"context"
	"database/sql"
	"mobile-legend-api/api/presenter"
	"mobile-legend-api/pkg/entities"
)

type HeroService interface {
	GetAllHero(ctx context.Context)(*[]presenter.AllHero,error)
	GetDetailHero(ctx context.Context, id int)(*entities.Hero,error)
	AddNewHero(ctx context.Context, hero *entities.Hero) (*entities.Hero,error)
	UpdateHero(ctx context.Context,hero *entities.Hero)(*entities.Hero,error)
	DeleteHero(ctx context.Context, id int)(error)
}

type heroService struct{
	repository HeroRepository
	DB *sql.DB
}

func NewHeroService(repository HeroRepository, DB *sql.DB) HeroService{
	return &heroService{
		repository: repository,
		DB: DB,
	}
}
func  (service *heroService) GetAllHero(ctx context.Context)(*[]presenter.AllHero,error){
	tx,err:=  service.DB.Begin()
	if (err!= nil){
		return nil,err
	}
	defer tx.Commit()
	heroes,err:= service.repository.FindAll(ctx,tx)
	if (err!= nil){
		return nil,err
	}
	return heroes,nil
}
func (service *heroService) 	GetDetailHero(ctx context.Context, id int)(*entities.Hero,error){
	tx,err:=  service.DB.Begin()
	if (err!= nil){
		return nil,err
	}
	defer tx.Commit()
	hero,err:= service.repository.FindById(ctx,tx,id)
	if (err!= nil){
		return nil,err
	}
	return hero,nil
}

func (service *heroService) 	AddNewHero(ctx context.Context, hero *entities.Hero) (*entities.Hero,error){
	tx,err:= service.DB.Begin()
	if err!=nil {
		return nil,err
	}
	defer tx.Commit()
	newHero,err:=  service.repository.AddNewData(ctx,tx,hero)
	if err!=nil {
		return nil,err
	}
	// newHeroRole,err:= herorole.HeroRoleRepository.AddNewData(ctx,tx,he)
	return newHero,nil
}

func (service *heroService) UpdateHero(ctx context.Context,hero *entities.Hero)(*entities.Hero,error){
	tx,err:= service.DB.Begin()
	if err!=nil {
		return nil,err
	}
	defer tx.Commit()
	_ ,err = service.repository.FindById(ctx,tx,hero.Id)
	if err!=nil {
		return nil,err
	}
	updatedHero,err:= service.repository.UpdateData(ctx,tx,hero)
	if err!=nil {
		return nil,err
	}
	return updatedHero,nil

}

func (service *heroService)	DeleteHero(ctx context.Context, id int)(error){
	tx,err:= service.DB.Begin()
	if err != nil {
		return err
}
	defer tx.Commit()
	deletedHero,err := service.repository.FindById(ctx,tx,id)
	if err != nil {
		return err
}
	err = service.repository.DeleteData(ctx,tx,deletedHero.Id)
	if err != nil {
		return err
}
	return nil
}