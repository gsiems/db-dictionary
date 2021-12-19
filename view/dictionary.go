package view

import (
	"os"

	m "github.com/gsiems/db-dictionary/model"
)

func CreateDictionary(md *m.MetaData) (err error) {

	err = initOutputDir(md)
	if err != nil {
		return err
	}

	err = makeCSS(md)
	if err != nil {
		return err
	}

	err = makeSchemaList(md)
	if err != nil {
		return err
	}

	err = makeTableList(md)
	if err != nil {
		return err
	}

	err = makeColumnList(md)
	if err != nil {
		return err
	}

	err = makeTablePages(md)
	if err != nil {
		return err
	}

	return err
}

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
