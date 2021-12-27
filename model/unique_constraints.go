package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

type UniqueConstraint struct {
	DBName     string
	SchemaName string
	TableName  string
	Name       string
	Columns    string
	Status     string
	Comment    string
}

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
	fmt.Printf("%d check constraints loaded\n", len(md.UniqueConstraints))
}

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
