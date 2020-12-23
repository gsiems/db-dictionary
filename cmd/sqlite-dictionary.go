package main

import (
	"flag"
	"fmt"
	"os"

	e "github.com/gsiems/db-dictionary/engine/sqlite"
	m "github.com/gsiems/db-dictionary/model"
	"github.com/gsiems/db-dictionary/util"
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
		fmt.Fprint(os.Stderr, `usage: pg_dictionary [flags]

Database connection flags

  -f       The database file to connect to.

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
	flag.StringVar(&file, "f", "", "")
	flag.StringVar(&base, "b", "", "")
	flag.StringVar(&schemas, "s", "", "")
	flag.StringVar(&xclude, "x", "", "")

	flag.Parse()

	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	//var osUser string
	//usr, err := user.Current()
	//if err == nil {
	//	osUser = usr.Username
	//}

	var c m.ConnectInfo
	c.File = file

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
