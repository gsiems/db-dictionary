package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gsiems/db-dictionary/config"
	"github.com/gsiems/db-dictionary/model"
	"github.com/gsiems/db-dictionary/util"
	"github.com/gsiems/db-dictionary/view"

	//"github.com/gsiems/db-dictionary/view"
	e "github.com/gsiems/go-db-meta/engine/sqlite"
	m "github.com/gsiems/go-db-meta/model"
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

	var c m.ConnectInfo
	c.File = cfg.File
	c.DbName = cfg.DbName
	//c.Debug = debug

	db, err := e.OpenDB(&c)
	util.FailOnErr(cfg.Quiet, err)
	defer func() {
		if cerr := db.CloseDB(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	md := model.Init("sqlite", cfg)

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
