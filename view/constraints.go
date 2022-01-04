package view

import (
	"sort"
	"strings"

	m "github.com/gsiems/db-dictionary-core/model"
	t "github.com/gsiems/db-dictionary-core/template"
)

// constraintsView contains the data used for generating the schema constraints page
type constraintsView struct {
	Title             string
	DBMSVersion       string
	DBName            string
	DBComment         string
	SchemaName        string
	SchemaComment     string
	TmspGenerated     string
	CheckConstraints  []m.CheckConstraint
	UniqueConstraints []m.UniqueConstraint
	ParentKeys        []m.ForeignKey
}

// sortCheckConstraints sets the default sort order for a list of check constraints
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

// sortUniqueConstraints sets the default sort order for a list of unique constraints
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

// sortForeignKeys sets the default sort order for a list of foreign key constraints
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

// makeConstraintsList marshals the data needed for, and then creates, a schema constraints page
func makeConstraintsList(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		context := constraintsView{
			Title:         "Constraints for " + md.Alias + "." + vs.Name,
			TmspGenerated: md.TmspGenerated,
			DBName:        md.Name,
			DBComment:     md.Comment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

		var tmplt t.T
		tmplt.AddPageHeader(1, md)
		tmplt.AddSnippet("SchemaConstraintsHeader")

		// check constraints
		context.CheckConstraints = md.FindCheckConstraints(vs.Name, "")
		if len(context.CheckConstraints) > 0 {
			tmplt.AddSectionHeader("Check constraints")
			tmplt.AddSnippet("SchemaCheckConstraints")
			sortCheckConstraints(context.CheckConstraints)
		}

		// unique constraints
		context.UniqueConstraints = md.FindUniqueConstraints(vs.Name, "")
		if len(context.UniqueConstraints) > 0 {
			tmplt.AddSectionHeader("Unique constraints")
			tmplt.AddSnippet("SchemaUniqueConstraints")
			sortUniqueConstraints(context.UniqueConstraints)
		}

		// foreign keys
		context.ParentKeys = md.FindParentKeys(vs.Name, "")
		if len(context.ParentKeys) > 0 {
			tmplt.AddSectionHeader("Foreign key constraints")
			tmplt.AddSnippet("SchemaFKConstraints")
			sortForeignKeys(context.ParentKeys)
		}

		if len(context.CheckConstraints) == 0 && len(context.UniqueConstraints) == 0 && len(context.ParentKeys) == 0 {
			tmplt.AddSnippet("      <p><b>No constraints extracted for this schema.</b></p>")
		}

		tmplt.AddPageFooter()

		dirName := md.OutputDir + "/" + vs.Name
		err = tmplt.RenderPage(dirName, "constraints", context)
		if err != nil {
			return err
		}
	}

	return err
}
