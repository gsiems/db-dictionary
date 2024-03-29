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

	"github.com/gsiems/db-dictionary/util"
)

// Config is the structure for the configuration
type Config struct {
	CommentsFormat string
	ConfigFile     string
	CSSFiles       string
	DbComment      string
	DbName         string
	DSN            string
	DbEngine       string
	ExcludeSchemas string
	File           string
	GraphvizCmd    string
	HideSQL        bool
	Host           string
	ImgFiles       string
	IncludeSchemas string
	JSFiles        string
	Minify         bool
	NoGraphviz     bool
	OutputDir      string
	Port           string
	SSLMode        string
	Username       string
	UserPass       string
	Verbose        bool
}

var envMap = map[string]string{
	"CommentsFormat": "comments_format",
	"CSSFiles":       "css_files",
	"DbComment":      "db_comment",
	"DbEngine":       "dbms",
	"DbName":         "db_name",
	"DSN":            "dsn",
	"ExcludeSchemas": "exclude_schemas",
	"File":           "file",
	"HideSQL":        "hide_sql",
	"Host":           "host",
	"GraphvizCmd":    "graphviz_cmd",
	"NoGraphviz":     "no_graphviz",
	"ImgFiles":       "img_files",
	"IncludeSchemas": "include_schemas",
	"JSFiles":        "js_files",
	"Minify":         "minify",
	"OutputDir":      "output_dir",
	"Port":           "port",
	"SSLMode":        "ssl_mode",
	"Username":       "username",
	"UserPass":       "user_pass",
	"Verbose":        "verbose",
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

	cp, err := readFile(cfgFile, fp.Verbose)
	if err != nil {
		return e, err
	}

	e.ConfigFile = cfgFile
	e.HideSQL = fp.HideSQL || ep.HideSQL || cp.HideSQL
	e.Minify = fp.Minify || ep.Minify || cp.Minify
	e.NoGraphviz = fp.NoGraphviz || ep.NoGraphviz || cp.NoGraphviz
	e.Verbose = fp.Verbose || ep.Verbose || cp.Verbose
	e.CommentsFormat = util.Coalesce(fp.CommentsFormat, ep.CommentsFormat, cp.CommentsFormat, "none")
	e.CSSFiles = util.Coalesce(fp.CSSFiles, ep.CSSFiles, cp.CSSFiles)
	e.DbComment = util.Coalesce(fp.DbComment, ep.DbComment, cp.DbComment)
	e.DbEngine = util.Coalesce(fp.DbEngine, ep.DbEngine, cp.DbEngine)
	e.DbName = util.Coalesce(fp.DbName, ep.DbName, cp.DbName)
	e.DSN = util.Coalesce(ep.DSN, cp.DSN) // no command line arg
	e.ExcludeSchemas = util.Coalesce(fp.ExcludeSchemas, ep.ExcludeSchemas, cp.ExcludeSchemas)
	e.File = util.Coalesce(fp.File, ep.File, cp.File)
	e.GraphvizCmd = util.Coalesce(fp.GraphvizCmd, ep.GraphvizCmd, cp.GraphvizCmd, "fdp")
	e.Host = util.Coalesce(fp.Host, ep.Host, cp.Host)
	e.ImgFiles = util.Coalesce(fp.ImgFiles, ep.ImgFiles, cp.ImgFiles)
	e.IncludeSchemas = util.Coalesce(fp.IncludeSchemas, ep.IncludeSchemas, cp.IncludeSchemas)
	e.JSFiles = util.Coalesce(fp.JSFiles, ep.JSFiles, cp.JSFiles)
	e.OutputDir = util.Coalesce(fp.OutputDir, ep.OutputDir, cp.OutputDir)
	e.Port = util.Coalesce(fp.Port, ep.Port, cp.Port)
	e.SSLMode = util.Coalesce(fp.SSLMode, ep.SSLMode, cp.SSLMode)
	e.Username = util.Coalesce(fp.Username, ep.Username, cp.Username)
	e.UserPass = util.Coalesce(ep.UserPass, cp.UserPass) // no command line arg

	return e, nil
}

// readFlags parses the command line arguments to the application
func readFlags() (e Config, err error) {

	flag.BoolVar(&e.NoGraphviz, "nogv", false, "")
	flag.BoolVar(&e.HideSQL, "nosql", false, "")
	flag.BoolVar(&e.Minify, "minify", false, "")
	flag.BoolVar(&e.Verbose, "v", false, "")

	flag.StringVar(&e.CommentsFormat, "f", "", "")
	flag.StringVar(&e.ConfigFile, "c", "", "")
	flag.StringVar(&e.CSSFiles, "css", "", "")
	flag.StringVar(&e.DbComment, "comment", "", "")
	flag.StringVar(&e.DbEngine, "dbms", "", "")
	flag.StringVar(&e.DbName, "db", "", "")
	flag.StringVar(&e.ExcludeSchemas, "x", "", "")
	flag.StringVar(&e.File, "file", "", "")
	flag.StringVar(&e.GraphvizCmd, "gv", "", "")
	flag.StringVar(&e.Host, "host", "", "")
	flag.StringVar(&e.ImgFiles, "img", "", "")
	flag.StringVar(&e.IncludeSchemas, "s", "", "")
	flag.StringVar(&e.JSFiles, "js", "", "")
	flag.StringVar(&e.OutputDir, "out", "", "")
	flag.StringVar(&e.Port, "port", "", "")
	flag.StringVar(&e.SSLMode, "sslmode", "", "")
	flag.StringVar(&e.Username, "user", "", "")

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
		case "CommentsFormat":
			e.CommentsFormat = n
		case "CSSFiles":
			e.CSSFiles = n
		case "DbComment":
			e.DbComment = n
		case "DbEngine":
			e.DbEngine = n
		case "DbName":
			e.DbName = n
		case "DSN":
			e.DSN = n
		case "ExcludeSchemas":
			e.ExcludeSchemas = n
		case "GraphvizCmd":
			e.GraphvizCmd = n
		case "Host":
			e.Host = n
		case "ImgFiles":
			e.ImgFiles = n
		case "IncludeSchemas":
			e.IncludeSchemas = n
		case "JSFiles":
			e.JSFiles = n
		case "OutputDir":
			e.OutputDir = n
		case "Port":
			e.Port = n
		case "SSLMode":
			e.SSLMode = n
		case "Username":
			e.Username = n
		case "UserPass":
			e.UserPass = n
		case "HideSQL":
			e.HideSQL = chkBoolCfgParm(n)
		case "Minify":
			e.Minify = chkBoolCfgParm(n)
		case "NoGraphviz":
			e.NoGraphviz = chkBoolCfgParm(n)
		case "Verbose":
			e.Verbose = chkBoolCfgParm(n)
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
func readFile(cfgFile string, verbose bool) (e Config, err error) {

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
		if verbose {
			log.Println("No config file specified or found")
		}
		return e, nil
	}

	if verbose {
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
			case "CommentsFormat":
				e.CommentsFormat = n
			case "CSSFiles":
				e.CSSFiles = n
			case "DbComment":
				e.DbComment = n
			case "DbEngine":
				e.DbEngine = n
			case "DbName":
				e.DbName = n
			case "DSN":
				e.DSN = n
			case "ExcludeSchemas":
				e.ExcludeSchemas = n
			case "GraphvizCmd":
				e.GraphvizCmd = n
			case "Host":
				e.Host = n
			case "ImgFiles":
				e.ImgFiles = n
			case "IncludeSchemas":
				e.IncludeSchemas = n
			case "JSFiles":
				e.JSFiles = n
			case "OutputDir":
				e.OutputDir = n
			case "Port":
				e.Port = n
			case "SSLMode":
				e.SSLMode = n
			case "Username":
				e.Username = n
			case "UserPass":
				e.UserPass = n
			case "HideSQL":
				e.HideSQL = chkBoolCfgParm(n)
			case "Minify":
				e.Minify = chkBoolCfgParm(n)
			case "NoGraphviz":
				e.NoGraphviz = chkBoolCfgParm(n)
			case "Verbose":
				e.Verbose = chkBoolCfgParm(n)
			}
		}
	}

	return e, nil
}

// chkBoolCfgParm checks the configuration parameter and returns true/false
// based on the value (or lack thereof)
func chkBoolCfgParm(s string) bool {
	switch s {
	case "", "0":
		return false
	}
	return true
}
