package config

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

func GetDB()(bd *sql.DB,err error){
	
	return sql.Open("sqlite3","db/forum.db")
}