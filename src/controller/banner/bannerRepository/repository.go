package bannerrepository

import (
	"context"
	"database/sql"
	"errors"
	"growdo/src/model"
	"log"
)

type Repository interface {
	Create(e *model.Banner) (*model.Banner, error)
	FindDetail(consisi string, f *model.FilterCari) (*model.Banner, error)
	All(condisi string, f *model.FilterCari) (*[]model.Banner, error)
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

func (r *repository) Create(e *model.Banner) (*model.Banner, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	errors := tx.QueryRow(model.NewBanner("", e, nil).Create()).Scan(&e.Id)
	if errors != nil {
		tx.Rollback()
		return nil, errors
	}
	tx.Commit()
	return e, nil
}

func (r *repository) FindDetail(consisi string, f *model.FilterCari) (*model.Banner, error) {

	query := model.NewBanner(consisi, nil, f).Detail()
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	category := model.Banner{}
	if rows.Next() {
		l := rows.Scan(&category.Id, &category.Images, &category.Status, &category.Url, &category.Created_at, &category.Updated_at)
		if l != nil {
			return nil, l
		}
		return &category, nil
	} else {
		return nil, errors.New("record is not found")
	}
}

func (r *repository) All(condisi string, f *model.FilterCari) (*[]model.Banner, error) {
	query := model.NewBanner(condisi, nil, f).Detail()
	rows, err := r.db.QueryContext(r.ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []model.Banner
	for rows.Next() {
		category := model.Banner{}
		l := rows.Scan(&category.Id, &category.Images, &category.Status, &category.Url, &category.Created_at, &category.Updated_at)
		log.Println(l)
		if l != nil {
			return nil, l
		}
		hasil = append(hasil, category)
	}

	if len(hasil) == 0 {
		kosong := make([]model.Banner, 0)
		return &kosong, nil
	}

	return &hasil, nil

}
