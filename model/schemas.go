package model

import (
	"log"

	m "github.com/gsiems/go-db-meta/model"
)

// Schema contains the (high level) metadata for a schema
type Schema struct {
	DBName       string
	Name         string
	Owner        string
	CharacterSet string
	Comment      string
}

// chkSchemaName checks to ensure that a schema name is specified. If no
// schema name was specified (i.e. SQLite doesn't have schemas) then the
// default "main" value is returned
func (md *MetaData) chkSchemaName(s string) string {
	if s == "" {
		return "main"
	}
	return s
}

// LoadSchemas loads the schema information from go-db-meta
// into the dictionary metadata structure
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
	if md.Cfg.Verbose {
		log.Printf("%d schemas loaded\n", len(md.Schemas))
	}
}
