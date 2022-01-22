package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/gsiems/db-dictionary/engine/mysql"
	"github.com/gsiems/db-dictionary/engine/postgresql"
	"github.com/gsiems/db-dictionary/engine/sqlite"
	"github.com/gsiems/db-dictionary/engine/oracle"

	"github.com/gsiems/db-dictionary/config"
	"github.com/gsiems/db-dictionary/runner"
	"github.com/gsiems/db-dictionary/util"

	d "github.com/gsiems/go-db-meta/dbms"
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `usage: db-dictionary [flags]

    -c           The configurations file to read, if any.

    -v           Verbose (default: false)

Database connection flags

    -dbms       The dbms to generate the dictionary for {oracle, postgresql, mariadb, mysql, sqlite}

    -db         The name of the database to connect to

    -host       The database host to connect to (default: localhost)

    -port       The port number to connect to (default depends on the DBMS)

    -user       The username to connect as (defaults to the current OS user)

    -sslmode    (PostgreSQL) Set the SSL mode to use {disable, require, verify-ca,
                verify-full} (default: require)

    -file       (SQLite) The database file to connect to.


Extract database/schema(s) DDL flags

    -s          The comma-separated list of schemas to include. Takes
                precedence over the -x flag (or exclude_schemas config
                file entry /environment variable). If neither are
                specified than all non-system schemas are included.

    -x          The comma-separated list of schemas to exclude (default: none)

Output/format flags

    -f          The formatter to use for rendering comments {none,
                markdown} (default: none)

    -css        The comma-separated list of CSS files to use in place
                of the default (default: none)

    -img        The comma-separated list of image files to include (for
                use with custom CSS) (default: none)

    -js         The comma-separated list of javascript files to include
                (default: none)

    -minify     Indicates if the output should be minified to reduce
                file sizes (default: false)

    -nosql      Do not show the queries used for views and materialized
                views (default is to show queries)

    -out        The directory to write the output files to (defaults to
                the current directory)

For further options please read the configuration documentation.

`)
	}

	cfg, err := config.LoadConfig()
	util.FailOnErr(true, err)

	var db *sql.DB
	var dbType int

	switch cfg.DbEngine {
	case "oracle", "ora":
		dbType = d.Oracle
		db, err = oracle.Connect(cfg)
	case "postgresql", "pg":
		dbType = d.PostgreSQL
		db, err = postgresql.Connect(cfg)
	case "mysql", "mariadb":
		dbType = d.MySQL
		db, err = mysql.Connect(cfg)
	case "sqlite":
		dbType = d.SQLite
		db, err = sqlite.Connect(cfg)
	default:
		err = fmt.Errorf("invalid database engine %q", cfg.DbEngine)
	}

	util.FailOnErr(true, err)
	defer func() {
		if cerr := db.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	md, err := d.Init(db, dbType)
	util.FailOnErr(true, err)

	err = runner.MakeDictionary(&md, cfg)
	util.FailOnErr(true, err)

}
