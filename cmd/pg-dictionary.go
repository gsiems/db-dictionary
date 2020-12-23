package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	e "github.com/gsiems/db-dictionary/engine/pg"
	m "github.com/gsiems/db-dictionary/model"
	"github.com/gsiems/db-dictionary/util"
)

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

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `usage: pg_dictionary [flags]

Database connection flags

  -db      The database to connect to.

  -host    The hostname that the database is on. Defaults to localhost.

  -port    The port that the database is listening on. Defaults to 5432.

  -user    The username to connect as. Defaults to the OS user.


Extract database/schema(s) DDL flags

  -b       The base directory to write the generated results to.
           Overrides the BASE_DIR environment variable. Defaults to the
           current directory.

  -s       The comma separated list of schemas to extract.

  -x       The comma separated list of schemas to exclude.
           Ignored if the -s flag is supplied.

Other flags

  -debug

  -q       Quiet mode. Do not print any error messages.

  -version Display the version information

`)
	}
	flag.BoolVar(&debug, "debug", false, "")
	flag.BoolVar(&quiet, "q", false, "")
	flag.BoolVar(&showVersion, "version", false, "")
	flag.StringVar(&dbName, "db", "", "")
	flag.StringVar(&host, "host", "", "")
	flag.StringVar(&port, "port", "", "")
	flag.StringVar(&userName, "user", "", "")
	flag.StringVar(&base, "b", "", "")
	flag.StringVar(&schemas, "s", "", "")
	flag.StringVar(&xclude, "x", "", "")

	flag.Parse()

	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	var osUser string
	usr, err := user.Current()
	if err == nil {
		osUser = usr.Username
	}

	var c m.ConnectInfo
	c.Username = util.Coalesce(userName, osUser)
	c.Host = util.Coalesce(host, "localhost")
	c.Port = util.Coalesce(port, "5432")
	c.DbName = dbName
	//c.Debug = debug

	db, err := e.OpenDB(&c)
	util.FailOnErr(quiet, err)
	defer func() {
		if cerr := db.CloseDB(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	dbInfo, err := e.DatabaseInfo(db)
	util.FailOnErr(quiet, err)

	fmt.Println(dbInfo)
}
