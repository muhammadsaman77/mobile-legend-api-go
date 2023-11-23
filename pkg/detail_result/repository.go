package detailresult

import (
	"context"
	"database/sql"
	"mobile-legend-api/api/presenter"
	"mobile-legend-api/pkg/entities"
)

type DetailResultRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) (*[]presenter.AllDetailResult, error)
	FindByResultId(ctx context.Context, tx *sql.Tx, id int)(*[]presenter.DetailResult,error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (*presenter.DetailResult, error)
	AddNewData(ctx context.Context, tx *sql.Tx, detailresult *entities.DetailResult) (*entities.DetailResult, error)
	DeleteData(ctx context.Context, tx *sql.Tx, id int) error
}

type detailResultRepository struct {
}

func NewDetailResultRepository() DetailResultRepository {
	return &detailResultRepository{}
}

func (repository *detailResultRepository) FindAll(ctx context.Context, tx *sql.Tx) (*[]presenter.AllDetailResult, error) {
	SQL := `SELECT drt.id , drt.result_id, ht.name AS hero ,ut.name AS user , drt .grade  FROM detail_result_tbl drt INNER JOIN hero_tbl ht ON (ht.id = drt.hero_id  ) INNER JOIN user_tbl ut ON (ut.id = drt.user_id);`
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var detailresults []presenter.AllDetailResult
	for rows.Next() {
		detailresult := presenter.AllDetailResult{}
		err := rows.Scan(&detailresult.Id, &detailresult.ResultId,&detailresult.Hero, &detailresult.User, &detailresult.Grade)
		if err != nil {
			return nil, err
		}
		detailresults = append(detailresults, detailresult)
	}
	return &detailresults, nil
}

func (repository *detailResultRepository) FindByResultId(ctx context.Context, tx *sql.Tx, id int) (*[]presenter.DetailResult, error) {
	SQL := "CALL GetDetailResultByResultId(?)	"
	row, err := tx.QueryContext(ctx, SQL,id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var detailresults []presenter.DetailResult
	for row.Next() {
		detailresult := presenter.DetailResult{}
		err := row.Scan(&detailresult.Id, &detailresult.ResultId, &detailresult.Hero, &detailresult.User, &detailresult.LevelUser, &detailresult.KillUser, &detailresult.DeathUser, &detailresult.AssistUser,&detailresult.Gold,&detailresult.Skor, &detailresult.Grade,)
		if err != nil {
			return nil, err
		}
		detailresults = append(detailresults, detailresult)
	}
	return &detailresults, nil
}

func (repository *detailResultRepository) FindById(ctx context.Context, tx *sql.Tx, id int) (*presenter.DetailResult, error) {
	SQL := "CALL GetDetailResultById(?)"
	row, err := tx.QueryContext(ctx, SQL,id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var detailresult presenter.DetailResult
	if row.Next() {
		err := row.Scan(&detailresult.Id, &detailresult.ResultId, &detailresult.Hero, &detailresult.User, &detailresult.LevelUser, &detailresult.KillUser, &detailresult.DeathUser, &detailresult.AssistUser,&detailresult.Gold,&detailresult.Skor, &detailresult.Grade,)
		if err != nil {
			return nil, err
		}
	}
	return &detailresult, nil
}

func (repository *detailResultRepository) AddNewData(ctx context.Context, tx *sql.Tx, detailresult *entities.DetailResult) (*entities.DetailResult, error) {
	SQl := "CALL InsertDetailResult(?,?,?,?,?,?)	"
	result, err := tx.ExecContext(ctx, SQl, &detailresult.ResultId, &detailresult.HeroId, &detailresult.LevelUser, &detailresult.KillUser, &detailresult.DeathUser, &detailresult.AssistUser,&detailresult.Gold,&detailresult.Skor,&detailresult.Grade)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	detailresult.Id = int(id)
	return detailresult, nil
}

func (repository *detailResultRepository) DeleteData(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := "CALL DeleteDetailResultById(?)"
	_, err := tx.ExecContext(ctx, SQL, id)
	if err != nil {
		return err
	}
	return nil
}