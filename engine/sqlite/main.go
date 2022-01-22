package sqlite

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gsiems/db-dictionary-core/config"
)

func Connect(cfg config.Config) (db *sql.DB, err error) {

	File := cfg.File
	_, err = os.Stat(File)

	if err != nil {
		return db, err
	}

	// Options can be given using the following format: KEYWORD=VALUE and
	// multiple options can be combined with the & ampersand.
	// mode=ro

	db, err = sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", File))
	return db, err
}
