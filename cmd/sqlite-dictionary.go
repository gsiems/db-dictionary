package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gsiems/db-dictionary/config"
	"github.com/gsiems/db-dictionary/dictionary"
	"github.com/gsiems/db-dictionary/util"

	d "github.com/gsiems/go-db-meta/dbms"
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `usage: sqlite_dictionary [flags]

Database connection flags

  -file    The database file to connect to.

Extract database/schema(s) DDL flags

  -b       The base directory to write the generated results to.
           Overrides the BASE_DIR environment variable. Defaults to the
           current directory.

  -s       The comma separated list of schemas to extract.

  -x       The comma separated list of schemas to exclude.
           Ignored if the -s flag is supplied.

Other flags

  -db      The name of the database to connect to. Overrides the
           DB_NAME environment parameter.

  -debug

  -q       Quiet mode. Do not print any error messages.

  -version Display the version information

`)
	}
	cfg, err := config.LoadConfig()
	util.FailOnErr(cfg.Quiet, err)

	File := cfg.File
	//DbName := cfg.DbName
	_, err = os.Stat(File)
	util.FailOnErr(cfg.Quiet, err)

	// Options can be given using the following format: KEYWORD=VALUE and
	// multiple options can be combined with the & ampersand.
	// mode=ro

	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", File))
	util.FailOnErr(cfg.Quiet, err)
	defer func() {
		if cerr := db.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	md, err := d.Init(db, d.SQLite)
	util.FailOnErr(cfg.Quiet, err)

	err = dictionary.MakeDictionary(&md, cfg)
	util.FailOnErr(cfg.Quiet, err)

}
