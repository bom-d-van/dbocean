package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bom-d-van/sqlkungfu"
	"github.com/golang/gddo/httputil/header"
)

func listDBs(rw http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("select * from dbs")
	if err != nil {
		panic(err)
	}
	var infos []DBInfo
	if err := sqlkungfu.Unmarshal(rows, &infos); err != nil {
		panic(err)
	}

	if returnJSONIfWanted(rw, req, infos) {
		return
	}

	rw.Write([]byte(Home(Index(infos))))
}

func getDB(rw http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get(":id")
	info, tables := retrieveTables(id)
	rw.Write([]byte(Home(Show(info, tables))))
}

func getDBTable(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	id := query.Get(":id")
	info, tables := retrieveTables(id)
	rows, err := info.db.Query("PRAGMA table_info(" + query.Get(":name") + ")")
	if err != nil {
		panic(err)
	}
	var table []map[string]interface{}
	if err := sqlkungfu.Unmarshal(rows, &table); err != nil && err != sqlkungfu.ErrNoRows {
		panic(err)
	}
	rows, err = info.db.Query("select * from " + query.Get(":name") + "")
	if err != nil {
		panic(err)
	}
	var data []map[string]interface{}
	if err := sqlkungfu.Unmarshal(rows, &data); err != nil && err != sqlkungfu.ErrNoRows {
		panic(err)
	}
	rw.Write([]byte(Home(Table(info, tables, table, data))))
}

func returnJSONIfWanted(rw http.ResponseWriter, req *http.Request, v interface{}) (wanted bool) {
	for _, spec := range header.ParseAccept(req.Header, "Accept") {
		if spec.Value == "application/json" {
			wanted = true
		}
	}

	if !wanted {
		return
	}

	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	if _, err := rw.Write(data); err != nil {
		panic(err)
	}

	return
}

func getDBTables(rw http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get(":id")
	_, tables := retrieveTables(id)
	if returnJSONIfWanted(rw, req, tables) {
		return
	}
}

func retrieveTables(id string) (info DBInfo, tables []map[string]string) {
	info = findDBInfo(id)
	rows, err := info.db.Query("SELECT * FROM sqlite_master WHERE type = 'table'")
	if err != nil {
		panic(err)
	}
	if err := sqlkungfu.Unmarshal(rows, &tables); err != nil {
		panic(err)
	}
	return
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

func (i *DBInfo) openDB() (err error) {
	switch i.Type {
	case "sqlite3":
		i.db, err = sql.Open("sqlite3", i.Path)
	case "mysql":
	case "postgresql":
	}

	return
}

func query(rw http.ResponseWriter, req *http.Request) {
	id, cmd := req.URL.Query().Get(":id"), req.FormValue("cmd")
	info := findDBInfo(id)
	rows, err := info.db.Query(cmd)
	if err != nil {
		panic(err)
	}
	var v interface{}
	if err := sqlkungfu.Unmarshal(rows, &v); err != nil && err != sqlkungfu.ErrNoRows {
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

func findDBInfo(id string) (info DBInfo) {
	rows, err := db.Query("select * from dbs where id = ?", id)
	if err != nil {
		panic(err)
	}
	if err := sqlkungfu.Unmarshal(rows, &info); err != nil {
		panic(err)
	}
	if err := info.openDB(); err != nil {
		panic(err)
	}
	return
}
