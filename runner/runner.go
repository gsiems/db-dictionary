// Package runner takes an instance of a go-db-meta DBMS and initializes/loads
// a metadata (model) structure then creates a data dictionary
package runner

import (
	"github.com/gsiems/db-dictionary-core/config"
	"github.com/gsiems/db-dictionary-core/model"
	"github.com/gsiems/db-dictionary-core/view"

	d "github.com/gsiems/go-db-meta/dbms"
)

// MakeDictionary does the initialize, load, and calls CreateDictionary to
// make a data dictionary
func MakeDictionary(db *d.DBMS, cfg config.Config) error {

	md := model.Init(cfg)

	////////////////////////////////////////////////////////////////////////////
	catalog, err := db.CurrentCatalog()
	if err != nil {
		return err
	}
	md.LoadCatalog(&catalog)

	schemata, err := db.Schemata(cfg.Schemas, cfg.ExcludeSchemas)
	if err != nil {
		return err
	}
	md.LoadSchemas(&schemata)

	tables, err := db.Tables("")
	if err != nil {
		return err
	}
	md.LoadTables(&tables)

	columns, err := db.Columns("", "")
	if err != nil {
		return err
	}
	md.LoadColumns(&columns)

	indexes, err := db.Indexes("", "")
	if err != nil {
		return err
	}
	md.LoadIndexes(&indexes)

	checkConstraints, err := db.CheckConstraints("", "")
	if err != nil {
		return err
	}
	md.LoadCheckConstraints(&checkConstraints)

	domains, err := db.Domains("")
	if err != nil {
		return err
	}
	md.LoadDomains(&domains)

	primaryKeys, err := db.PrimaryKeys("", "")
	if err != nil {
		return err
	}
	md.LoadPrimaryKeys(&primaryKeys)

	foreignKeys, err := db.ReferentialConstraints("", "")
	if err != nil {
		return err
	}
	md.LoadForeignKeys(&foreignKeys)

	uniqueConstraints, err := db.UniqueConstraints("", "")
	if err != nil {
		return err
	}
	md.LoadUniqueConstraints(&uniqueConstraints)

	dependencies, err := db.Dependencies("", "")
	if err != nil {
		return err
	}
	md.LoadDependencies(&dependencies)

	userTypes, err := db.Types("")
	if err != nil {
		return err
	}
	md.LoadUserTypes(&userTypes)

	//////////////////////////////////////////////////////////////////////////////
	err = view.CreateDictionary(md)
	if err != nil {
		return err
	}

	return err
}
