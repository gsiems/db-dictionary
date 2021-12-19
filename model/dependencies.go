package model

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

type Dependency struct {
	DBName          string
	ObjectSchema    string
	ObjectName      string
	ObjectOwner     string
	ObjectType      string
	DepDBName       string
	DepObjectSchema string
	DepObjectName   string
	DepObjectOwner  string
	DepObjectType   string
}

func (md *MetaData) LoadDependencies(x *[]m.Dependency) {
	for _, v := range *x {
		dependency := Dependency{

			DBName:          v.ObjectCatalog.String,
			ObjectSchema:    md.chkSchemaName(v.ObjectSchema.String),
			ObjectName:      v.ObjectName.String,
			ObjectOwner:     v.ObjectOwner.String,
			ObjectType:      v.ObjectType.String,
			DepDBName:       v.DepObjectCatalog.String,
			DepObjectSchema: md.chkSchemaName(v.DepObjectSchema.String),
			DepObjectName:   v.DepObjectName.String,
			DepObjectOwner:  v.DepObjectOwner.String,
			DepObjectType:   v.DepObjectType.String,
		}
		md.Dependencies = append(md.Dependencies, dependency)
	}
	fmt.Printf("%d dependencies loaded\n", len(md.Dependencies))
}

func (md *MetaData) FindDependencies(schemaName string, objectName string) (d []Dependency) {

	for _, v := range md.Dependencies {
		if schemaName == "" || schemaName == v.ObjectSchema {
			if objectName == "" || objectName == v.ObjectName {
				d = append(d, v)
			}
		}
	}

	return d
}

func (md *MetaData) FindDependents(schemaName string, objectName string) (d []Dependency) {

	for _, v := range md.Dependencies {
		if schemaName == "" || schemaName == v.DepObjectSchema {
			if objectName == "" || objectName == v.DepObjectName {
				d = append(d, v)
			}
		}
	}

	return d
}
