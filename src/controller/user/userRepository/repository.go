package userRepository

import (
	"database/sql"
	"fmt"
	"growdo/src/model"
)

type Repository interface {
	Create(e *model.Users) (*model.Users, error)
	FindDetail(consisi, parameter string) (*model.Users, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(e *model.Users) (*model.Users, error) {
	err := r.db.QueryRow(model.NewUsers("", e).Create()).Err()
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *repository) FindDetail(consisi, parameter string) (*model.Users, error) {

	query := model.NewUsers(consisi, nil).Detail(parameter)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	var category model.Users
	for rows.Next() {
		l := rows.Scan(&category.Id, &category.Name, &category.Email, &category.Password, &category.Date_of_birth, &category.Phone, &category.Roles, &category.Created_at, &category.Updated_at)
		fmt.Println(l)
	}
	return &category, nil
}
