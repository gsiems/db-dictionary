package model

import (
	"database/sql"
)

type ConnectInfo struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
	File     string
}

// DB contains an database/sql connection
type DB struct {
	*sql.DB
}

// BindDB
//func BindDB(db *sql.DB) DB {
//	return DB{db}
//}

// CloseDB closes a DB reference.
func (db *DB) CloseDB() error {
	return db.DB.Close()
}
