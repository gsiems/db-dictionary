package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

type Table struct {
	DBName     string
	SchemaName string
	Name       string
	Owner      string
	TableType  string
	RowCount   int64
	Comment    string
	Query      string
}

func (md *MetaData) LoadTables(x *[]m.Table) {
	for _, v := range *x {
		table := Table{
			DBName:     v.TableCatalog.String,
			SchemaName: md.chkSchemaName(v.TableSchema.String),
			Name:       v.TableName.String,
			Owner:      v.TableOwner.String,
			TableType:  v.TableType.String,
			RowCount:   v.RowCount.Int64,
			Query:      v.ViewDefinition.String,
			Comment:    md.renderComment(v.Comment.String),
		}
		md.Tables = append(md.Tables, table)
	}
	fmt.Printf("%d tables loaded\n", len(md.Tables))
}

func (md *MetaData) FindTables(schemaName string) (d []Table) {

	for _, v := range md.Tables {
		if schemaName == "" || schemaName == v.SchemaName {
			d = append(d, v)
		}
	}

	return d
}
