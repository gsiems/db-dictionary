package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

type UserType struct {
	DBName     string
	SchemaName string
	Name       string
	Owner      string
	//DataType    string
	Comment string
}

func (md *MetaData) LoadUserTypes(x *[]m.Type) {
	for _, v := range *x {
		udt := UserType{
			DBName:     v.TypeCatalog.String,
			SchemaName: md.chkSchemaName(v.TypeSchema.String),
			Name:       v.TypeName.String,
			Owner:      v.TypeOwner.String,
			Comment:    md.renderComment(v.Comment.String),
		}
		md.UserTypes = append(md.UserTypes, udt)
	}
	fmt.Printf("%d types loaded\n", len(md.UserTypes))
}

func (md *MetaData) FindUserTypes(schemaName string) (d []UserType) {

	for _, v := range md.UserTypes {
		if schemaName == "" || schemaName == v.SchemaName {
			d = append(d, v)
		}
	}

	return d
}
