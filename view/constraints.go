package view

import (
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
)

type constraintsView struct {
	Title             string
	PathPrefix        string
	DBMSVersion       string
	DBName            string
	DBComment         string
	SchemaName        string
	SchemaComment     string
	TmspGenerated     string
	CheckConstraints  []m.CheckConstraint
	UniqueConstraints []m.UniqueConstraint
	ParentKeys        []m.ForeignKey
	//Indexes           []m.Index
	//ChildKeys         []m.ForeignKey
}

// Sort functions for Check Constraints
func sortCheckConstraints(x []m.CheckConstraint) {
	sort.Slice(x, func(i, j int) bool {

		switch strings.Compare(x[i].SchemaName, x[j].SchemaName) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].TableName, x[j].TableName) {
		case -1:
			return true
		case 1:
			return false
		}

		return x[i].Name < x[j].Name
	})
}

// Sort function for Unique Constraints
func sortUniqueConstraints(x []m.UniqueConstraint) {
	sort.Slice(x, func(i, j int) bool {

		switch strings.Compare(x[i].SchemaName, x[j].SchemaName) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].TableName, x[j].TableName) {
		case -1:
			return true
		case 1:
			return false
		}

		return x[i].Name < x[j].Name
	})
}

// Sort function for Foreign Keys
func sortForeignKeys(x []m.ForeignKey) {
	sort.Slice(x, func(i, j int) bool {

		switch strings.Compare(x[i].SchemaName, x[j].SchemaName) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].TableName, x[j].TableName) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].Name, x[j].Name) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].RefSchemaName, x[j].RefSchemaName) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].RefTableName, x[j].RefTableName) {
		case -1:
			return true
		case 1:
			return false
		}

		return x[i].RefConstraintName > x[j].RefConstraintName
	})
}
