package postgresql

import (
	"database/sql"
	"fmt"
	"os/user"

	_ "github.com/lib/pq"

	"github.com/gsiems/db-dictionary-core/config"
	"github.com/gsiems/db-dictionary-core/util"
)

func Connect(cfg config.Config) (db *sql.DB, err error) {

	dsn := cfg.DSN
	if dsn == "" {

		var osUser string
		usr, cerr := user.Current()
		if cerr == nil {
			osUser = usr.Username
		}

		Username := util.Coalesce(cfg.Username, osUser)
		Host := util.Coalesce(cfg.Host, "localhost")
		Port := util.Coalesce(cfg.Port, "5432")
		DbName := cfg.DbName
		SSLMode := cfg.SSLMode

		dsn = fmt.Sprintf("user=%s dbname=%s host=%s port=%s", Username, DbName, Host, Port)

		if SSLMode != "" {
			dsn += fmt.Sprintf(" sslmode=%s", SSLMode)
		}
	}
	db, err = sql.Open("postgres", dsn)
	return db, err
}
