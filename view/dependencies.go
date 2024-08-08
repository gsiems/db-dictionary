package view

import (
	"fmt"
	"os"
	"path"

	"github.com/gsiems/db-dictionary/graph"
	m "github.com/gsiems/db-dictionary/model"
	t "github.com/gsiems/db-dictionary/template"
)

type DepFile struct {
	File   string
	Format string
}

// dependenciesView contains the data used for generating a dependencies information page
type dependenciesView struct {
	Title         string
	DBMSVersion   string
	DBName        string
	DBComment     string
	SchemaName    string
	SchemaComment string
	TmspGenerated string
	Files         []DepFile
}

func makeDependencyPages(md *m.MetaData) (err error) {

	if len(md.Dependencies) == 0 {
		return err
	}

	for _, vs := range md.Schemas {

		context := dependenciesView{
			Title:         "Dependencies for " + md.Alias + "." + vs.Name,
			TmspGenerated: md.TmspGenerated,
			DBName:        md.Name,
			DBComment:     md.Comment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

		err = graph.MakeDepenencyGraphs(&vs, md)
		if err != nil {
			return err
		}

		dirName := path.Join(md.OutputDir, vs.Name)
		// Add links to files
		context.Files, err = listDependencyFiles(dirName)
		if err != nil {
			return err
		}

		var tmplt t.T
		tmplt.AddPageHeader(1, md)
		tmplt.AddSnippet("SchemaDependencies")

		if len(context.Files) == 0 {
			tmplt.AddSnippet("      <p><b>No dependencies extracted for this schema.</b></p>")
		}

		tmplt.AddPageFooter(1, md)

		err = tmplt.RenderPage(dirName, "dependencies", context, md.Cfg.Minify)
		if err != nil {
			return err
		}
	}
	return err
}

func listDependencyFiles(dirName string) (d []DepFile, err error) {
	dir, cerr := os.Open(dirName)
	if cerr != nil {
		err = fmt.Errorf("failed opening directory: %s", cerr)
		return d, err
	}
	defer dir.Close()

	list, _ := dir.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		fileName := path.Base(name)

		switch fileName {
		case "dependencies.dot", "dependencies.gv":
			d = append(d, DepFile{File: fileName, Format: "Graphviz file"})
		case "dependencies.gml":
			d = append(d, DepFile{File: fileName, Format: "GraphML file (with columns)"})
		case "dependencies-min.gml":
			d = append(d, DepFile{File: fileName, Format: "GraphML file (no columns)"})
		}

	}
	return d, err
}
