package view

import (
	m "github.com/gsiems/db-dictionary-core/model"
	t "github.com/gsiems/db-dictionary-core/template"
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

	md.SortSchemas(context.Schemas)

	var tmplt t.T
	tmplt.AddPageHeader(0, md)
	tmplt.AddSnippet("Schemas")
	tmplt.AddPageFooter(0, md)

	err = tmplt.RenderPage(md.OutputDir, "index", context, md.Cfg.Minify)
	if err != nil {
		return err
	}

	return err
}
