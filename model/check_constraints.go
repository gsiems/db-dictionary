package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// CheckConstraint contains the metadata for a check constraint
type CheckConstraint struct {
	DBName      string
	SchemaName  string
	TableName   string
	Name        string
	CheckClause string
	Status      string
	Comment     string
}

// LoadCheckConstraints loads the check constraint information from go-db-meta
// into the dictionary metadata structure
func (md *MetaData) LoadCheckConstraints(x *[]m.CheckConstraint) {
	for _, v := range *x {
		chk := CheckConstraint{
			DBName:      v.TableCatalog.String,
			SchemaName:  md.chkSchemaName(v.TableSchema.String),
			TableName:   v.TableName.String,
			Name:        v.ConstraintName.String,
			CheckClause: v.CheckClause.String,
			Status:      v.Status.String,
			Comment:     md.renderComment(v.Comment.String),
		}
		md.CheckConstraints = append(md.CheckConstraints, chk)
	}
	fmt.Printf("%d check constraints loaded\n", len(md.CheckConstraints))
}

// FindCheckConstraints returns the check contraint metadata that matches the
// specified schema/table name. If no schema is specified then all check
// constraints are returned. If only the schema is specified then all check
// constraints for that schema are returned.
func (md *MetaData) FindCheckConstraints(schemaName string, tableName string) (d []CheckConstraint) {

	for _, v := range md.CheckConstraints {
		if schemaName == "" || schemaName == v.SchemaName {
			if tableName == "" || tableName == v.TableName {
				d = append(d, v)
			}
		}
	}

	return d
}
