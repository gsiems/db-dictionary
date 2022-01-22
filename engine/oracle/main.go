package oracle

import (
	"database/sql"
	"fmt"
	"os/user"

	_ "github.com/godror/godror"

	orap "github.com/gsiems/orapass"

	"github.com/gsiems/db-dictionary/config"
	"github.com/gsiems/db-dictionary/util"
)

func Connect(cfg config.Config) (db *sql.DB, err error) {

	dsn := cfg.DSN
	if dsn == "" {

		var osUser string
		usr, err := user.Current()
		if err == nil {
			osUser = usr.Username
		}

		username := util.Coalesce(cfg.Username, osUser)
		host := util.Coalesce(cfg.Host, "localhost")
		port := util.Coalesce(cfg.Port, "1521")
		dbName := cfg.DbName

		var p orap.Parser
		p.Username = username
		p.Host = host
		p.Port = port
		p.DbName = dbName
		//p.OrapassFile = orapassFile

		// todo add config option for orapass file

		cp, err := p.GetPasswd()

		if cp.Username == "" || cp.Password == "" {
			err = fmt.Errorf("invalid username/password specified")
			util.FailOnErr(true, err)
		}

		dsn = fmt.Sprintf("%s/%s@%s", cp.Username, cp.Password, cp.DbName)
	}
	db, err = sql.Open("godror", dsn)
	return db, err
}
