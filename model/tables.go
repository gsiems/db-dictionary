package model

import (
	m "github.com/gsiems/go-db-meta/model"
)

type Table struct {
	DBName         string
	SchemaName     string
	Name           string
	Owner          string
	TableType      string
	RowCount       int64
	Comment        string
	ViewDefinition string
}

func Tables(s *[]m.Table) (r []Table, err error) {

	for _, v := range *s {
		r = append(r, Table{
			DBName:         v.TableCatalog.String,
			SchemaName:     v.TableSchema.String,
			Name:           v.TableName.String,
			Owner:          v.TableOwner.String,
			TableType:      v.TableType.String,
			RowCount:       v.RowCount.Int64,
			ViewDefinition: v.ViewDefinition.String,
			Comment:        v.Comment.String,
		})
	}

	return r, err
}
