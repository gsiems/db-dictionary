package view

import (
	"sort"

	m "github.com/gsiems/db-dictionary/model"
)

func sortIndexes(indexes []m.Index) {
	sort.Slice(indexes, func(i, j int) bool {
		return indexes[i].TableName < indexes[j].TableName && indexes[i].Name < indexes[j].Name
	})
}
