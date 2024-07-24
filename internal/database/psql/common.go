package psql

import (
	"database/sql"
	"go_psql/internal/config"

	_ "github.com/lib/pq"
)

func getConnect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:123@localhost/"+config.DataBaseName+"?sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
