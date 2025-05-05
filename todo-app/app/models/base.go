package models

import (
	"app/todo-app/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)

	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`create table if not exist %s (
	id INTEGER PRIMARY KRY AUTOINCREMENT,
	uuid STRING NOT NULL UNIQUE,
	name STRING,
	email STRING,
	password STRING,
	created?at DATETIME)`, tableNameUser)
}
