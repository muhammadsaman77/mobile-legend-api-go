package result

import (
	"context"
	"database/sql"
	"mobile-legend-api/api/presenter"
	"mobile-legend-api/pkg/entities"
)

type ResultRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx)(*[]entities.Result,error)
	FindResultByUserId(ctx context.Context, tx *sql.Tx,id int) (*[]presenter.ResultByUserId,error)
	FindById(ctx context.Context, tx *sql.Tx,id int) (*[]presenter.ResultByUserId,error)
	UpdateData(ctx context.Context, tx *sql.Tx, result *entities.Result) (*entities.Result, error)
	DeleteData(ctx context.Context, tx *sql.Tx, id int) error
}
type resultRepository struct {

}

func NewResultRepository() ResultRepository {
	return &resultRepository{}
}
func (repository *resultRepository) FindAll(ctx context.Context, tx *sql.Tx) (*[]entities.Result, error){
	SQL:= "SELECT id,status,durasi,mode FROM result_tbl"
		rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []entities.Result
	for rows.Next() {
		result := entities.Result{}
		err := rows.Scan(&result.Id, &result.Status,&result.Durasi, &result.Mode)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return &results, nil
}

func (repository *resultRepository) 	FindResultByUserId(ctx context.Context, tx *sql.Tx,id int) (*[]presenter.ResultByUserId,error){
	SQL:= `SELECT
	rt.id ,
	rt.status,
	rt.durasi,
	rt.mode,
	ht.name AS hero,
	drt.level_user,
	drt.kill_user,
	drt.death_user,
	drt.asisst_user,
	drt.gold,
	drt.skor,
	drt.grade
FROM
	result_tbl rt
INNER JOIN
	detail_result_tbl drt ON drt.result_id = rt.id
INNER JOIN
	hero_tbl ht ON drt.hero_id = ht.id
INNER JOIN
	user_tbl ut ON ut.id = drt.user_id
WHERE
	ut .id  = ?;
`
		rows, err := tx.QueryContext(ctx, SQL,id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []presenter.ResultByUserId
	for rows.Next() {
		result := presenter.ResultByUserId{}
		err := rows.Scan(&result.Id, &result.Status,&result.Durasi, &result.Mode,&result.Hero,&result.Level,&result.Kill,&result.Death,&result.Assist,&result.Gold,&result.Skor,&result.Grade)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return &results, nil
}

func (repository *resultRepository) 	FindById(ctx context.Context, tx *sql.Tx,id int) (*[]presenter.ResultByUserId,error){
	SQL:= `SELECT
	rt.id ,
	rt.status,
	rt.durasi,
	rt.mode,
	ht.name AS hero,
	drt.level_user,
	drt.kill_user,
	drt.death_user,
	drt.asisst_user,
	drt.gold,
	drt.skor,
	drt.grade
FROM
	result_tbl rt
INNER JOIN
	detail_result_tbl drt ON drt.result_id = rt.id
INNER JOIN
	hero_tbl ht ON drt.hero_id = ht.id
INNER JOIN
	user_tbl ut ON ut.id = drt.user_id
WHERE
	rt .id  = ?;
`
		rows, err := tx.QueryContext(ctx, SQL,id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []presenter.ResultByUserId
	for rows.Next() {
		result := presenter.ResultByUserId{}
		err := rows.Scan(&result.Id, &result.Status,&result.Durasi, &result.Mode,&result.Hero,&result.Level,&result.Kill,&result.Death,&result.Assist,&result.Gold,&result.Skor,&result.Grade)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return &results, nil
}



func (repository *resultRepository) UpdateData(ctx context.Context, tx *sql.Tx, result *entities.Result) (*entities.Result, error) {
	SQL := "CALL UpdateResultData('?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL,result.Status,result.Durasi,result.Mode)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repository *resultRepository) DeleteData(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := "CALL DeleteResultById(?)"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	return nil
}