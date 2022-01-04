package view

import (
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary-core/model"
)

// sortIndexes sets the default sort order for a list of indices
func sortIndexes(x []m.Index) {
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
