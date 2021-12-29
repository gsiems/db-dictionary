package view

import (
	"sort"

	m "github.com/gsiems/db-dictionary/model"
	t "github.com/gsiems/db-dictionary/template"
)

type schemasView struct {
	Title         string
	TmspGenerated string
	DBMSVersion   string
	DBName        string
	DBComment     string
	Schemas       []m.Schema
}

func sortSchemas(schemas []m.Schema) {
	sort.Slice(schemas, func(i, j int) bool {
		return schemas[i].Name < schemas[j].Name
	})
}

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

	err = tmplt.RenderPage(md.OutputDir, "index", context)
	if err != nil {
		return err
	}

	return err
}
