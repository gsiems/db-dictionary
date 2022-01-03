package view

import (
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
	t "github.com/gsiems/db-dictionary/template"
)

// tablesView contains the data used for generating the schema tables page
type tablesView struct {
	Title         string
	DBMSVersion   string
	DBName        string
	DBComment     string
	SchemaName    string
	SchemaComment string
	TmspGenerated string
	Tables        []m.Table
}

// sortTables sets the default sort order for a list of tables
func sortTables(x []m.Table) {
	sort.Slice(x, func(i, j int) bool {
		switch strings.Compare(x[i].SchemaName, x[j].SchemaName) {
		case -1:
			return true
		case 1:
			return false
		}

		return x[i].Name < x[j].Name
	})
}

// makeTableList marshals the data needed for, and then creates, a schema list of
// tables/views/materialized views page
func makeTableList(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		context := tablesView{
			Title:         "Tables for " + md.Alias + "." + vs.Name,
			TmspGenerated: md.TmspGenerated,
			DBName:        md.Name,
			DBComment:     md.Comment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

		var tmplt t.T
		tmplt.AddPageHeader(1, md)

		context.Tables = md.FindTables(vs.Name)
		if len(context.Tables) > 0 {
			tmplt.AddSnippet("SchemaTables")
			sortTables(context.Tables)
		} else {
			tmplt.AddSnippet("      <p><b>No tables extracted for this schema.</b></p>")
		}

		tmplt.AddPageFooter()

		dirName := md.OutputDir + "/" + vs.Name

		err = tmplt.RenderPage(dirName, "tables", context)
		if err != nil {
			return err
		}

	}

	return err
}
