package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

type Schema struct {
	DBName       string
	Name         string
	Owner        string
	CharacterSet string
	Comment      string
}

func (md *MetaData) chkSchemaName(s string) string {
	if s == "" {
		return "main"
	}
	return s
}

func (md *MetaData) LoadSchemas(x *[]m.Schema) {
	for _, v := range *x {
		schema := Schema{
			DBName:       v.CatalogName.String,
			Name:         md.chkSchemaName(v.SchemaName.String),
			Owner:        v.SchemaOwner.String,
			CharacterSet: v.DefaultCharacterSetName.String,
			Comment:      md.renderComment(v.Comment.String),
		}
		md.Schemas = append(md.Schemas, schema)
	}
	fmt.Printf("%d schemas loaded\n", len(md.Schemas))
}
