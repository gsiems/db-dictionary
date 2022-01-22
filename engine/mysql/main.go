package mysql

import (
	"database/sql"
	"fmt"
	"os/user"

	_ "github.com/go-sql-driver/mysql"

	config "github.com/gsiems/db-dictionary/config"
	util "github.com/gsiems/db-dictionary/util"
)

func Connect(cfg config.Config) (db *sql.DB, err error) {

	var osUser string
	usr, err := user.Current()
	if err == nil {
		osUser = usr.Username
	}

	Username := util.Coalesce(cfg.Username, osUser)
	//Host := util.Coalesce(cfg.Host, "localhost")
	//Port := util.Coalesce(cfg.Port, "3306")
	DbName := cfg.DbName
	UserPass := util.Coalesce(cfg.UserPass, Username)

	//dsn := fmt.Sprintf("Server=%s;Port=%s;Database=%s;Uid=%s;Pwd=%s;", Host, Port, DbName, Username, UserPass)
	//dsn := fmt.Sprintf("%s:%s@%s:%s/%s?timeout=30s", Username, UserPass, Host, Port, DbName)
	dsn := fmt.Sprintf("%s:%s@/%s", Username, UserPass, DbName)
	//fmt.Println (dsn)
	db, err = sql.Open("mysql", dsn)
	return db, err
}
