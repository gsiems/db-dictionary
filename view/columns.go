package view

import (
	"html/template"
	"os"
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
)

type columnsView struct {
	Title         string
	DBMSVersion   string
	DBName        string
	DBComment     string
	SchemaName    string
	SchemaComment string
	TmspGenerated string
	Columns       []m.Column
}

func sortColumns(x []m.Column) {
	sort.Slice(x, func(i, j int) bool {

		switch strings.Compare(x[i].SchemaName, x[j].SchemaName) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].TableName, x[j].TableName) {
		case -1:
			return true
		case 1:
			return false
		}

		switch {
		case x[i].OrdinalPosition < x[j].OrdinalPosition:
			return true
		case x[i].OrdinalPosition > x[j].OrdinalPosition:
			return false
		}

		return x[i].Name < x[j].Name
	})
}

func makeColumnList(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		context := columnsView{
			Title:         "Columns for " + md.Alias + "." + vs.Name,
			TmspGenerated: md.TmspGenerated,
			DBName:        md.Name,
			DBComment:     md.Comment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

		context.Columns = md.FindColumns(vs.Name, "")
		sortColumns(context.Columns)

		var pageParts []string
		pageParts = append(pageParts, pageHeader(1, md))
		pageParts = append(pageParts, tpltSchemaColumns())
		pageParts = append(pageParts, pageFooter())

		templates, err := template.New("doc").Funcs(template.FuncMap{
			"safeHTML": func(u string) template.HTML { return template.HTML(u) },
		}).Parse(strings.Join(pageParts, ""))
		if err != nil {
			return err
		}

		dirName := md.OutputDir + "/" + vs.Name
		_, err = os.Stat(dirName)
		if os.IsNotExist(err) {
			err = os.Mkdir(dirName, 0744)
			if err != nil {
				return err
			}
		}

		outfile, err := os.Create(dirName + "/columns.html")
		if err != nil {
			return err
		}
		defer outfile.Close()

		err = templates.Lookup("doc").Execute(outfile, context)
		if err != nil {
			return err
		}
	}

	return err
}
