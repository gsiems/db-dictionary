package view

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	m "github.com/gsiems/db-dictionary-core/model"
)

// CreateDictionary orchestrates the creation of a data dictionary
func CreateDictionary(md *m.MetaData) (err error) {

	err = initOutputDir(md)
	if err != nil {
		return err
	}

	err = makeCSS(md)
	if err != nil {
		return err
	}

	err = makeImg(md)
	if err != nil {
		return err
	}

	err = makeJS(md)
	if err != nil {
		return err
	}

	err = makeSchemaList(md)
	if err != nil {
		return err
	}
	if md.Cfg.Verbose {
		log.Println("finished generating schema page")
	}

	err = makeTableList(md)
	if err != nil {
		return err
	}
	if md.Cfg.Verbose {
		log.Println("finished generating table list pages")
	}

	err = makeColumnList(md)
	if err != nil {
		return err
	}
	if md.Cfg.Verbose {
		log.Println("finished generating column list pages")
	}

	err = makeConstraintsList(md)
	if err != nil {
		return err
	}
	if md.Cfg.Verbose {
		log.Println("finished generating constraint list pages")
	}

	err = makeDomainsList(md)
	if err != nil {
		return err
	}
	if md.Cfg.Verbose {
		log.Println("finished generating domain list pages")
	}

	err = makeTablePages(md)
	if err != nil {
		return err
	}
	if md.Cfg.Verbose {
		log.Println("finished generating table pages")
	}

	return err
}

// initOutputDir ensures that the target directory for a data dictionary exists
func initOutputDir(md *m.MetaData) (err error) {
	if md.OutputDir != "." {
		err = ensurePath(md.OutputDir)
	}
	return err
}

// ensurePath checks if a directory exists and attempts to create the directory if it does not
func ensurePath(dirName string) (err error) {
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0745)
	}
	return err
}

// copyFileList copies the comma-separated list of files to the specified directory
func copyFileList(dirName, files string) (err error) {

	if files == "" {
		return err
	}

	f := strings.Split(files, ",")
	for _, source := range f {

		input, err := ioutil.ReadFile(source)
		if err != nil {
			return err
		}

		target := dirName + "/" + path.Base(source)
		err = ioutil.WriteFile(target, input, 0644)
		if err != nil {
			return err
		}
	}

	return err
}

// writeFile writes the contents string to the specified filename
func writeFile(fileName, contents string) (err error) {

	outfile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outfile.Close()

	_, err = outfile.WriteString(contents)
	return err
}
