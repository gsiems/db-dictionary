package view

import (
	"path"

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
			md.SortTables(context.Tables)
		} else {
			tmplt.AddSnippet("      <p><b>No tables extracted for this schema.</b></p>")
		}

		tmplt.AddPageFooter(1, md)

		dirName := path.Join(md.OutputDir, vs.Name)
		err = tmplt.RenderPage(dirName, "tables", context, md.Cfg.Minify)
		if err != nil {
			return err
		}
	}

	return err
}
