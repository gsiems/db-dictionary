package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// Index contains the metadata for an index
type Index struct {
	DBName       string
	SchemaName   string
	Name         string
	IndexType    string
	IndexColumns string
	TableSchema  string
	TableName    string
	IsUnique     string
	Comment      string
}

// LoadIndexes loads the index information from go-db-meta
// into the dictionary metadata structure
func (md *MetaData) LoadIndexes(x *[]m.Index) {
	for _, v := range *x {
		idx := Index{
			DBName:       v.TableCatalog.String,
			TableSchema:  md.chkSchemaName(v.TableSchema.String),
			TableName:    v.TableName.String,
			SchemaName:   md.chkSchemaName(v.IndexSchema.String),
			Name:         v.IndexName.String,
			IndexType:    v.IndexType.String,
			IndexColumns: v.IndexColumns.String,
			IsUnique:     v.IsUnique.String,
			Comment:      md.renderComment(v.Comment.String),
		}
		md.Indexes = append(md.Indexes, idx)
	}
	fmt.Printf("%d indexes loaded\n", len(md.Indexes))
}

// FindIndexes returns the index metadata that matches the specified
// schema/table name. If no schema is specified then all indices are returned.
// If only the schema is specified then all indices for that schema are returned.
func (md *MetaData) FindIndexes(schemaName string, tableName string) (d []Index) {

	for _, v := range md.Indexes {
		if schemaName == "" || schemaName == v.SchemaName {
			if tableName == "" || tableName == v.TableName {
				d = append(d, v)
			}
		}
	}

	return d
}
