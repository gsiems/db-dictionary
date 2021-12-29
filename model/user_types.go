package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// UserType contains the metadata for a user defined datatype
type UserType struct {
	DBName     string
	SchemaName string
	Name       string
	Owner      string
	//DataType    string
	Comment string
}

// LoadUserTypes loads the user defined datatype information from go-db-meta
// into the dictionary metadata structure
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

// FindUserTypes returns the user-defined datatype metadata that matches the
// specified schema. If no schema is specified then all user-defined datatypes
// are returned.
func (md *MetaData) FindUserTypes(schemaName string) (d []UserType) {

	for _, v := range md.UserTypes {
		if schemaName == "" || schemaName == v.SchemaName {
			d = append(d, v)
		}
	}

	return d
}
