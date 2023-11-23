package hero

import (
	"context"
	"database/sql"
	"mobile-legend-api/api/presenter"
	"mobile-legend-api/pkg/entities"
)

type HeroRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx)(*[]presenter.AllHero, error)
	FindById(ctx context.Context, tx *sql.Tx, id int)(*entities.Hero,error)
	AddNewData(ctx context.Context, tx *sql.Tx,hero *entities.Hero)(*entities.Hero,error)
	UpdateData(ctx context.Context, tx *sql.Tx, hero *entities.Hero)(*entities.Hero,error)
	DeleteData(ctx context.Context, tx *sql.Tx, id int)(error)
}

type heroRepository struct{

}

func NewHeroRepository()HeroRepository{
	return &heroRepository{}
}

func (repository *heroRepository) FindAll(ctx context.Context, tx *sql.Tx  )(*[]presenter.AllHero,error){
	SQL := "CALL CALL GetHeroData()	"
	rows, err:= tx.QueryContext(ctx,SQL)
	if(err!=nil){
		return nil, err
	}
	defer rows.Close()
	var heros []presenter.AllHero
	for rows.Next(){
		hero:= presenter.AllHero{}
		err:=rows.Scan(&hero.Id,&hero.Name)
		if(err!=nil){
			return nil, err
		}
		heros = append(heros, hero)
	}
	return &heros,nil
} 

func (repository *heroRepository) 	FindById(ctx context.Context, tx *sql.Tx, id int)(*entities.Hero,error){
	SQL := "CALL GetHeroDataById(?)"
	row, err:= tx.QueryContext(ctx,SQL,id)
	if(err!=nil){
		return nil, err
	}
	defer row.Close()
	var hero entities.Hero
	if row.Next(){
		err:= row.Scan(&hero.Id,&hero.Name,&hero.BasicHp,&hero.PhysicalAttack,&hero.MagicalAttack,&hero.PhysicalDefense,&hero.MagicalDefense,&hero.Roles)
		if(err!=nil){
			return nil, err
		}
	}
	return &hero,nil
}

func (repository *heroRepository) AddNewData(ctx context.Context, tx *sql.Tx,hero *entities.Hero)(*entities.Hero,error){
	SQl := "CALL InsertHero(?,?,?,?,?,?)"
	result,err:= tx.ExecContext(ctx,SQl,&hero.Name,&hero.BasicHp,&hero.PhysicalAttack,&hero.MagicalAttack,&hero.PhysicalDefense,&hero.MagicalDefense)
	if (err!=nil){
		return nil,err
	}
	id,err := result.LastInsertId()
	if err!=nil {
		return nil,err
	}
	hero.Id = int(id)
	return hero,nil
}

func (repository *heroRepository) UpdateData(ctx context.Context, tx *sql.Tx, hero *entities.Hero)(*entities.Hero,error){
	SQL := "CALL UpdateHero(?,?,?,?,?,?,?)"
	_,err:= tx.ExecContext(ctx,SQL,&hero.Name,&hero.BasicHp,&hero.PhysicalAttack,&hero.MagicalAttack,&hero.PhysicalDefense,&hero.MagicalDefense, &hero.Id)
	if err!=nil{
		return nil,err
	}
	return hero,nil
}

func (repository *heroRepository) DeleteData(ctx context.Context, tx *sql.Tx, id int)(error){
	SQL := "CALL DeleteHero(?)"
	_,err :=tx.ExecContext(ctx,SQL,id)
	if (err!=nil){
		return err
	}
	return nil
}