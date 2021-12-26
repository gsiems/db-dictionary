package view

import (
	"html/template"
	"os"
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

func makeConstraintsList(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		context := constraintsView{
			Title:         "Constraints for " + md.Alias + "." + vs.Name,
			PathPrefix:    "../",
			TmspGenerated: md.TmspGenerated,
			DBName:        md.Name,
			DBComment:     md.Comment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

		var pageParts []string

		pageParts = append(pageParts, pageHeader())

		pageParts = append(pageParts, tpltSchemaConstraintsHeader())

		// check constraints
		context.CheckConstraints = md.FindCheckConstraints(vs.Name, "")
		if len(context.CheckConstraints) > 0 {
			pageParts = append(pageParts, sectionHeader("Check constraints"))
			pageParts = append(pageParts, tpltSchemaCheckConstraints())
			sortCheckConstraints(context.CheckConstraints)
		}

		// unique constraints
		context.UniqueConstraints = md.FindUniqueConstraints(vs.Name, "")
		if len(context.UniqueConstraints) > 0 {
			pageParts = append(pageParts, sectionHeader("Unique constraints"))
			pageParts = append(pageParts, tpltSchemaUniqueConstraints())
			sortUniqueConstraints(context.UniqueConstraints)
		}

		// foreign keys
		context.ParentKeys = md.FindParentKeys(vs.Name, "")
		if len(context.ParentKeys) > 0 {
			pageParts = append(pageParts, sectionHeader("Foreign key constraints"))
			pageParts = append(pageParts, tpltSchemaFKConstraints())
			sortForeignKeys(context.ParentKeys)
		}

		if len(context.CheckConstraints) == 0 && len(context.UniqueConstraints) == 0 && len(context.ParentKeys) == 0 {
			pageParts = append(pageParts, "      <p><b>No constraints extracted for this schema.</b></p>")
		}

		pageParts = append(pageParts, pageFooter())

		templates, err := template.New("doc").Parse(strings.Join(pageParts, ""))
		if err != nil {
			return err
		}

		dirName := md.OutputDir + "/" + vs.Name
		_, err = os.Stat(dirName)
		if os.IsNotExist(err) {
			err = os.Mkdir(dirName, 0744)
			if err != nil {
				return err
			}
		}

		outfile, err := os.Create(dirName + "/constraints.html")
		if err != nil {
			return err
		}
		defer outfile.Close()

		err = templates.Lookup("doc").Execute(outfile, context)
		if err != nil {
			return err
		}
	}

	return err

}
