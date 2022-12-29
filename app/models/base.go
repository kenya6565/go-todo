package models

import (
	"database/sql"
	"fmt"
	"go-todo/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatal(err)
	}

	// generate sql query with Sprintf
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
    password STRING,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP)`, tableNameUser)

	Db.Exec(cmdU)
}
