package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

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
