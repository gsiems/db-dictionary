package model

import (
	"log"

	m "github.com/gsiems/go-db-meta/model"
)

// UniqueConstraint contains the metadata for a unique constraint
type UniqueConstraint struct {
	DBName     string
	SchemaName string
	TableName  string
	Name       string
	Columns    string
	Status     string
	Comment    string
}

// LoadUniqueConstraints loads the unique constraint information from go-db-meta
// into the dictionary metadata structure
func (md *MetaData) LoadUniqueConstraints(x *[]m.UniqueConstraint) {
	for _, v := range *x {
		chk := UniqueConstraint{
			DBName:     v.TableCatalog.String,
			SchemaName: md.chkSchemaName(v.TableSchema.String),
			TableName:  v.TableName.String,
			Name:       v.ConstraintName.String,
			Columns:    v.ConstraintColumns.String,
			Status:     v.Status.String,
			Comment:    md.renderComment(v.Comment.String),
		}
		md.UniqueConstraints = append(md.UniqueConstraints, chk)
	}
	if !md.Cfg.Quiet {
		log.Printf("%d unique constraints loaded\n", len(md.UniqueConstraints))
	}
}

// FindUniqueConstraints returns the unique contraint metadata that matches the
// specified schema/table name. If no schema is specified then all unique
// constraints are returned. If only the schema is specified then all unique
// constraints for that schema are returned.
func (md *MetaData) FindUniqueConstraints(schemaName string, tableName string) (d []UniqueConstraint) {

	for _, v := range md.UniqueConstraints {
		if schemaName == "" || schemaName == v.SchemaName {
			if tableName == "" || tableName == v.TableName {
				d = append(d, v)
			}
		}
	}

	return d
}
