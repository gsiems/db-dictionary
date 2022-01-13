package view

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	m "github.com/gsiems/db-dictionary-core/model"
)

// makeImg copies any specified image files
func makeImg(md *m.MetaData) (err error) {

	if md.Cfg.ImgFiles == "" {
		return
	}

	dirName := md.OutputDir + "/img"
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0744)
		if err != nil {
			return err
		}
	}

	err = copyImgFiles(dirName, md.Cfg.ImgFiles)

	return err
}

func copyImgFiles(dirName, files string) (err error) {

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
