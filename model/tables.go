package model

import (
	m "github.com/gsiems/go-db-meta/model"
)

type Column struct {
	Name            string
	OrdinalPosition int32
	IsNullable      string
	DataType        string
	Default         string
	DomainCatalog   string
	DomainSchema    string
	DomainName      string
	Comment         string
}

type Table struct {
	DBName         string
	SchemaName     string
	Name           string
	Owner          string
	TableType      string
	RowCount       int64
	Comment        string
	ViewDefinition string
	Columns        []Column
}

func Tables(t *[]m.Table, c *[]m.Column) (r []Table, err error) {

	for _, vt := range *t {

		table := Table{
			DBName:         vt.TableCatalog.String,
			SchemaName:     vt.TableSchema.String,
			Name:           vt.TableName.String,
			Owner:          vt.TableOwner.String,
			TableType:      vt.TableType.String,
			RowCount:       vt.RowCount.Int64,
			ViewDefinition: vt.ViewDefinition.String,
			Comment:        vt.Comment.String,
		}

		for _, vc := range *c {
			if vc.TableCatalog.String == vt.TableCatalog.String && vc.TableSchema.String == vt.TableSchema.String && vc.TableName.String == vt.TableName.String {

				table.Columns = append(table.Columns, Column{
					Name:            vc.ColumnName.String,
					OrdinalPosition: vc.OrdinalPosition.Int32,
					IsNullable:      vc.IsNullable.String,
					DataType:        vc.DataType.String,
					Default:         vc.ColumnDefault.String,
					DomainCatalog:   vc.DomainCatalog.String,
					DomainSchema:    vc.DomainSchema.String,
					DomainName:      vc.DomainName.String,
					Comment:         vc.Comment.String,
				})
			}
		}

		r = append(r, table)
	}

	return r, err
}
