package view

import (
	"log"
	"os"

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
		_, err = os.Stat(md.OutputDir)
		if os.IsNotExist(err) {
			err = os.MkdirAll(md.OutputDir, 0744)
			if err != nil {
				return err
			}
		}

	}
	return err
}
