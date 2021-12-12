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

/*
var (
	showVersion bool
	version     = "0.1"
	base        string
	dbName      string
	debug       bool
	host        string
	port        string
	quiet       bool
	schemas     string
	userName    string
	xclude      string
)
*/
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
	d, err := model.DBDictionary("pg", cfg, catalog)
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

	/*
		catalogName := catalog.CatalogName.String
		catalogOwner := catalog.CatalogOwner.String
		catalogComment := catalog.Comment.String

		fmt.Printf("Catalog Name: %q\n", catalogName)
		fmt.Printf("Catalog Owner: %q\n", catalogOwner)
		fmt.Printf("Catalog Comment: %q\n", catalogComment)
	*/

	/*
		for _, table := range tables {
			fmt.Printf("        Table Schema: %q\n", table.TableSchema.String)
			fmt.Printf("        Table Name: %q\n", table.TableName.String)
			fmt.Printf("        Table Owner: %q\n", table.TableOwner.String)
			fmt.Printf("        Table Type: %q\n", table.TableType.String)
			fmt.Printf("        Table Comment: %q\n", table.Comment.String)
		}

		for _, column := range columns {
			fmt.Printf("        Table Schema: %q\n", column.TableSchema.String)
			fmt.Printf("        Table Name: %q\n", column.TableName.String)
			fmt.Printf("        Column Name: %q\n", column.ColumnName.String)
			fmt.Printf("        Data Type: %q\n", column.DataType.String)
			fmt.Printf("        Column Comment: %q\n", column.Comment.String)
		}
	*/

}
