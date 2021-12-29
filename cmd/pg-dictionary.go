package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/user"

	_ "github.com/lib/pq"

	"github.com/gsiems/db-dictionary/config"
	"github.com/gsiems/db-dictionary/dictionary"
	"github.com/gsiems/db-dictionary/util"

	d "github.com/gsiems/go-db-meta/dbms"
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `usage: pg_dictionary [flags]

Database connection flags

  -db      The database to connect to. Overrides the DB_NAME environment
           parameter.

  -host    The hostname that the database is on. Overrides the DB_HOST
           environment parameter. Defaults to localhost.

  -port    The port that the database is listening on. Overrides the DB_PORT
           environment parameter. Defaults to 5432.

  -user    The username to connect as. Overrides the DB_USER environment
           parameter. Defaults to the OS user.

Extract database/schema(s) DDL flags

  -b       The base directory to write the generated results to.
           Overrides the BASE_DIR environment variable. Defaults to the
           current directory.

  -f       The format that comments should be rendered as (none, markdown).
           Overrides the COMMENT_FORMAT environment parameter.
           Defaults to none.

  -s       The comma separated list of schemas to extract. Overrides the
           DB_SCHEMAS environment parameter.

  -x       The comma separated list of schemas to exclude. Overrides the
           EXCLUDE_SCHEMAS environment parameter. Ignored if the -s flag
           is supplied. If no -s or -x flags are specified then all schemas
           are extracted.

Other flags

  -debug

  -q       Quiet mode. Do not print any error messages.

  -version Display the version information

`)
	}

	cfg, err := config.LoadConfig()
	util.FailOnErr(cfg.Quiet, err)

	var osUser string
	usr, err := user.Current()
	if err == nil {
		osUser = usr.Username
	}

	Username := util.Coalesce(cfg.Username, osUser)
	Host := util.Coalesce(cfg.Host, "localhost")
	Port := util.Coalesce(cfg.Port, "5432")
	DbName := cfg.DbName

	dsn := fmt.Sprintf("user=%s dbname=%s host=%s port=%s", Username, DbName, Host, Port)
	db, err := sql.Open("postgres", dsn)
	util.FailOnErr(cfg.Quiet, err)
	defer func() {
		if cerr := db.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	md, err := d.Init(db, d.PostgreSQL)
	util.FailOnErr(cfg.Quiet, err)

	err = dictionary.MakeDictionary(&md, cfg)
	util.FailOnErr(cfg.Quiet, err)

}
