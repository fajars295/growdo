package model

import (
	"fmt"
	"growdo/src/helpers/componen"
	"log"
	"time"
)

// "name" varchar NOT NULL,
// 	"email" VARCHAR NOT NULL UNIQUE,
// 	"password" text NOT NULL,
// 	"date_of_birth" DATE NOT NULL,
// 	"phone" VARCHAR NOT NULL,
// 	"role" BIGINT,

type Users struct {
	Id            uint64    `json:"id"`
	Name          string    `json:"name" validate:"required,max=100"`
	Email         string    `json:"email" validate:"required,max=250"`
	Password      string    `json:"password" validate:"required,min=8"`
	Date_of_birth string    `json:"date_of_birth" validate:"required"`
	Phone         string    `json:"phone" validate:"required"`
	Roles         int64     `json:"roles"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

type Login struct {
	Email    string `json:"email" validate:"required,max=250"`
	Password string `json:"password" validate:"required,min=8"`
}

type funcUsers struct {
	Data  string
	Model *Users
}

func NewUsers(data string, M *Users) *funcUsers {
	return &funcUsers{
		Data:  data,
		Model: M,
	}
}

var (
	columUsers = "name, email, password, date_of_birth,phone, role"
	tabelUsers = "users"
)

func (r *funcUsers) Create() string {
	// data := "users () VALUES ($1, $2, $3, $4, $5, $6)"

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s', '%s', '%s', '%s', '%s', '%d') RETURNING ID",
		tabelUsers,
		columUsers,
		r.Model.Name,
		r.Model.Email,
		componen.HashPassword(r.Model.Password),
		r.Model.Date_of_birth,
		r.Model.Phone,
		r.Model.Roles,
	)
	log.Println(query)
	return query
}

// Userfindemail = "email=$1"
func (r *funcUsers) Detail(parameter string) string {

	var hasil string
	var data = fmt.Sprintf("SELECT id, %s, %s FROM %s WHERE", columUsers, DateGlobal(), tabelUsers)

	switch r.Data {
	case "email":
		hasil = fmt.Sprintf("%s email='%s'", data, parameter)
	case "id":
		hasil = fmt.Sprintf("%s id='%s'", data, parameter)
	}

	log.Println(hasil)
	return hasil
}
