package view

import (
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary-core/model"
	t "github.com/gsiems/db-dictionary-core/template"
)

// columnsView contains the data used for generating the schema columns page
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

// sortColumns sets the default sort order for a list of columns
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

// makeColumnList marshals the data needed for, and then creates, a schema columns page
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

		var tmplt t.T
		tmplt.AddPageHeader(1, md)
		tmplt.AddSnippet("SchemaColumns")
		tmplt.AddPageFooter()

		dirName := md.OutputDir + "/" + vs.Name
		err = tmplt.RenderPage(dirName, "columns", context)
		if err != nil {
			return err
		}
	}

	return err
}
