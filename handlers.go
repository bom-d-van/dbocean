package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bom-d-van/sqlkungfu"
)

func getDBs(rw http.ResponseWriter, req *http.Request) {

}

func addDB(rw http.ResponseWriter, req *http.Request) {

}

func removeDB(rw http.ResponseWriter, req *http.Request) {

}

type DBInfo struct {
	ID       uint
	Name     string
	Type     string
	User     string
	Password string
	Path     string

	db *sql.DB
}

func (i DBInfo) openDB() (err error) {
	switch i.Type {
	case "sqlite":
		i.db, err = sql.Open("sqlite3", i.Path)
	case "mysql":
	case "postgresql":
	}

	return
}

func execCmd(rw http.ResponseWriter, req *http.Request) {
	id, cmd := req.FormValue("db"), req.FormValue("cmd")
	rows, err := db.Query("select * from dbs where id = ?", id)
	if err != nil {
		panic(err)
	}
	var info DBInfo
	if err := sqlkungfu.Unmarshal(rows, &info); err != nil {
		panic(err)
	}
	if err := info.openDB(); err != nil {
		panic(err)
	}
	rows, err = info.db.Query(cmd)
	if err != nil {
		panic(err)
	}
	var v interface{}
	if err := sqlkungfu.Unmarshal(rows, v); err != nil {
		panic(err)
	}
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	if _, err := rw.Write(data); err != nil {
		panic(err)
	}
}
