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
	PathPrefix    string
	DBMSVersion   string
	DBName        string
	DBComment     string
	SchemaName    string
	SchemaComment string
	TmspGenerated string
	Columns       []m.Column
}

func sortColumns(columns []m.Column) {
	sort.Slice(columns, func(i, j int) bool {
		return columns[i].TableName < columns[j].TableName && columns[i].OrdinalPosition < columns[j].OrdinalPosition && columns[i].Name < columns[j].Name
	})
}

func makeColumnList(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		context := columnsView{
			Title:         "Columns for " + md.Alias + "." + vs.Name,
			PathPrefix:    "../",
			TmspGenerated: md.TmspGenerated,
			DBName:        md.Name,
			DBComment:     md.Comment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

		context.Columns = md.FindColumns(vs.Name, "")
		sortColumns(context.Columns)

		var pageParts []string
		pageParts = append(pageParts, pageHeader())
		pageParts = append(pageParts, tpltSchemaColumns())
		pageParts = append(pageParts, pageFooter())

		templates, err := template.New("doc").Parse(strings.Join(pageParts, ""))
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
