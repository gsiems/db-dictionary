package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

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
	//is_deferrable
	//initially_deferred
	Comment string
}

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
			Comment:           v.Comment.String,
		}
		md.ForeignKeys = append(md.ForeignKeys, fk)
	}
	fmt.Printf("%d foreign loaded\n", len(md.ForeignKeys))
}

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
