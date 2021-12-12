package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gsiems/db-dictionary/config"
	"github.com/gsiems/db-dictionary/model"
	"github.com/gsiems/db-dictionary/util"
	"github.com/gsiems/db-dictionary/view"
	e "github.com/gsiems/go-db-meta/engine/sqlite"
	m "github.com/gsiems/go-db-meta/model"
)

var (
	showVersion bool
	version     = "0.1"
	base        string
	file        string
	debug       bool
	quiet       bool
	schemas     string
	xclude      string
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

	////////////////////////////////////////////////////////////////////////////
	catalog, err := e.CurrentCatalog(db)
	util.FailOnErr(cfg.Quiet, err)

	schemata, err := e.Schemata(db, cfg.Schemas, cfg.Xclude)
	util.FailOnErr(cfg.Quiet, err)

	tables, err := e.Tables(db, "")
	util.FailOnErr(cfg.Quiet, err)

	columns, err := e.Columns(db, "", "")
	util.FailOnErr(cfg.Quiet, err)

	////////////////////////////////////////////////////////////////////////////
	d, err := model.DBDictionary("sqlite", cfg, catalog)
	util.FailOnErr(cfg.Quiet, err)

	s, err := model.Schemas(&schemata)
	util.FailOnErr(cfg.Quiet, err)

	t, err := model.Tables(&tables, &columns)
	util.FailOnErr(cfg.Quiet, err)

	////////////////////////////////////////////////////////////////////////////
	view.RenderSchemaList(&d, &s)
	util.FailOnErr(cfg.Quiet, err)

	view.RenderTableList(&d, &s, &t)
	util.FailOnErr(cfg.Quiet, err)

	view.RenderTables(&d, &s, &t)
	util.FailOnErr(cfg.Quiet, err)

	view.RenderColumns(&d, &s, &t)
	util.FailOnErr(cfg.Quiet, err)

}
