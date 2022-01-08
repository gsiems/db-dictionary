package model

import (
	"log"

	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKey contains the metadata for a primary key
type PrimaryKey struct {
	DBName     string
	SchemaName string
	TableName  string
	Name       string
	Columns    string
	Status     string
	Comment    string
}

// LoadPrimaryKeys loads the primary key information from go-db-meta
// into the dictionary metadata structure
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
	if !md.Cfg.Quiet {
		log.Printf("%d primary keys loaded\n", len(md.PrimaryKeys))
	}
}

// FindPrimaryKeys returns the primary key contraint metadata that matches the
// specified schema/table name. If no schema is specified then all primary
// key constraints are returned. If only the schema is specified then all primary
// key constraints for that schema are returned.
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
