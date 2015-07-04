// +build ignore

package main

import (
	"database/sql"
	"flag"
	"os"

	"github.com/bom-d-van/sqlkungfu"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var force bool
	var dev bool
	flag.BoolVar(&force, "f", false, "recreate database")
	flag.BoolVar(&dev, "dev", false, "create development data")
	flag.Parse()

	if force {
		if err := os.Remove("tmp/dbocean.db"); err != nil {
			panic(err)
		}
	}

	db, err := sql.Open("sqlite3", "tmp/dbocean.db")
	if err != nil {
		panic(err)
	}

	if err := sqlkungfu.InitMigration(db); err != nil {
		panic(err)
	}

	sqlkungfu.MustExec(db, `
		CREATE TABLE dbs (
			id       INTEGER PRIMARY KEY AUTOINCREMENT,
			name     VARCHAR(255) DEFAULT '',
			user     VARCHAR(255) DEFAULT '',
			password VARCHAR(255) DEFAULT '',
			path     VARCHAR(255) DEFAULT ''
		);
	`)

	sqlkungfu.MustExec(db, `
		ALTER TABLE dbs ADD type varchar(255);
	`)

	if dev {
		sqlkungfu.MustExec(db, `
			INSERT INTO dbs (name, type, path) VALUES ('devdb', 'sqlite3', 'tmp/devdb');
		`)
	}

	sqlkungfu.MustExec(db, `
		CREATE TABLE commands (
			id        INTEGER PRIMARY KEY AUTOINCREMENT,
			content   STRING DEFAULT '',
			log       STRING DEFAULT '',
			createdat DATE DEFAULT GETDATE
		)
	`)
}
