package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// ForeignKey contains the metadata for a foreign key constraint
type ForeignKey struct {
	DBName            string
	SchemaName        string
	TableName         string
	TableColumns      string
	Name              string
	RefDBName         string
	RefSchemaName     string
	RefTableName      string
	RefTableColumns   string
	RefConstraintName string
	MatchOption       string
	UpdateRule        string
	DeleteRule        string
	IsEnforced        string
	IsIndexed         string
	Comment           string
	//is_deferrable
	//initially_deferred
}

// LoadForeignKeys loads the foreign key relationship information from
// go-db-meta into the dictionary metadata structure
func (md *MetaData) LoadForeignKeys(x *[]m.ReferentialConstraint) {
	for _, v := range *x {
		fk := ForeignKey{
			DBName:            v.TableCatalog.String,
			SchemaName:        md.chkSchemaName(v.TableSchema.String),
			TableName:         v.TableName.String,
			TableColumns:      v.TableColumns.String,
			Name:              v.ConstraintName.String,
			RefDBName:         v.RefTableCatalog.String,
			RefSchemaName:     md.chkSchemaName(v.RefTableSchema.String),
			RefTableName:      v.RefTableName.String,
			RefTableColumns:   v.RefTableColumns.String,
			RefConstraintName: v.RefConstraintName.String,
			MatchOption:       v.MatchOption.String,
			UpdateRule:        v.UpdateRule.String,
			DeleteRule:        v.DeleteRule.String,
			IsEnforced:        v.IsEnforced.String,
			Comment:           md.renderComment(v.Comment.String),
		}
		md.ForeignKeys = append(md.ForeignKeys, fk)
	}
	fmt.Printf("%d foreign loaded\n", len(md.ForeignKeys))
	md.tagIndexedFKs()
}

// tagIndexedFKs compares the foreign key metadata with the index metadata to
// determine which child keys are indexed
func (md *MetaData) tagIndexedFKs() {

	idxs := make(map[string]int)
	for _, i := range md.Indexes {
		mk := i.TableSchema + "." + i.TableName + "." + i.IndexColumns
		idxs[mk] = 0
	}

	for k, v := range md.ForeignKeys {
		mk := v.SchemaName + "." + v.TableName + "." + v.TableColumns
		_, ok := idxs[mk]
		if ok {
			md.ForeignKeys[k].IsIndexed = "YES"
		}
	}

}

// FindChildKeys returns the foreign key contraint metadata for all child
// tables of the specified schema/table name. If no schema is specified then all
// foreign key constraints are returned. If only the schema is specified then
// all foreign keys for the child tables for that schema are returned.
func (md *MetaData) FindChildKeys(schemaName string, tableName string) (d []ForeignKey) {

	for _, v := range md.ForeignKeys {
		if schemaName == "" || schemaName == v.RefSchemaName {
			if tableName == "" || tableName == v.RefTableName {
				d = append(d, v)
			}
		}
	}

	return d
}

// FindParentKeys returns the foreign key contraint metadata for all parent keys
// that match the specified schema/table name. If no schema is specified then all
// foreign key are returned. If only the schema is specified then all foreign key
// constraints for that parent tables for that schema are returned.
func (md *MetaData) FindParentKeys(schemaName string, tableName string) (d []ForeignKey) {

	for _, v := range md.ForeignKeys {
		if schemaName == "" || schemaName == v.SchemaName {
			if tableName == "" || tableName == v.TableName {
				d = append(d, v)
			}
		}
	}

	return d
}
