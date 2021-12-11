package model

import (
	m "github.com/gsiems/go-db-meta/model"
)

type Schema struct {
	DBName       string
	Name         string
	Owner        string
	CharacterSet string
	Comment      string
}

func Schemas(s *[]m.Schema) (r []Schema, err error) {

	for _, v := range *s {
		r = append(r, Schema{
			DBName:       v.CatalogName.String,
			Name:         v.SchemaName.String,
			Owner:        v.SchemaOwner.String,
			CharacterSet: v.DefaultCharacterSetName.String,
			Comment:      v.Comment.String,
		})
	}

	return r, err
}
