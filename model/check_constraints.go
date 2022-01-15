package model

import (
	"log"
	"sort"
	"strings"

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

		// Don't want not null constraints (they're in the column description and
		// they really aren't interesting compared to the other check constraints
		// plus too many not null constraints can hide the other check constraints)
		if strings.Contains(strings.ToUpper(v.CheckClause.String), "IS NOT NULL") {
			continue
		}

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
	if md.Cfg.Verbose {
		log.Printf("%d check constraints loaded\n", len(md.CheckConstraints))
	}
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

// SortCheckConstraints sets the default sort order for a list of check constraints
func (md *MetaData) SortCheckConstraints(x []CheckConstraint) {
	sort.Slice(x, func(i, j int) bool {

		switch strings.Compare(x[i].SchemaName, x[j].SchemaName) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].TableName, x[j].TableName) {
		case -1:
			return true
		case 1:
			return false
		}

		return x[i].Name < x[j].Name
	})
}
