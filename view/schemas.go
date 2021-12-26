package view

import (
	"html/template"
	"os"
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
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

	var pageParts []string
	pageParts = append(pageParts, pageHeader(0))
	pageParts = append(pageParts, tpltSchemas())
	pageParts = append(pageParts, pageFooter())

	templates, err := template.New("doc").Parse(strings.Join(pageParts, ""))
	if err != nil {
		return err
	}

	//
	outfile, err := os.Create(md.OutputDir + "/index.html")
	if err != nil {
		return err
	}
	defer outfile.Close()

	err = templates.Lookup("doc").Execute(outfile, context)
	if err != nil {
		return err
	}

	return err
}
