package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/gsiems/db-dictionary/util"
)

type Config struct {
	ShowVersion    bool
	Debug          bool
	Quiet          bool
	Version        string
	OutputDir      string
	DbName         string
	File           string
	Host           string
	Port           string
	UserName       string
	UserPass       string
	Schemas        string
	Xclude         string
	DictionaryDate string
	BinFile        string
	ConfigFile     string
	CommentsFormat string
}

var envMap = map[string]string{
	"OutputDir":      "BASE_DIR",
	"DbName":         "DB_NAME",
	"Host":           "DB_HOST",
	"Port":           "DB_PORT",
	"UserName":       "DB_USER",
	"UserPass":       "DB_USER_PASSWORD",
	"Schemas":        "DB_SCHEMAS",
	"Xclude":         "EXCLUDE_SCHEMAS",
	"CommentsFormat": "COMMENT_FORMAT",
}

func LoadConfig() (e Config, err error) {

	var cfgFile string

	flag.BoolVar(&e.Debug, "debug", false, "")
	flag.BoolVar(&e.Quiet, "q", false, "")
	flag.BoolVar(&e.ShowVersion, "v", false, "")
	flag.StringVar(&e.DbName, "db", "", "")
	flag.StringVar(&e.File, "file", "", "")
	flag.StringVar(&e.Host, "host", "", "")
	flag.StringVar(&e.Port, "port", "", "")
	flag.StringVar(&e.UserName, "user", "", "")
	flag.StringVar(&e.OutputDir, "b", "", "")
	flag.StringVar(&e.Schemas, "s", "", "")
	flag.StringVar(&e.Xclude, "x", "", "")
	flag.StringVar(&e.CommentsFormat, "f", "", "")
	flag.StringVar(&cfgFile, "c", "", "")

	flag.Parse()

	t := time.Now()
	e.DictionaryDate = t.Format(time.RFC1123)

	if e.File != "" {
		e.File, err = filepath.Abs(e.File)
		if err != nil {
			return e, err
		}
	}

	var myEnv map[string]string
	var keys []string
	for _, v := range envMap {
		keys = append(keys, v)
	}
	myEnv, errc := ReadEnv(cfgFile, keys)
	if errc != nil {
		return e, errc
	}

	for cName, eName := range envMap {
		eVal, _ := myEnv[eName]

		switch cName {
		case "OutputDir":
			e.OutputDir = util.Coalesce(e.OutputDir, eVal)

			if "" == e.OutputDir {
				p, errc := os.Getwd()
				if errc != nil {
					return e, fmt.Errorf("Error determining current directory: ", errc)
				}
				p, errc = filepath.EvalSymlinks(p)
				if errc != nil {
					return e, fmt.Errorf("Error resolving current directory: ", errc)
				}
				e.OutputDir = p
			}

		case "DbName":
			e.DbName = util.Coalesce(e.DbName, eVal)

		case "Host":
			e.Host = util.Coalesce(e.Host, eVal)

		case "Port":
			e.Port = util.Coalesce(e.Port, eVal)

		case "UserName":
			e.UserName = util.Coalesce(e.UserName, eVal)

			if "" == e.UserName {
				usr, errc := user.Current()
				if errc == nil {
					e.UserName = usr.Username
				}
			}

		case "UserPass":
			e.UserPass = util.Coalesce(e.UserPass, eVal)

		case "Schemas":
			e.Schemas = util.Coalesce(e.Schemas, eVal)

		case "Xclude":
			e.Xclude = util.Coalesce(e.Xclude, eVal)

		case "CommentsFormat":
			e.CommentsFormat = util.Coalesce(e.CommentsFormat, eVal, "none")
		}
	}

	return e, err
}

/* ReadEnv initializes the environment variables for the application

If a configuration file is specified then ensure that the file is valid
and raise an error if not. If a configuration file is not specified
then look in the directory of the executable for a similarly named .cfg
file. If there is a valid configuration file then use it to initialize
the environment.

For each value in the configuration file, if that environment variable
is already set then the pre-existing value is used (the config file
does not overwrite the value).

*/
func ReadEnv(cfgFile string, keys []string) (myEnv map[string]string, err error) {

	// No file specified, look for one
	if "" == cfgFile {

		appPath, errc := os.Executable()
		if errc != nil {
			return myEnv, fmt.Errorf("Error determining executable: ", errc)
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
		log.Println("No config file specified or found")
		return myEnv, nil
	}

	log.Printf("Using config file (%s)\n", cfgFile)
	var dotEnv map[string]string
	dotEnv, errc := godotenv.Read(cfgFile)

	if errc != nil {
		return myEnv, fmt.Errorf("Error loading config file (%s): %s", cfgFile, errc)
	}

	// Check the environment first and use that value if found. Next check the
	// config file and use that value if found.
	for _, k := range keys {
		eVal := os.Getenv(k)
		if "" != eVal {
			myEnv[k] = eVal
		} else {
			n, ok := dotEnv[k]
			if ok {
				myEnv[k] = n
			}
		}
	}

	return myEnv, nil
}
