package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

type CheckConstraint struct {
	DBName      string
	SchemaName  string
	TableName   string
	Name        string
	CheckClause string
	Status      string
	Comment     string
}

func (md *MetaData) LoadCheckConstraints(x *[]m.CheckConstraint) {
	for _, v := range *x {
		chk := CheckConstraint{
			DBName:      v.TableCatalog.String,
			SchemaName:  md.chkSchemaName(v.TableSchema.String),
			TableName:   v.TableName.String,
			Name:        v.ConstraintName.String,
			CheckClause: v.CheckClause.String,
			Status:      v.Status.String,
			Comment:     v.Comment.String,
		}
		md.CheckConstraints = append(md.CheckConstraints, chk)
	}
	fmt.Printf("%d check constraints loaded\n", len(md.CheckConstraints))
}

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
