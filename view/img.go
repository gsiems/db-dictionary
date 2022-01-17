package view

import (
	m "github.com/gsiems/db-dictionary-core/model"
)

// makeImg copies any specified image files
func makeImg(md *m.MetaData) (err error) {

	if md.Cfg.ImgFiles == "" {
		return
	}

	dirName := md.OutputDir + "/img"
	err = ensurePath(dirName)
	if err != nil {
		return err
	}

	err = copyFileList(dirName, md.Cfg.ImgFiles)

	return err
}
