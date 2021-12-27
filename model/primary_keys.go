package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

type PrimaryKey struct {
	DBName     string
	SchemaName string
	TableName  string
	Name       string
	Columns    string
	Status     string
	Comment    string
}

func (md *MetaData) LoadPrimaryKeys(x *[]m.PrimaryKey) {
	for _, v := range *x {
		pk := PrimaryKey{
			DBName:     v.TableCatalog.String,
			SchemaName: md.chkSchemaName(v.TableSchema.String),
			TableName:  v.TableName.String,
			Name:       v.ConstraintName.String,
			Columns:    v.ConstraintColumns.String,
			Status:     v.ConstraintStatus.String,
			Comment:    md.renderComment(v.Comment.String),
		}
		md.PrimaryKeys = append(md.PrimaryKeys, pk)
	}
	fmt.Printf("%d primary keys loaded\n", len(md.PrimaryKeys))
}

func (md *MetaData) FindPrimaryKeys(schemaName string, tableName string) (d []PrimaryKey) {

	for _, v := range md.PrimaryKeys {
		if schemaName == "" || schemaName == v.SchemaName {
			if tableName == "" || tableName == v.TableName {
				d = append(d, v)
			}
		}
	}

	return d
}
