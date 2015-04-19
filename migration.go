// +build ignore

package main

import (
	"database/sql"

	"github.com/bom-d-van/sqlkungfu"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "tmp/dbocean.db")
	if err != nil {
		panic(err)
	}

	if err := sqlkungfu.InitMigration(db); err != nil {
		panic(err)
	}

	sqlkungfu.MustExec(db, `
		CREATE TABLE dbs (
			id       int,
			name     varchar(255),
			user     varchar(255),
			password varchar(255),
			path     varchar(255)
		);
	`)

	sqlkungfu.MustExec(db, `
		CREATE TABLE commands (
			id        int,
			content   string,
			log       string,
			createdat date
		)
	`)
}
