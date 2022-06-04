package masterDatarepository

import (
	"context"
	"database/sql"
	"errors"
	"growdo/src/helpers/componen"
	"growdo/src/model"
)

type Repository interface {
	Create(e *model.MasterData) (*model.MasterData, error)
	FindDetail(consisi string, f *model.FilterCari) (*model.MasterData, error)
	FindAll(f *model.FilterCari) (*[]model.MasterData, error)
	All(condisi string, f *model.FilterCari) (*[]model.MasterData, error)
}

type repository struct {
	db  *sql.DB
	ctx context.Context
}

func NewRepository(db *sql.DB, ctx context.Context) *repository {
	return &repository{
		db:  db,
		ctx: ctx,
	}
}

func (r *repository) Create(e *model.MasterData) (*model.MasterData, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	id, err := componen.NewQueryLogSlow(r.db, tx, r.ctx, model.NewMasterData("", e, nil).Create()).Create()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	e.Id = uint64(id)
	tx.Commit()
	return e, nil
}

func (r *repository) FindDetail(consisi string, f *model.FilterCari) (*model.MasterData, error) {

	query := model.NewMasterData(consisi, nil, f).Detail()

	rows, err := componen.NewQueryLogSlow(r.db, nil, r.ctx, query).Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	category := model.MasterData{}
	if rows.Next() {
		l := rows.Scan(&category.Id, &category.Type, &category.Status, &category.Value, &category.Created_at, &category.Updated_at)
		if l != nil {
			return nil, l
		}
		return &category, nil
	} else {
		return nil, errors.New("record is not found")
	}
}
func (r *repository) FindAll(f *model.FilterCari) (*[]model.MasterData, error) {

	query := model.NewMasterData(f.Kolom, nil, f).Detail()
	rows, err := componen.NewQueryLogSlow(r.db, nil, r.ctx, query).Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var hasil []model.MasterData
	for rows.Next() {
		category := model.MasterData{}
		l := rows.Scan(&category.Id, &category.Type, &category.Status, &category.Value, &category.Created_at, &category.Updated_at)
		if l != nil {
			return nil, l
		}
		hasil = append(hasil, category)
	}

	if len(hasil) == 0 {
		kosong := make([]model.MasterData, 0)
		return &kosong, nil
	}

	return &hasil, nil
}

func (r *repository) All(condisi string, f *model.FilterCari) (*[]model.MasterData, error) {
	query := model.NewMasterData(condisi, nil, f).Detail()
	rows, err := componen.NewQueryLogSlow(r.db, nil, r.ctx, query).Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var hasil []model.MasterData
	for rows.Next() {
		category := model.MasterData{}
		l := rows.Scan(&category.Id, &category.Type, &category.Status, &category.Value, &category.Created_at, &category.Updated_at)
		if l != nil {
			return nil, l
		}
		hasil = append(hasil, category)
	}

	if len(hasil) == 0 {
		kosong := make([]model.MasterData, 0)
		return &kosong, nil
	}

	return &hasil, nil

}
