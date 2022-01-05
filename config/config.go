// Package config contains the configuration structure along the the functions
// for resolving the configuration to use when creating a data dictionary
package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"

	"github.com/gsiems/db-dictionary-core/util"
)

// Config is the structure for the configuration
type Config struct {
	ShowVersion    bool
	Debug          bool
	Quiet          bool
	Version        string
	OutputDir      string
	DbName         string
	DbComment      string
	DSN            string
	File           string
	Host           string
	Port           string
	Username       string
	UserPass       string
	Schemas        string
	Xclude         string
	ConfigFile     string
	CommentsFormat string
}

var envMap = map[string]string{
	"OutputDir":      "target_dir",
	"DbName":         "db_name",
	"DbComment":      "db_comment",
	"DSN":            "db_dsn",
	"Host":           "db_host",
	"Port":           "db_port",
	"Username":       "db_user",
	"UserPass":       "user_pass",
	"Schemas":        "schemas",
	"Xclude":         "exclude_schemas",
	"CommentsFormat": "comment_format",
}

// LoadConfig loads a configuration by using a configuration file (if
// found/specifed) combined with environmental values and run-time arguments
// to the application.
func LoadConfig() (e Config, err error) {

	fp, err := readFlags()
	if err != nil {
		return e, err
	}

	ep, err := readEnv()
	if err != nil {
		return e, err
	}

	cfgFile := util.Coalesce(fp.ConfigFile, ep.ConfigFile)

	cp, err := readFile(cfgFile, fp.Quiet)
	if err != nil {
		return e, err
	}

	e.ShowVersion = fp.ShowVersion
	e.Debug = fp.Debug
	e.Quiet = fp.Quiet
	//e.Version        = util.Coalesce(fp. , ep.  , cp.  )
	e.OutputDir = util.Coalesce(fp.OutputDir, ep.OutputDir, cp.OutputDir)
	e.DbName = util.Coalesce(fp.DbName, ep.DbName, cp.DbName)
	e.File = util.Coalesce(fp.File, ep.File, cp.File)
	e.Host = util.Coalesce(fp.Host, ep.Host, cp.Host)
	e.Port = util.Coalesce(fp.Port, ep.Port, cp.Port)
	e.Username = util.Coalesce(fp.Username, ep.Username, cp.Username)
	e.UserPass = util.Coalesce(fp.UserPass, ep.UserPass, cp.UserPass)
	e.Schemas = util.Coalesce(fp.Schemas, ep.Schemas, cp.Schemas)
	e.Xclude = util.Coalesce(fp.Xclude, ep.Xclude, cp.Xclude)
	e.ConfigFile = cfgFile
	e.CommentsFormat = util.Coalesce(fp.CommentsFormat, ep.CommentsFormat, cp.CommentsFormat, "none")

	return e, nil
}

// readFlags parses the command line arguments to the application
func readFlags() (e Config, err error) {

	flag.BoolVar(&e.Debug, "debug", false, "")
	flag.BoolVar(&e.Quiet, "q", false, "")
	flag.BoolVar(&e.ShowVersion, "v", false, "")
	flag.StringVar(&e.DbName, "db", "", "")
	flag.StringVar(&e.File, "file", "", "")
	flag.StringVar(&e.Host, "host", "", "")
	flag.StringVar(&e.Port, "port", "", "")
	flag.StringVar(&e.Username, "user", "", "")
	flag.StringVar(&e.OutputDir, "b", "", "")
	flag.StringVar(&e.Schemas, "s", "", "")
	flag.StringVar(&e.Xclude, "x", "", "")
	flag.StringVar(&e.CommentsFormat, "f", "", "")
	flag.StringVar(&e.ConfigFile, "c", "", "")

	flag.Parse()

	if e.File != "" {
		e.File, err = filepath.Abs(e.File)
		if err != nil {
			return e, err
		}
	}
	return e, nil

}

// readEnv reads the environment variables for configuration information
func readEnv() (e Config, err error) {

	for k, v := range envMap {
		n := os.Getenv(v)
		switch k {

		case "OutputDir":
			e.OutputDir = n
		case "DbName":
			e.DbName = n
		case "DSN":
			e.DSN = n
		case "DbComment":
			e.DbComment = n
		case "Host":
			e.Host = n
		case "Port":
			e.Port = n
		case "Username":
			e.Username = n
		case "UserPass":
			e.UserPass = n
		case "Schemas":
			e.Schemas = n
		case "Xclude":
			e.Xclude = n
		case "CommentsFormat":
			e.CommentsFormat = n
		}
	}

	return e, nil

}

// readFile reads a configuration file for the application
//
// If a configuration file is specified then ensure that the file is valid
// and raise an error if not. If a configuration file is not specified
// then look in the directory of the executable for a similarly named .cfg
// file. If there is a valid configuration file then use it to initialize
// the environment.
func readFile(cfgFile string, quiet bool) (e Config, err error) {

	// No file specified, look for one
	if "" == cfgFile {

		appPath, errc := os.Executable()
		if errc != nil {
			return e, fmt.Errorf("Error determining executable: ", errc)
		}

		appPath = path.Clean(appPath)
		tmpFile := path.Base(appPath)
		if "" != path.Ext(appPath) {
			// strip the extension
			tmpFile = strings.TrimSuffix(tmpFile, "."+path.Ext(appPath))
		}

		cfgFile = path.Join(path.Dir(appPath), tmpFile+".cfg")
		_, errc = os.Lstat(cfgFile)
		if errc != nil {
			cfgFile = ""
		}
	}

	if "" == cfgFile {
		if !quiet {
			log.Println("No config file specified or found")
		}
		return e, nil
	}

	if !quiet {
		log.Printf("Using config file (%s)\n", cfgFile)
	}
	var dotEnv map[string]string
	dotEnv, err = godotenv.Read(cfgFile)

	if err != nil {
		return e, fmt.Errorf("Error loading config file (%s): %s", cfgFile, err)
	}

	for k, v := range envMap {
		n, ok := dotEnv[v]
		if ok {
			switch k {
			case "OutputDir":
				e.OutputDir = n
			case "DbName":
				e.DbName = n
			case "DbComment":
				e.DbComment = n
			case "DSN":
				e.DSN = n
			case "Host":
				e.Host = n
			case "Port":
				e.Port = n
			case "Username":
				e.Username = n
			case "UserPass":
				e.UserPass = n
			case "Schemas":
				e.Schemas = n
			case "Xclude":
				e.Xclude = n
			case "CommentsFormat":
				e.CommentsFormat = n
			}
		}
	}

	return e, nil
}
