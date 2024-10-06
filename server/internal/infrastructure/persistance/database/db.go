package database

import (
	"database/sql"
	"fmt"

	"github.com/gotired/real-estate-admin/server/internal/domain/model"
)

func Init(config *model.DATABASE) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.HOST, config.PORT, config.USERNAME, config.PASSWORD, config.DB_NAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db

}
