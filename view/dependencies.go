package view

import (
	//"html/template"
	//"os"
	"sort"

	m "github.com/gsiems/db-dictionary/model"
)

func sortDependencies(dependencies []m.Dependency) {
	sort.Slice(dependencies, func(i, j int) bool {
		return dependencies[i].ObjectSchema < dependencies[j].ObjectSchema && dependencies[i].ObjectName < dependencies[j].ObjectName && dependencies[i].DepObjectSchema < dependencies[j].DepObjectSchema && dependencies[i].DepObjectName < dependencies[j].DepObjectName
	})
}
