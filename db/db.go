package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var Db *sql.DB
var err error

func InitDb(){
	Db, err = sql.Open("sqlite","klipa_db.sql")

	if err != nil {
		panic(err)
	}

	Db.SetMaxOpenConns(3)
	Db.SetMaxIdleConns(2)

	createTable()
}

func createTable () {
	tableQuery := `
	
	CREATE TABLE IF NOT EXISTS klipa (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT NOT NULL,
		value TEXT NOT NULL
	)`
	_, err := Db.Exec(tableQuery)

	if err != nil{
		panic(err)
	}
}