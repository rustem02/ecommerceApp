package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Связка с ДБ
func DBConn() (db *sql.DB, err error) {

	//dbDriver := "mysql"
	//dbUser := "root"
	//dbPass := "root"
	//dbName := "golang"
	//
	//db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	// если у вас ОС Виндоус, то вместо 8889 пишем 3306
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
	return
}
