package view

import (
	"sort"

	m "github.com/gsiems/db-dictionary/model"
	t "github.com/gsiems/db-dictionary/template"
)

// schemasView contains the data used for generating the schemas list page
type schemasView struct {
	Title         string
	TmspGenerated string
	DBMSVersion   string
	DBName        string
	DBComment     string
	Schemas       []m.Schema
}

// sortSchemas sets the default sort order for a list of schemas
func sortSchemas(schemas []m.Schema) {
	sort.Slice(schemas, func(i, j int) bool {
		return schemas[i].Name < schemas[j].Name
	})
}

// makeSchemaList marshals the data needed for, and then creates, a database schemas page
func makeSchemaList(md *m.MetaData) (err error) {

	context := schemasView{
		Title:         "Schemas for " + md.Alias,
		TmspGenerated: md.TmspGenerated,
		DBMSVersion:   md.Version,
		DBName:        md.Name,
		DBComment:     md.Comment,
		Schemas:       md.Schemas,
	}

	sortSchemas(context.Schemas)

	var tmplt t.T
	tmplt.AddPageHeader(0, md)
	tmplt.AddSnippet("Schemas")
	tmplt.AddPageFooter()

	err = tmplt.RenderPage(md.OutputDir, "index", context)
	if err != nil {
		return err
	}

	return err
}
