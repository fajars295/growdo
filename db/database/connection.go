package database

import (
	"database/sql"
	"fmt"
	"growdo/src/helpers/componen"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		componen.GodotEnv("DB_HOST"), componen.GodotEnv("DB_PORT"), componen.GodotEnv("DB_USERNAME"), componen.GodotEnv("DB_PASSWORD"), componen.GodotEnv("DB_DATABASE"))
	db, err := sql.Open(componen.GodotEnv("DB_CONNECTION"), psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
