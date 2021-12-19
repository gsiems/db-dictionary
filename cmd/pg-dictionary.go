package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gsiems/db-dictionary/config"
	"github.com/gsiems/db-dictionary/model"
	"github.com/gsiems/db-dictionary/util"
	"github.com/gsiems/db-dictionary/view"

	e "github.com/gsiems/go-db-meta/engine/pg"
	m "github.com/gsiems/go-db-meta/model"
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

	var c m.ConnectInfo
	c.Username = cfg.UserName
	c.Host = cfg.Host
	c.Port = cfg.Port
	c.DbName = cfg.DbName
	//c.Debug = debug

	db, err := e.OpenDB(&c)
	util.FailOnErr(cfg.Quiet, err)
	defer func() {
		if cerr := db.CloseDB(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	md := model.Init("pg", cfg)

	////////////////////////////////////////////////////////////////////////////
	catalog, err := e.CurrentCatalog(db)
	util.FailOnErr(cfg.Quiet, err)
	md.LoadCatalog(&catalog)

	schemata, err := e.Schemata(db, cfg.Schemas, cfg.Xclude)
	util.FailOnErr(cfg.Quiet, err)
	md.LoadSchemas(&schemata)

	tables, err := e.Tables(db, "")
	util.FailOnErr(cfg.Quiet, err)
	md.LoadTables(&tables)

	columns, err := e.Columns(db, "", "")
	util.FailOnErr(cfg.Quiet, err)
	md.LoadColumns(&columns)

	indexes, err := e.Indexes(db, "", "")
	util.FailOnErr(cfg.Quiet, err)
	md.LoadIndexes(&indexes)

	checkConstraints, err := e.CheckConstraints(db, "", "")
	util.FailOnErr(cfg.Quiet, err)
	md.LoadCheckConstraints(&checkConstraints)

	domains, err := e.Domains(db, "")
	util.FailOnErr(cfg.Quiet, err)
	md.LoadDomains(&domains)

	primaryKeys, err := e.PrimaryKeys(db, "", "")
	util.FailOnErr(cfg.Quiet, err)
	md.LoadPrimaryKeys(&primaryKeys)

	foreignKeys, err := e.ReferentialConstraints(db, "", "")
	util.FailOnErr(cfg.Quiet, err)
	md.LoadForeignKeys(&foreignKeys)

	uniqueConstraints, err := e.UniqueConstraints(db, "", "")
	util.FailOnErr(cfg.Quiet, err)
	md.LoadUniqueConstraints(&uniqueConstraints)

	dependencies, err := e.Dependencies(db, "", "")
	util.FailOnErr(cfg.Quiet, err)
	md.LoadDependencies(&dependencies)

	userTypes, err := e.Types(db, "")
	util.FailOnErr(cfg.Quiet, err)
	md.LoadUserTypes(&userTypes)

	//////////////////////////////////////////////////////////////////////////////
	err = view.CreateDictionary(md)
	util.FailOnErr(cfg.Quiet, err)

}
