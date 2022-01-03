package view

import (
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary/model"
)

// sortDependencies sets the default sort order for a list of object dependencies
func sortDependencies(x []m.Dependency) {
	sort.Slice(x, func(i, j int) bool {

		switch strings.Compare(x[i].ObjectSchema, x[j].ObjectSchema) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].ObjectName, x[j].ObjectName) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].DepObjectSchema, x[j].DepObjectSchema) {
		case -1:
			return true
		case 1:
			return false
		}

		switch strings.Compare(x[i].ObjectName, x[j].ObjectName) {
		case -1:
			return true
		case 1:
			return false
		}

		return x[i].DepObjectName > x[j].DepObjectName
	})
}
