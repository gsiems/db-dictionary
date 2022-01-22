package view

import (
	"path"

	m "github.com/gsiems/db-dictionary/model"
	t "github.com/gsiems/db-dictionary/template"
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
			md.SortCheckConstraints(context.CheckConstraints)
		}

		// unique constraints
		context.UniqueConstraints = md.FindUniqueConstraints(vs.Name, "")
		if len(context.UniqueConstraints) > 0 {
			tmplt.AddSectionHeader("Unique constraints")
			tmplt.AddSnippet("SchemaUniqueConstraints")
			md.SortUniqueConstraints(context.UniqueConstraints)
		}

		// foreign keys
		context.ParentKeys = md.FindParentKeys(vs.Name, "")
		if len(context.ParentKeys) > 0 {
			tmplt.AddSectionHeader("Foreign key constraints")
			tmplt.AddSnippet("SchemaFKConstraints")
			md.SortForeignKeys(context.ParentKeys)
		}

		if len(context.CheckConstraints) == 0 && len(context.UniqueConstraints) == 0 && len(context.ParentKeys) == 0 {
			tmplt.AddSnippet("      <p><b>No constraints extracted for this schema.</b></p>")
		}

		tmplt.AddPageFooter(1, md)

		dirName := path.Join(md.OutputDir, vs.Name)
		err = tmplt.RenderPage(dirName, "constraints", context, md.Cfg.Minify)
		if err != nil {
			return err
		}
	}

	return err
}
