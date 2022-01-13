package model

import (
	"log"

	m "github.com/gsiems/go-db-meta/model"
)

// Table contains the (hight level) metadata for a table/view/materialized view.
// Other attributes (i.e. columns, indices, etc.) will have their own metadata
// structures.
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

// LoadTables loads the table information from go-db-meta
// into the dictionary metadata structure
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
	if md.Cfg.Verbose {
		log.Printf("%d tables loaded\n", len(md.Tables))
	}
}

// FindTables returns the table metadata that matches the specified schema.
// If no schema is specified then all tables are returned.
func (md *MetaData) FindTables(schemaName string) (d []Table) {

	for _, v := range md.Tables {
		if schemaName == "" || schemaName == v.SchemaName {
			d = append(d, v)
		}
	}

	return d
}
