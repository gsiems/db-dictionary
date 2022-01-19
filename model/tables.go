package model

import (
	"log"
	"sort"
	"strings"

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
			Comment:    md.renderComment(v.Comment.String),
		}
		if !md.Cfg.HideSQL {
			table.Query = v.ViewDefinition.String
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

// SortTables sets the default sort order for a list of tables
func (md *MetaData) SortTables(x []Table) {
	sort.Slice(x, func(i, j int) bool {
		switch strings.Compare(x[i].SchemaName, x[j].SchemaName) {
		case -1:
			return true
		case 1:
			return false
		}

		return x[i].Name < x[j].Name
	})
}
