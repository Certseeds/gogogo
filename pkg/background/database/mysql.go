package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var (
	mysql *sql.DB
	once  sync.Once
)

func readConfig(connection string) {
	once.Do(func() {
		db, err := sql.Open("mysql", connection)
		if err != nil {
			mysql = nil
			return
		}
		err = db.Ping()
		if err != nil {
			mysql = nil
			return
		}
		mysql = db
		mysql.SetMaxOpenConns(100)
		mysql.SetMaxIdleConns(100)
	})
}
