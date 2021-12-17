package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

type Column struct {
	DBName          string
	SchemaName      string
	TableName       string
	Name            string
	OrdinalPosition int32
	IsNullable      string
	DataType        string
	Default         string
	DomainDBName    string
	DomainSchema    string
	DomainName      string
	Comment         string
}

func (md *MetaData) LoadColumns(x *[]m.Column) {
	for _, v := range *x {
		column := Column{
			DBName:          v.TableCatalog.String,
			SchemaName:      md.chkSchemaName(v.TableSchema.String),
			TableName:       v.TableName.String,
			Name:            v.ColumnName.String,
			OrdinalPosition: v.OrdinalPosition.Int32,
			IsNullable:      v.IsNullable.String,
			DataType:        v.DataType.String,
			Default:         v.ColumnDefault.String,
			DomainDBName:    v.DomainCatalog.String,
			DomainSchema:    md.chkSchemaName(v.DomainSchema.String),
			DomainName:      v.DomainName.String,
			Comment:         v.Comment.String,
		}
		md.Columns = append(md.Columns, column)
	}
	fmt.Printf("%d columns loaded\n", len(md.Columns))
}

func (md *MetaData) FindColumns(schemaName string, tableName string) (d []Column) {

	for _, v := range md.Columns {
		if schemaName == "" || schemaName == v.SchemaName {
			if tableName == "" || tableName == v.TableName {
				d = append(d, v)
			}
		}
	}

	return d
}
