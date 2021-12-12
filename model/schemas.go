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
		schema := Schema{
			DBName:       v.CatalogName.String,
			Name:         v.SchemaName.String,
			Owner:        v.SchemaOwner.String,
			CharacterSet: v.DefaultCharacterSetName.String,
			Comment:      v.Comment.String,
		}

		if schema.Name == "" {
			schema.Name = "default"
		}

		r = append(r, schema)
	}

	return r, err
}
