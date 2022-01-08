package model

import (
	"log"

	m "github.com/gsiems/go-db-meta/model"
)

// Dependency contains the metadata for a database object dependency
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

// LoadDependencies loads the object dependency information from go-db-meta
// into the dictionary metadata structure
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
	if !md.Cfg.Quiet {
		log.Printf("%d dependencies loaded\n", len(md.Dependencies))
	}
}

// FindDependencies returns the object metadata that the specified schema/object
// is dependent on. If no schema is specified then all dependencies are returned.
// If only the schema is specified then all dependencies for objects in that
// schema are returned.
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

// FindDependents returns the object metadata for the database objects that
// depend on the specified schema/object. If no schema is specified then all
// dependencies are returned. If only the schema is specified then all objects
// that depend on objects in that schema are returned.
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
